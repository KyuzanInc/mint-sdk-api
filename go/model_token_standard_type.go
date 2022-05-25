/*
 * sdk_api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TokenStandardType the model 'TokenStandardType'
type TokenStandardType string

// List of TokenStandardType
const (
	ERC721 TokenStandardType = "ERC721"
)

var allowedTokenStandardTypeEnumValues = []TokenStandardType{
	"ERC721",
}

func (v *TokenStandardType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenStandardType(value)
	for _, existing := range allowedTokenStandardTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TokenStandardType", value)
}

// NewTokenStandardTypeFromValue returns a pointer to a valid TokenStandardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenStandardTypeFromValue(v string) (*TokenStandardType, error) {
	ev := TokenStandardType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenStandardType: valid values are %v", v, allowedTokenStandardTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenStandardType) IsValid() bool {
	for _, existing := range allowedTokenStandardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TokenStandardType value
func (v TokenStandardType) Ptr() *TokenStandardType {
	return &v
}

type NullableTokenStandardType struct {
	value *TokenStandardType
	isSet bool
}

func (v NullableTokenStandardType) Get() *TokenStandardType {
	return v.value
}

func (v *NullableTokenStandardType) Set(val *TokenStandardType) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenStandardType) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenStandardType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenStandardType(val *TokenStandardType) *NullableTokenStandardType {
	return &NullableTokenStandardType{value: val, isSet: true}
}

func (v NullableTokenStandardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenStandardType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

