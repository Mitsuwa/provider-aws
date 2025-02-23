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

type GlobalClusterMembersObservation struct {

	// Amazon Resource Name (ARN) of member DB Cluster.
	DBClusterArn *string `json:"dbClusterArn,omitempty" tf:"db_cluster_arn,omitempty"`

	// Whether the member is the primary DB Cluster.
	IsWriter *bool `json:"isWriter,omitempty" tf:"is_writer,omitempty"`
}

type GlobalClusterMembersParameters struct {
}

type GlobalClusterObservation struct {

	// Global Cluster Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Set of objects containing Global Cluster members.
	GlobalClusterMembers []GlobalClusterMembersObservation `json:"globalClusterMembers,omitempty" tf:"global_cluster_members,omitempty"`

	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.
	GlobalClusterResourceID *string `json:"globalClusterResourceId,omitempty" tf:"global_cluster_resource_id,omitempty"`

	// Neptune Global Cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type GlobalClusterParameters struct {

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Name of the database engine to be used for this DB cluster. Current Valid values: neptune. Conflicts with source_db_cluster_identifier.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/neptune/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifier *string `json:"sourceDbClusterIdentifier,omitempty" tf:"source_db_cluster_identifier,omitempty"`

	// Reference to a Cluster in neptune to populate sourceDbClusterIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifierRef *v1.Reference `json:"sourceDbClusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in neptune to populate sourceDbClusterIdentifier.
	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifierSelector *v1.Selector `json:"sourceDbClusterIdentifierSelector,omitempty" tf:"-"`

	// Specifies whether the DB cluster is encrypted. The default is false unless source_db_cluster_identifier is specified and encrypted.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

// GlobalClusterSpec defines the desired state of GlobalCluster
type GlobalClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalClusterParameters `json:"forProvider"`
}

// GlobalClusterStatus defines the observed state of GlobalCluster.
type GlobalClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalCluster is the Schema for the GlobalClusters API. Provides an Neptune Global Cluster Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalClusterSpec   `json:"spec"`
	Status            GlobalClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalClusterList contains a list of GlobalClusters
type GlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalCluster `json:"items"`
}

// Repository type metadata.
var (
	GlobalCluster_Kind             = "GlobalCluster"
	GlobalCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalCluster_Kind}.String()
	GlobalCluster_KindAPIVersion   = GlobalCluster_Kind + "." + CRDGroupVersion.String()
	GlobalCluster_GroupVersionKind = CRDGroupVersion.WithKind(GlobalCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalCluster{}, &GlobalClusterList{})
}
