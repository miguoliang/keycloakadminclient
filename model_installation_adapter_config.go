/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keycloakadminclient

import (
	"encoding/json"
)

// checks if the InstallationAdapterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationAdapterConfig{}

// InstallationAdapterConfig struct for InstallationAdapterConfig
type InstallationAdapterConfig struct {
	Realm *string `json:"realm,omitempty"`
	RealmPublicKey *string `json:"realm-public-key,omitempty"`
	AuthServerUrl *string `json:"auth-server-url,omitempty"`
	SslRequired *string `json:"ssl-required,omitempty"`
	BearerOnly *bool `json:"bearer-only,omitempty"`
	Resource *string `json:"resource,omitempty"`
	PublicClient *bool `json:"public-client,omitempty"`
	VerifyTokenAudience *bool `json:"verify-token-audience,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	UseResourceRoleMappings *bool `json:"use-resource-role-mappings,omitempty"`
	ConfidentialPort *int32 `json:"confidential-port,omitempty"`
	PolicyEnforcer *PolicyEnforcerConfig `json:"policy-enforcer,omitempty"`
}

// NewInstallationAdapterConfig instantiates a new InstallationAdapterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationAdapterConfig() *InstallationAdapterConfig {
	this := InstallationAdapterConfig{}
	return &this
}

// NewInstallationAdapterConfigWithDefaults instantiates a new InstallationAdapterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationAdapterConfigWithDefaults() *InstallationAdapterConfig {
	this := InstallationAdapterConfig{}
	return &this
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *InstallationAdapterConfig) SetRealm(v string) {
	o.Realm = &v
}

// GetRealmPublicKey returns the RealmPublicKey field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetRealmPublicKey() string {
	if o == nil || IsNil(o.RealmPublicKey) {
		var ret string
		return ret
	}
	return *o.RealmPublicKey
}

// GetRealmPublicKeyOk returns a tuple with the RealmPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetRealmPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RealmPublicKey) {
		return nil, false
	}
	return o.RealmPublicKey, true
}

// HasRealmPublicKey returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasRealmPublicKey() bool {
	if o != nil && !IsNil(o.RealmPublicKey) {
		return true
	}

	return false
}

// SetRealmPublicKey gets a reference to the given string and assigns it to the RealmPublicKey field.
func (o *InstallationAdapterConfig) SetRealmPublicKey(v string) {
	o.RealmPublicKey = &v
}

// GetAuthServerUrl returns the AuthServerUrl field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetAuthServerUrl() string {
	if o == nil || IsNil(o.AuthServerUrl) {
		var ret string
		return ret
	}
	return *o.AuthServerUrl
}

// GetAuthServerUrlOk returns a tuple with the AuthServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetAuthServerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthServerUrl) {
		return nil, false
	}
	return o.AuthServerUrl, true
}

// HasAuthServerUrl returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasAuthServerUrl() bool {
	if o != nil && !IsNil(o.AuthServerUrl) {
		return true
	}

	return false
}

// SetAuthServerUrl gets a reference to the given string and assigns it to the AuthServerUrl field.
func (o *InstallationAdapterConfig) SetAuthServerUrl(v string) {
	o.AuthServerUrl = &v
}

// GetSslRequired returns the SslRequired field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetSslRequired() string {
	if o == nil || IsNil(o.SslRequired) {
		var ret string
		return ret
	}
	return *o.SslRequired
}

// GetSslRequiredOk returns a tuple with the SslRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetSslRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.SslRequired) {
		return nil, false
	}
	return o.SslRequired, true
}

// HasSslRequired returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasSslRequired() bool {
	if o != nil && !IsNil(o.SslRequired) {
		return true
	}

	return false
}

// SetSslRequired gets a reference to the given string and assigns it to the SslRequired field.
func (o *InstallationAdapterConfig) SetSslRequired(v string) {
	o.SslRequired = &v
}

// GetBearerOnly returns the BearerOnly field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetBearerOnly() bool {
	if o == nil || IsNil(o.BearerOnly) {
		var ret bool
		return ret
	}
	return *o.BearerOnly
}

// GetBearerOnlyOk returns a tuple with the BearerOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetBearerOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.BearerOnly) {
		return nil, false
	}
	return o.BearerOnly, true
}

// HasBearerOnly returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasBearerOnly() bool {
	if o != nil && !IsNil(o.BearerOnly) {
		return true
	}

	return false
}

// SetBearerOnly gets a reference to the given bool and assigns it to the BearerOnly field.
func (o *InstallationAdapterConfig) SetBearerOnly(v bool) {
	o.BearerOnly = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *InstallationAdapterConfig) SetResource(v string) {
	o.Resource = &v
}

// GetPublicClient returns the PublicClient field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetPublicClient() bool {
	if o == nil || IsNil(o.PublicClient) {
		var ret bool
		return ret
	}
	return *o.PublicClient
}

// GetPublicClientOk returns a tuple with the PublicClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetPublicClientOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicClient) {
		return nil, false
	}
	return o.PublicClient, true
}

// HasPublicClient returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasPublicClient() bool {
	if o != nil && !IsNil(o.PublicClient) {
		return true
	}

	return false
}

// SetPublicClient gets a reference to the given bool and assigns it to the PublicClient field.
func (o *InstallationAdapterConfig) SetPublicClient(v bool) {
	o.PublicClient = &v
}

// GetVerifyTokenAudience returns the VerifyTokenAudience field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetVerifyTokenAudience() bool {
	if o == nil || IsNil(o.VerifyTokenAudience) {
		var ret bool
		return ret
	}
	return *o.VerifyTokenAudience
}

// GetVerifyTokenAudienceOk returns a tuple with the VerifyTokenAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetVerifyTokenAudienceOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyTokenAudience) {
		return nil, false
	}
	return o.VerifyTokenAudience, true
}

// HasVerifyTokenAudience returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasVerifyTokenAudience() bool {
	if o != nil && !IsNil(o.VerifyTokenAudience) {
		return true
	}

	return false
}

// SetVerifyTokenAudience gets a reference to the given bool and assigns it to the VerifyTokenAudience field.
func (o *InstallationAdapterConfig) SetVerifyTokenAudience(v bool) {
	o.VerifyTokenAudience = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetCredentials() map[string]interface{} {
	if o == nil || IsNil(o.Credentials) {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *InstallationAdapterConfig) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetUseResourceRoleMappings returns the UseResourceRoleMappings field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetUseResourceRoleMappings() bool {
	if o == nil || IsNil(o.UseResourceRoleMappings) {
		var ret bool
		return ret
	}
	return *o.UseResourceRoleMappings
}

// GetUseResourceRoleMappingsOk returns a tuple with the UseResourceRoleMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetUseResourceRoleMappingsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseResourceRoleMappings) {
		return nil, false
	}
	return o.UseResourceRoleMappings, true
}

// HasUseResourceRoleMappings returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasUseResourceRoleMappings() bool {
	if o != nil && !IsNil(o.UseResourceRoleMappings) {
		return true
	}

	return false
}

// SetUseResourceRoleMappings gets a reference to the given bool and assigns it to the UseResourceRoleMappings field.
func (o *InstallationAdapterConfig) SetUseResourceRoleMappings(v bool) {
	o.UseResourceRoleMappings = &v
}

// GetConfidentialPort returns the ConfidentialPort field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetConfidentialPort() int32 {
	if o == nil || IsNil(o.ConfidentialPort) {
		var ret int32
		return ret
	}
	return *o.ConfidentialPort
}

// GetConfidentialPortOk returns a tuple with the ConfidentialPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetConfidentialPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfidentialPort) {
		return nil, false
	}
	return o.ConfidentialPort, true
}

// HasConfidentialPort returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasConfidentialPort() bool {
	if o != nil && !IsNil(o.ConfidentialPort) {
		return true
	}

	return false
}

// SetConfidentialPort gets a reference to the given int32 and assigns it to the ConfidentialPort field.
func (o *InstallationAdapterConfig) SetConfidentialPort(v int32) {
	o.ConfidentialPort = &v
}

// GetPolicyEnforcer returns the PolicyEnforcer field value if set, zero value otherwise.
func (o *InstallationAdapterConfig) GetPolicyEnforcer() PolicyEnforcerConfig {
	if o == nil || IsNil(o.PolicyEnforcer) {
		var ret PolicyEnforcerConfig
		return ret
	}
	return *o.PolicyEnforcer
}

// GetPolicyEnforcerOk returns a tuple with the PolicyEnforcer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationAdapterConfig) GetPolicyEnforcerOk() (*PolicyEnforcerConfig, bool) {
	if o == nil || IsNil(o.PolicyEnforcer) {
		return nil, false
	}
	return o.PolicyEnforcer, true
}

// HasPolicyEnforcer returns a boolean if a field has been set.
func (o *InstallationAdapterConfig) HasPolicyEnforcer() bool {
	if o != nil && !IsNil(o.PolicyEnforcer) {
		return true
	}

	return false
}

// SetPolicyEnforcer gets a reference to the given PolicyEnforcerConfig and assigns it to the PolicyEnforcer field.
func (o *InstallationAdapterConfig) SetPolicyEnforcer(v PolicyEnforcerConfig) {
	o.PolicyEnforcer = &v
}

func (o InstallationAdapterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationAdapterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.RealmPublicKey) {
		toSerialize["realm-public-key"] = o.RealmPublicKey
	}
	if !IsNil(o.AuthServerUrl) {
		toSerialize["auth-server-url"] = o.AuthServerUrl
	}
	if !IsNil(o.SslRequired) {
		toSerialize["ssl-required"] = o.SslRequired
	}
	if !IsNil(o.BearerOnly) {
		toSerialize["bearer-only"] = o.BearerOnly
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.PublicClient) {
		toSerialize["public-client"] = o.PublicClient
	}
	if !IsNil(o.VerifyTokenAudience) {
		toSerialize["verify-token-audience"] = o.VerifyTokenAudience
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.UseResourceRoleMappings) {
		toSerialize["use-resource-role-mappings"] = o.UseResourceRoleMappings
	}
	if !IsNil(o.ConfidentialPort) {
		toSerialize["confidential-port"] = o.ConfidentialPort
	}
	if !IsNil(o.PolicyEnforcer) {
		toSerialize["policy-enforcer"] = o.PolicyEnforcer
	}
	return toSerialize, nil
}

type NullableInstallationAdapterConfig struct {
	value *InstallationAdapterConfig
	isSet bool
}

func (v NullableInstallationAdapterConfig) Get() *InstallationAdapterConfig {
	return v.value
}

func (v *NullableInstallationAdapterConfig) Set(val *InstallationAdapterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationAdapterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationAdapterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationAdapterConfig(val *InstallationAdapterConfig) *NullableInstallationAdapterConfig {
	return &NullableInstallationAdapterConfig{value: val, isSet: true}
}

func (v NullableInstallationAdapterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationAdapterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


