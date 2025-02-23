//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedMetricSpecificationObservation) DeepCopyInto(out *CustomizedMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedMetricSpecificationObservation.
func (in *CustomizedMetricSpecificationObservation) DeepCopy() *CustomizedMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizedMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedMetricSpecificationParameters) DeepCopyInto(out *CustomizedMetricSpecificationParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedMetricSpecificationParameters.
func (in *CustomizedMetricSpecificationParameters) DeepCopy() *CustomizedMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizedMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsObservation) DeepCopyInto(out *DimensionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsObservation.
func (in *DimensionsObservation) DeepCopy() *DimensionsObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsParameters) DeepCopyInto(out *DimensionsParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsParameters.
func (in *DimensionsParameters) DeepCopy() *DimensionsParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObservation) DeepCopyInto(out *PolicyObservation) {
	*out = *in
	if in.AlarmArns != nil {
		in, out := &in.AlarmArns, &out.AlarmArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObservation.
func (in *PolicyObservation) DeepCopy() *PolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyParameters) DeepCopyInto(out *PolicyParameters) {
	*out = *in
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDRef != nil {
		in, out := &in.ResourceIDRef, &out.ResourceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceIDSelector != nil {
		in, out := &in.ResourceIDSelector, &out.ResourceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimensionRef != nil {
		in, out := &in.ScalableDimensionRef, &out.ScalableDimensionRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableDimensionSelector != nil {
		in, out := &in.ScalableDimensionSelector, &out.ScalableDimensionSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespaceRef != nil {
		in, out := &in.ServiceNamespaceRef, &out.ServiceNamespaceRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceNamespaceSelector != nil {
		in, out := &in.ServiceNamespaceSelector, &out.ServiceNamespaceSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StepScalingPolicyConfiguration != nil {
		in, out := &in.StepScalingPolicyConfiguration, &out.StepScalingPolicyConfiguration
		*out = make([]StepScalingPolicyConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetTrackingScalingPolicyConfiguration != nil {
		in, out := &in.TargetTrackingScalingPolicyConfiguration, &out.TargetTrackingScalingPolicyConfiguration
		*out = make([]TargetTrackingScalingPolicyConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyParameters.
func (in *PolicyParameters) DeepCopy() *PolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedMetricSpecificationObservation) DeepCopyInto(out *PredefinedMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedMetricSpecificationObservation.
func (in *PredefinedMetricSpecificationObservation) DeepCopy() *PredefinedMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(PredefinedMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedMetricSpecificationParameters) DeepCopyInto(out *PredefinedMetricSpecificationParameters) {
	*out = *in
	if in.PredefinedMetricType != nil {
		in, out := &in.PredefinedMetricType, &out.PredefinedMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedMetricSpecificationParameters.
func (in *PredefinedMetricSpecificationParameters) DeepCopy() *PredefinedMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(PredefinedMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetActionObservation) DeepCopyInto(out *ScalableTargetActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetActionObservation.
func (in *ScalableTargetActionObservation) DeepCopy() *ScalableTargetActionObservation {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetActionParameters) DeepCopyInto(out *ScalableTargetActionParameters) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(string)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetActionParameters.
func (in *ScalableTargetActionParameters) DeepCopy() *ScalableTargetActionParameters {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledAction) DeepCopyInto(out *ScheduledAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledAction.
func (in *ScheduledAction) DeepCopy() *ScheduledAction {
	if in == nil {
		return nil
	}
	out := new(ScheduledAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledActionList) DeepCopyInto(out *ScheduledActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledActionList.
func (in *ScheduledActionList) DeepCopy() *ScheduledActionList {
	if in == nil {
		return nil
	}
	out := new(ScheduledActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledActionObservation) DeepCopyInto(out *ScheduledActionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledActionObservation.
func (in *ScheduledActionObservation) DeepCopy() *ScheduledActionObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduledActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledActionParameters) DeepCopyInto(out *ScheduledActionParameters) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDRef != nil {
		in, out := &in.ResourceIDRef, &out.ResourceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceIDSelector != nil {
		in, out := &in.ResourceIDSelector, &out.ResourceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimensionRef != nil {
		in, out := &in.ScalableDimensionRef, &out.ScalableDimensionRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableDimensionSelector != nil {
		in, out := &in.ScalableDimensionSelector, &out.ScalableDimensionSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableTargetAction != nil {
		in, out := &in.ScalableTargetAction, &out.ScalableTargetAction
		*out = make([]ScalableTargetActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespaceRef != nil {
		in, out := &in.ServiceNamespaceRef, &out.ServiceNamespaceRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceNamespaceSelector != nil {
		in, out := &in.ServiceNamespaceSelector, &out.ServiceNamespaceSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledActionParameters.
func (in *ScheduledActionParameters) DeepCopy() *ScheduledActionParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduledActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledActionSpec) DeepCopyInto(out *ScheduledActionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledActionSpec.
func (in *ScheduledActionSpec) DeepCopy() *ScheduledActionSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledActionStatus) DeepCopyInto(out *ScheduledActionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledActionStatus.
func (in *ScheduledActionStatus) DeepCopy() *ScheduledActionStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepAdjustmentObservation) DeepCopyInto(out *StepAdjustmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepAdjustmentObservation.
func (in *StepAdjustmentObservation) DeepCopy() *StepAdjustmentObservation {
	if in == nil {
		return nil
	}
	out := new(StepAdjustmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepAdjustmentParameters) DeepCopyInto(out *StepAdjustmentParameters) {
	*out = *in
	if in.MetricIntervalLowerBound != nil {
		in, out := &in.MetricIntervalLowerBound, &out.MetricIntervalLowerBound
		*out = new(string)
		**out = **in
	}
	if in.MetricIntervalUpperBound != nil {
		in, out := &in.MetricIntervalUpperBound, &out.MetricIntervalUpperBound
		*out = new(string)
		**out = **in
	}
	if in.ScalingAdjustment != nil {
		in, out := &in.ScalingAdjustment, &out.ScalingAdjustment
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepAdjustmentParameters.
func (in *StepAdjustmentParameters) DeepCopy() *StepAdjustmentParameters {
	if in == nil {
		return nil
	}
	out := new(StepAdjustmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepScalingPolicyConfigurationObservation) DeepCopyInto(out *StepScalingPolicyConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepScalingPolicyConfigurationObservation.
func (in *StepScalingPolicyConfigurationObservation) DeepCopy() *StepScalingPolicyConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(StepScalingPolicyConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepScalingPolicyConfigurationParameters) DeepCopyInto(out *StepScalingPolicyConfigurationParameters) {
	*out = *in
	if in.AdjustmentType != nil {
		in, out := &in.AdjustmentType, &out.AdjustmentType
		*out = new(string)
		**out = **in
	}
	if in.Cooldown != nil {
		in, out := &in.Cooldown, &out.Cooldown
		*out = new(float64)
		**out = **in
	}
	if in.MetricAggregationType != nil {
		in, out := &in.MetricAggregationType, &out.MetricAggregationType
		*out = new(string)
		**out = **in
	}
	if in.MinAdjustmentMagnitude != nil {
		in, out := &in.MinAdjustmentMagnitude, &out.MinAdjustmentMagnitude
		*out = new(float64)
		**out = **in
	}
	if in.StepAdjustment != nil {
		in, out := &in.StepAdjustment, &out.StepAdjustment
		*out = make([]StepAdjustmentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepScalingPolicyConfigurationParameters.
func (in *StepScalingPolicyConfigurationParameters) DeepCopy() *StepScalingPolicyConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(StepScalingPolicyConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Target) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetList) DeepCopyInto(out *TargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Target, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetList.
func (in *TargetList) DeepCopy() *TargetList {
	if in == nil {
		return nil
	}
	out := new(TargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetObservation) DeepCopyInto(out *TargetObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetObservation.
func (in *TargetObservation) DeepCopy() *TargetObservation {
	if in == nil {
		return nil
	}
	out := new(TargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParameters) DeepCopyInto(out *TargetParameters) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(float64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleArnRef != nil {
		in, out := &in.RoleArnRef, &out.RoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArnSelector != nil {
		in, out := &in.RoleArnSelector, &out.RoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParameters.
func (in *TargetParameters) DeepCopy() *TargetParameters {
	if in == nil {
		return nil
	}
	out := new(TargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetStatus) DeepCopyInto(out *TargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetStatus.
func (in *TargetStatus) DeepCopy() *TargetStatus {
	if in == nil {
		return nil
	}
	out := new(TargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingScalingPolicyConfigurationObservation) DeepCopyInto(out *TargetTrackingScalingPolicyConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingScalingPolicyConfigurationObservation.
func (in *TargetTrackingScalingPolicyConfigurationObservation) DeepCopy() *TargetTrackingScalingPolicyConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingScalingPolicyConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingScalingPolicyConfigurationParameters) DeepCopyInto(out *TargetTrackingScalingPolicyConfigurationParameters) {
	*out = *in
	if in.CustomizedMetricSpecification != nil {
		in, out := &in.CustomizedMetricSpecification, &out.CustomizedMetricSpecification
		*out = make([]CustomizedMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisableScaleIn != nil {
		in, out := &in.DisableScaleIn, &out.DisableScaleIn
		*out = new(bool)
		**out = **in
	}
	if in.PredefinedMetricSpecification != nil {
		in, out := &in.PredefinedMetricSpecification, &out.PredefinedMetricSpecification
		*out = make([]PredefinedMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScaleInCooldown != nil {
		in, out := &in.ScaleInCooldown, &out.ScaleInCooldown
		*out = new(float64)
		**out = **in
	}
	if in.ScaleOutCooldown != nil {
		in, out := &in.ScaleOutCooldown, &out.ScaleOutCooldown
		*out = new(float64)
		**out = **in
	}
	if in.TargetValue != nil {
		in, out := &in.TargetValue, &out.TargetValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingScalingPolicyConfigurationParameters.
func (in *TargetTrackingScalingPolicyConfigurationParameters) DeepCopy() *TargetTrackingScalingPolicyConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingScalingPolicyConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
