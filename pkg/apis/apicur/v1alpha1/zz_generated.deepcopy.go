// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistry) DeepCopyInto(out *ApicurioRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistry.
func (in *ApicurioRegistry) DeepCopy() *ApicurioRegistry {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApicurioRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryList) DeepCopyInto(out *ApicurioRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApicurioRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryList.
func (in *ApicurioRegistryList) DeepCopy() *ApicurioRegistryList {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApicurioRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpec) DeepCopyInto(out *ApicurioRegistrySpec) {
	*out = *in
	out.Image = in.Image
	out.Configuration = in.Configuration
	out.Deployment = in.Deployment
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpec.
func (in *ApicurioRegistrySpec) DeepCopy() *ApicurioRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfiguration) DeepCopyInto(out *ApicurioRegistrySpecConfiguration) {
	*out = *in
	out.DataSource = in.DataSource
	out.Kafka = in.Kafka
	out.Streams = in.Streams
	out.Infinispan = in.Infinispan
	out.UI = in.UI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfiguration.
func (in *ApicurioRegistrySpecConfiguration) DeepCopy() *ApicurioRegistrySpecConfiguration {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationDataSource) DeepCopyInto(out *ApicurioRegistrySpecConfigurationDataSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationDataSource.
func (in *ApicurioRegistrySpecConfigurationDataSource) DeepCopy() *ApicurioRegistrySpecConfigurationDataSource {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationInfinispan) DeepCopyInto(out *ApicurioRegistrySpecConfigurationInfinispan) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationInfinispan.
func (in *ApicurioRegistrySpecConfigurationInfinispan) DeepCopy() *ApicurioRegistrySpecConfigurationInfinispan {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationInfinispan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationKafka) DeepCopyInto(out *ApicurioRegistrySpecConfigurationKafka) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationKafka.
func (in *ApicurioRegistrySpecConfigurationKafka) DeepCopy() *ApicurioRegistrySpecConfigurationKafka {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationStreams) DeepCopyInto(out *ApicurioRegistrySpecConfigurationStreams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationStreams.
func (in *ApicurioRegistrySpecConfigurationStreams) DeepCopy() *ApicurioRegistrySpecConfigurationStreams {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationStreams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecConfigurationUI) DeepCopyInto(out *ApicurioRegistrySpecConfigurationUI) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecConfigurationUI.
func (in *ApicurioRegistrySpecConfigurationUI) DeepCopy() *ApicurioRegistrySpecConfigurationUI {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecConfigurationUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecDeployment) DeepCopyInto(out *ApicurioRegistrySpecDeployment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecDeployment.
func (in *ApicurioRegistrySpecDeployment) DeepCopy() *ApicurioRegistrySpecDeployment {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecDeploymentResources) DeepCopyInto(out *ApicurioRegistrySpecDeploymentResources) {
	*out = *in
	out.Cpu = in.Cpu
	out.Memory = in.Memory
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecDeploymentResources.
func (in *ApicurioRegistrySpecDeploymentResources) DeepCopy() *ApicurioRegistrySpecDeploymentResources {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecDeploymentResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecDeploymentResourcesRequestsLimit) DeepCopyInto(out *ApicurioRegistrySpecDeploymentResourcesRequestsLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecDeploymentResourcesRequestsLimit.
func (in *ApicurioRegistrySpecDeploymentResourcesRequestsLimit) DeepCopy() *ApicurioRegistrySpecDeploymentResourcesRequestsLimit {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecDeploymentResourcesRequestsLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistrySpecImage) DeepCopyInto(out *ApicurioRegistrySpecImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistrySpecImage.
func (in *ApicurioRegistrySpecImage) DeepCopy() *ApicurioRegistrySpecImage {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistrySpecImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicurioRegistryStatus) DeepCopyInto(out *ApicurioRegistryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicurioRegistryStatus.
func (in *ApicurioRegistryStatus) DeepCopy() *ApicurioRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(ApicurioRegistryStatus)
	in.DeepCopyInto(out)
	return out
}
