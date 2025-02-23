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
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateObservation) DeepCopyInto(out *CertificateObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DomainValidationOptions != nil {
		in, out := &in.DomainValidationOptions, &out.DomainValidationOptions
		*out = make([]DomainValidationOptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = new(string)
		**out = **in
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = new(string)
		**out = **in
	}
	if in.PendingRenewal != nil {
		in, out := &in.PendingRenewal, &out.PendingRenewal
		*out = new(bool)
		**out = **in
	}
	if in.RenewalEligibility != nil {
		in, out := &in.RenewalEligibility, &out.RenewalEligibility
		*out = new(string)
		**out = **in
	}
	if in.RenewalSummary != nil {
		in, out := &in.RenewalSummary, &out.RenewalSummary
		*out = make([]RenewalSummaryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationEmails != nil {
		in, out := &in.ValidationEmails, &out.ValidationEmails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateObservation.
func (in *CertificateObservation) DeepCopy() *CertificateObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateParameters) DeepCopyInto(out *CertificateParameters) {
	*out = *in
	if in.CertificateAuthorityArn != nil {
		in, out := &in.CertificateAuthorityArn, &out.CertificateAuthorityArn
		*out = new(string)
		**out = **in
	}
	if in.CertificateBody != nil {
		in, out := &in.CertificateBody, &out.CertificateBody
		*out = new(string)
		**out = **in
	}
	if in.CertificateChain != nil {
		in, out := &in.CertificateChain, &out.CertificateChain
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.EarlyRenewalDuration != nil {
		in, out := &in.EarlyRenewalDuration, &out.EarlyRenewalDuration
		*out = new(string)
		**out = **in
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.ValidationOption != nil {
		in, out := &in.ValidationOption, &out.ValidationOption
		*out = make([]ValidationOptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateParameters.
func (in *CertificateParameters) DeepCopy() *CertificateParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidation) DeepCopyInto(out *CertificateValidation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidation.
func (in *CertificateValidation) DeepCopy() *CertificateValidation {
	if in == nil {
		return nil
	}
	out := new(CertificateValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateValidation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidationList) DeepCopyInto(out *CertificateValidationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateValidation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidationList.
func (in *CertificateValidationList) DeepCopy() *CertificateValidationList {
	if in == nil {
		return nil
	}
	out := new(CertificateValidationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateValidationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidationObservation) DeepCopyInto(out *CertificateValidationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidationObservation.
func (in *CertificateValidationObservation) DeepCopy() *CertificateValidationObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateValidationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidationParameters) DeepCopyInto(out *CertificateValidationParameters) {
	*out = *in
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
	if in.CertificateArnRef != nil {
		in, out := &in.CertificateArnRef, &out.CertificateArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateArnSelector != nil {
		in, out := &in.CertificateArnSelector, &out.CertificateArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ValidationRecordFqdns != nil {
		in, out := &in.ValidationRecordFqdns, &out.ValidationRecordFqdns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidationParameters.
func (in *CertificateValidationParameters) DeepCopy() *CertificateValidationParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateValidationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidationSpec) DeepCopyInto(out *CertificateValidationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidationSpec.
func (in *CertificateValidationSpec) DeepCopy() *CertificateValidationSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateValidationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateValidationStatus) DeepCopyInto(out *CertificateValidationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateValidationStatus.
func (in *CertificateValidationStatus) DeepCopy() *CertificateValidationStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateValidationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidationOptionsObservation) DeepCopyInto(out *DomainValidationOptionsObservation) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ResourceRecordName != nil {
		in, out := &in.ResourceRecordName, &out.ResourceRecordName
		*out = new(string)
		**out = **in
	}
	if in.ResourceRecordType != nil {
		in, out := &in.ResourceRecordType, &out.ResourceRecordType
		*out = new(string)
		**out = **in
	}
	if in.ResourceRecordValue != nil {
		in, out := &in.ResourceRecordValue, &out.ResourceRecordValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidationOptionsObservation.
func (in *DomainValidationOptionsObservation) DeepCopy() *DomainValidationOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DomainValidationOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidationOptionsParameters) DeepCopyInto(out *DomainValidationOptionsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidationOptionsParameters.
func (in *DomainValidationOptionsParameters) DeepCopy() *DomainValidationOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DomainValidationOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsObservation) DeepCopyInto(out *OptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsObservation.
func (in *OptionsObservation) DeepCopy() *OptionsObservation {
	if in == nil {
		return nil
	}
	out := new(OptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsParameters) DeepCopyInto(out *OptionsParameters) {
	*out = *in
	if in.CertificateTransparencyLoggingPreference != nil {
		in, out := &in.CertificateTransparencyLoggingPreference, &out.CertificateTransparencyLoggingPreference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsParameters.
func (in *OptionsParameters) DeepCopy() *OptionsParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenewalSummaryObservation) DeepCopyInto(out *RenewalSummaryObservation) {
	*out = *in
	if in.RenewalStatus != nil {
		in, out := &in.RenewalStatus, &out.RenewalStatus
		*out = new(string)
		**out = **in
	}
	if in.RenewalStatusReason != nil {
		in, out := &in.RenewalStatusReason, &out.RenewalStatusReason
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenewalSummaryObservation.
func (in *RenewalSummaryObservation) DeepCopy() *RenewalSummaryObservation {
	if in == nil {
		return nil
	}
	out := new(RenewalSummaryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenewalSummaryParameters) DeepCopyInto(out *RenewalSummaryParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenewalSummaryParameters.
func (in *RenewalSummaryParameters) DeepCopy() *RenewalSummaryParameters {
	if in == nil {
		return nil
	}
	out := new(RenewalSummaryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationOptionObservation) DeepCopyInto(out *ValidationOptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationOptionObservation.
func (in *ValidationOptionObservation) DeepCopy() *ValidationOptionObservation {
	if in == nil {
		return nil
	}
	out := new(ValidationOptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationOptionParameters) DeepCopyInto(out *ValidationOptionParameters) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ValidationDomain != nil {
		in, out := &in.ValidationDomain, &out.ValidationDomain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationOptionParameters.
func (in *ValidationOptionParameters) DeepCopy() *ValidationOptionParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationOptionParameters)
	in.DeepCopyInto(out)
	return out
}
