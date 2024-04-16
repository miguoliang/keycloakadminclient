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

// checks if the UserProfileAttributeMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfileAttributeMetadata{}

// UserProfileAttributeMetadata struct for UserProfileAttributeMetadata
type UserProfileAttributeMetadata struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Required *bool `json:"required,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	Validators *map[string]map[string]interface{} `json:"validators,omitempty"`
	Group *string `json:"group,omitempty"`
	Multivalued *bool `json:"multivalued,omitempty"`
}

// NewUserProfileAttributeMetadata instantiates a new UserProfileAttributeMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfileAttributeMetadata() *UserProfileAttributeMetadata {
	this := UserProfileAttributeMetadata{}
	return &this
}

// NewUserProfileAttributeMetadataWithDefaults instantiates a new UserProfileAttributeMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileAttributeMetadataWithDefaults() *UserProfileAttributeMetadata {
	this := UserProfileAttributeMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserProfileAttributeMetadata) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserProfileAttributeMetadata) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UserProfileAttributeMetadata) SetRequired(v bool) {
	o.Required = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *UserProfileAttributeMetadata) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *UserProfileAttributeMetadata) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetValidators returns the Validators field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetValidators() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Validators) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Validators
}

// GetValidatorsOk returns a tuple with the Validators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetValidatorsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Validators) {
		return nil, false
	}
	return o.Validators, true
}

// HasValidators returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasValidators() bool {
	if o != nil && !IsNil(o.Validators) {
		return true
	}

	return false
}

// SetValidators gets a reference to the given map[string]map[string]interface{} and assigns it to the Validators field.
func (o *UserProfileAttributeMetadata) SetValidators(v map[string]map[string]interface{}) {
	o.Validators = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *UserProfileAttributeMetadata) SetGroup(v string) {
	o.Group = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *UserProfileAttributeMetadata) GetMultivalued() bool {
	if o == nil || IsNil(o.Multivalued) {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileAttributeMetadata) GetMultivaluedOk() (*bool, bool) {
	if o == nil || IsNil(o.Multivalued) {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *UserProfileAttributeMetadata) HasMultivalued() bool {
	if o != nil && !IsNil(o.Multivalued) {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *UserProfileAttributeMetadata) SetMultivalued(v bool) {
	o.Multivalued = &v
}

func (o UserProfileAttributeMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfileAttributeMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.Validators) {
		toSerialize["validators"] = o.Validators
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Multivalued) {
		toSerialize["multivalued"] = o.Multivalued
	}
	return toSerialize, nil
}

type NullableUserProfileAttributeMetadata struct {
	value *UserProfileAttributeMetadata
	isSet bool
}

func (v NullableUserProfileAttributeMetadata) Get() *UserProfileAttributeMetadata {
	return v.value
}

func (v *NullableUserProfileAttributeMetadata) Set(val *UserProfileAttributeMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileAttributeMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileAttributeMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileAttributeMetadata(val *UserProfileAttributeMetadata) *NullableUserProfileAttributeMetadata {
	return &NullableUserProfileAttributeMetadata{value: val, isSet: true}
}

func (v NullableUserProfileAttributeMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileAttributeMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


