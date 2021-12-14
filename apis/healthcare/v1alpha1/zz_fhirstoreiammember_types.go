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

type FhirStoreIamMemberConditionObservation struct {
}

type FhirStoreIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type FhirStoreIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FhirStoreIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []FhirStoreIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	FhirStoreID *string `json:"fhirStoreId" tf:"fhir_store_id,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// FhirStoreIamMemberSpec defines the desired state of FhirStoreIamMember
type FhirStoreIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FhirStoreIamMemberParameters `json:"forProvider"`
}

// FhirStoreIamMemberStatus defines the observed state of FhirStoreIamMember.
type FhirStoreIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FhirStoreIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStoreIamMember is the Schema for the FhirStoreIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FhirStoreIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FhirStoreIamMemberSpec   `json:"spec"`
	Status            FhirStoreIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStoreIamMemberList contains a list of FhirStoreIamMembers
type FhirStoreIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FhirStoreIamMember `json:"items"`
}

// Repository type metadata.
var (
	FhirStoreIamMember_Kind             = "FhirStoreIamMember"
	FhirStoreIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FhirStoreIamMember_Kind}.String()
	FhirStoreIamMember_KindAPIVersion   = FhirStoreIamMember_Kind + "." + CRDGroupVersion.String()
	FhirStoreIamMember_GroupVersionKind = CRDGroupVersion.WithKind(FhirStoreIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&FhirStoreIamMember{}, &FhirStoreIamMemberList{})
}
