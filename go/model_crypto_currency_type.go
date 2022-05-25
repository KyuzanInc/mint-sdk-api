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

// CryptoCurrencyType the model 'CryptoCurrencyType'
type CryptoCurrencyType string

// List of CryptoCurrencyType
const (
	ETH CryptoCurrencyType = "eth"
	MATIC CryptoCurrencyType = "matic"
)

var allowedCryptoCurrencyTypeEnumValues = []CryptoCurrencyType{
	"eth",
	"matic",
}

func (v *CryptoCurrencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CryptoCurrencyType(value)
	for _, existing := range allowedCryptoCurrencyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CryptoCurrencyType", value)
}

// NewCryptoCurrencyTypeFromValue returns a pointer to a valid CryptoCurrencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCryptoCurrencyTypeFromValue(v string) (*CryptoCurrencyType, error) {
	ev := CryptoCurrencyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CryptoCurrencyType: valid values are %v", v, allowedCryptoCurrencyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CryptoCurrencyType) IsValid() bool {
	for _, existing := range allowedCryptoCurrencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CryptoCurrencyType value
func (v CryptoCurrencyType) Ptr() *CryptoCurrencyType {
	return &v
}

type NullableCryptoCurrencyType struct {
	value *CryptoCurrencyType
	isSet bool
}

func (v NullableCryptoCurrencyType) Get() *CryptoCurrencyType {
	return v.value
}

func (v *NullableCryptoCurrencyType) Set(val *CryptoCurrencyType) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoCurrencyType) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoCurrencyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoCurrencyType(val *CryptoCurrencyType) *NullableCryptoCurrencyType {
	return &NullableCryptoCurrencyType{value: val, isSet: true}
}

func (v NullableCryptoCurrencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoCurrencyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

