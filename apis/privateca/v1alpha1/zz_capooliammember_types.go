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

type CaPoolIamMemberConditionObservation struct {
}

type CaPoolIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type CaPoolIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CaPoolIamMemberParameters struct {

	// +kubebuilder:validation:Required
	CaPool *string `json:"caPool" tf:"ca_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []CaPoolIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// CaPoolIamMemberSpec defines the desired state of CaPoolIamMember
type CaPoolIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CaPoolIamMemberParameters `json:"forProvider"`
}

// CaPoolIamMemberStatus defines the observed state of CaPoolIamMember.
type CaPoolIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CaPoolIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CaPoolIamMember is the Schema for the CaPoolIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type CaPoolIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CaPoolIamMemberSpec   `json:"spec"`
	Status            CaPoolIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaPoolIamMemberList contains a list of CaPoolIamMembers
type CaPoolIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaPoolIamMember `json:"items"`
}

// Repository type metadata.
var (
	CaPoolIamMember_Kind             = "CaPoolIamMember"
	CaPoolIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CaPoolIamMember_Kind}.String()
	CaPoolIamMember_KindAPIVersion   = CaPoolIamMember_Kind + "." + CRDGroupVersion.String()
	CaPoolIamMember_GroupVersionKind = CRDGroupVersion.WithKind(CaPoolIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&CaPoolIamMember{}, &CaPoolIamMemberList{})
}
