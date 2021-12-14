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

type ObjectAclObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ObjectAclParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Object *string `json:"object" tf:"object,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedACL *string `json:"predefinedAcl,omitempty" tf:"predefined_acl,omitempty"`

	// +kubebuilder:validation:Optional
	RoleEntity []*string `json:"roleEntity,omitempty" tf:"role_entity,omitempty"`
}

// ObjectAclSpec defines the desired state of ObjectAcl
type ObjectAclSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectAclParameters `json:"forProvider"`
}

// ObjectAclStatus defines the observed state of ObjectAcl.
type ObjectAclStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectAcl is the Schema for the ObjectAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ObjectAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectAclSpec   `json:"spec"`
	Status            ObjectAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectAclList contains a list of ObjectAcls
type ObjectAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectAcl `json:"items"`
}

// Repository type metadata.
var (
	ObjectAcl_Kind             = "ObjectAcl"
	ObjectAcl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectAcl_Kind}.String()
	ObjectAcl_KindAPIVersion   = ObjectAcl_Kind + "." + CRDGroupVersion.String()
	ObjectAcl_GroupVersionKind = CRDGroupVersion.WithKind(ObjectAcl_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectAcl{}, &ObjectAclList{})
}
