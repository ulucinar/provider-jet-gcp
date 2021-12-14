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

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type DiskIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskIamBindingSpec defines the desired state of DiskIamBinding
type DiskIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskIamBindingParameters `json:"forProvider"`
}

// DiskIamBindingStatus defines the observed state of DiskIamBinding.
type DiskIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamBinding is the Schema for the DiskIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DiskIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskIamBindingSpec   `json:"spec"`
	Status            DiskIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamBindingList contains a list of DiskIamBindings
type DiskIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskIamBinding `json:"items"`
}

// Repository type metadata.
var (
	DiskIamBinding_Kind             = "DiskIamBinding"
	DiskIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskIamBinding_Kind}.String()
	DiskIamBinding_KindAPIVersion   = DiskIamBinding_Kind + "." + CRDGroupVersion.String()
	DiskIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(DiskIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskIamBinding{}, &DiskIamBindingList{})
}
