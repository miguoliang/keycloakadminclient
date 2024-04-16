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

// checks if the GlobalRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRequestResult{}

// GlobalRequestResult struct for GlobalRequestResult
type GlobalRequestResult struct {
	SuccessRequests []string `json:"successRequests,omitempty"`
	FailedRequests []string `json:"failedRequests,omitempty"`
}

// NewGlobalRequestResult instantiates a new GlobalRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRequestResult() *GlobalRequestResult {
	this := GlobalRequestResult{}
	return &this
}

// NewGlobalRequestResultWithDefaults instantiates a new GlobalRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRequestResultWithDefaults() *GlobalRequestResult {
	this := GlobalRequestResult{}
	return &this
}

// GetSuccessRequests returns the SuccessRequests field value if set, zero value otherwise.
func (o *GlobalRequestResult) GetSuccessRequests() []string {
	if o == nil || IsNil(o.SuccessRequests) {
		var ret []string
		return ret
	}
	return o.SuccessRequests
}

// GetSuccessRequestsOk returns a tuple with the SuccessRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRequestResult) GetSuccessRequestsOk() ([]string, bool) {
	if o == nil || IsNil(o.SuccessRequests) {
		return nil, false
	}
	return o.SuccessRequests, true
}

// HasSuccessRequests returns a boolean if a field has been set.
func (o *GlobalRequestResult) HasSuccessRequests() bool {
	if o != nil && !IsNil(o.SuccessRequests) {
		return true
	}

	return false
}

// SetSuccessRequests gets a reference to the given []string and assigns it to the SuccessRequests field.
func (o *GlobalRequestResult) SetSuccessRequests(v []string) {
	o.SuccessRequests = v
}

// GetFailedRequests returns the FailedRequests field value if set, zero value otherwise.
func (o *GlobalRequestResult) GetFailedRequests() []string {
	if o == nil || IsNil(o.FailedRequests) {
		var ret []string
		return ret
	}
	return o.FailedRequests
}

// GetFailedRequestsOk returns a tuple with the FailedRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRequestResult) GetFailedRequestsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailedRequests) {
		return nil, false
	}
	return o.FailedRequests, true
}

// HasFailedRequests returns a boolean if a field has been set.
func (o *GlobalRequestResult) HasFailedRequests() bool {
	if o != nil && !IsNil(o.FailedRequests) {
		return true
	}

	return false
}

// SetFailedRequests gets a reference to the given []string and assigns it to the FailedRequests field.
func (o *GlobalRequestResult) SetFailedRequests(v []string) {
	o.FailedRequests = v
}

func (o GlobalRequestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuccessRequests) {
		toSerialize["successRequests"] = o.SuccessRequests
	}
	if !IsNil(o.FailedRequests) {
		toSerialize["failedRequests"] = o.FailedRequests
	}
	return toSerialize, nil
}

type NullableGlobalRequestResult struct {
	value *GlobalRequestResult
	isSet bool
}

func (v NullableGlobalRequestResult) Get() *GlobalRequestResult {
	return v.value
}

func (v *NullableGlobalRequestResult) Set(val *GlobalRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRequestResult(val *GlobalRequestResult) *NullableGlobalRequestResult {
	return &NullableGlobalRequestResult{value: val, isSet: true}
}

func (v NullableGlobalRequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

