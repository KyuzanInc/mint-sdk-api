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

// GetItemStockPhysicalShippingInfoStatusByItemStockId200Response struct for GetItemStockPhysicalShippingInfoStatusByItemStockId200Response
type GetItemStockPhysicalShippingInfoStatusByItemStockId200Response struct {
	Data ItemStockPhysicalShippingInfoStatus `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

// NewGetItemStockPhysicalShippingInfoStatusByItemStockId200Response instantiates a new GetItemStockPhysicalShippingInfoStatusByItemStockId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemStockPhysicalShippingInfoStatusByItemStockId200Response(data ItemStockPhysicalShippingInfoStatus, meta map[string]interface{}) *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response {
	this := GetItemStockPhysicalShippingInfoStatusByItemStockId200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetItemStockPhysicalShippingInfoStatusByItemStockId200ResponseWithDefaults instantiates a new GetItemStockPhysicalShippingInfoStatusByItemStockId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemStockPhysicalShippingInfoStatusByItemStockId200ResponseWithDefaults() *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response {
	this := GetItemStockPhysicalShippingInfoStatusByItemStockId200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) GetData() ItemStockPhysicalShippingInfoStatus {
	if o == nil {
		var ret ItemStockPhysicalShippingInfoStatus
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) GetDataOk() (*ItemStockPhysicalShippingInfoStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) SetData(v ItemStockPhysicalShippingInfoStatus) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response struct {
	value *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response
	isSet bool
}

func (v NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) Get() *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response {
	return v.value
}

func (v *NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) Set(val *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response(val *GetItemStockPhysicalShippingInfoStatusByItemStockId200Response) *NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response {
	return &NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response{value: val, isSet: true}
}

func (v NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetItemStockPhysicalShippingInfoStatusByItemStockId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


