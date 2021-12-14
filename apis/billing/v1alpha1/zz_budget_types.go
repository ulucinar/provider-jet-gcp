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

type AllUpdatesRuleObservation struct {
}

type AllUpdatesRuleParameters struct {

	// Boolean. When set to true, disables default notifications sent
	// when a threshold is exceeded. Default recipients are
	// those with Billing Account Administrators and Billing
	// Account Users IAM roles for the target account.
	// +kubebuilder:validation:Optional
	DisableDefaultIamRecipients *bool `json:"disableDefaultIamRecipients,omitempty" tf:"disable_default_iam_recipients,omitempty"`

	// The full resource name of a monitoring notification
	// channel in the form
	// projects/{project_id}/notificationChannels/{channel_id}.
	// A maximum of 5 channels are allowed.
	// +kubebuilder:validation:Optional
	MonitoringNotificationChannels []*string `json:"monitoringNotificationChannels,omitempty" tf:"monitoring_notification_channels,omitempty"`

	// The name of the Cloud Pub/Sub topic where budget related
	// messages will be published, in the form
	// projects/{project_id}/topics/{topic_id}. Updates are sent
	// at regular intervals to the topic.
	// +kubebuilder:validation:Optional
	PubsubTopic *string `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// The schema version of the notification. Only "1.0" is
	// accepted. It represents the JSON schema as defined in
	// https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
	// +kubebuilder:validation:Optional
	SchemaVersion *string `json:"schemaVersion,omitempty" tf:"schema_version,omitempty"`
}

type AmountObservation struct {
}

type AmountParameters struct {

	// Configures a budget amount that is automatically set to 100% of
	// last period's spend.
	// Boolean. Set value to true to use. Do not set to false, instead
	// use the 'specified_amount' block.
	// +kubebuilder:validation:Optional
	LastPeriodAmount *bool `json:"lastPeriodAmount,omitempty" tf:"last_period_amount,omitempty"`

	// A specified amount to use as the budget. currencyCode is
	// optional. If specified, it must match the currency of the
	// billing account. The currencyCode is provided on output.
	// +kubebuilder:validation:Optional
	SpecifiedAmount []SpecifiedAmountParameters `json:"specifiedAmount,omitempty" tf:"specified_amount,omitempty"`
}

type BudgetFilterObservation struct {
}

type BudgetFilterParameters struct {

	// A set of subaccounts of the form billingAccounts/{account_id},
	// specifying that usage from only this set of subaccounts should
	// be included in the budget. If a subaccount is set to the name of
	// the parent account, usage from the parent account will be included.
	// If the field is omitted, the report will include usage from the parent
	// account and all subaccounts, if they exist.
	// +kubebuilder:validation:Optional
	CreditTypes []*string `json:"creditTypes,omitempty" tf:"credit_types,omitempty"`

	// Specifies how credits should be treated when determining spend
	// for threshold calculations. Default value: "INCLUDE_ALL_CREDITS" Possible values: ["INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS"]
	// +kubebuilder:validation:Optional
	CreditTypesTreatment *string `json:"creditTypesTreatment,omitempty" tf:"credit_types_treatment,omitempty"`

	// A single label and value pair specifying that usage from only
	// this set of labeled resources should be included in the budget.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A set of projects of the form projects/{project_number},
	// specifying that usage from only this set of projects should be
	// included in the budget. If omitted, the report will include
	// all usage for the billing account, regardless of which project
	// the usage occurred on.
	// +kubebuilder:validation:Optional
	Projects []*string `json:"projects,omitempty" tf:"projects,omitempty"`

	// A set of services of the form services/{service_id},
	// specifying that usage from only this set of services should be
	// included in the budget. If omitted, the report will include
	// usage for all the services. The service names are available
	// through the Catalog API:
	// https://cloud.google.com/billing/v1/how-tos/catalog-api.
	// +kubebuilder:validation:Optional
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`

	// A set of subaccounts of the form billingAccounts/{account_id},
	// specifying that usage from only this set of subaccounts should
	// be included in the budget. If a subaccount is set to the name of
	// the parent account, usage from the parent account will be included.
	// If the field is omitted, the report will include usage from the parent
	// account and all subaccounts, if they exist.
	// +kubebuilder:validation:Optional
	Subaccounts []*string `json:"subaccounts,omitempty" tf:"subaccounts,omitempty"`
}

type BudgetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BudgetParameters struct {

	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// +kubebuilder:validation:Optional
	AllUpdatesRule []AllUpdatesRuleParameters `json:"allUpdatesRule,omitempty" tf:"all_updates_rule,omitempty"`

	// The budgeted amount for each usage period.
	// +kubebuilder:validation:Required
	Amount []AmountParameters `json:"amount" tf:"amount,omitempty"`

	// ID of the billing account to set a budget on.
	// +kubebuilder:validation:Required
	BillingAccount *string `json:"billingAccount" tf:"billing_account,omitempty"`

	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// +kubebuilder:validation:Optional
	BudgetFilter []BudgetFilterParameters `json:"budgetFilter,omitempty" tf:"budget_filter,omitempty"`

	// User data for display name in UI. Must be <= 60 chars.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// +kubebuilder:validation:Required
	ThresholdRules []ThresholdRulesParameters `json:"thresholdRules" tf:"threshold_rules,omitempty"`
}

type SpecifiedAmountObservation struct {
}

type SpecifiedAmountParameters struct {

	// The 3-letter currency code defined in ISO 4217.
	// +kubebuilder:validation:Optional
	CurrencyCode *string `json:"currencyCode,omitempty" tf:"currency_code,omitempty"`

	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999
	// inclusive. If units is positive, nanos must be positive or
	// zero. If units is zero, nanos can be positive, zero, or
	// negative. If units is negative, nanos must be negative or
	// zero. For example $-1.75 is represented as units=-1 and
	// nanos=-750,000,000.
	// +kubebuilder:validation:Optional
	Nanos *int64 `json:"nanos,omitempty" tf:"nanos,omitempty"`

	// The whole units of the amount. For example if currencyCode
	// is "USD", then 1 unit is one US dollar.
	// +kubebuilder:validation:Optional
	Units *string `json:"units,omitempty" tf:"units,omitempty"`
}

type ThresholdRulesObservation struct {
}

type ThresholdRulesParameters struct {

	// The type of basis used to determine if spend has passed
	// the threshold. Default value: "CURRENT_SPEND" Possible values: ["CURRENT_SPEND", "FORECASTED_SPEND"]
	// +kubebuilder:validation:Optional
	SpendBasis *string `json:"spendBasis,omitempty" tf:"spend_basis,omitempty"`

	// Send an alert when this threshold is exceeded. This is a
	// 1.0-based percentage, so 0.5 = 50%. Must be >= 0.
	// +kubebuilder:validation:Required
	ThresholdPercent *float64 `json:"thresholdPercent" tf:"threshold_percent,omitempty"`
}

// BudgetSpec defines the desired state of Budget
type BudgetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetParameters `json:"forProvider"`
}

// BudgetStatus defines the observed state of Budget.
type BudgetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Budget is the Schema for the Budgets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Budget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetSpec   `json:"spec"`
	Status            BudgetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetList contains a list of Budgets
type BudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Budget `json:"items"`
}

// Repository type metadata.
var (
	Budget_Kind             = "Budget"
	Budget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Budget_Kind}.String()
	Budget_KindAPIVersion   = Budget_Kind + "." + CRDGroupVersion.String()
	Budget_GroupVersionKind = CRDGroupVersion.WithKind(Budget_Kind)
)

func init() {
	SchemeBuilder.Register(&Budget{}, &BudgetList{})
}
