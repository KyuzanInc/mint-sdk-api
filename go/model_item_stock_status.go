/*
sdk_api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ItemStockStatus the model 'ItemStockStatus'
type ItemStockStatus string

// List of ItemStockStatus
const (
	CREATED ItemStockStatus = "created"
	MINTED ItemStockStatus = "minted"
)

// All allowed values of ItemStockStatus enum
var AllowedItemStockStatusEnumValues = []ItemStockStatus{
	"created",
	"minted",
}

func (v *ItemStockStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemStockStatus(value)
	for _, existing := range AllowedItemStockStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemStockStatus", value)
}

// NewItemStockStatusFromValue returns a pointer to a valid ItemStockStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemStockStatusFromValue(v string) (*ItemStockStatus, error) {
	ev := ItemStockStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemStockStatus: valid values are %v", v, AllowedItemStockStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemStockStatus) IsValid() bool {
	for _, existing := range AllowedItemStockStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemStockStatus value
func (v ItemStockStatus) Ptr() *ItemStockStatus {
	return &v
}

type NullableItemStockStatus struct {
	value *ItemStockStatus
	isSet bool
}

func (v NullableItemStockStatus) Get() *ItemStockStatus {
	return v.value
}

func (v *NullableItemStockStatus) Set(val *ItemStockStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStockStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStockStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStockStatus(val *ItemStockStatus) *NullableItemStockStatus {
	return &NullableItemStockStatus{value: val, isSet: true}
}

func (v NullableItemStockStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStockStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

