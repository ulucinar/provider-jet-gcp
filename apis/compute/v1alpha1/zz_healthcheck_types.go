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

type GrpcHealthCheckObservation struct {
}

type GrpcHealthCheckParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	// - Empty serviceName means the overall status of all services at the backend.
	// - Non-empty serviceName means the health of that gRPC service, as defined by the owner of the service.
	// The grpcServiceName can only be ASCII.
	// +kubebuilder:validation:Optional
	GrpcServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, gRPC health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type HTTPHealthCheckObservation struct {
}

type HTTPHealthCheckParameters struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, HTTP health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPSHealthCheckObservation struct {
}

type HTTPSHealthCheckParameters struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, HTTPS health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HealthCheckLogConfigObservation struct {
}

type HealthCheckLogConfigParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type HealthCheckObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckParameters struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *int64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	GrpcHealthCheck []GrpcHealthCheckParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	HTTPHealthCheck []HTTPHealthCheckParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	HTTPSHealthCheck []HTTPSHealthCheckParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	Http2HealthCheck []Http2HealthCheckParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// +kubebuilder:validation:Optional
	LogConfig []HealthCheckLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	SslHealthCheck []SslHealthCheckParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	TCPHealthCheck []TCPHealthCheckParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +kubebuilder:validation:Optional
	TimeoutSec *int64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type Http2HealthCheckObservation struct {
}

type Http2HealthCheckParameters struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, HTTP2 health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type SslHealthCheckObservation struct {
}

type SslHealthCheckParameters struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, SSL health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type TCPHealthCheckObservation struct {
}

type TCPHealthCheckParameters struct {

	// The TCP port number for the TCP health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, TCP health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

// HealthCheckSpec defines the desired state of HealthCheck
type HealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthCheckParameters `json:"forProvider"`
}

// HealthCheckStatus defines the observed state of HealthCheck.
type HealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheck is the Schema for the HealthChecks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthCheckSpec   `json:"spec"`
	Status            HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheckList contains a list of HealthChecks
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HealthCheck_Kind             = "HealthCheck"
	HealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthCheck_Kind}.String()
	HealthCheck_KindAPIVersion   = HealthCheck_Kind + "." + CRDGroupVersion.String()
	HealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
