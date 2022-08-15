/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ReplicaObservation struct {

	// Date that you last accessed the secret in the Region.
	LastAccessedDate *string `json:"lastAccessedDate,omitempty" tf:"last_accessed_date,omitempty"`

	// Status can be InProgress, Failed, or InSync.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Message such as Replication succeeded or Secret with this name already exists in this region.
	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message,omitempty"`
}

type ReplicaParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Region for replicating the secret.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type RotationRulesObservation struct {
}

type RotationRulesParameters struct {

	// Specifies the number of days between automatic scheduled rotations of the secret.
	// +kubebuilder:validation:Required
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays" tf:"automatically_after_days,omitempty"`
}

type SecretObservation struct {

	// ARN of the secret.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the secret.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Replica []ReplicaObservation `json:"replica,omitempty" tf:"replica,omitempty"`

	// Whether automatic rotation is enabled for this secret.
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SecretParameters struct {

	// Description of the secret.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Accepts boolean value to specify whether to overwrite a secret with the same name in the destination Region.
	// +kubebuilder:validation:Optional
	ForceOverwriteReplicaSecret *bool `json:"forceOverwriteReplicaSecret,omitempty" tf:"force_overwrite_replica_secret,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: /_+=.@- Conflicts with name_prefix.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Valid JSON document representing a resource policy. For more information about building AWS IAM policy documents with Terraform, see the AWS IAM Policy Document Guide. Removing policy from your configuration or setting policy to null or an empty string  will not delete the policy since it could have been set by aws_secretsmanager_secret_policy. To delete the policy, set it to "{}" .
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be 0 to force deletion without recovery or range from 7 to 30 days. The default value is 30.
	// +kubebuilder:validation:Optional
	RecoveryWindowInDays *float64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days,omitempty"`

	// Region for replicating the secret.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Replica []ReplicaParameters `json:"replica,omitempty" tf:"replica,omitempty"`

	// ARN of the Lambda function that can rotate the secret. Use the aws_secretsmanager_secret_rotation resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RotationLambdaArnRef *v1.Reference `json:"rotationLambdaArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RotationLambdaArnSelector *v1.Selector `json:"rotationLambdaArnSelector,omitempty" tf:"-"`

	// Configuration block for the rotation configuration of this secret. Defined below. Use the aws_secretsmanager_secret_rotation resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
	// +kubebuilder:validation:Optional
	RotationRules []RotationRulesParameters `json:"rotationRules,omitempty" tf:"rotation_rules,omitempty"`

	// Key-value map of user-defined tags that are attached to the secret. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API. Provides a resource to manage AWS Secrets Manager secret metadata
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
