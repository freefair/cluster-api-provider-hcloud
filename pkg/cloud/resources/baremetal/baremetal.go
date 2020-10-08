package baremetal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cluster-api-provider-hcloud/cluster-api-provider-hcloud/pkg/cloud/resources/server/userdata"
	"github.com/cluster-api-provider-hcloud/cluster-api-provider-hcloud/pkg/cloud/scope"
	"github.com/nl2go/hrobot-go/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"
	capierrors "sigs.k8s.io/cluster-api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	delimiter                      string        = "--"
	hoursBeforeDeletion            time.Duration = 36
	maxWaitingTimeForSSHConnection int           = 300
)

type Service struct {
	scope *scope.BareMetalMachineScope
}

type blockDeviceData struct {
	Name     string `json:"name"`
	Size     string `json:"size"`
	Rotation bool   `json:"rota"`
	Type     string `json:"fstype"`
	Label    string `json:"label"`
}

type blockDevice struct {
	Name     string            `json:"name"`
	Size     string            `json:"size"`
	Rotation bool              `json:"rota"`
	Type     string            `json:"fstype"`
	Label    string            `json:"label"`
	Children []blockDeviceData `json:"children"`
}

type blockDevices struct {
	Devices []blockDevice `json:"blockdevices"`
}

func NewService(scope *scope.BareMetalMachineScope) *Service {
	return &Service{
		scope: scope,
	}
}

var errNotImplemented = errors.New("Not implemented")

const etcdMountPath = "/var/lib/etcd"

func stringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s *Service) Reconcile(ctx context.Context) (_ *ctrl.Result, err error) {

	// If not token information has been given, the server cannot be successfully reconciled
	if s.scope.HcloudCluster.Spec.HrobotTokenRef == nil {
		return nil, errors.Errorf("ERROR: No token for Hetzner Robot provided: Cannot reconcile server %s", s.scope.BareMetalMachine.Name)
	}

	// Default port is 22
	port := 22
	if s.scope.BareMetalMachine.Spec.Port != nil {
		port = *s.scope.BareMetalMachine.Spec.Port
	}

	sshKeyName, _, privateSSHKey, err := s.retrieveSSHSecret(ctx)
	if err != nil {
		return nil, errors.Errorf("Unable to retrieve SSH secret: ", err)
	}

	sshFingerprint, err := s.getSSHFingerprintFromName(sshKeyName)
	if err != nil {
		return nil, errors.Errorf("Unable to get SSH fingerprint for the SSH key %s: %s ", sshKeyName, err)
	}

	// Get the Kubeadm config from the bootstrap provider
	userDataInitial, err := s.scope.GetRawBootstrapData(ctx)
	if err == scope.ErrBootstrapDataNotReady {
		s.scope.V(1).Info("Bootstrap data is not ready yet")
		return &reconcile.Result{RequeueAfter: 15 * time.Second}, nil
	} else if err != nil {
		return nil, err
	}
	userData, err := userdata.NewFromReader(bytes.NewReader(userDataInitial))
	if err != nil {
		return nil, err
	}

	kubeadmConfig, err := userData.GetKubeadmConfig()
	if err != nil {
		return nil, err
	}

	cloudProviderKey := "cloud-provider"
	cloudProviderValue := "external"

	if j := kubeadmConfig.JoinConfiguration; j != nil {
		if j.NodeRegistration.KubeletExtraArgs == nil {
			j.NodeRegistration.KubeletExtraArgs = make(map[string]string)
		}
		if _, ok := j.NodeRegistration.KubeletExtraArgs[cloudProviderKey]; !ok {
			j.NodeRegistration.KubeletExtraArgs[cloudProviderKey] = cloudProviderValue
		}
	}

	if err := userData.SetKubeadmConfig(kubeadmConfig); err != nil {
		return nil, err
	}

	kubeadmConfigString, err := userData.GetContentInformation()
	if err != nil {
		return nil, errors.Errorf("Error while obtaining the kubeadm config for the baremetal server", err)
	}

	// update list of servers which have the right type and are not taken in other clusters
	actualServers, err := s.actualStatus(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh server status")
	}

	var actualServer models.Server
	// isAttached is true if actualServer has the prefix clusterName in its name. This means it is
	// already attached to the cluster.
	var isAttached bool
	if len(actualServers) == 0 {
		// If no servers are in the list, we have to return an error
		return nil, errors.Errorf("No bare metal server found with type %s", *s.scope.BareMetalMachine.Spec.ServerType)
	} else if len(actualServers) == 1 {
		// If there is just one server, then it either not attached to a cluster or already attached to the cluster
		// If it is attached, then it has to have the right name. If it doesn't, we have to give an error, as this is
		// then not the right server and there are no free ones left to provision
		actualServer = actualServers[0]
		splitName := strings.Split(actualServer.ServerName, delimiter)

		if splitName[0] == s.scope.Cluster.Name {
			if len(splitName) == 3 && splitName[2] == s.scope.BareMetalMachine.Name {
				isAttached = true
			} else {
				return nil, errors.Errorf(
					"Found one bare metal server with type %s, but it is already attached with name %s",
					*s.scope.BareMetalMachine.Spec.ServerType, actualServer.ServerName,
				)
			}
		}

	} else if len(actualServers) > 1 {

		var check int
		var newServer *models.Server
		for _, server := range actualServers {
			splitName := strings.Split(server.ServerName, delimiter)

			// If server is attached to the cluster
			if len(splitName) == 3 {
				// If server is the one we are looking for
				if splitName[2] == s.scope.BareMetalMachine.Name {
					isAttached = true
					actualServer = server
					check++
				}
			} else {
				// If server is not attached, then take it
				newServer = &server
			}
		}

		if check > 1 {
			return nil, errors.Errorf("There are %s servers which are attached to the cluster with name %s", check, actualServer.ServerName)
		} else if check == 0 {
			// No attached server with the correct name found, so we have to take a new one
			if newServer == nil {
				return nil, errors.Errorf("There are no servers of type %s which are not yet attached to the cluster", *s.scope.BareMetalMachine.Spec.ServerType)
			}
			actualServer = *newServer
		}
	}

	// check if server has been cancelled
	if actualServer.Cancelled {
		s.scope.V(1).Info("server has been cancelled", "server", actualServer.ServerName, "cancelled",
			actualServer.Cancelled, "paid until", actualServer.PaidUntil)
		paidUntil, err := time.Parse("2006-01-02", actualServer.PaidUntil)
		if err != nil {
			return nil, errors.Errorf("ERROR: Failed to parse paidUntil date. Error: %s", err)
		}
		// If the server will be cancelled soon but is still in the list of actualServers, it means that
		// it is already attached since we don't add new servers that get deleted soon to actualServers
		// This means we have to set a failureReason so that the infrastructure object gets deleted
		if paidUntil.Before(time.Now().Add(hoursBeforeDeletion * time.Hour)) {
			er := capierrors.UpdateMachineError
			s.scope.BareMetalMachine.Status.FailureReason = &er
			s.scope.BareMetalMachine.Status.FailureMessage = pointer.StringPtr("Machine has been cancelled and is paid until less than" + hoursBeforeDeletion.String() + "hours")
			s.scope.BareMetalMachine.Status.Ready = false
		}
	}

	providerID := fmt.Sprintf("hetzner://%d", actualServer.ServerNumber)

	s.scope.BareMetalMachine.Spec.ProviderID = &providerID

	if isAttached == true {
		return nil, nil
	}

	// If the server is not already attached, then we have to configure it

	// wait for server being running
	if actualServer.Status != "ready" {
		s.scope.V(1).Info("server not in running state", "server", actualServer.ServerName, "status", actualServer.Status)
		return &reconcile.Result{RequeueAfter: 300 * time.Second}, nil
	}

	_, err = s.scope.HrobotClient().ActivateRescue(actualServer.ServerIP, sshFingerprint)
	if err != nil {
		return nil, errors.Errorf("Unable to activate rescue system: ", err)
	}

	_, err = s.scope.HrobotClient().ResetBMServer(actualServer.ServerIP, "hw")
	if err != nil {
		return nil, errors.Errorf("Unable to reset server: ", err)
	}

	stdout, stderr, err := runSSH("hostname", actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("SSH command hostname returned the error %s. The output of stderr is %s", err, stderr)
	}
	if !strings.Contains(stdout, "rescue") {
		return nil, errors.Errorf("Rescue system not successfully started. Output of command hostname is %s", stdout)
	}

	var blockDevices blockDevices
	blockDeviceCommand := "lsblk -o name,size,rota,fstype,label -e1 -e7 --json"
	stdout, stderr, err = runSSH(blockDeviceCommand, actualServer.ServerIP, port, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", blockDeviceCommand, err, stderr)
	}

	err = json.Unmarshal([]byte(stdout), &blockDevices)
	if err != nil {
		return nil, errors.Errorf("Error while unmarshaling: %s", err)
	}

	drive, err := findCorrectDevice(blockDevices)
	if err != nil {
		return nil, errors.Errorf("Error while finding correct device: %s", err)
	}

	partitionString := `PART /boot ext3 512M
PART / ext4 all`
	if s.scope.BareMetalMachine.Spec.Partition != nil {
		partitionString = *s.scope.BareMetalMachine.Spec.Partition
	}
	autoSetup := fmt.Sprintf(
		`cat > /autosetup << EOF
DRIVE1 /dev/%s
BOOTLOADER grub
HOSTNAME %s
%s
IMAGE %s
EOF`,
		drive, actualServer.ServerName, partitionString, *s.scope.BareMetalMachine.Spec.ImagePath)

	// Send autosetup file to server
	stdout, stderr, err = runSSH(autoSetup, actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("SSH command autosetup returned the error %s. The output of stderr is %s", err, stderr)
	}

	// install image
	stdout, stderr, err = runSSH("bash /root/.oldroot/nfs/install/installimage", actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("SSH command installimage returned the error %s. The output of stderr is %s", err, stderr)
	}
	// get again list of block devices and label children of our drive
	stdout, stderr, err = runSSH(blockDeviceCommand, actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", blockDeviceCommand, err, stderr)
	}

	err = json.Unmarshal([]byte(stdout), &blockDevices)
	if err != nil {
		return nil, errors.Errorf("Error while unmarshaling: %s", err)
	}

	command, err := labelChildrenCommand(blockDevices, drive)
	if err != nil {
		return nil, errors.Errorf("Error while constructing labeling children command of device %s: %s", drive, err)
	}

	// label children of our drive
	stdout, stderr, err = runSSH(command, actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", command, err, stderr)
	}

	// reboot system
	stdout, stderr, err = runSSH("reboot", actualServer.ServerIP, 22, privateSSHKey)
	if err != nil {
		if !strings.Contains(err.Error(), "remote command exited without exit status or exit signal") {
			return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", "reboot", err, stderr)
		}
	}

	// We cannot create the file in the tmp folder right after rebooting, as it gets deleted again
	// so we have to wait for a bit
	_, _, _ = runSSH("sleep 60", actualServer.ServerIP, port, privateSSHKey)

	kubeadmCommand := fmt.Sprintf(
		`cat > /tmp/kubeadm-join-config.yaml << EOF
%s
EOF`, kubeadmConfigString)

	// send kubeadmConfig to server
	stdout, stderr, err = runSSH(kubeadmCommand, actualServer.ServerIP, port, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", kubeadmCommand, err, stderr)
	}

	command = "chmod 640 /tmp/kubeadm-join-config.yaml && chown root:root /tmp/kubeadm-join-config.yaml"

	stdout, stderr, err = runSSH(command, actualServer.ServerIP, port, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", command, err, stderr)
	}

	command = "kubeadm join --config /tmp/kubeadm-join-config.yaml"

	stdout, stderr, err = runSSH(command, actualServer.ServerIP, port, privateSSHKey)
	if err != nil {
		return nil, errors.Errorf("Error running the ssh command %s: Error: %s, stderr: %s", command, err, stderr)
	}

	_, err = s.scope.HrobotClient().SetBMServerName(actualServer.ServerIP,
		s.scope.Cluster.Name+delimiter+*s.scope.BareMetalMachine.Spec.ServerType+delimiter+s.scope.BareMetalMachine.Name)
	if err != nil {
		return nil, errors.Errorf("Unable to change bare metal server name: ", err)
	}
	s.scope.V(1).Info("Finished provisioning bare metal server %s", s.scope.BareMetalMachine.Name)
	s.scope.BareMetalMachine.Status.Ready = true
	return nil, nil
}

