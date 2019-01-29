// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfig) DeepCopyInto(out *KubeSchedulerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfig.
func (in *KubeSchedulerConfig) DeepCopy() *KubeSchedulerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerOperatorConfig) DeepCopyInto(out *KubeSchedulerOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerOperatorConfig.
func (in *KubeSchedulerOperatorConfig) DeepCopy() *KubeSchedulerOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerOperatorConfigList) DeepCopyInto(out *KubeSchedulerOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeSchedulerOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerOperatorConfigList.
func (in *KubeSchedulerOperatorConfigList) DeepCopy() *KubeSchedulerOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerOperatorConfigSpec) DeepCopyInto(out *KubeSchedulerOperatorConfigSpec) {
	*out = *in
	in.StaticPodOperatorSpec.DeepCopyInto(&out.StaticPodOperatorSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerOperatorConfigSpec.
func (in *KubeSchedulerOperatorConfigSpec) DeepCopy() *KubeSchedulerOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerOperatorConfigStatus) DeepCopyInto(out *KubeSchedulerOperatorConfigStatus) {
	*out = *in
	in.StaticPodOperatorStatus.DeepCopyInto(&out.StaticPodOperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerOperatorConfigStatus.
func (in *KubeSchedulerOperatorConfigStatus) DeepCopy() *KubeSchedulerOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}
