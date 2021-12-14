/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigObservation struct {
	CloudUpdateTime *string `json:"cloudUpdateTime,omitempty" tf:"cloud_update_time,omitempty"`

	DeviceAckTime *string `json:"deviceAckTime,omitempty" tf:"device_ack_time,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ConfigParameters struct {

	// The device configuration data.
	// +kubebuilder:validation:Optional
	BinaryData *string `json:"binaryData,omitempty" tf:"binary_data,omitempty"`
}

type CredentialsObservation struct {
}

type CredentialsParameters struct {

	// The time at which this credential becomes invalid.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// A public key used to verify the signature of JSON Web Tokens (JWTs).
	// +kubebuilder:validation:Required
	PublicKey []PublicKeyParameters `json:"publicKey" tf:"public_key,omitempty"`
}

type DeviceObservation struct {
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastConfigAckTime *string `json:"lastConfigAckTime,omitempty" tf:"last_config_ack_time,omitempty"`

	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty" tf:"last_config_send_time,omitempty"`

	LastErrorStatus []LastErrorStatusObservation `json:"lastErrorStatus,omitempty" tf:"last_error_status,omitempty"`

	LastErrorTime *string `json:"lastErrorTime,omitempty" tf:"last_error_time,omitempty"`

	LastEventTime *string `json:"lastEventTime,omitempty" tf:"last_event_time,omitempty"`

	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty" tf:"last_heartbeat_time,omitempty"`

	LastStateTime *string `json:"lastStateTime,omitempty" tf:"last_state_time,omitempty"`

	NumID *string `json:"numId,omitempty" tf:"num_id,omitempty"`

	State []StateObservation `json:"state,omitempty" tf:"state,omitempty"`
}

type DeviceParameters struct {

	// If a device is blocked, connections or requests from this device will fail.
	// +kubebuilder:validation:Optional
	Blocked *bool `json:"blocked,omitempty" tf:"blocked,omitempty"`

	// The credentials used to authenticate this device.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Gateway-related configuration and state.
	// +kubebuilder:validation:Optional
	GatewayConfig []GatewayConfigParameters `json:"gatewayConfig,omitempty" tf:"gateway_config,omitempty"`

	// The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// The metadata key-value pairs assigned to the device.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the resource.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the device registry where this device should be created.
	// +kubebuilder:validation:Required
	Registry *string `json:"registry" tf:"registry,omitempty"`
}

type GatewayConfigObservation struct {
	LastAccessedGatewayID *string `json:"lastAccessedGatewayId,omitempty" tf:"last_accessed_gateway_id,omitempty"`

	LastAccessedGatewayTime *string `json:"lastAccessedGatewayTime,omitempty" tf:"last_accessed_gateway_time,omitempty"`
}

type GatewayConfigParameters struct {

	// Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"]
	// +kubebuilder:validation:Optional
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty" tf:"gateway_auth_method,omitempty"`

	// Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"]
	// +kubebuilder:validation:Optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`
}

type LastErrorStatusObservation struct {
}

type LastErrorStatusParameters struct {

	// A list of messages that carry the error details.
	// +kubebuilder:validation:Optional
	Details []string `json:"details,omitempty" tf:"details,omitempty"`

	// A developer-facing error message, which should be in English.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The status code, which should be an enum value of google.rpc.Code.
	// +kubebuilder:validation:Optional
	Number *int64 `json:"number,omitempty" tf:"number,omitempty"`
}

type PublicKeyObservation struct {
}

type PublicKeyParameters struct {

	// The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"]
	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// The key data.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type StateObservation struct {
}

type StateParameters struct {

	// The device state data.
	// +kubebuilder:validation:Optional
	BinaryData *string `json:"binaryData,omitempty" tf:"binary_data,omitempty"`

	// The time at which this state version was updated in Cloud IoT Core.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Device is the Schema for the Devices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSpec   `json:"spec"`
	Status            DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
