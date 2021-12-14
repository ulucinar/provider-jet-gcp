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

type ClusterObservation struct {
}

type ClusterParameters struct {

	// The ID of the Cloud Bigtable cluster.
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster. 3) All clusters within an instance must use the same CMEK key. Values are of the form projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}
	// +kubebuilder:validation:Optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The number of nodes in your Cloud Bigtable cluster. Required, with a minimum of 1 for a PRODUCTION instance. Must be left unset for a DEVELOPMENT instance.
	// +kubebuilder:validation:Optional
	NumNodes *int64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The storage type to use. One of "SSD" or "HDD". Defaults to "SSD".
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// The zone to create the Cloud Bigtable cluster in. Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the Cloud Bigtable locations page.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InstanceParameters struct {

	// A block of cluster configuration options. This can be specified at least once.
	// +kubebuilder:validation:Optional
	Cluster []ClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The human-readable display name of the Bigtable instance. Defaults to the instance name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// A mapping of labels to assign to the resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