func (s *Service) Delete(ctx context.Context) (_ *ctrl.Result, err error) {
	// update current servers
	actualServers, err := s.actualStatus(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh server status")
	}

	var attachedServers []models.Server

	for _, serverData := range actualServers {

		splitName := strings.Split(serverData.ServerName, delimiter)
		if splitName[0] == s.scope.Cluster.Name {
			if len(splitName) == 3 && splitName[2] == s.scope.BareMetalMachine.Name {
				attachedServers = append(attachedServers, serverData)
			}
		}
	}

	for _, server := range attachedServers {
		_, err = s.scope.HrobotClient().SetBMServerName(server.ServerIP,
			*s.scope.BareMetalMachine.Spec.ServerType+delimiter+"unused-"+s.scope.BareMetalMachine.Name)
		if err != nil {
			return nil, errors.Errorf("Unable to change bare metal server name: ", err)
		}

	}
	return nil, nil
}

func (s *Service) getSSHFingerprintFromName(name string) (fingerprint string, err error) {

	sshKeys, err := s.scope.HrobotClient().ListBMKeys()
	if len(sshKeys) == 0 {
		return "", errors.New("No SSH Keys given")
	}

	var found bool
	for _, key := range sshKeys {
		if name == key.Name {
			fingerprint = key.Fingerprint
			found = true
		}
	}

	if found == false {
		return "", errors.Errorf("No SSH key with name %s found", name)
	}

	return fingerprint, nil
}

