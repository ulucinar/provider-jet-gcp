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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NodePoolAutoscalingObservation struct {
}

type NodePoolAutoscalingParameters struct {

	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	// +kubebuilder:validation:Required
	MaxNodeCount *float64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// Minimum number of nodes in the NodePool. Must be >=0 and <= max_node_count.
	// +kubebuilder:validation:Required
	MinNodeCount *float64 `json:"minNodeCount" tf:"min_node_count,omitempty"`
}

type NodePoolManagementObservation struct {
}

type NodePoolManagementParameters struct {

	// Whether the nodes will be automatically repaired.
	// +kubebuilder:validation:Optional
	AutoRepair *bool `json:"autoRepair,omitempty" tf:"auto_repair,omitempty"`

	// Whether the nodes will be automatically upgraded.
	// +kubebuilder:validation:Optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`
}

type NodePoolNodeConfigGuestAcceleratorObservation struct {
}

type NodePoolNodeConfigGuestAcceleratorParameters struct {

	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count"`

	// +kubebuilder:validation:Optional
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty" tf:"gpu_partition_size"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type NodePoolNodeConfigObservation_2 struct {
}

type NodePoolNodeConfigParameters_2 struct {

	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Type of the disk attached to each node.
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	// +kubebuilder:validation:Optional
	GuestAccelerator []NodePoolNodeConfigGuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// The image type to use for this node. Note that for a given image type, the latest version of it will be used.
	// +kubebuilder:validation:Optional
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The number of local SSD disks to be attached to the node.
	// +kubebuilder:validation:Optional
	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`

	// The name of a Google Compute Engine machine type.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The metadata key/value pairs assigned to instances in the cluster.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// The set of Google API scopes to be made available on all of the node VMs.
	// +kubebuilder:validation:Optional
	OAuthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`

	// Whether the nodes are created as preemptible VM instances.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs.
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Shielded Instance options.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []NodePoolNodeConfigShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// The list of instance tags applied to all nodes.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// List of Kubernetes taints to be applied to each node.
	// +kubebuilder:validation:Optional
	Taint []NodePoolNodeConfigTaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// The workload metadata configuration for this node.
	// +kubebuilder:validation:Optional
	WorkloadMetadataConfig []NodePoolNodeConfigWorkloadMetadataConfigParameters `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type NodePoolNodeConfigShieldedInstanceConfigObservation struct {
}

type NodePoolNodeConfigShieldedInstanceConfigParameters struct {

	// Defines whether the instance has integrity monitoring enabled.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// Defines whether the instance has Secure Boot enabled.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`
}

type NodePoolNodeConfigTaintObservation struct {
}

type NodePoolNodeConfigTaintParameters struct {

	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type NodePoolNodeConfigWorkloadMetadataConfigObservation struct {
}

type NodePoolNodeConfigWorkloadMetadataConfigParameters struct {

	// Mode is the configuration for how to expose metadata to workloads running on the node.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type NodePoolObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceGroupUrls []*string `json:"instanceGroupUrls,omitempty" tf:"instance_group_urls,omitempty"`

	ManagedInstanceGroupUrls []*string `json:"managedInstanceGroupUrls,omitempty" tf:"managed_instance_group_urls,omitempty"`

	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`
}

type NodePoolParameters_2 struct {

	// Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage.
	// +kubebuilder:validation:Optional
	Autoscaling []NodePoolAutoscalingParameters `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// The cluster to create the node pool for. Cluster must be present in location provided for zonal clusters.
	// +crossplane:generate:reference:type=Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	// +kubebuilder:validation:Optional
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// The location (region or zone) of the cluster.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Node management configuration, wherein auto-repair and auto-upgrade is configured.
	// +kubebuilder:validation:Optional
	Management []NodePoolManagementParameters `json:"management,omitempty" tf:"management,omitempty"`

	// The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// The configuration of the nodepool
	// +kubebuilder:validation:Optional
	NodeConfig []NodePoolNodeConfigParameters_2 `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	// +kubebuilder:validation:Optional
	NodeLocations []*string `json:"nodeLocations,omitempty" tf:"node_locations,omitempty"`

	// The ID of the project in which to create the node pool. If blank, the provider-configured project will be used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.
	// +kubebuilder:validation:Optional
	UpgradeSettings []NodePoolUpgradeSettingsParameters `json:"upgradeSettings,omitempty" tf:"upgrade_settings,omitempty"`

	// The Kubernetes version for the nodes in this pool. Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodePoolUpgradeSettingsObservation struct {
}

type NodePoolUpgradeSettingsParameters struct {

	// The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.
	// +kubebuilder:validation:Required
	MaxSurge *float64 `json:"maxSurge" tf:"max_surge,omitempty"`

	// The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.
	// +kubebuilder:validation:Required
	MaxUnavailable *float64 `json:"maxUnavailable" tf:"max_unavailable,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters_2 `json:"forProvider"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePool is the Schema for the NodePools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodePoolSpec   `json:"spec"`
	Status            NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}
