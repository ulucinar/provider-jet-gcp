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

type NetworkSettingsObservation struct {
}

type NetworkSettingsParameters struct {

	// The ingress settings for version or service. Default value: "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED" Possible values: ["INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB"]
	// +kubebuilder:validation:Optional
	IngressTrafficAllowed *string `json:"ingressTrafficAllowed,omitempty" tf:"ingress_traffic_allowed,omitempty"`
}

type ServiceNetworkSettingsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceNetworkSettingsParameters struct {

	// Ingress settings for this service. Will apply to all versions.
	// +kubebuilder:validation:Required
	NetworkSettings []NetworkSettingsParameters `json:"networkSettings" tf:"network_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the service these settings apply to.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ServiceNetworkSettingsSpec defines the desired state of ServiceNetworkSettings
type ServiceNetworkSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceNetworkSettingsParameters `json:"forProvider"`
}

// ServiceNetworkSettingsStatus defines the observed state of ServiceNetworkSettings.
type ServiceNetworkSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceNetworkSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkSettings is the Schema for the ServiceNetworkSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ServiceNetworkSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceNetworkSettingsSpec   `json:"spec"`
	Status            ServiceNetworkSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkSettingsList contains a list of ServiceNetworkSettingss
type ServiceNetworkSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkSettings `json:"items"`
}

// Repository type metadata.
var (
	ServiceNetworkSettings_Kind             = "ServiceNetworkSettings"
	ServiceNetworkSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceNetworkSettings_Kind}.String()
	ServiceNetworkSettings_KindAPIVersion   = ServiceNetworkSettings_Kind + "." + CRDGroupVersion.String()
	ServiceNetworkSettings_GroupVersionKind = CRDGroupVersion.WithKind(ServiceNetworkSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceNetworkSettings{}, &ServiceNetworkSettingsList{})
}
