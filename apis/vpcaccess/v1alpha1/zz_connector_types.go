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

type ConnectorObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ConnectorParameters struct {

	// The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.
	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300.
	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	// +kubebuilder:validation:Optional
	MinThroughput *int64 `json:"minThroughput,omitempty" tf:"min_throughput,omitempty"`

	// The name of the resource (Max 25 characters).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Name of the VPC network. Required if 'ip_cidr_range' is set.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ConnectorSpec defines the desired state of Connector
type ConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectorParameters `json:"forProvider"`
}

// ConnectorStatus defines the observed state of Connector.
type ConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connector is the Schema for the Connectors API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Connector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectorSpec   `json:"spec"`
	Status            ConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorList contains a list of Connectors
type ConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connector `json:"items"`
}

// Repository type metadata.
var (
	Connector_Kind             = "Connector"
	Connector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connector_Kind}.String()
	Connector_KindAPIVersion   = Connector_Kind + "." + CRDGroupVersion.String()
	Connector_GroupVersionKind = CRDGroupVersion.WithKind(Connector_Kind)
)

func init() {
	SchemeBuilder.Register(&Connector{}, &ConnectorList{})
}