// actualStatus gathers all matching server instances, matched by tag
func (s *Service) actualStatus(ctx context.Context) ([]models.Server, error) {

	serverList, err := s.scope.HrobotClient().ListBMServers()
	if err != nil {
		return nil, errors.Errorf("unable to list bare metal servers: %s", err)
	}

	var actualServers []models.Server

	for _, server := range serverList {
		splitName := strings.Split(server.ServerName, delimiter)
		if len(splitName) == 2 && splitName[0] == *s.scope.BareMetalMachine.Spec.ServerType {
			if !server.Cancelled {
				actualServers = append(actualServers, server)
			} else {
				paidUntil, err := time.Parse("2006-01-02", server.PaidUntil)
				if err != nil {
					return nil, errors.Errorf("ERROR: Failed to parse paidUntil date. Error: %s", err)
				}
				if !paidUntil.Before(time.Now().Add(hoursBeforeDeletion * time.Hour)) {
					actualServers = append(actualServers, server)
				}
			}
		}
		if len(splitName) == 3 && splitName[0] == s.scope.Cluster.Name && splitName[1] == *s.scope.BareMetalMachine.Spec.ServerType {
			actualServers = append(actualServers, server)
		}
	}

	return actualServers, nil
}

func labelChildrenCommand(blockDevices blockDevices, drive string) (string, error) {

	var device blockDevice
	var check bool
	for _, d := range blockDevices.Devices {
		if drive == d.Name {
			device = d
			check = true
		}
	}

	if check == false {
		return "", errors.Errorf("no device with name %s found", drive)
	}

	if device.Children == nil {
		return "", errors.Errorf("no children for device with name %s found, instalimage did not work properly", drive)
	}

	var command string
	for _, child := range device.Children {
		str := fmt.Sprintf(`e2label /dev/%s os
`, child.Name)
		command = command + str
	}
	command = command[:len(command)-1]
	return command, nil
}

func findCorrectDevice(blockDevices blockDevices) (drive string, err error) {
	// If no blockdevice has correctly labeled children, we follow a certain set of rules to find the right one
	var numLabels int
	var hasChildren int
	for _, device := range blockDevices.Devices {
		if device.Children != nil && strings.HasPrefix(device.Name, "sd") {
			hasChildren++
			for _, child := range device.Children {
				if child.Label == "os" {
					numLabels++
					drive = device.Name
					break
				}
			}
		}
	}

	// if numLabels == 1 then finished (drive has been set already)
	// if numLabels == 0 then start sorting
	// if numLabels > 1 then throw error
	var filteredDevices []blockDevice

	if numLabels > 1 {
		return "", errors.Errorf("Found %v devices with the correct labels", numLabels)
	} else if numLabels == 0 {
		// If every device has children, then there is none left for us
		if hasChildren == len(blockDevices.Devices) {
			return "", errors.New("No device is left for installing the operating system")
		}

		// Choose devices with no children, whose name starts with "sd" and which are SSDs (i.e. rota=false)
		for _, device := range blockDevices.Devices {
			if device.Children == nil && strings.HasPrefix(device.Name, "sd") && device.Rotation == false {
				filteredDevices = append(filteredDevices, device)
			}
		}

		// This means that there is no SSD available. Then we have to include HDD as well
		if len(filteredDevices) == 0 {
			for _, device := range blockDevices.Devices {
				if device.Children == nil && strings.HasPrefix(device.Name, "sd") {
					filteredDevices = append(filteredDevices, device)
				}
			}
			// If there is only one device which satisfies the requirements then we choose it
		} else if len(filteredDevices) == 1 {
			drive = filteredDevices[0].Name
			// If there are more devices then we need to sort them according to our specifications
		} else {
			// First change the data type of size, so that we can compare it
			type reducedBlockDevice struct {
				Name string
				Size int
			}
			var reducedDevices []reducedBlockDevice
			for _, device := range filteredDevices {
				size, err := convertSizeToInt(device.Size)
				if err != nil {
					return "", errors.Errorf("Could not convert size %s to integer", device.Size)
				}
				reducedDevices = append(reducedDevices, reducedBlockDevice{
					Name: device.Name,
					Size: size,
				})
			}
			// Sort the devices with respect to size
			sort.SliceStable(reducedDevices, func(i, j int) bool {
				return reducedDevices[i].Size < reducedDevices[j].Size
			})

			// Look whether there is more than one device with the same size
			var filteredReducedDevices []reducedBlockDevice
			if reducedDevices[0].Size < reducedDevices[1].Size {
				drive = reducedDevices[0].Name
			} else {
				for _, device := range reducedDevices {
					if device.Size == reducedDevices[0].Size {
						filteredReducedDevices = append(filteredReducedDevices, device)
					}
				}
				// Sort the devices with respect to name
				sort.SliceStable(filteredReducedDevices, func(i, j int) bool {
					return filteredReducedDevices[i].Name > filteredReducedDevices[j].Name
				})
				drive = filteredReducedDevices[0].Name
			}
		}
	}
	return drive, nil
}

