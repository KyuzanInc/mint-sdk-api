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

// SignatureDomain struct for SignatureDomain
type SignatureDomain struct {
	ChainId string `json:"chainId"`
	Name string `json:"name"`
	Version string `json:"version"`
}

// NewSignatureDomain instantiates a new SignatureDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureDomain(chainId string, name string, version string) *SignatureDomain {
	this := SignatureDomain{}
	this.ChainId = chainId
	this.Name = name
	this.Version = version
	return &this
}

// NewSignatureDomainWithDefaults instantiates a new SignatureDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureDomainWithDefaults() *SignatureDomain {
	this := SignatureDomain{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *SignatureDomain) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *SignatureDomain) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *SignatureDomain) SetChainId(v string) {
	o.ChainId = v
}

// GetName returns the Name field value
func (o *SignatureDomain) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignatureDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignatureDomain) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *SignatureDomain) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SignatureDomain) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SignatureDomain) SetVersion(v string) {
	o.Version = v
}

func (o SignatureDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chainId"] = o.ChainId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureDomain struct {
	value *SignatureDomain
	isSet bool
}

func (v NullableSignatureDomain) Get() *SignatureDomain {
	return v.value
}

func (v *NullableSignatureDomain) Set(val *SignatureDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureDomain(val *SignatureDomain) *NullableSignatureDomain {
	return &NullableSignatureDomain{value: val, isSet: true}
}

func (v NullableSignatureDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


