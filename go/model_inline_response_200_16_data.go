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

// InlineResponse20016Data struct for InlineResponse20016Data
type InlineResponse20016Data struct {
	Profile *WalletAddressProfile `json:"profile,omitempty"`
}

// NewInlineResponse20016Data instantiates a new InlineResponse20016Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016Data() *InlineResponse20016Data {
	this := InlineResponse20016Data{}
	return &this
}

// NewInlineResponse20016DataWithDefaults instantiates a new InlineResponse20016Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016DataWithDefaults() *InlineResponse20016Data {
	this := InlineResponse20016Data{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *InlineResponse20016Data) GetProfile() WalletAddressProfile {
	if o == nil || o.Profile == nil {
		var ret WalletAddressProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Data) GetProfileOk() (*WalletAddressProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *InlineResponse20016Data) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given WalletAddressProfile and assigns it to the Profile field.
func (o *InlineResponse20016Data) SetProfile(v WalletAddressProfile) {
	o.Profile = &v
}

func (o InlineResponse20016Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016Data struct {
	value *InlineResponse20016Data
	isSet bool
}

func (v NullableInlineResponse20016Data) Get() *InlineResponse20016Data {
	return v.value
}

func (v *NullableInlineResponse20016Data) Set(val *InlineResponse20016Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016Data(val *InlineResponse20016Data) *NullableInlineResponse20016Data {
	return &NullableInlineResponse20016Data{value: val, isSet: true}
}

func (v NullableInlineResponse20016Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


