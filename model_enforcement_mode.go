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

// EnforcementMode the model 'EnforcementMode'
type EnforcementMode string

// List of EnforcementMode
const (
	PERMISSIVE EnforcementMode = "PERMISSIVE"
	ENFORCING EnforcementMode = "ENFORCING"
	DISABLED EnforcementMode = "DISABLED"
)

// All allowed values of EnforcementMode enum
var AllowedEnforcementModeEnumValues = []EnforcementMode{
	"PERMISSIVE",
	"ENFORCING",
	"DISABLED",
}

func (v *EnforcementMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnforcementMode(value)
	for _, existing := range AllowedEnforcementModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnforcementMode", value)
}

// NewEnforcementModeFromValue returns a pointer to a valid EnforcementMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnforcementModeFromValue(v string) (*EnforcementMode, error) {
	ev := EnforcementMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnforcementMode: valid values are %v", v, AllowedEnforcementModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnforcementMode) IsValid() bool {
	for _, existing := range AllowedEnforcementModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnforcementMode value
func (v EnforcementMode) Ptr() *EnforcementMode {
	return &v
}

type NullableEnforcementMode struct {
	value *EnforcementMode
	isSet bool
}

func (v NullableEnforcementMode) Get() *EnforcementMode {
	return v.value
}

func (v *NullableEnforcementMode) Set(val *EnforcementMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEnforcementMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEnforcementMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnforcementMode(val *EnforcementMode) *NullableEnforcementMode {
	return &NullableEnforcementMode{value: val, isSet: true}
}

func (v NullableEnforcementMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnforcementMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
