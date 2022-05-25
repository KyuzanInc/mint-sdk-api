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
)

// SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage struct for SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage
type SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage struct {
	WalletAddress *string `json:"walletAddress,omitempty"`
	RequestTimestamp *string `json:"requestTimestamp,omitempty"`
}

// NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage instantiates a new SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage() *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage {
	this := SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage{}
	return &this
}

// NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessageWithDefaults instantiates a new SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessageWithDefaults() *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage {
	this := SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage{}
	return &this
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) GetWalletAddress() string {
	if o == nil || o.WalletAddress == nil {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) GetWalletAddressOk() (*string, bool) {
	if o == nil || o.WalletAddress == nil {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) HasWalletAddress() bool {
	if o != nil && o.WalletAddress != nil {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) GetRequestTimestamp() string {
	if o == nil || o.RequestTimestamp == nil {
		var ret string
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) GetRequestTimestampOk() (*string, bool) {
	if o == nil || o.RequestTimestamp == nil {
		return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) HasRequestTimestamp() bool {
	if o != nil && o.RequestTimestamp != nil {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given string and assigns it to the RequestTimestamp field.
func (o *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) SetRequestTimestamp(v string) {
	o.RequestTimestamp = &v
}

func (o SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WalletAddress != nil {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if o.RequestTimestamp != nil {
		toSerialize["requestTimestamp"] = o.RequestTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage struct {
	value *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage
	isSet bool
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) Get() *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage {
	return v.value
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) Set(val *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage(val *SdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) *NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage {
	return &NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage{value: val, isSet: true}
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

