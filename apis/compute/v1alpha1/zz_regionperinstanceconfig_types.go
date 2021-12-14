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

type RegionPerInstanceConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegionPerInstanceConfigParameters struct {

	// +kubebuilder:validation:Optional
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action,omitempty"`

	// +kubebuilder:validation:Optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The preserved state for this instance.
	// +kubebuilder:validation:Optional
	PreservedState []RegionPerInstanceConfigPreservedStateParameters `json:"preservedState,omitempty" tf:"preserved_state,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the containing instance group manager is located
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The region instance group manager this instance config is part of.
	// +kubebuilder:validation:Required
	RegionInstanceGroupManager *string `json:"regionInstanceGroupManager" tf:"region_instance_group_manager,omitempty"`

	// +kubebuilder:validation:Optional
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy,omitempty"`
}

type RegionPerInstanceConfigPreservedStateDiskObservation struct {
}

type RegionPerInstanceConfigPreservedStateDiskParameters struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
	// 'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
	// 'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
	// deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	// +kubebuilder:validation:Optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// The mode of the disk. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URI of an existing persistent disk to attach under the specified device-name in the format
	// 'projects/project-id/zones/zone/disks/disk-name'.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type RegionPerInstanceConfigPreservedStateObservation struct {
}

type RegionPerInstanceConfigPreservedStateParameters struct {

	// Stateful disks for the instance.
	// +kubebuilder:validation:Optional
	Disk []RegionPerInstanceConfigPreservedStateDiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

// RegionPerInstanceConfigSpec defines the desired state of RegionPerInstanceConfig
type RegionPerInstanceConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionPerInstanceConfigParameters `json:"forProvider"`
}

// RegionPerInstanceConfigStatus defines the observed state of RegionPerInstanceConfig.
type RegionPerInstanceConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionPerInstanceConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionPerInstanceConfig is the Schema for the RegionPerInstanceConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RegionPerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionPerInstanceConfigSpec   `json:"spec"`
	Status            RegionPerInstanceConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionPerInstanceConfigList contains a list of RegionPerInstanceConfigs
type RegionPerInstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionPerInstanceConfig `json:"items"`
}

// Repository type metadata.
var (
	RegionPerInstanceConfig_Kind             = "RegionPerInstanceConfig"
	RegionPerInstanceConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionPerInstanceConfig_Kind}.String()
	RegionPerInstanceConfig_KindAPIVersion   = RegionPerInstanceConfig_Kind + "." + CRDGroupVersion.String()
	RegionPerInstanceConfig_GroupVersionKind = CRDGroupVersion.WithKind(RegionPerInstanceConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionPerInstanceConfig{}, &RegionPerInstanceConfigList{})
}
