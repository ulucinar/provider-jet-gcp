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

type JobIamMemberConditionObservation struct {
}

type JobIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type JobIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type JobIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []JobIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	JobID *string `json:"jobId" tf:"job_id,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// JobIamMemberSpec defines the desired state of JobIamMember
type JobIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobIamMemberParameters `json:"forProvider"`
}

// JobIamMemberStatus defines the observed state of JobIamMember.
type JobIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobIamMember is the Schema for the JobIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type JobIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobIamMemberSpec   `json:"spec"`
	Status            JobIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobIamMemberList contains a list of JobIamMembers
type JobIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobIamMember `json:"items"`
}

// Repository type metadata.
var (
	JobIamMember_Kind             = "JobIamMember"
	JobIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobIamMember_Kind}.String()
	JobIamMember_KindAPIVersion   = JobIamMember_Kind + "." + CRDGroupVersion.String()
	JobIamMember_GroupVersionKind = CRDGroupVersion.WithKind(JobIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&JobIamMember{}, &JobIamMemberList{})
}
