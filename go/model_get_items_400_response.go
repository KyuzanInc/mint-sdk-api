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

// GetItems400Response struct for GetItems400Response
type GetItems400Response struct {
	Message string `json:"message"`
}

// NewGetItems400Response instantiates a new GetItems400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetItems400Response(message string) *GetItems400Response {
	this := GetItems400Response{}
	this.Message = message
	return &this
}

// NewGetItems400ResponseWithDefaults instantiates a new GetItems400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetItems400ResponseWithDefaults() *GetItems400Response {
	this := GetItems400Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *GetItems400Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetItems400Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetItems400Response) SetMessage(v string) {
	o.Message = v
}

func (o GetItems400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableGetItems400Response struct {
	value *GetItems400Response
	isSet bool
}

func (v NullableGetItems400Response) Get() *GetItems400Response {
	return v.value
}

func (v *NullableGetItems400Response) Set(val *GetItems400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetItems400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetItems400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetItems400Response(val *GetItems400Response) *NullableGetItems400Response {
	return &NullableGetItems400Response{value: val, isSet: true}
}

func (v NullableGetItems400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetItems400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


