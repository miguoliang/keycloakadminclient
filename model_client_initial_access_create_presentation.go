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

// checks if the ClientInitialAccessCreatePresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientInitialAccessCreatePresentation{}

// ClientInitialAccessCreatePresentation struct for ClientInitialAccessCreatePresentation
type ClientInitialAccessCreatePresentation struct {
	Expiration *int32 `json:"expiration,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewClientInitialAccessCreatePresentation instantiates a new ClientInitialAccessCreatePresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientInitialAccessCreatePresentation() *ClientInitialAccessCreatePresentation {
	this := ClientInitialAccessCreatePresentation{}
	return &this
}

// NewClientInitialAccessCreatePresentationWithDefaults instantiates a new ClientInitialAccessCreatePresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientInitialAccessCreatePresentationWithDefaults() *ClientInitialAccessCreatePresentation {
	this := ClientInitialAccessCreatePresentation{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ClientInitialAccessCreatePresentation) GetExpiration() int32 {
	if o == nil || IsNil(o.Expiration) {
		var ret int32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInitialAccessCreatePresentation) GetExpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ClientInitialAccessCreatePresentation) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int32 and assigns it to the Expiration field.
func (o *ClientInitialAccessCreatePresentation) SetExpiration(v int32) {
	o.Expiration = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ClientInitialAccessCreatePresentation) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInitialAccessCreatePresentation) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ClientInitialAccessCreatePresentation) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ClientInitialAccessCreatePresentation) SetCount(v int32) {
	o.Count = &v
}

func (o ClientInitialAccessCreatePresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientInitialAccessCreatePresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableClientInitialAccessCreatePresentation struct {
	value *ClientInitialAccessCreatePresentation
	isSet bool
}

func (v NullableClientInitialAccessCreatePresentation) Get() *ClientInitialAccessCreatePresentation {
	return v.value
}

func (v *NullableClientInitialAccessCreatePresentation) Set(val *ClientInitialAccessCreatePresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientInitialAccessCreatePresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientInitialAccessCreatePresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientInitialAccessCreatePresentation(val *ClientInitialAccessCreatePresentation) *NullableClientInitialAccessCreatePresentation {
	return &NullableClientInitialAccessCreatePresentation{value: val, isSet: true}
}

func (v NullableClientInitialAccessCreatePresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientInitialAccessCreatePresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


