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

// CreditCardStripeCurrencyType Stripeで利用する通貨型
type CreditCardStripeCurrencyType string

// List of CreditCardStripeCurrencyType
const (
	JPY CreditCardStripeCurrencyType = "jpy"
	USD CreditCardStripeCurrencyType = "usd"
	EUR CreditCardStripeCurrencyType = "eur"
)

var allowedCreditCardStripeCurrencyTypeEnumValues = []CreditCardStripeCurrencyType{
	"jpy",
	"usd",
	"eur",
}

func (v *CreditCardStripeCurrencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreditCardStripeCurrencyType(value)
	for _, existing := range allowedCreditCardStripeCurrencyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreditCardStripeCurrencyType", value)
}

// NewCreditCardStripeCurrencyTypeFromValue returns a pointer to a valid CreditCardStripeCurrencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditCardStripeCurrencyTypeFromValue(v string) (*CreditCardStripeCurrencyType, error) {
	ev := CreditCardStripeCurrencyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreditCardStripeCurrencyType: valid values are %v", v, allowedCreditCardStripeCurrencyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditCardStripeCurrencyType) IsValid() bool {
	for _, existing := range allowedCreditCardStripeCurrencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditCardStripeCurrencyType value
func (v CreditCardStripeCurrencyType) Ptr() *CreditCardStripeCurrencyType {
	return &v
}

type NullableCreditCardStripeCurrencyType struct {
	value *CreditCardStripeCurrencyType
	isSet bool
}

func (v NullableCreditCardStripeCurrencyType) Get() *CreditCardStripeCurrencyType {
	return v.value
}

func (v *NullableCreditCardStripeCurrencyType) Set(val *CreditCardStripeCurrencyType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardStripeCurrencyType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardStripeCurrencyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardStripeCurrencyType(val *CreditCardStripeCurrencyType) *NullableCreditCardStripeCurrencyType {
	return &NullableCreditCardStripeCurrencyType{value: val, isSet: true}
}

func (v NullableCreditCardStripeCurrencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardStripeCurrencyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
