/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keycloakadminclient

import (
	"encoding/json"
	"fmt"
)

// KeyUse the model 'KeyUse'
type KeyUse string

// List of KeyUse
const (
	KEYUSE_SIG KeyUse = "SIG"
	KEYUSE_ENC KeyUse = "ENC"
)

// All allowed values of KeyUse enum
var AllowedKeyUseEnumValues = []KeyUse{
	"SIG",
	"ENC",
}

func (v *KeyUse) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyUse(value)
	for _, existing := range AllowedKeyUseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyUse", value)
}

// NewKeyUseFromValue returns a pointer to a valid KeyUse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyUseFromValue(v string) (*KeyUse, error) {
	ev := KeyUse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyUse: valid values are %v", v, AllowedKeyUseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyUse) IsValid() bool {
	for _, existing := range AllowedKeyUseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyUse value
func (v KeyUse) Ptr() *KeyUse {
	return &v
}

type NullableKeyUse struct {
	value *KeyUse
	isSet bool
}

func (v NullableKeyUse) Get() *KeyUse {
	return v.value
}

func (v *NullableKeyUse) Set(val *KeyUse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyUse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyUse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyUse(val *KeyUse) *NullableKeyUse {
	return &NullableKeyUse{value: val, isSet: true}
}

func (v NullableKeyUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyUse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
