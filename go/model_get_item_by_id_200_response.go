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

// GetItemById200Response struct for GetItemById200Response
type GetItemById200Response struct {
	Data Item `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

// NewGetItemById200Response instantiates a new GetItemById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItemById200Response(data Item, meta map[string]interface{}) *GetItemById200Response {
	this := GetItemById200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetItemById200ResponseWithDefaults instantiates a new GetItemById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItemById200ResponseWithDefaults() *GetItemById200Response {
	this := GetItemById200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetItemById200Response) GetData() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetItemById200Response) GetDataOk() (*Item, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetItemById200Response) SetData(v Item) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetItemById200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetItemById200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetItemById200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetItemById200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetItemById200Response struct {
	value *GetItemById200Response
	isSet bool
}

func (v NullableGetItemById200Response) Get() *GetItemById200Response {
	return v.value
}

func (v *NullableGetItemById200Response) Set(val *GetItemById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItemById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItemById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItemById200Response(val *GetItemById200Response) *NullableGetItemById200Response {
	return &NullableGetItemById200Response{value: val, isSet: true}
}

func (v NullableGetItemById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetItemById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


