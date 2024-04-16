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

// checks if the ClientTemplateRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientTemplateRepresentation{}

// ClientTemplateRepresentation struct for ClientTemplateRepresentation
type ClientTemplateRepresentation struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	FullScopeAllowed *bool `json:"fullScopeAllowed,omitempty"`
	BearerOnly *bool `json:"bearerOnly,omitempty"`
	ConsentRequired *bool `json:"consentRequired,omitempty"`
	StandardFlowEnabled *bool `json:"standardFlowEnabled,omitempty"`
	ImplicitFlowEnabled *bool `json:"implicitFlowEnabled,omitempty"`
	DirectAccessGrantsEnabled *bool `json:"directAccessGrantsEnabled,omitempty"`
	ServiceAccountsEnabled *bool `json:"serviceAccountsEnabled,omitempty"`
	PublicClient *bool `json:"publicClient,omitempty"`
	FrontchannelLogout *bool `json:"frontchannelLogout,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
	ProtocolMappers []ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
}

// NewClientTemplateRepresentation instantiates a new ClientTemplateRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientTemplateRepresentation() *ClientTemplateRepresentation {
	this := ClientTemplateRepresentation{}
	return &this
}

// NewClientTemplateRepresentationWithDefaults instantiates a new ClientTemplateRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientTemplateRepresentationWithDefaults() *ClientTemplateRepresentation {
	this := ClientTemplateRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientTemplateRepresentation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientTemplateRepresentation) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientTemplateRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ClientTemplateRepresentation) SetProtocol(v string) {
	o.Protocol = &v
}

// GetFullScopeAllowed returns the FullScopeAllowed field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetFullScopeAllowed() bool {
	if o == nil || IsNil(o.FullScopeAllowed) {
		var ret bool
		return ret
	}
	return *o.FullScopeAllowed
}

// GetFullScopeAllowedOk returns a tuple with the FullScopeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetFullScopeAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.FullScopeAllowed) {
		return nil, false
	}
	return o.FullScopeAllowed, true
}

// HasFullScopeAllowed returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasFullScopeAllowed() bool {
	if o != nil && !IsNil(o.FullScopeAllowed) {
		return true
	}

	return false
}

// SetFullScopeAllowed gets a reference to the given bool and assigns it to the FullScopeAllowed field.
func (o *ClientTemplateRepresentation) SetFullScopeAllowed(v bool) {
	o.FullScopeAllowed = &v
}

// GetBearerOnly returns the BearerOnly field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetBearerOnly() bool {
	if o == nil || IsNil(o.BearerOnly) {
		var ret bool
		return ret
	}
	return *o.BearerOnly
}

// GetBearerOnlyOk returns a tuple with the BearerOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetBearerOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.BearerOnly) {
		return nil, false
	}
	return o.BearerOnly, true
}

// HasBearerOnly returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasBearerOnly() bool {
	if o != nil && !IsNil(o.BearerOnly) {
		return true
	}

	return false
}

// SetBearerOnly gets a reference to the given bool and assigns it to the BearerOnly field.
func (o *ClientTemplateRepresentation) SetBearerOnly(v bool) {
	o.BearerOnly = &v
}

// GetConsentRequired returns the ConsentRequired field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetConsentRequired() bool {
	if o == nil || IsNil(o.ConsentRequired) {
		var ret bool
		return ret
	}
	return *o.ConsentRequired
}

// GetConsentRequiredOk returns a tuple with the ConsentRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetConsentRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentRequired) {
		return nil, false
	}
	return o.ConsentRequired, true
}

// HasConsentRequired returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasConsentRequired() bool {
	if o != nil && !IsNil(o.ConsentRequired) {
		return true
	}

	return false
}

// SetConsentRequired gets a reference to the given bool and assigns it to the ConsentRequired field.
func (o *ClientTemplateRepresentation) SetConsentRequired(v bool) {
	o.ConsentRequired = &v
}

// GetStandardFlowEnabled returns the StandardFlowEnabled field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetStandardFlowEnabled() bool {
	if o == nil || IsNil(o.StandardFlowEnabled) {
		var ret bool
		return ret
	}
	return *o.StandardFlowEnabled
}

// GetStandardFlowEnabledOk returns a tuple with the StandardFlowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetStandardFlowEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StandardFlowEnabled) {
		return nil, false
	}
	return o.StandardFlowEnabled, true
}

// HasStandardFlowEnabled returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasStandardFlowEnabled() bool {
	if o != nil && !IsNil(o.StandardFlowEnabled) {
		return true
	}

	return false
}

// SetStandardFlowEnabled gets a reference to the given bool and assigns it to the StandardFlowEnabled field.
func (o *ClientTemplateRepresentation) SetStandardFlowEnabled(v bool) {
	o.StandardFlowEnabled = &v
}

// GetImplicitFlowEnabled returns the ImplicitFlowEnabled field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetImplicitFlowEnabled() bool {
	if o == nil || IsNil(o.ImplicitFlowEnabled) {
		var ret bool
		return ret
	}
	return *o.ImplicitFlowEnabled
}

// GetImplicitFlowEnabledOk returns a tuple with the ImplicitFlowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetImplicitFlowEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitFlowEnabled) {
		return nil, false
	}
	return o.ImplicitFlowEnabled, true
}

// HasImplicitFlowEnabled returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasImplicitFlowEnabled() bool {
	if o != nil && !IsNil(o.ImplicitFlowEnabled) {
		return true
	}

	return false
}

// SetImplicitFlowEnabled gets a reference to the given bool and assigns it to the ImplicitFlowEnabled field.
func (o *ClientTemplateRepresentation) SetImplicitFlowEnabled(v bool) {
	o.ImplicitFlowEnabled = &v
}

// GetDirectAccessGrantsEnabled returns the DirectAccessGrantsEnabled field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetDirectAccessGrantsEnabled() bool {
	if o == nil || IsNil(o.DirectAccessGrantsEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectAccessGrantsEnabled
}

// GetDirectAccessGrantsEnabledOk returns a tuple with the DirectAccessGrantsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetDirectAccessGrantsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectAccessGrantsEnabled) {
		return nil, false
	}
	return o.DirectAccessGrantsEnabled, true
}

// HasDirectAccessGrantsEnabled returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasDirectAccessGrantsEnabled() bool {
	if o != nil && !IsNil(o.DirectAccessGrantsEnabled) {
		return true
	}

	return false
}

// SetDirectAccessGrantsEnabled gets a reference to the given bool and assigns it to the DirectAccessGrantsEnabled field.
func (o *ClientTemplateRepresentation) SetDirectAccessGrantsEnabled(v bool) {
	o.DirectAccessGrantsEnabled = &v
}

// GetServiceAccountsEnabled returns the ServiceAccountsEnabled field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetServiceAccountsEnabled() bool {
	if o == nil || IsNil(o.ServiceAccountsEnabled) {
		var ret bool
		return ret
	}
	return *o.ServiceAccountsEnabled
}

// GetServiceAccountsEnabledOk returns a tuple with the ServiceAccountsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetServiceAccountsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ServiceAccountsEnabled) {
		return nil, false
	}
	return o.ServiceAccountsEnabled, true
}

// HasServiceAccountsEnabled returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasServiceAccountsEnabled() bool {
	if o != nil && !IsNil(o.ServiceAccountsEnabled) {
		return true
	}

	return false
}

// SetServiceAccountsEnabled gets a reference to the given bool and assigns it to the ServiceAccountsEnabled field.
func (o *ClientTemplateRepresentation) SetServiceAccountsEnabled(v bool) {
	o.ServiceAccountsEnabled = &v
}

// GetPublicClient returns the PublicClient field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetPublicClient() bool {
	if o == nil || IsNil(o.PublicClient) {
		var ret bool
		return ret
	}
	return *o.PublicClient
}

// GetPublicClientOk returns a tuple with the PublicClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetPublicClientOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicClient) {
		return nil, false
	}
	return o.PublicClient, true
}

// HasPublicClient returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasPublicClient() bool {
	if o != nil && !IsNil(o.PublicClient) {
		return true
	}

	return false
}

// SetPublicClient gets a reference to the given bool and assigns it to the PublicClient field.
func (o *ClientTemplateRepresentation) SetPublicClient(v bool) {
	o.PublicClient = &v
}

// GetFrontchannelLogout returns the FrontchannelLogout field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetFrontchannelLogout() bool {
	if o == nil || IsNil(o.FrontchannelLogout) {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogout
}

// GetFrontchannelLogoutOk returns a tuple with the FrontchannelLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetFrontchannelLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontchannelLogout) {
		return nil, false
	}
	return o.FrontchannelLogout, true
}

// HasFrontchannelLogout returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasFrontchannelLogout() bool {
	if o != nil && !IsNil(o.FrontchannelLogout) {
		return true
	}

	return false
}

// SetFrontchannelLogout gets a reference to the given bool and assigns it to the FrontchannelLogout field.
func (o *ClientTemplateRepresentation) SetFrontchannelLogout(v bool) {
	o.FrontchannelLogout = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ClientTemplateRepresentation) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetProtocolMappers returns the ProtocolMappers field value if set, zero value otherwise.
func (o *ClientTemplateRepresentation) GetProtocolMappers() []ProtocolMapperRepresentation {
	if o == nil || IsNil(o.ProtocolMappers) {
		var ret []ProtocolMapperRepresentation
		return ret
	}
	return o.ProtocolMappers
}

// GetProtocolMappersOk returns a tuple with the ProtocolMappers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTemplateRepresentation) GetProtocolMappersOk() ([]ProtocolMapperRepresentation, bool) {
	if o == nil || IsNil(o.ProtocolMappers) {
		return nil, false
	}
	return o.ProtocolMappers, true
}

// HasProtocolMappers returns a boolean if a field has been set.
func (o *ClientTemplateRepresentation) HasProtocolMappers() bool {
	if o != nil && !IsNil(o.ProtocolMappers) {
		return true
	}

	return false
}

// SetProtocolMappers gets a reference to the given []ProtocolMapperRepresentation and assigns it to the ProtocolMappers field.
func (o *ClientTemplateRepresentation) SetProtocolMappers(v []ProtocolMapperRepresentation) {
	o.ProtocolMappers = v
}

func (o ClientTemplateRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientTemplateRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.FullScopeAllowed) {
		toSerialize["fullScopeAllowed"] = o.FullScopeAllowed
	}
	if !IsNil(o.BearerOnly) {
		toSerialize["bearerOnly"] = o.BearerOnly
	}
	if !IsNil(o.ConsentRequired) {
		toSerialize["consentRequired"] = o.ConsentRequired
	}
	if !IsNil(o.StandardFlowEnabled) {
		toSerialize["standardFlowEnabled"] = o.StandardFlowEnabled
	}
	if !IsNil(o.ImplicitFlowEnabled) {
		toSerialize["implicitFlowEnabled"] = o.ImplicitFlowEnabled
	}
	if !IsNil(o.DirectAccessGrantsEnabled) {
		toSerialize["directAccessGrantsEnabled"] = o.DirectAccessGrantsEnabled
	}
	if !IsNil(o.ServiceAccountsEnabled) {
		toSerialize["serviceAccountsEnabled"] = o.ServiceAccountsEnabled
	}
	if !IsNil(o.PublicClient) {
		toSerialize["publicClient"] = o.PublicClient
	}
	if !IsNil(o.FrontchannelLogout) {
		toSerialize["frontchannelLogout"] = o.FrontchannelLogout
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.ProtocolMappers) {
		toSerialize["protocolMappers"] = o.ProtocolMappers
	}
	return toSerialize, nil
}

type NullableClientTemplateRepresentation struct {
	value *ClientTemplateRepresentation
	isSet bool
}

func (v NullableClientTemplateRepresentation) Get() *ClientTemplateRepresentation {
	return v.value
}

func (v *NullableClientTemplateRepresentation) Set(val *ClientTemplateRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientTemplateRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientTemplateRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientTemplateRepresentation(val *ClientTemplateRepresentation) *NullableClientTemplateRepresentation {
	return &NullableClientTemplateRepresentation{value: val, isSet: true}
}

func (v NullableClientTemplateRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientTemplateRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

