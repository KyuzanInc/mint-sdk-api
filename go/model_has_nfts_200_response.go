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

// HasNfts200Response struct for HasNfts200Response
type HasNfts200Response struct {
	Data bool `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

// NewHasNfts200Response instantiates a new HasNfts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasNfts200Response(data bool, meta map[string]interface{}) *HasNfts200Response {
	this := HasNfts200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewHasNfts200ResponseWithDefaults instantiates a new HasNfts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasNfts200ResponseWithDefaults() *HasNfts200Response {
	this := HasNfts200Response{}
	return &this
}

// GetData returns the Data field value
func (o *HasNfts200Response) GetData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *HasNfts200Response) GetDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *HasNfts200Response) SetData(v bool) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *HasNfts200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *HasNfts200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *HasNfts200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o HasNfts200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableHasNfts200Response struct {
	value *HasNfts200Response
	isSet bool
}

func (v NullableHasNfts200Response) Get() *HasNfts200Response {
	return v.value
}

func (v *NullableHasNfts200Response) Set(val *HasNfts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableHasNfts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableHasNfts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasNfts200Response(val *HasNfts200Response) *NullableHasNfts200Response {
	return &NullableHasNfts200Response{value: val, isSet: true}
}

func (v NullableHasNfts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasNfts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