// converts the size of Hetzner drives, e.g. 3,5T to an int, here 3,500,000 (MB)
func convertSizeToInt(str string) (x int, err error) {
	s := str
	var m float64
	var z float64
	strings.ReplaceAll(s, ",", ".")
	if strings.HasSuffix(s, "T") {
		m = 1000000
	} else if strings.HasSuffix(s, "G") {
		m = 1000
	} else if strings.HasSuffix(s, "M") {
		m = 1
	} else {
		return 0, errors.Errorf("Unknown unit in size %s", s)
	}

	s = s[:len(s)-1]

	z, err = strconv.ParseFloat(s, 32)

	if err != nil {
		return 0, errors.Errorf("Error while converting string %s to integer: %s", s, err)
	}
	x = int(z * m)
	return x, nil
}

func (s *Service) retrieveSSHSecret(ctx context.Context) (sshKeyName string, publicKey string, privateKey string, err error) {
	// retrieve token secret
	var tokenSecret corev1.Secret
	tokenSecretName := types.NamespacedName{Namespace: s.scope.HcloudCluster.Namespace, Name: s.scope.BareMetalMachine.Spec.SSHTokenRef.TokenName}
	if err := s.scope.Client.Get(ctx, tokenSecretName, &tokenSecret); err != nil {
		return "", "", "", errors.Errorf("error getting referenced token secret/%s: %s", tokenSecretName, err)
	}

	publicKeyTokenBytes, keyExists := tokenSecret.Data[s.scope.BareMetalMachine.Spec.SSHTokenRef.PublicKey]
	if !keyExists {
		return "", "", "", errors.Errorf("error key %s does not exist in secret/%s", s.scope.BareMetalMachine.Spec.SSHTokenRef.PublicKey, tokenSecretName)
	}
	privateKeyTokenBytes, keyExists := tokenSecret.Data[s.scope.BareMetalMachine.Spec.SSHTokenRef.PrivateKey]
	if !keyExists {
		return "", "", "", errors.Errorf("error key %s does not exist in secret/%s", s.scope.BareMetalMachine.Spec.SSHTokenRef.PrivateKey, tokenSecretName)
	}
	sshKeyNameTokenBytes, keyExists := tokenSecret.Data[s.scope.BareMetalMachine.Spec.SSHTokenRef.SSHKeyName]
	if !keyExists {
		return "", "", "", errors.Errorf("error key %s does not exist in secret/%s", s.scope.BareMetalMachine.Spec.SSHTokenRef.SSHKeyName, tokenSecretName)
	}

	sshKeyName = string(sshKeyNameTokenBytes)
	privateKey = string(privateKeyTokenBytes)
	publicKey = string(publicKeyTokenBytes)
	return sshKeyName, publicKey, privateKey, nil
}

func runSSH(command, ip string, port int, privateSSHKey string) (stdout string, stderr string, err error) {

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey([]byte(privateSSHKey))
	if err != nil {
		return "", "", errors.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ssh.FixedHostKey(hostKey),
	}

	// Connect to the remote server and perform the SSH handshake.
	var client *ssh.Client
	var check bool
	for i := 0; i < (maxWaitingTimeForSSHConnection / 15); i++ {
		client, err = ssh.Dial("tcp", ip+":"+strconv.Itoa(port), config)
		if err != nil {
			// If the SSH connection could not be established, then retry 15 sec later
			time.Sleep(15 * time.Second)
			continue
		}
		check = true
		break
	}

	if check == false {
		return "", "", errors.Errorf("Unable to establish connection to remote server: %s", err)
	}

	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer

	sess.Stdout = &stdoutBuffer
	sess.Stderr = &stderrBuffer

	err = sess.Run(command)

	stdout = stdoutBuffer.String()
	stderr = stderrBuffer.String()
	return stdout, stderr, err
}
