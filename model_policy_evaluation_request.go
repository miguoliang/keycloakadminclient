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

// checks if the PolicyEvaluationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvaluationRequest{}

// PolicyEvaluationRequest struct for PolicyEvaluationRequest
type PolicyEvaluationRequest struct {
	Context *map[string]map[string]string `json:"context,omitempty"`
	Resources []ResourceRepresentation `json:"resources,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	RoleIds []string `json:"roleIds,omitempty"`
	Entitlements *bool `json:"entitlements,omitempty"`
}

// NewPolicyEvaluationRequest instantiates a new PolicyEvaluationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationRequest() *PolicyEvaluationRequest {
	this := PolicyEvaluationRequest{}
	return &this
}

// NewPolicyEvaluationRequestWithDefaults instantiates a new PolicyEvaluationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationRequestWithDefaults() *PolicyEvaluationRequest {
	this := PolicyEvaluationRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetContext() map[string]map[string]string {
	if o == nil || IsNil(o.Context) {
		var ret map[string]map[string]string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetContextOk() (*map[string]map[string]string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]map[string]string and assigns it to the Context field.
func (o *PolicyEvaluationRequest) SetContext(v map[string]map[string]string) {
	o.Context = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetResources() []ResourceRepresentation {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourceRepresentation
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetResourcesOk() ([]ResourceRepresentation, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceRepresentation and assigns it to the Resources field.
func (o *PolicyEvaluationRequest) SetResources(v []ResourceRepresentation) {
	o.Resources = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PolicyEvaluationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PolicyEvaluationRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *PolicyEvaluationRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *PolicyEvaluationRequest) GetEntitlements() bool {
	if o == nil || IsNil(o.Entitlements) {
		var ret bool
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationRequest) GetEntitlementsOk() (*bool, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *PolicyEvaluationRequest) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given bool and assigns it to the Entitlements field.
func (o *PolicyEvaluationRequest) SetEntitlements(v bool) {
	o.Entitlements = &v
}

func (o PolicyEvaluationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvaluationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	return toSerialize, nil
}

type NullablePolicyEvaluationRequest struct {
	value *PolicyEvaluationRequest
	isSet bool
}

func (v NullablePolicyEvaluationRequest) Get() *PolicyEvaluationRequest {
	return v.value
}

func (v *NullablePolicyEvaluationRequest) Set(val *PolicyEvaluationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationRequest(val *PolicyEvaluationRequest) *NullablePolicyEvaluationRequest {
	return &NullablePolicyEvaluationRequest{value: val, isSet: true}
}

func (v NullablePolicyEvaluationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

