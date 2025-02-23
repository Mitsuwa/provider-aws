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

type AsyncInferenceConfigObservation struct {
}

type AsyncInferenceConfigOutputConfigObservation struct {
}

type AsyncInferenceConfigOutputConfigParameters struct {

	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Specifies the configuration for notifications of inference results for asynchronous inference.
	// +kubebuilder:validation:Optional
	NotificationConfig []NotificationConfigParameters `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// The Amazon S3 location to upload inference responses to.
	// +kubebuilder:validation:Required
	S3OutputPath *string `json:"s3OutputPath" tf:"s3_output_path,omitempty"`
}

type AsyncInferenceConfigParameters struct {

	// Configures the behavior of the client used by Amazon SageMaker to interact with the model container during asynchronous inference.
	// +kubebuilder:validation:Optional
	ClientConfig []ClientConfigParameters `json:"clientConfig,omitempty" tf:"client_config,omitempty"`

	// Specifies the configuration for asynchronous inference invocation outputs.
	// +kubebuilder:validation:Required
	OutputConfig []AsyncInferenceConfigOutputConfigParameters `json:"outputConfig" tf:"output_config,omitempty"`
}

type CaptureContentTypeHeaderObservation struct {
}

type CaptureContentTypeHeaderParameters struct {

	// The CSV content type headers to capture.
	// +kubebuilder:validation:Optional
	CsvContentTypes []*string `json:"csvContentTypes,omitempty" tf:"csv_content_types,omitempty"`

	// The JSON content type headers to capture.
	// +kubebuilder:validation:Optional
	JSONContentTypes []*string `json:"jsonContentTypes,omitempty" tf:"json_content_types,omitempty"`
}

type CaptureOptionsObservation struct {
}

type CaptureOptionsParameters struct {

	// Specifies the data to be captured. Should be one of Input or Output.
	// +kubebuilder:validation:Required
	CaptureMode *string `json:"captureMode" tf:"capture_mode,omitempty"`
}

type ClientConfigObservation struct {
}

type ClientConfigParameters struct {

	// The maximum number of concurrent requests sent by the SageMaker client to the model container. If no value is provided, Amazon SageMaker will choose an optimal value for you.
	// +kubebuilder:validation:Optional
	MaxConcurrentInvocationsPerInstance *float64 `json:"maxConcurrentInvocationsPerInstance,omitempty" tf:"max_concurrent_invocations_per_instance,omitempty"`
}

type CoreDumpConfigObservation struct {
}

type CoreDumpConfigParameters struct {

	// The Amazon S3 bucket to send the core dump to.
	// +kubebuilder:validation:Required
	DestinationS3URI *string `json:"destinationS3Uri" tf:"destination_s3_uri,omitempty"`

	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type DataCaptureConfigObservation struct {
}

type DataCaptureConfigParameters struct {

	// The content type headers to capture. Fields are documented below.
	// +kubebuilder:validation:Optional
	CaptureContentTypeHeader []CaptureContentTypeHeaderParameters `json:"captureContentTypeHeader,omitempty" tf:"capture_content_type_header,omitempty"`

	// Specifies what data to capture. Fields are documented below.
	// +kubebuilder:validation:Required
	CaptureOptions []CaptureOptionsParameters `json:"captureOptions" tf:"capture_options,omitempty"`

	// The URL for S3 location where the captured data is stored.
	// +kubebuilder:validation:Required
	DestinationS3URI *string `json:"destinationS3Uri" tf:"destination_s3_uri,omitempty"`

	// Flag to enable data capture. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableCapture *bool `json:"enableCapture,omitempty" tf:"enable_capture,omitempty"`

	// Portion of data to capture. Should be between 0 and 100.
	// +kubebuilder:validation:Required
	InitialSamplingPercentage *float64 `json:"initialSamplingPercentage" tf:"initial_sampling_percentage,omitempty"`

	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt the captured data on Amazon S3.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type EndpointConfigurationObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EndpointConfigurationParameters struct {

	// Specifies configuration for how an endpoint performs asynchronous inference.
	// +kubebuilder:validation:Optional
	AsyncInferenceConfig []AsyncInferenceConfigParameters `json:"asyncInferenceConfig,omitempty" tf:"async_inference_config,omitempty"`

	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	// +kubebuilder:validation:Optional
	DataCaptureConfig []DataCaptureConfigParameters `json:"dataCaptureConfig,omitempty" tf:"data_capture_config,omitempty"`

	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	// +kubebuilder:validation:Required
	ProductionVariants []ProductionVariantsParameters `json:"productionVariants" tf:"production_variants,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	// +kubebuilder:validation:Optional
	ShadowProductionVariants []ShadowProductionVariantsParameters `json:"shadowProductionVariants,omitempty" tf:"shadow_production_variants,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NotificationConfigObservation struct {
}

type NotificationConfigParameters struct {

	// Amazon SNS topic to post a notification to when inference fails. If no topic is provided, no notification is sent on failure.
	// +kubebuilder:validation:Optional
	ErrorTopic *string `json:"errorTopic,omitempty" tf:"error_topic,omitempty"`

	// Amazon SNS topic to post a notification to when inference completes successfully. If no topic is provided, no notification is sent on success.
	// +kubebuilder:validation:Optional
	SuccessTopic *string `json:"successTopic,omitempty" tf:"success_topic,omitempty"`
}

type ProductionVariantsObservation struct {
}

type ProductionVariantsParameters struct {

	// The size of the Elastic Inference (EI) instance to use for the production variant.
	// +kubebuilder:validation:Optional
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting. For more information about health check, see How Your Container Should Respond to Health Check (Ping) Requests. Valid values between 60 and 3600.
	// +kubebuilder:validation:Optional
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `json:"containerStartupHealthCheckTimeoutInSeconds,omitempty" tf:"container_startup_health_check_timeout_in_seconds,omitempty"`

	// Specifies configuration for a core dump from the model container when the process crashes. Fields are documented below.
	// +kubebuilder:validation:Optional
	CoreDumpConfig []CoreDumpConfigParameters `json:"coreDumpConfig,omitempty" tf:"core_dump_config,omitempty"`

	// Initial number of instances used for auto-scaling.
	// +kubebuilder:validation:Optional
	InitialInstanceCount *float64 `json:"initialInstanceCount,omitempty" tf:"initial_instance_count,omitempty"`

	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration. If unspecified, it defaults to 1.0.
	// +kubebuilder:validation:Optional
	InitialVariantWeight *float64 `json:"initialVariantWeight,omitempty" tf:"initial_variant_weight,omitempty"`

	// The type of instance to start.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant. Valid values between 60 and 3600.
	// +kubebuilder:validation:Optional
	ModelDataDownloadTimeoutInSeconds *float64 `json:"modelDataDownloadTimeoutInSeconds,omitempty" tf:"model_data_download_timeout_in_seconds,omitempty"`

	// The name of the model to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Model
	// +kubebuilder:validation:Optional
	ModelName *string `json:"modelName,omitempty" tf:"model_name,omitempty"`

	// Reference to a Model in sagemaker to populate modelName.
	// +kubebuilder:validation:Optional
	ModelNameRef *v1.Reference `json:"modelNameRef,omitempty" tf:"-"`

	// Selector for a Model in sagemaker to populate modelName.
	// +kubebuilder:validation:Optional
	ModelNameSelector *v1.Selector `json:"modelNameSelector,omitempty" tf:"-"`

	// Specifies configuration for how an endpoint performs asynchronous inference.
	// +kubebuilder:validation:Optional
	ServerlessConfig []ServerlessConfigParameters `json:"serverlessConfig,omitempty" tf:"serverless_config,omitempty"`

	// The name of the variant.
	// +kubebuilder:validation:Optional
	VariantName *string `json:"variantName,omitempty" tf:"variant_name,omitempty"`

	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Valid values between 1 and 512.
	// +kubebuilder:validation:Optional
	VolumeSizeInGb *float64 `json:"volumeSizeInGb,omitempty" tf:"volume_size_in_gb,omitempty"`
}

type ServerlessConfigObservation struct {
}

type ServerlessConfigParameters struct {

	// The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between 1 and 200.
	// +kubebuilder:validation:Required
	MaxConcurrency *float64 `json:"maxConcurrency" tf:"max_concurrency,omitempty"`

	// The memory size of your serverless endpoint. Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	// +kubebuilder:validation:Required
	MemorySizeInMb *float64 `json:"memorySizeInMb" tf:"memory_size_in_mb,omitempty"`
}

type ShadowProductionVariantsCoreDumpConfigObservation struct {
}

type ShadowProductionVariantsCoreDumpConfigParameters struct {

	// The Amazon S3 bucket to send the core dump to.
	// +kubebuilder:validation:Required
	DestinationS3URI *string `json:"destinationS3Uri" tf:"destination_s3_uri,omitempty"`

	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption.
	// +kubebuilder:validation:Required
	KMSKeyID *string `json:"kmsKeyId" tf:"kms_key_id,omitempty"`
}

type ShadowProductionVariantsObservation struct {
}

type ShadowProductionVariantsParameters struct {

	// The size of the Elastic Inference (EI) instance to use for the production variant.
	// +kubebuilder:validation:Optional
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting. For more information about health check, see How Your Container Should Respond to Health Check (Ping) Requests. Valid values between 60 and 3600.
	// +kubebuilder:validation:Optional
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `json:"containerStartupHealthCheckTimeoutInSeconds,omitempty" tf:"container_startup_health_check_timeout_in_seconds,omitempty"`

	// Specifies configuration for a core dump from the model container when the process crashes. Fields are documented below.
	// +kubebuilder:validation:Optional
	CoreDumpConfig []ShadowProductionVariantsCoreDumpConfigParameters `json:"coreDumpConfig,omitempty" tf:"core_dump_config,omitempty"`

	// Initial number of instances used for auto-scaling.
	// +kubebuilder:validation:Optional
	InitialInstanceCount *float64 `json:"initialInstanceCount,omitempty" tf:"initial_instance_count,omitempty"`

	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration. If unspecified, it defaults to 1.0.
	// +kubebuilder:validation:Optional
	InitialVariantWeight *float64 `json:"initialVariantWeight,omitempty" tf:"initial_variant_weight,omitempty"`

	// The type of instance to start.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant. Valid values between 60 and 3600.
	// +kubebuilder:validation:Optional
	ModelDataDownloadTimeoutInSeconds *float64 `json:"modelDataDownloadTimeoutInSeconds,omitempty" tf:"model_data_download_timeout_in_seconds,omitempty"`

	// The name of the model to use.
	// +kubebuilder:validation:Required
	ModelName *string `json:"modelName" tf:"model_name,omitempty"`

	// Specifies configuration for how an endpoint performs asynchronous inference.
	// +kubebuilder:validation:Optional
	ServerlessConfig []ShadowProductionVariantsServerlessConfigParameters `json:"serverlessConfig,omitempty" tf:"serverless_config,omitempty"`

	// The name of the variant.
	// +kubebuilder:validation:Optional
	VariantName *string `json:"variantName,omitempty" tf:"variant_name,omitempty"`

	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Valid values between 1 and 512.
	// +kubebuilder:validation:Optional
	VolumeSizeInGb *float64 `json:"volumeSizeInGb,omitempty" tf:"volume_size_in_gb,omitempty"`
}

type ShadowProductionVariantsServerlessConfigObservation struct {
}

type ShadowProductionVariantsServerlessConfigParameters struct {

	// The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between 1 and 200.
	// +kubebuilder:validation:Required
	MaxConcurrency *float64 `json:"maxConcurrency" tf:"max_concurrency,omitempty"`

	// The memory size of your serverless endpoint. Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	// +kubebuilder:validation:Required
	MemorySizeInMb *float64 `json:"memorySizeInMb" tf:"memory_size_in_mb,omitempty"`
}

// EndpointConfigurationSpec defines the desired state of EndpointConfiguration
type EndpointConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointConfigurationParameters `json:"forProvider"`
}

// EndpointConfigurationStatus defines the observed state of EndpointConfiguration.
type EndpointConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointConfiguration is the Schema for the EndpointConfigurations API. Provides a SageMaker Endpoint Configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointConfigurationSpec   `json:"spec"`
	Status            EndpointConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointConfigurationList contains a list of EndpointConfigurations
type EndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointConfiguration `json:"items"`
}

// Repository type metadata.
var (
	EndpointConfiguration_Kind             = "EndpointConfiguration"
	EndpointConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointConfiguration_Kind}.String()
	EndpointConfiguration_KindAPIVersion   = EndpointConfiguration_Kind + "." + CRDGroupVersion.String()
	EndpointConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(EndpointConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointConfiguration{}, &EndpointConfigurationList{})
}
