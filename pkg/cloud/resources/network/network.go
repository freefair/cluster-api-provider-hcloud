package network

import (
	"context"
	"net"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/pkg/errors"

	infrav1 "github.com/freefair/cluster-api-provider-hcloud/api/v1beta1"
	"github.com/freefair/cluster-api-provider-hcloud/pkg/cloud/utils"
	"github.com/freefair/cluster-api-provider-hcloud/pkg/scope"
)

type Service struct {
	scope *scope.ClusterScope
}

func NewService(scope *scope.ClusterScope) *Service {
	return &Service{
		scope: scope,
	}
}

var errNotImplemented = errors.New("Not implemented")

func apiToStatus(network *hcloud.Network) *infrav1.HcloudNetworkStatus {
	var subnets = make([]infrav1.HcloudNetworkSubnetSpec, len(network.Subnets))
	for pos, n := range network.Subnets {
		subnets[pos].NetworkZone = infrav1.HcloudNetworkZone(n.NetworkZone)
		subnets[pos].CIDRBlock = n.IPRange.String()
	}

	var status infrav1.HcloudNetworkStatus
	status.ID = network.ID
	status.CIDRBlock = network.IPRange.String()
	status.Subnets = subnets
	status.Labels = network.Labels
	return &status
}

func (s *Service) defaults() *infrav1.HcloudNetworkSpec {
	n := infrav1.HcloudNetworkSpec{}
	n.CIDRBlock = "10.0.0.0/16"
	n.Subnets = []infrav1.HcloudNetworkSubnetSpec{
		{
			NetworkZone: s.scope.HcloudCluster.Status.NetworkZone,
			HcloudNetwork: infrav1.HcloudNetwork{
				CIDRBlock: "10.0.0.0/24",
			},
		},
	}
	return &n
}

func (s *Service) labels() map[string]string {
	clusterTagKey := infrav1.ClusterTagKey(s.scope.HcloudCluster.Name)
	return map[string]string{
		clusterTagKey: string(infrav1.ResourceLifecycleOwned),
	}
}

func (s *Service) Reconcile(ctx context.Context) (err error) {
	// set defaults if nothing is set
	if s.scope.HcloudCluster.Spec.Network == nil {
		s.scope.HcloudCluster.Spec.Network = s.defaults()
	}

	// No network requested, end here
	if s.scope.HcloudCluster.Spec.Network.IsZero() {
		return nil
	}

	// update current status
	networkStatus, err := s.actualStatus(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to refresh networks")
	}
	s.scope.HcloudCluster.Status.Network = networkStatus

	if networkStatus != nil {
		// TODO: Check if the actual values are matching
		return nil
	}

	if s.scope.HcloudCluster.Spec.Network == nil {
		s.scope.HcloudCluster.Spec.Network = s.defaults()
	}

	networkStatus, err = s.createNetwork(ctx, s.scope.HcloudCluster.Spec.Network)
	if err != nil {
		return errors.Wrap(err, "failed to create network")
	}
	s.scope.HcloudCluster.Status.Network = networkStatus

	return nil
}

func (s *Service) createNetwork(ctx context.Context, spec *infrav1.HcloudNetworkSpec) (*infrav1.HcloudNetworkStatus, error) {
	hc := s.scope.HcloudCluster

	s.scope.V(2).Info("Create a new network", "cidrBlock", spec.CIDRBlock, "subnets", spec.Subnets)
	_, network, err := net.ParseCIDR(spec.CIDRBlock)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid network '%s'", spec.CIDRBlock)
	}

	var subnets = make([]hcloud.NetworkSubnet, len(spec.Subnets))
	for pos, sn := range spec.Subnets {
		_, subnet, err := net.ParseCIDR(sn.CIDRBlock)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid network '%s'", sn.CIDRBlock)
		}
		subnets[pos].IPRange = subnet
		subnets[pos].NetworkZone = hcloud.NetworkZone(sn.NetworkZone)
		subnets[pos].Type = hcloud.NetworkSubnetTypeServer
	}

	opts := hcloud.NetworkCreateOpts{
		Name:    hc.Name,
		IPRange: network,
		Labels:  s.labels(),
		Subnets: subnets,
	}

	s.scope.V(1).Info("Create a new network", "opts", opts)

	respNetworkCreate, _, err := s.scope.HcloudClient().CreateNetwork(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "error creating network")
	}

	return apiToStatus(respNetworkCreate), nil
}

func (s *Service) deleteNetwork(ctx context.Context, status *infrav1.HcloudNetworkStatus) error {
	// ensure deleted network is actually owned by us
	clusterTagKey := infrav1.ClusterTagKey(s.scope.HcloudCluster.Name)
	if status.Labels == nil || infrav1.ResourceLifecycle(status.Labels[clusterTagKey]) != infrav1.ResourceLifecycleOwned {
		s.scope.V(3).Info("Ignore request to delete network, as it is not owned", "id", status.ID, "cidrBlock", status.CIDRBlock)
		return nil
	}
	_, err := s.scope.HcloudClient().DeleteNetwork(ctx, &hcloud.Network{ID: status.ID})
	s.scope.V(2).Info("Delete network", "id", status.ID, "cidrBlock", status.CIDRBlock)
	return err
}

func (s *Service) Delete(ctx context.Context) (err error) {
	// update current status
	networkStatus, err := s.actualStatus(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to refresh networks")
	}

	if networkStatus == nil {
		return nil
	}

	if err := s.deleteNetwork(ctx, networkStatus); err != nil {
		return errors.Wrap(err, "failed to delete network")
	}

	return nil
}

func (s *Service) actualStatus(ctx context.Context) (*infrav1.HcloudNetworkStatus, error) {
	opts := hcloud.NetworkListOpts{}
	opts.LabelSelector = utils.LabelsToLabelSelector(s.labels())
	networks, err := s.scope.HcloudClient().ListNetworks(s.scope.Ctx, opts)
	if err != nil {
		return nil, err
	}

	for _, n := range networks {
		return apiToStatus(n), nil
	}

	return nil, nil
}
