// +build !ignore_autogenerated

// Copyright (c) VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZone) DeepCopyInto(out *AvailabilityZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZone.
func (in *AvailabilityZone) DeepCopy() *AvailabilityZone {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvailabilityZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZoneList) DeepCopyInto(out *AvailabilityZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AvailabilityZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZoneList.
func (in *AvailabilityZoneList) DeepCopy() *AvailabilityZoneList {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvailabilityZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZoneSpec) DeepCopyInto(out *AvailabilityZoneSpec) {
	*out = *in
	if in.ClusterComputeResourceMoIDs != nil {
		in, out := &in.ClusterComputeResourceMoIDs, &out.ClusterComputeResourceMoIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(map[string]NamespaceInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZoneSpec.
func (in *AvailabilityZoneSpec) DeepCopy() *AvailabilityZoneSpec {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZoneStatus) DeepCopyInto(out *AvailabilityZoneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZoneStatus.
func (in *AvailabilityZoneStatus) DeepCopy() *AvailabilityZoneStatus {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceInfo) DeepCopyInto(out *NamespaceInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceInfo.
func (in *NamespaceInfo) DeepCopy() *NamespaceInfo {
	if in == nil {
		return nil
	}
	out := new(NamespaceInfo)
	in.DeepCopyInto(out)
	return out
}
