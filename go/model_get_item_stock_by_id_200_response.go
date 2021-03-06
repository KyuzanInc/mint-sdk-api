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

// GetItemStockById200Response struct for GetItemStockById200Response
type GetItemStockById200Response struct {
	Data ItemStock `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

// NewGetItemStockById200Response instantiates a new GetItemStockById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemStockById200Response(data ItemStock, meta map[string]interface{}) *GetItemStockById200Response {
	this := GetItemStockById200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetItemStockById200ResponseWithDefaults instantiates a new GetItemStockById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemStockById200ResponseWithDefaults() *GetItemStockById200Response {
	this := GetItemStockById200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetItemStockById200Response) GetData() ItemStock {
	if o == nil {
		var ret ItemStock
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetItemStockById200Response) GetDataOk() (*ItemStock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetItemStockById200Response) SetData(v ItemStock) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetItemStockById200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetItemStockById200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetItemStockById200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetItemStockById200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetItemStockById200Response struct {
	value *GetItemStockById200Response
	isSet bool
}

func (v NullableGetItemStockById200Response) Get() *GetItemStockById200Response {
	return v.value
}

func (v *NullableGetItemStockById200Response) Set(val *GetItemStockById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemStockById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemStockById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemStockById200Response(val *GetItemStockById200Response) *NullableGetItemStockById200Response {
	return &NullableGetItemStockById200Response{value: val, isSet: true}
}

func (v NullableGetItemStockById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetItemStockById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


