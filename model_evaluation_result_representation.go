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

// checks if the EvaluationResultRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvaluationResultRepresentation{}

// EvaluationResultRepresentation struct for EvaluationResultRepresentation
type EvaluationResultRepresentation struct {
	Resource *ResourceRepresentation `json:"resource,omitempty"`
	Scopes []ScopeRepresentation `json:"scopes,omitempty"`
	Policies []PolicyResultRepresentation `json:"policies,omitempty"`
	Status *DecisionEffect `json:"status,omitempty"`
	AllowedScopes []ScopeRepresentation `json:"allowedScopes,omitempty"`
}

// NewEvaluationResultRepresentation instantiates a new EvaluationResultRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluationResultRepresentation() *EvaluationResultRepresentation {
	this := EvaluationResultRepresentation{}
	return &this
}

// NewEvaluationResultRepresentationWithDefaults instantiates a new EvaluationResultRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluationResultRepresentationWithDefaults() *EvaluationResultRepresentation {
	this := EvaluationResultRepresentation{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *EvaluationResultRepresentation) GetResource() ResourceRepresentation {
	if o == nil || IsNil(o.Resource) {
		var ret ResourceRepresentation
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationResultRepresentation) GetResourceOk() (*ResourceRepresentation, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *EvaluationResultRepresentation) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ResourceRepresentation and assigns it to the Resource field.
func (o *EvaluationResultRepresentation) SetResource(v ResourceRepresentation) {
	o.Resource = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *EvaluationResultRepresentation) GetScopes() []ScopeRepresentation {
	if o == nil || IsNil(o.Scopes) {
		var ret []ScopeRepresentation
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationResultRepresentation) GetScopesOk() ([]ScopeRepresentation, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *EvaluationResultRepresentation) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ScopeRepresentation and assigns it to the Scopes field.
func (o *EvaluationResultRepresentation) SetScopes(v []ScopeRepresentation) {
	o.Scopes = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *EvaluationResultRepresentation) GetPolicies() []PolicyResultRepresentation {
	if o == nil || IsNil(o.Policies) {
		var ret []PolicyResultRepresentation
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationResultRepresentation) GetPoliciesOk() ([]PolicyResultRepresentation, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *EvaluationResultRepresentation) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []PolicyResultRepresentation and assigns it to the Policies field.
func (o *EvaluationResultRepresentation) SetPolicies(v []PolicyResultRepresentation) {
	o.Policies = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EvaluationResultRepresentation) GetStatus() DecisionEffect {
	if o == nil || IsNil(o.Status) {
		var ret DecisionEffect
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationResultRepresentation) GetStatusOk() (*DecisionEffect, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EvaluationResultRepresentation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DecisionEffect and assigns it to the Status field.
func (o *EvaluationResultRepresentation) SetStatus(v DecisionEffect) {
	o.Status = &v
}

// GetAllowedScopes returns the AllowedScopes field value if set, zero value otherwise.
func (o *EvaluationResultRepresentation) GetAllowedScopes() []ScopeRepresentation {
	if o == nil || IsNil(o.AllowedScopes) {
		var ret []ScopeRepresentation
		return ret
	}
	return o.AllowedScopes
}

// GetAllowedScopesOk returns a tuple with the AllowedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationResultRepresentation) GetAllowedScopesOk() ([]ScopeRepresentation, bool) {
	if o == nil || IsNil(o.AllowedScopes) {
		return nil, false
	}
	return o.AllowedScopes, true
}

// HasAllowedScopes returns a boolean if a field has been set.
func (o *EvaluationResultRepresentation) HasAllowedScopes() bool {
	if o != nil && !IsNil(o.AllowedScopes) {
		return true
	}

	return false
}

// SetAllowedScopes gets a reference to the given []ScopeRepresentation and assigns it to the AllowedScopes field.
func (o *EvaluationResultRepresentation) SetAllowedScopes(v []ScopeRepresentation) {
	o.AllowedScopes = v
}

func (o EvaluationResultRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvaluationResultRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AllowedScopes) {
		toSerialize["allowedScopes"] = o.AllowedScopes
	}
	return toSerialize, nil
}

type NullableEvaluationResultRepresentation struct {
	value *EvaluationResultRepresentation
	isSet bool
}

func (v NullableEvaluationResultRepresentation) Get() *EvaluationResultRepresentation {
	return v.value
}

func (v *NullableEvaluationResultRepresentation) Set(val *EvaluationResultRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluationResultRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluationResultRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluationResultRepresentation(val *EvaluationResultRepresentation) *NullableEvaluationResultRepresentation {
	return &NullableEvaluationResultRepresentation{value: val, isSet: true}
}

func (v NullableEvaluationResultRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluationResultRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

