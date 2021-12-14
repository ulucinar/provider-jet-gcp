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

type AutoscalingPolicyCPUUtilizationObservation struct {
}

type AutoscalingPolicyCPUUtilizationParameters struct {

	// Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are:
	//
	// - NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics.
	//
	// - OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand.
	// +kubebuilder:validation:Optional
	PredictiveMethod *string `json:"predictiveMethod,omitempty" tf:"predictive_method,omitempty"`

	// The target CPU utilization that the autoscaler should maintain.
	// Must be a float value in the range (0, 1]. If not specified, the
	// default is 0.6.
	//
	// If the CPU level is below the target utilization, the autoscaler
	// scales down the number of instances until it reaches the minimum
	// number of instances you specified or until the average CPU of
	// your instances reaches the target utilization.
	//
	// If the average CPU is above the target utilization, the autoscaler
	// scales up until it reaches the maximum number of instances you
	// specified or until the average utilization reaches the target
	// utilization.
	// +kubebuilder:validation:Required
	Target *float64 `json:"target" tf:"target,omitempty"`
}

type AutoscalingPolicyLoadBalancingUtilizationObservation struct {
}

type AutoscalingPolicyLoadBalancingUtilizationParameters struct {

	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	// +kubebuilder:validation:Required
	Target *float64 `json:"target" tf:"target,omitempty"`
}

type AutoscalingPolicyMetricObservation struct {
}

type AutoscalingPolicyMetricParameters struct {

	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	//
	// The metric must have a value type of INT64 or DOUBLE.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The target value of the metric that autoscaler should
	// maintain. This must be a positive value. A utilization
	// metric scales number of virtual machines handling requests
	// to increase or decrease proportionally to the metric.
	//
	// For example, a good metric to use as a utilizationTarget is
	// www.googleapis.com/compute/instance/network/received_bytes_count.
	// The autoscaler will work to keep this value constant for each
	// of the instances.
	// +kubebuilder:validation:Optional
	Target *float64 `json:"target,omitempty" tf:"target,omitempty"`

	// Defines how target utilization value is expressed for a
	// Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AutoscalingPolicyScaleInControlObservation struct {
}

type AutoscalingPolicyScaleInControlParameters struct {

	// A nested object resource
	// +kubebuilder:validation:Optional
	MaxScaledInReplicas []ScaleInControlMaxScaledInReplicasParameters `json:"maxScaledInReplicas,omitempty" tf:"max_scaled_in_replicas,omitempty"`

	// How long back autoscaling should look when computing recommendations
	// to include directives regarding slower scale down, as described above.
	// +kubebuilder:validation:Optional
	TimeWindowSec *int64 `json:"timeWindowSec,omitempty" tf:"time_window_sec,omitempty"`
}

type AutoscalingPolicyScalingSchedulesObservation struct {
}

type AutoscalingPolicyScalingSchedulesParameters struct {

	// A description of a scaling schedule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A boolean value that specifies if a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The duration of time intervals (in seconds) for which this scaling schedule will be running. The minimum allowed value is 300.
	// +kubebuilder:validation:Required
	DurationSec *int64 `json:"durationSec" tf:"duration_sec,omitempty"`

	// Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule.
	// +kubebuilder:validation:Required
	MinRequiredReplicas *int64 `json:"minRequiredReplicas" tf:"min_required_replicas,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The start timestamps of time intervals when this scaling schedule should provide a scaling signal. This field uses the extended cron format (with an optional year field).
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`

	// The time zone to be used when interpreting the schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type RegionAutoscalerAutoscalingPolicyObservation struct {
}

