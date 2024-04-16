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

// checks if the Confirmation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Confirmation{}

// Confirmation struct for Confirmation
type Confirmation struct {
	X5tS256 *string `json:"x5t#S256,omitempty"`
	Jkt *string `json:"jkt,omitempty"`
}

// NewConfirmation instantiates a new Confirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmation() *Confirmation {
	this := Confirmation{}
	return &this
}

// NewConfirmationWithDefaults instantiates a new Confirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmationWithDefaults() *Confirmation {
	this := Confirmation{}
	return &this
}

// GetX5tS256 returns the X5tS256 field value if set, zero value otherwise.
func (o *Confirmation) GetX5tS256() string {
	if o == nil || IsNil(o.X5tS256) {
		var ret string
		return ret
	}
	return *o.X5tS256
}

// GetX5tS256Ok returns a tuple with the X5tS256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Confirmation) GetX5tS256Ok() (*string, bool) {
	if o == nil || IsNil(o.X5tS256) {
		return nil, false
	}
	return o.X5tS256, true
}

// HasX5tS256 returns a boolean if a field has been set.
func (o *Confirmation) HasX5tS256() bool {
	if o != nil && !IsNil(o.X5tS256) {
		return true
	}

	return false
}

// SetX5tS256 gets a reference to the given string and assigns it to the X5tS256 field.
func (o *Confirmation) SetX5tS256(v string) {
	o.X5tS256 = &v
}

// GetJkt returns the Jkt field value if set, zero value otherwise.
func (o *Confirmation) GetJkt() string {
	if o == nil || IsNil(o.Jkt) {
		var ret string
		return ret
	}
	return *o.Jkt
}

// GetJktOk returns a tuple with the Jkt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Confirmation) GetJktOk() (*string, bool) {
	if o == nil || IsNil(o.Jkt) {
		return nil, false
	}
	return o.Jkt, true
}

// HasJkt returns a boolean if a field has been set.
func (o *Confirmation) HasJkt() bool {
	if o != nil && !IsNil(o.Jkt) {
		return true
	}

	return false
}

// SetJkt gets a reference to the given string and assigns it to the Jkt field.
func (o *Confirmation) SetJkt(v string) {
	o.Jkt = &v
}

func (o Confirmation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Confirmation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X5tS256) {
		toSerialize["x5t#S256"] = o.X5tS256
	}
	if !IsNil(o.Jkt) {
		toSerialize["jkt"] = o.Jkt
	}
	return toSerialize, nil
}

type NullableConfirmation struct {
	value *Confirmation
	isSet bool
}

func (v NullableConfirmation) Get() *Confirmation {
	return v.value
}

func (v *NullableConfirmation) Set(val *Confirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmation(val *Confirmation) *NullableConfirmation {
	return &NullableConfirmation{value: val, isSet: true}
}

func (v NullableConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

