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

type AppEngineVersionIamMemberConditionObservation struct {
}

type AppEngineVersionIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type AppEngineVersionIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppEngineVersionIamMemberParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []AppEngineVersionIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`

	// +kubebuilder:validation:Required
	VersionID *string `json:"versionId" tf:"version_id,omitempty"`
}

// AppEngineVersionIamMemberSpec defines the desired state of AppEngineVersionIamMember
type AppEngineVersionIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppEngineVersionIamMemberParameters `json:"forProvider"`
}

// AppEngineVersionIamMemberStatus defines the observed state of AppEngineVersionIamMember.
type AppEngineVersionIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppEngineVersionIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineVersionIamMember is the Schema for the AppEngineVersionIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type AppEngineVersionIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppEngineVersionIamMemberSpec   `json:"spec"`
	Status            AppEngineVersionIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineVersionIamMemberList contains a list of AppEngineVersionIamMembers
type AppEngineVersionIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineVersionIamMember `json:"items"`
}

// Repository type metadata.
var (
	AppEngineVersionIamMember_Kind             = "AppEngineVersionIamMember"
	AppEngineVersionIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppEngineVersionIamMember_Kind}.String()
	AppEngineVersionIamMember_KindAPIVersion   = AppEngineVersionIamMember_Kind + "." + CRDGroupVersion.String()
	AppEngineVersionIamMember_GroupVersionKind = CRDGroupVersion.WithKind(AppEngineVersionIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&AppEngineVersionIamMember{}, &AppEngineVersionIamMemberList{})
}
