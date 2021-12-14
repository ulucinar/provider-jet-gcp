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

type TenantOauthIdpConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TenantOauthIdpConfigParameters struct {

	// The client id of an OAuth client.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The client secret of the OAuth client, to enable OIDC code flow.
	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// Human friendly display name.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with 'oidc.'.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the tenant where this OIDC IDP configuration resource exists
	// +kubebuilder:validation:Required
	Tenant *string `json:"tenant" tf:"tenant,omitempty"`
}

// TenantOauthIdpConfigSpec defines the desired state of TenantOauthIdpConfig
type TenantOauthIdpConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantOauthIdpConfigParameters `json:"forProvider"`
}

// TenantOauthIdpConfigStatus defines the observed state of TenantOauthIdpConfig.
type TenantOauthIdpConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantOauthIdpConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TenantOauthIdpConfig is the Schema for the TenantOauthIdpConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TenantOauthIdpConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantOauthIdpConfigSpec   `json:"spec"`
	Status            TenantOauthIdpConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantOauthIdpConfigList contains a list of TenantOauthIdpConfigs
type TenantOauthIdpConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantOauthIdpConfig `json:"items"`
}

// Repository type metadata.
var (
	TenantOauthIdpConfig_Kind             = "TenantOauthIdpConfig"
	TenantOauthIdpConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TenantOauthIdpConfig_Kind}.String()
	TenantOauthIdpConfig_KindAPIVersion   = TenantOauthIdpConfig_Kind + "." + CRDGroupVersion.String()
	TenantOauthIdpConfig_GroupVersionKind = CRDGroupVersion.WithKind(TenantOauthIdpConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&TenantOauthIdpConfig{}, &TenantOauthIdpConfigList{})
}
