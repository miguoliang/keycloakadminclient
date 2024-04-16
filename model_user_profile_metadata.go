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

// checks if the UserProfileMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfileMetadata{}

// UserProfileMetadata struct for UserProfileMetadata
type UserProfileMetadata struct {
	Attributes []UserProfileAttributeMetadata `json:"attributes,omitempty"`
	Groups []UserProfileAttributeGroupMetadata `json:"groups,omitempty"`
}

// NewUserProfileMetadata instantiates a new UserProfileMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfileMetadata() *UserProfileMetadata {
	this := UserProfileMetadata{}
	return &this
}

// NewUserProfileMetadataWithDefaults instantiates a new UserProfileMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileMetadataWithDefaults() *UserProfileMetadata {
	this := UserProfileMetadata{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserProfileMetadata) GetAttributes() []UserProfileAttributeMetadata {
	if o == nil || IsNil(o.Attributes) {
		var ret []UserProfileAttributeMetadata
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileMetadata) GetAttributesOk() ([]UserProfileAttributeMetadata, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserProfileMetadata) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []UserProfileAttributeMetadata and assigns it to the Attributes field.
func (o *UserProfileMetadata) SetAttributes(v []UserProfileAttributeMetadata) {
	o.Attributes = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserProfileMetadata) GetGroups() []UserProfileAttributeGroupMetadata {
	if o == nil || IsNil(o.Groups) {
		var ret []UserProfileAttributeGroupMetadata
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileMetadata) GetGroupsOk() ([]UserProfileAttributeGroupMetadata, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserProfileMetadata) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []UserProfileAttributeGroupMetadata and assigns it to the Groups field.
func (o *UserProfileMetadata) SetGroups(v []UserProfileAttributeGroupMetadata) {
	o.Groups = v
}

func (o UserProfileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfileMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableUserProfileMetadata struct {
	value *UserProfileMetadata
	isSet bool
}

func (v NullableUserProfileMetadata) Get() *UserProfileMetadata {
	return v.value
}

func (v *NullableUserProfileMetadata) Set(val *UserProfileMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileMetadata(val *UserProfileMetadata) *NullableUserProfileMetadata {
	return &NullableUserProfileMetadata{value: val, isSet: true}
}

func (v NullableUserProfileMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

