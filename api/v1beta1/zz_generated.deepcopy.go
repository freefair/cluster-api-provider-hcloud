//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachine) DeepCopyInto(out *BareMetalMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachine.
func (in *BareMetalMachine) DeepCopy() *BareMetalMachine {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineList) DeepCopyInto(out *BareMetalMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineList.
func (in *BareMetalMachineList) DeepCopy() *BareMetalMachineList {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineSpec) DeepCopyInto(out *BareMetalMachineSpec) {
	*out = *in
	out.SSHTokenRef = in.SSHTokenRef
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(string)
		**out = **in
	}
	if in.ImagePath != nil {
		in, out := &in.ImagePath, &out.ImagePath
		*out = new(string)
		**out = **in
	}
	if in.ServerType != nil {
		in, out := &in.ServerType, &out.ServerType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineSpec.
func (in *BareMetalMachineSpec) DeepCopy() *BareMetalMachineSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineStatus) DeepCopyInto(out *BareMetalMachineStatus) {
	*out = *in
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineStatus.
func (in *BareMetalMachineStatus) DeepCopy() *BareMetalMachineStatus {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineTemplate) DeepCopyInto(out *BareMetalMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineTemplate.
func (in *BareMetalMachineTemplate) DeepCopy() *BareMetalMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineTemplateList) DeepCopyInto(out *BareMetalMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineTemplateList.
func (in *BareMetalMachineTemplateList) DeepCopy() *BareMetalMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineTemplateResource) DeepCopyInto(out *BareMetalMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineTemplateResource.
func (in *BareMetalMachineTemplateResource) DeepCopy() *BareMetalMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineTemplateSpec) DeepCopyInto(out *BareMetalMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineTemplateSpec.
func (in *BareMetalMachineTemplateSpec) DeepCopy() *BareMetalMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudCluster) DeepCopyInto(out *HcloudCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudCluster.
func (in *HcloudCluster) DeepCopy() *HcloudCluster {
	if in == nil {
		return nil
	}
	out := new(HcloudCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudClusterList) DeepCopyInto(out *HcloudClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HcloudCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudClusterList.
func (in *HcloudClusterList) DeepCopy() *HcloudClusterList {
	if in == nil {
		return nil
	}
	out := new(HcloudClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudClusterSpec) DeepCopyInto(out *HcloudClusterSpec) {
	*out = *in
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]HcloudLocation, len(*in))
		copy(*out, *in)
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]HcloudSSHKeySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ControlPlaneLoadBalancer.DeepCopyInto(&out.ControlPlaneLoadBalancer)
	if in.ControlPlaneEndpoint != nil {
		in, out := &in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint
		*out = new(apiv1beta1.APIEndpoint)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(HcloudNetworkSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.HcloudTokenRef != nil {
		in, out := &in.HcloudTokenRef, &out.HcloudTokenRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.HrobotTokenRef != nil {
		in, out := &in.HrobotTokenRef, &out.HrobotTokenRef
		*out = new(hrobotTokenRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudClusterSpec.
func (in *HcloudClusterSpec) DeepCopy() *HcloudClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudClusterStatus) DeepCopyInto(out *HcloudClusterStatus) {
	*out = *in
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]HcloudLocation, len(*in))
		copy(*out, *in)
	}
	in.ControlPlaneLoadBalancer.DeepCopyInto(&out.ControlPlaneLoadBalancer)
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(HcloudNetworkStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = new(HcloudClusterStatusManifests)
		(*in).DeepCopyInto(*out)
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudClusterStatus.
func (in *HcloudClusterStatus) DeepCopy() *HcloudClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HcloudClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudClusterStatusManifests) DeepCopyInto(out *HcloudClusterStatusManifests) {
	*out = *in
	if in.Initialized != nil {
		in, out := &in.Initialized, &out.Initialized
		*out = new(bool)
		**out = **in
	}
	if in.AppliedHash != nil {
		in, out := &in.AppliedHash, &out.AppliedHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudClusterStatusManifests.
func (in *HcloudClusterStatusManifests) DeepCopy() *HcloudClusterStatusManifests {
	if in == nil {
		return nil
	}
	out := new(HcloudClusterStatusManifests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudLoadBalancerSpec) DeepCopyInto(out *HcloudLoadBalancerSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]LoadBalancerServiceSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudLoadBalancerSpec.
func (in *HcloudLoadBalancerSpec) DeepCopy() *HcloudLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudLoadBalancerStatus) DeepCopyInto(out *HcloudLoadBalancerStatus) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudLoadBalancerStatus.
func (in *HcloudLoadBalancerStatus) DeepCopy() *HcloudLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(HcloudLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachine) DeepCopyInto(out *HcloudMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachine.
func (in *HcloudMachine) DeepCopy() *HcloudMachine {
	if in == nil {
		return nil
	}
	out := new(HcloudMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineList) DeepCopyInto(out *HcloudMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HcloudMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineList.
func (in *HcloudMachineList) DeepCopy() *HcloudMachineList {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineSpec) DeepCopyInto(out *HcloudMachineSpec) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]HcloudSSHKeySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]HcloudMachineVolume, len(*in))
		copy(*out, *in)
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineSpec.
func (in *HcloudMachineSpec) DeepCopy() *HcloudMachineSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineStatus) DeepCopyInto(out *HcloudMachineStatus) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(HcloudImageID)
		**out = **in
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineStatus.
func (in *HcloudMachineStatus) DeepCopy() *HcloudMachineStatus {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineTemplate) DeepCopyInto(out *HcloudMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineTemplate.
func (in *HcloudMachineTemplate) DeepCopy() *HcloudMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineTemplateList) DeepCopyInto(out *HcloudMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HcloudMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineTemplateList.
func (in *HcloudMachineTemplateList) DeepCopy() *HcloudMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineTemplateResource) DeepCopyInto(out *HcloudMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineTemplateResource.
func (in *HcloudMachineTemplateResource) DeepCopy() *HcloudMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineTemplateSpec) DeepCopyInto(out *HcloudMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineTemplateSpec.
func (in *HcloudMachineTemplateSpec) DeepCopy() *HcloudMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudMachineVolume) DeepCopyInto(out *HcloudMachineVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudMachineVolume.
func (in *HcloudMachineVolume) DeepCopy() *HcloudMachineVolume {
	if in == nil {
		return nil
	}
	out := new(HcloudMachineVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudNetwork) DeepCopyInto(out *HcloudNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudNetwork.
func (in *HcloudNetwork) DeepCopy() *HcloudNetwork {
	if in == nil {
		return nil
	}
	out := new(HcloudNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudNetworkSpec) DeepCopyInto(out *HcloudNetworkSpec) {
	*out = *in
	out.HcloudNetwork = in.HcloudNetwork
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]HcloudNetworkSubnetSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudNetworkSpec.
func (in *HcloudNetworkSpec) DeepCopy() *HcloudNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudNetworkStatus) DeepCopyInto(out *HcloudNetworkStatus) {
	*out = *in
	in.HcloudNetworkSpec.DeepCopyInto(&out.HcloudNetworkSpec)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudNetworkStatus.
func (in *HcloudNetworkStatus) DeepCopy() *HcloudNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(HcloudNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudNetworkSubnetSpec) DeepCopyInto(out *HcloudNetworkSubnetSpec) {
	*out = *in
	out.HcloudNetwork = in.HcloudNetwork
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudNetworkSubnetSpec.
func (in *HcloudNetworkSubnetSpec) DeepCopy() *HcloudNetworkSubnetSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudNetworkSubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudSSHKeySpec) DeepCopyInto(out *HcloudSSHKeySpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudSSHKeySpec.
func (in *HcloudSSHKeySpec) DeepCopy() *HcloudSSHKeySpec {
	if in == nil {
		return nil
	}
	out := new(HcloudSSHKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudVolume) DeepCopyInto(out *HcloudVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudVolume.
func (in *HcloudVolume) DeepCopy() *HcloudVolume {
	if in == nil {
		return nil
	}
	out := new(HcloudVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudVolumeList) DeepCopyInto(out *HcloudVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HcloudVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudVolumeList.
func (in *HcloudVolumeList) DeepCopy() *HcloudVolumeList {
	if in == nil {
		return nil
	}
	out := new(HcloudVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HcloudVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudVolumeSpec) DeepCopyInto(out *HcloudVolumeSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudVolumeSpec.
func (in *HcloudVolumeSpec) DeepCopy() *HcloudVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(HcloudVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HcloudVolumeStatus) DeepCopyInto(out *HcloudVolumeStatus) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(HcloudVolumeID)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HcloudVolumeStatus.
func (in *HcloudVolumeStatus) DeepCopy() *HcloudVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(HcloudVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerServiceSpec) DeepCopyInto(out *LoadBalancerServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerServiceSpec.
func (in *LoadBalancerServiceSpec) DeepCopy() *LoadBalancerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerServiceSpec)
	in.DeepCopyInto(out)
	return out
}
