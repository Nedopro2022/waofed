//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedDeploymentSelector) DeepCopyInto(out *FederatedDeploymentSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedDeploymentSelector.
func (in *FederatedDeploymentSelector) DeepCopy() *FederatedDeploymentSelector {
	if in == nil {
		return nil
	}
	out := new(FederatedDeploymentSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancingSettings) DeepCopyInto(out *LoadBalancingSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancingSettings.
func (in *LoadBalancingSettings) DeepCopy() *LoadBalancingSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancingSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RSPOptimizerSettings) DeepCopyInto(out *RSPOptimizerSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RSPOptimizerSettings.
func (in *RSPOptimizerSettings) DeepCopy() *RSPOptimizerSettings {
	if in == nil {
		return nil
	}
	out := new(RSPOptimizerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSettings) DeepCopyInto(out *SchedulingSettings) {
	*out = *in
	out.Selector = in.Selector
	out.Optimizer = in.Optimizer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSettings.
func (in *SchedulingSettings) DeepCopy() *SchedulingSettings {
	if in == nil {
		return nil
	}
	out := new(SchedulingSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WAOFedConfig) DeepCopyInto(out *WAOFedConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WAOFedConfig.
func (in *WAOFedConfig) DeepCopy() *WAOFedConfig {
	if in == nil {
		return nil
	}
	out := new(WAOFedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WAOFedConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WAOFedConfigList) DeepCopyInto(out *WAOFedConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WAOFedConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WAOFedConfigList.
func (in *WAOFedConfigList) DeepCopy() *WAOFedConfigList {
	if in == nil {
		return nil
	}
	out := new(WAOFedConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WAOFedConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WAOFedConfigSpec) DeepCopyInto(out *WAOFedConfigSpec) {
	*out = *in
	if in.Scheduling != nil {
		in, out := &in.Scheduling, &out.Scheduling
		*out = new(SchedulingSettings)
		**out = **in
	}
	if in.LoadBalancing != nil {
		in, out := &in.LoadBalancing, &out.LoadBalancing
		*out = new(LoadBalancingSettings)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WAOFedConfigSpec.
func (in *WAOFedConfigSpec) DeepCopy() *WAOFedConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WAOFedConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WAOFedConfigStatus) DeepCopyInto(out *WAOFedConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WAOFedConfigStatus.
func (in *WAOFedConfigStatus) DeepCopy() *WAOFedConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WAOFedConfigStatus)
	in.DeepCopyInto(out)
	return out
}
