/*
sdk_api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetItemStockPhysicalShippingInfoByItemStockIdRequest struct for GetItemStockPhysicalShippingInfoByItemStockIdRequest
type GetItemStockPhysicalShippingInfoByItemStockIdRequest struct {
	Data GetItemStockPhysicalShippingInfoByItemStockIdRequestData `json:"data"`
	Signature string `json:"signature"`
}

// NewGetItemStockPhysicalShippingInfoByItemStockIdRequest instantiates a new GetItemStockPhysicalShippingInfoByItemStockIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemStockPhysicalShippingInfoByItemStockIdRequest(data GetItemStockPhysicalShippingInfoByItemStockIdRequestData, signature string) *GetItemStockPhysicalShippingInfoByItemStockIdRequest {
	this := GetItemStockPhysicalShippingInfoByItemStockIdRequest{}
	this.Data = data
	this.Signature = signature
	return &this
}

// NewGetItemStockPhysicalShippingInfoByItemStockIdRequestWithDefaults instantiates a new GetItemStockPhysicalShippingInfoByItemStockIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemStockPhysicalShippingInfoByItemStockIdRequestWithDefaults() *GetItemStockPhysicalShippingInfoByItemStockIdRequest {
	this := GetItemStockPhysicalShippingInfoByItemStockIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) GetData() GetItemStockPhysicalShippingInfoByItemStockIdRequestData {
	if o == nil {
		var ret GetItemStockPhysicalShippingInfoByItemStockIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) GetDataOk() (*GetItemStockPhysicalShippingInfoByItemStockIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) SetData(v GetItemStockPhysicalShippingInfoByItemStockIdRequestData) {
	o.Data = v
}

// GetSignature returns the Signature field value
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *GetItemStockPhysicalShippingInfoByItemStockIdRequest) SetSignature(v string) {
	o.Signature = v
}

func (o GetItemStockPhysicalShippingInfoByItemStockIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest struct {
	value *GetItemStockPhysicalShippingInfoByItemStockIdRequest
	isSet bool
}

func (v NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) Get() *GetItemStockPhysicalShippingInfoByItemStockIdRequest {
	return v.value
}

func (v *NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) Set(val *GetItemStockPhysicalShippingInfoByItemStockIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemStockPhysicalShippingInfoByItemStockIdRequest(val *GetItemStockPhysicalShippingInfoByItemStockIdRequest) *NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest {
	return &NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest{value: val, isSet: true}
}

func (v NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetItemStockPhysicalShippingInfoByItemStockIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


