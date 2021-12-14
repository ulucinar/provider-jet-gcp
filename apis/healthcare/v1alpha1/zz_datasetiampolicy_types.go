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

type DatasetIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatasetIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`
}

// DatasetIamPolicySpec defines the desired state of DatasetIamPolicy
type DatasetIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetIamPolicyParameters `json:"forProvider"`
}

// DatasetIamPolicyStatus defines the observed state of DatasetIamPolicy.
type DatasetIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIamPolicy is the Schema for the DatasetIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DatasetIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetIamPolicySpec   `json:"spec"`
	Status            DatasetIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIamPolicyList contains a list of DatasetIamPolicys
type DatasetIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasetIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	DatasetIamPolicy_Kind             = "DatasetIamPolicy"
	DatasetIamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatasetIamPolicy_Kind}.String()
	DatasetIamPolicy_KindAPIVersion   = DatasetIamPolicy_Kind + "." + CRDGroupVersion.String()
	DatasetIamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DatasetIamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DatasetIamPolicy{}, &DatasetIamPolicyList{})
}
