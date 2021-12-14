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

type WebIamBindingConditionObservation struct {
}

type WebIamBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type WebIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WebIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []WebIamBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// WebIamBindingSpec defines the desired state of WebIamBinding
type WebIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebIamBindingParameters `json:"forProvider"`
}

// WebIamBindingStatus defines the observed state of WebIamBinding.
type WebIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebIamBinding is the Schema for the WebIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type WebIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebIamBindingSpec   `json:"spec"`
	Status            WebIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebIamBindingList contains a list of WebIamBindings
type WebIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebIamBinding `json:"items"`
}

// Repository type metadata.
var (
	WebIamBinding_Kind             = "WebIamBinding"
	WebIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebIamBinding_Kind}.String()
	WebIamBinding_KindAPIVersion   = WebIamBinding_Kind + "." + CRDGroupVersion.String()
	WebIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(WebIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&WebIamBinding{}, &WebIamBindingList{})
}
