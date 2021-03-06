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

// UpdateProfile200Response struct for UpdateProfile200Response
type UpdateProfile200Response struct {
	Data *UpdateProfile200ResponseData `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// NewUpdateProfile200Response instantiates a new UpdateProfile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfile200Response() *UpdateProfile200Response {
	this := UpdateProfile200Response{}
	return &this
}

// NewUpdateProfile200ResponseWithDefaults instantiates a new UpdateProfile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfile200ResponseWithDefaults() *UpdateProfile200Response {
	this := UpdateProfile200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateProfile200Response) GetData() UpdateProfile200ResponseData {
	if o == nil || o.Data == nil {
		var ret UpdateProfile200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfile200Response) GetDataOk() (*UpdateProfile200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateProfile200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateProfile200ResponseData and assigns it to the Data field.
func (o *UpdateProfile200Response) SetData(v UpdateProfile200ResponseData) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UpdateProfile200Response) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfile200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UpdateProfile200Response) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *UpdateProfile200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o UpdateProfile200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfile200Response struct {
	value *UpdateProfile200Response
	isSet bool
}

func (v NullableUpdateProfile200Response) Get() *UpdateProfile200Response {
	return v.value
}

func (v *NullableUpdateProfile200Response) Set(val *UpdateProfile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfile200Response(val *UpdateProfile200Response) *NullableUpdateProfile200Response {
	return &NullableUpdateProfile200Response{value: val, isSet: true}
}

func (v NullableUpdateProfile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


