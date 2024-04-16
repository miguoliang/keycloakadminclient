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

// checks if the MappingsRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MappingsRepresentation{}

// MappingsRepresentation struct for MappingsRepresentation
type MappingsRepresentation struct {
	RealmMappings []RoleRepresentation `json:"realmMappings,omitempty"`
	ClientMappings *map[string]ClientMappingsRepresentation `json:"clientMappings,omitempty"`
}

// NewMappingsRepresentation instantiates a new MappingsRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingsRepresentation() *MappingsRepresentation {
	this := MappingsRepresentation{}
	return &this
}

// NewMappingsRepresentationWithDefaults instantiates a new MappingsRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingsRepresentationWithDefaults() *MappingsRepresentation {
	this := MappingsRepresentation{}
	return &this
}

// GetRealmMappings returns the RealmMappings field value if set, zero value otherwise.
func (o *MappingsRepresentation) GetRealmMappings() []RoleRepresentation {
	if o == nil || IsNil(o.RealmMappings) {
		var ret []RoleRepresentation
		return ret
	}
	return o.RealmMappings
}

// GetRealmMappingsOk returns a tuple with the RealmMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingsRepresentation) GetRealmMappingsOk() ([]RoleRepresentation, bool) {
	if o == nil || IsNil(o.RealmMappings) {
		return nil, false
	}
	return o.RealmMappings, true
}

// HasRealmMappings returns a boolean if a field has been set.
func (o *MappingsRepresentation) HasRealmMappings() bool {
	if o != nil && !IsNil(o.RealmMappings) {
		return true
	}

	return false
}

// SetRealmMappings gets a reference to the given []RoleRepresentation and assigns it to the RealmMappings field.
func (o *MappingsRepresentation) SetRealmMappings(v []RoleRepresentation) {
	o.RealmMappings = v
}

// GetClientMappings returns the ClientMappings field value if set, zero value otherwise.
func (o *MappingsRepresentation) GetClientMappings() map[string]ClientMappingsRepresentation {
	if o == nil || IsNil(o.ClientMappings) {
		var ret map[string]ClientMappingsRepresentation
		return ret
	}
	return *o.ClientMappings
}

// GetClientMappingsOk returns a tuple with the ClientMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingsRepresentation) GetClientMappingsOk() (*map[string]ClientMappingsRepresentation, bool) {
	if o == nil || IsNil(o.ClientMappings) {
		return nil, false
	}
	return o.ClientMappings, true
}

// HasClientMappings returns a boolean if a field has been set.
func (o *MappingsRepresentation) HasClientMappings() bool {
	if o != nil && !IsNil(o.ClientMappings) {
		return true
	}

	return false
}

// SetClientMappings gets a reference to the given map[string]ClientMappingsRepresentation and assigns it to the ClientMappings field.
func (o *MappingsRepresentation) SetClientMappings(v map[string]ClientMappingsRepresentation) {
	o.ClientMappings = &v
}

func (o MappingsRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MappingsRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RealmMappings) {
		toSerialize["realmMappings"] = o.RealmMappings
	}
	if !IsNil(o.ClientMappings) {
		toSerialize["clientMappings"] = o.ClientMappings
	}
	return toSerialize, nil
}

type NullableMappingsRepresentation struct {
	value *MappingsRepresentation
	isSet bool
}

func (v NullableMappingsRepresentation) Get() *MappingsRepresentation {
	return v.value
}

func (v *NullableMappingsRepresentation) Set(val *MappingsRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingsRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingsRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingsRepresentation(val *MappingsRepresentation) *NullableMappingsRepresentation {
	return &NullableMappingsRepresentation{value: val, isSet: true}
}

func (v NullableMappingsRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingsRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