type RegionAutoscalerAutoscalingPolicyParameters struct {

	// Defines the CPU utilization policy that allows the autoscaler to
	// scale based on the average CPU utilization of a managed instance
	// group.
	// +kubebuilder:validation:Optional
	CPUUtilization []AutoscalingPolicyCPUUtilizationParameters `json:"cpuUtilization,omitempty" tf:"cpu_utilization,omitempty"`

	// The number of seconds that the autoscaler should wait before it
	// starts collecting information from a new instance. This prevents
	// the autoscaler from collecting information when the instance is
	// initializing, during which the collected usage would not be
	// reliable. The default time autoscaler waits is 60 seconds.
	//
	// Virtual machine initialization times might vary because of
	// numerous factors. We recommend that you test how long an
	// instance may take to initialize. To do this, create an instance
	// and time the startup process.
	// +kubebuilder:validation:Optional
	CooldownPeriod *int64 `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`

	// Configuration parameters of autoscaling based on a load balancer.
	// +kubebuilder:validation:Optional
	LoadBalancingUtilization []AutoscalingPolicyLoadBalancingUtilizationParameters `json:"loadBalancingUtilization,omitempty" tf:"load_balancing_utilization,omitempty"`

	// The maximum number of instances that the autoscaler can scale up
	// to. This is required when creating or updating an autoscaler. The
	// maximum number of replicas should not be lower than minimal number
	// of replicas.
	// +kubebuilder:validation:Required
	MaxReplicas *int64 `json:"maxReplicas" tf:"max_replicas,omitempty"`

	// Configuration parameters of autoscaling based on a custom metric.
	// +kubebuilder:validation:Optional
	Metric []AutoscalingPolicyMetricParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// The minimum number of replicas that the autoscaler can scale down
	// to. This cannot be less than 0. If not provided, autoscaler will
	// choose a default value depending on maximum number of instances
	// allowed.
	// +kubebuilder:validation:Required
	MinReplicas *int64 `json:"minReplicas" tf:"min_replicas,omitempty"`

	// Defines operating mode for this policy. Default value: "ON" Possible values: ["OFF", "ONLY_UP", "ON"]
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Defines scale in controls to reduce the risk of response latency
	// and outages due to abrupt scale-in events
	// +kubebuilder:validation:Optional
	ScaleInControl []AutoscalingPolicyScaleInControlParameters `json:"scaleInControl,omitempty" tf:"scale_in_control,omitempty"`

	// Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap.
	// +kubebuilder:validation:Optional
	ScalingSchedules []AutoscalingPolicyScalingSchedulesParameters `json:"scalingSchedules,omitempty" tf:"scaling_schedules,omitempty"`
}

type RegionAutoscalerObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RegionAutoscalerParameters struct {

	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// +kubebuilder:validation:Required
	AutoscalingPolicy []RegionAutoscalerAutoscalingPolicyParameters `json:"autoscalingPolicy" tf:"autoscaling_policy,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource. The name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URL of the region where the instance group resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// URL of the managed instance group that this autoscaler will scale.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`
}

type ScaleInControlMaxScaledInReplicasObservation struct {
}

type ScaleInControlMaxScaledInReplicasParameters struct {

	// Specifies a fixed number of VM instances. This must be a positive
	// integer.
	// +kubebuilder:validation:Optional
	Fixed *int64 `json:"fixed,omitempty" tf:"fixed,omitempty"`

	// Specifies a percentage of instances between 0 to 100%, inclusive.
	// For example, specify 80 for 80%.
	// +kubebuilder:validation:Optional
	Percent *int64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

// RegionAutoscalerSpec defines the desired state of RegionAutoscaler
type RegionAutoscalerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionAutoscalerParameters `json:"forProvider"`
}

// RegionAutoscalerStatus defines the observed state of RegionAutoscaler.
type RegionAutoscalerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionAutoscalerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionAutoscaler is the Schema for the RegionAutoscalers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RegionAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionAutoscalerSpec   `json:"spec"`
	Status            RegionAutoscalerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionAutoscalerList contains a list of RegionAutoscalers
type RegionAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionAutoscaler `json:"items"`
}

// Repository type metadata.
var (
	RegionAutoscaler_Kind             = "RegionAutoscaler"
	RegionAutoscaler_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionAutoscaler_Kind}.String()
	RegionAutoscaler_KindAPIVersion   = RegionAutoscaler_Kind + "." + CRDGroupVersion.String()
	RegionAutoscaler_GroupVersionKind = CRDGroupVersion.WithKind(RegionAutoscaler_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionAutoscaler{}, &RegionAutoscalerList{})
}
