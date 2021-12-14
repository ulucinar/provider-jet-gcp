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

type AdditionalExtensionsObjectIDObservation struct {
}

type AdditionalExtensionsObjectIDParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type AuthorityKeyIDObservation struct {
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type AuthorityKeyIDParameters struct {
}

type CertFingerprintObservation struct {
	Sha256Hash *string `json:"sha256Hash,omitempty" tf:"sha256_hash,omitempty"`
}

type CertFingerprintParameters struct {
}

type CertificateDescriptionObservation struct {
	AiaIssuingCertificateUrls []*string `json:"aiaIssuingCertificateUrls,omitempty" tf:"aia_issuing_certificate_urls,omitempty"`

	AuthorityKeyID []AuthorityKeyIDObservation `json:"authorityKeyId,omitempty" tf:"authority_key_id,omitempty"`

	CertFingerprint []CertFingerprintObservation `json:"certFingerprint,omitempty" tf:"cert_fingerprint,omitempty"`

	ConfigValues []ConfigValuesObservation `json:"configValues,omitempty" tf:"config_values,omitempty"`

	CrlDistributionPoints []*string `json:"crlDistributionPoints,omitempty" tf:"crl_distribution_points,omitempty"`

	PublicKey []PublicKeyObservation `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	SubjectDescription []SubjectDescriptionObservation `json:"subjectDescription,omitempty" tf:"subject_description,omitempty"`

	SubjectKeyID []SubjectKeyIDObservation `json:"subjectKeyId,omitempty" tf:"subject_key_id,omitempty"`
}

type CertificateDescriptionParameters struct {
}

type CertificateObservation struct {
	CertificateDescription []CertificateDescriptionObservation `json:"certificateDescription,omitempty" tf:"certificate_description,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PemCertificate *string `json:"pemCertificate,omitempty" tf:"pem_certificate,omitempty"`

	PemCertificates []*string `json:"pemCertificates,omitempty" tf:"pem_certificates,omitempty"`

	RevocationDetails []RevocationDetailsObservation `json:"revocationDetails,omitempty" tf:"revocation_details,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateParameters struct {

	// Certificate Authority name.
	// +kubebuilder:validation:Optional
	CertificateAuthority *string `json:"certificateAuthority,omitempty" tf:"certificate_authority,omitempty"`

	// The resource name for a CertificateTemplate used to issue this certificate,
	// in the format 'projects/*/locations/*/certificateTemplates/*'. If this is specified,
	// the caller must have the necessary permission to use this template. If this is
	// omitted, no template will be used. This template must be in the same location
	// as the Certificate.
	// +kubebuilder:validation:Optional
	CertificateTemplate *string `json:"certificateTemplate,omitempty" tf:"certificate_template,omitempty"`

	// The config used to create a self-signed X.509 certificate or CSR.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Labels with user-defined metadata to apply to this resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	Lifetime *string `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	// Location of the Certificate. A full list of valid locations can be found by
	// running 'gcloud privateca locations list'.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name for this Certificate.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	// +kubebuilder:validation:Optional
	PemCsr *string `json:"pemCsr,omitempty" tf:"pem_csr,omitempty"`

	// The name of the CaPool this Certificate belongs to.
	// +kubebuilder:validation:Required
	Pool *string `json:"pool" tf:"pool,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ConfigObservation struct {
}

type ConfigParameters struct {

	// A PublicKey describes a public key.
	// +kubebuilder:validation:Required
	PublicKey []ConfigPublicKeyParameters `json:"publicKey" tf:"public_key,omitempty"`

	// Specifies some of the values in a certificate that are related to the subject.
	// +kubebuilder:validation:Required
	SubjectConfig []SubjectConfigParameters `json:"subjectConfig" tf:"subject_config,omitempty"`

	// Describes how some of the technical X.509 fields in a certificate should be populated.
	// +kubebuilder:validation:Required
	X509Config []X509ConfigParameters `json:"x509Config" tf:"x509_config,omitempty"`
}

type ConfigPublicKeyObservation struct {
}

type ConfigPublicKeyParameters struct {

	// The format of the public key. Currently, only PEM format is supported. Possible values: ["KEY_TYPE_UNSPECIFIED", "PEM"]
	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// Required. A public key. When this is specified in a request, the padding and encoding can be any of the options described by the respective 'KeyType' value. When this is generated by the service, it will always be an RFC 5280 SubjectPublicKeyInfo structure containing an algorithm identifier and a key. A base64-encoded string.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ConfigValuesKeyUsageObservation struct {
	BaseKeyUsage []KeyUsageBaseKeyUsageObservation `json:"baseKeyUsage,omitempty" tf:"base_key_usage,omitempty"`

	ExtendedKeyUsage []KeyUsageExtendedKeyUsageObservation `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	UnknownExtendedKeyUsages []KeyUsageUnknownExtendedKeyUsagesObservation `json:"unknownExtendedKeyUsages,omitempty" tf:"unknown_extended_key_usages,omitempty"`
}

type ConfigValuesKeyUsageParameters struct {
}

type ConfigValuesObservation struct {
	KeyUsage []ConfigValuesKeyUsageObservation `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`
}

type ConfigValuesParameters struct {
}

type CustomSansObectIDObservation struct {
	ObjectIDPath []*int64 `json:"objectIdPath,omitempty" tf:"object_id_path,omitempty"`
}

type CustomSansObectIDParameters struct {
}

type CustomSansObservation struct {
	Critical *bool `json:"critical,omitempty" tf:"critical,omitempty"`

	ObectID []CustomSansObectIDObservation `json:"obectId,omitempty" tf:"obect_id,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomSansParameters struct {
}

type KeyUsageBaseKeyUsageObservation struct {
	KeyUsageOptions []KeyUsageOptionsObservation `json:"keyUsageOptions,omitempty" tf:"key_usage_options,omitempty"`
}

type KeyUsageBaseKeyUsageParameters struct {
}

type KeyUsageExtendedKeyUsageObservation struct {
	ClientAuth *bool `json:"clientAuth,omitempty" tf:"client_auth,omitempty"`

	CodeSigning *bool `json:"codeSigning,omitempty" tf:"code_signing,omitempty"`

	EmailProtection *bool `json:"emailProtection,omitempty" tf:"email_protection,omitempty"`

	OcspSigning *bool `json:"ocspSigning,omitempty" tf:"ocsp_signing,omitempty"`

	ServerAuth *bool `json:"serverAuth,omitempty" tf:"server_auth,omitempty"`

	TimeStamping *bool `json:"timeStamping,omitempty" tf:"time_stamping,omitempty"`
}

type KeyUsageExtendedKeyUsageParameters struct {
}

type KeyUsageOptionsObservation struct {
	CertSign *bool `json:"certSign,omitempty" tf:"cert_sign,omitempty"`

	ContentCommitment *bool `json:"contentCommitment,omitempty" tf:"content_commitment,omitempty"`

	CrlSign *bool `json:"crlSign,omitempty" tf:"crl_sign,omitempty"`

	DataEncipherment *bool `json:"dataEncipherment,omitempty" tf:"data_encipherment,omitempty"`

	DecipherOnly *bool `json:"decipherOnly,omitempty" tf:"decipher_only,omitempty"`

	DigitalSignature *bool `json:"digitalSignature,omitempty" tf:"digital_signature,omitempty"`

	EncipherOnly *bool `json:"encipherOnly,omitempty" tf:"encipher_only,omitempty"`

	KeyAgreement *bool `json:"keyAgreement,omitempty" tf:"key_agreement,omitempty"`

	KeyEncipherment *bool `json:"keyEncipherment,omitempty" tf:"key_encipherment,omitempty"`
}

type KeyUsageOptionsParameters struct {
}

type KeyUsageUnknownExtendedKeyUsagesObservation struct {
	ObectID []ObectIDObservation `json:"obectId,omitempty" tf:"obect_id,omitempty"`
}

type KeyUsageUnknownExtendedKeyUsagesParameters struct {
}

type ObectIDObservation struct {
	ObjectIDPath []*int64 `json:"objectIdPath,omitempty" tf:"object_id_path,omitempty"`
}

type ObectIDParameters struct {
}

type PublicKeyObservation struct {
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type PublicKeyParameters struct {
}

type RevocationDetailsObservation struct {
	RevocationState *string `json:"revocationState,omitempty" tf:"revocation_state,omitempty"`

	RevocationTime *string `json:"revocationTime,omitempty" tf:"revocation_time,omitempty"`
}

type RevocationDetailsParameters struct {
}

type SubjectAltNameObservation struct {
	CustomSans []CustomSansObservation `json:"customSans,omitempty" tf:"custom_sans,omitempty"`

	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type SubjectAltNameParameters struct {
}

type SubjectConfigObservation struct {
}

type SubjectConfigParameters struct {

	// Contains distinguished name fields such as the location and organization.
	// +kubebuilder:validation:Required
	Subject []SubjectConfigSubjectParameters `json:"subject" tf:"subject,omitempty"`

	// The subject alternative name fields.
	// +kubebuilder:validation:Optional
	SubjectAltName []SubjectConfigSubjectAltNameParameters `json:"subjectAltName,omitempty" tf:"subject_alt_name,omitempty"`
}

type SubjectConfigSubjectAltNameObservation struct {
}

type SubjectConfigSubjectAltNameParameters struct {

	// Contains only valid, fully-qualified host names.
	// +kubebuilder:validation:Optional
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// Contains only valid RFC 2822 E-mail addresses.
	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// Contains only valid RFC 3986 URIs.
	// +kubebuilder:validation:Optional
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type SubjectConfigSubjectObservation struct {
}

type SubjectConfigSubjectParameters struct {

	// The common name of the distinguished name.
	// +kubebuilder:validation:Required
	CommonName *string `json:"commonName" tf:"common_name,omitempty"`

	// The country code of the subject.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The locality or city of the subject.
	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The organization of the subject.
	// +kubebuilder:validation:Required
	Organization *string `json:"organization" tf:"organization,omitempty"`

	// The organizational unit of the subject.
	// +kubebuilder:validation:Optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// The postal code of the subject.
	// +kubebuilder:validation:Optional
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The province, territory, or regional state of the subject.
	// +kubebuilder:validation:Optional
	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	// The street address of the subject.
	// +kubebuilder:validation:Optional
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`
}

type SubjectDescriptionObservation struct {
	HexSerialNumber *string `json:"hexSerialNumber,omitempty" tf:"hex_serial_number,omitempty"`

	Lifetime *string `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	NotAfterTime *string `json:"notAfterTime,omitempty" tf:"not_after_time,omitempty"`

	NotBeforeTime *string `json:"notBeforeTime,omitempty" tf:"not_before_time,omitempty"`

	Subject []SubjectObservation `json:"subject,omitempty" tf:"subject,omitempty"`

	SubjectAltName []SubjectAltNameObservation `json:"subjectAltName,omitempty" tf:"subject_alt_name,omitempty"`
}

type SubjectDescriptionParameters struct {
}

type SubjectKeyIDObservation struct {
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SubjectKeyIDParameters struct {
}

type SubjectObservation struct {
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`
}

type SubjectParameters struct {
}

type X509ConfigAdditionalExtensionsObservation struct {
}

type X509ConfigAdditionalExtensionsParameters struct {

	// Indicates whether or not this extension is critical (i.e., if the client does not know how to
	// handle this extension, the client should consider this to be an error).
	// +kubebuilder:validation:Required
	Critical *bool `json:"critical" tf:"critical,omitempty"`

	// Describes values that are relevant in a CA certificate.
	// +kubebuilder:validation:Required
	ObjectID []AdditionalExtensionsObjectIDParameters `json:"objectId" tf:"object_id,omitempty"`

	// The value of this X.509 extension. A base64-encoded string.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type X509ConfigCaOptionsObservation struct {
}

type X509ConfigCaOptionsParameters struct {

	// Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing,
	// the extension will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	IsCa *bool `json:"isCa,omitempty" tf:"is_ca,omitempty"`

	// Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this
	// value is missing, the max path length will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty" tf:"max_issuer_path_length,omitempty"`
}

type X509ConfigKeyUsageBaseKeyUsageObservation struct {
}

type X509ConfigKeyUsageBaseKeyUsageParameters struct {

	// The key may be used to sign certificates.
	// +kubebuilder:validation:Optional
	CertSign *bool `json:"certSign,omitempty" tf:"cert_sign,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	// +kubebuilder:validation:Optional
	ContentCommitment *bool `json:"contentCommitment,omitempty" tf:"content_commitment,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kubebuilder:validation:Optional
	CrlSign *bool `json:"crlSign,omitempty" tf:"crl_sign,omitempty"`

	// The key may be used to encipher data.
	// +kubebuilder:validation:Optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty" tf:"data_encipherment,omitempty"`

	// The key may be used to decipher only.
	// +kubebuilder:validation:Optional
	DecipherOnly *bool `json:"decipherOnly,omitempty" tf:"decipher_only,omitempty"`

	// The key may be used for digital signatures.
	// +kubebuilder:validation:Optional
	DigitalSignature *bool `json:"digitalSignature,omitempty" tf:"digital_signature,omitempty"`

	// The key may be used to encipher only.
	// +kubebuilder:validation:Optional
	EncipherOnly *bool `json:"encipherOnly,omitempty" tf:"encipher_only,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kubebuilder:validation:Optional
	KeyAgreement *bool `json:"keyAgreement,omitempty" tf:"key_agreement,omitempty"`

	// The key may be used to encipher other keys.
	// +kubebuilder:validation:Optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty" tf:"key_encipherment,omitempty"`
}

type X509ConfigKeyUsageExtendedKeyUsageObservation struct {
}

type X509ConfigKeyUsageExtendedKeyUsageParameters struct {

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ClientAuth *bool `json:"clientAuth,omitempty" tf:"client_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	// +kubebuilder:validation:Optional
	CodeSigning *bool `json:"codeSigning,omitempty" tf:"code_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	// +kubebuilder:validation:Optional
	EmailProtection *bool `json:"emailProtection,omitempty" tf:"email_protection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	// +kubebuilder:validation:Optional
	OcspSigning *bool `json:"ocspSigning,omitempty" tf:"ocsp_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ServerAuth *bool `json:"serverAuth,omitempty" tf:"server_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	// +kubebuilder:validation:Optional
	TimeStamping *bool `json:"timeStamping,omitempty" tf:"time_stamping,omitempty"`
}

type X509ConfigKeyUsageObservation struct {
}

type X509ConfigKeyUsageParameters struct {

	// Describes high-level ways in which a key may be used.
	// +kubebuilder:validation:Required
	BaseKeyUsage []X509ConfigKeyUsageBaseKeyUsageParameters `json:"baseKeyUsage" tf:"base_key_usage,omitempty"`

	// Describes high-level ways in which a key may be used.
	// +kubebuilder:validation:Required
	ExtendedKeyUsage []X509ConfigKeyUsageExtendedKeyUsageParameters `json:"extendedKeyUsage" tf:"extended_key_usage,omitempty"`

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Optional
	UnknownExtendedKeyUsages []X509ConfigKeyUsageUnknownExtendedKeyUsagesParameters `json:"unknownExtendedKeyUsages,omitempty" tf:"unknown_extended_key_usages,omitempty"`
}

type X509ConfigKeyUsageUnknownExtendedKeyUsagesObservation struct {
}

type X509ConfigKeyUsageUnknownExtendedKeyUsagesParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type X509ConfigObservation struct {
}

type X509ConfigParameters struct {

	// Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.
	// +kubebuilder:validation:Optional
	AdditionalExtensions []X509ConfigAdditionalExtensionsParameters `json:"additionalExtensions,omitempty" tf:"additional_extensions,omitempty"`

	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the
	// "Authority Information Access" extension in the certificate.
	// +kubebuilder:validation:Optional
	AiaOcspServers []*string `json:"aiaOcspServers,omitempty" tf:"aia_ocsp_servers,omitempty"`

	// Describes values that are relevant in a CA certificate.
	// +kubebuilder:validation:Optional
	CaOptions []X509ConfigCaOptionsParameters `json:"caOptions,omitempty" tf:"ca_options,omitempty"`

	// Indicates the intended use for keys that correspond to a certificate.
	// +kubebuilder:validation:Required
	KeyUsage []X509ConfigKeyUsageParameters `json:"keyUsage" tf:"key_usage,omitempty"`

	// Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kubebuilder:validation:Optional
	PolicyIds []X509ConfigPolicyIdsParameters `json:"policyIds,omitempty" tf:"policy_ids,omitempty"`
}

type X509ConfigPolicyIdsObservation struct {
}

type X509ConfigPolicyIdsParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
