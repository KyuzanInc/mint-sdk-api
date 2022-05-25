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

// SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData struct for SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData
type SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData struct {
	Domain SignatureDomain `json:"domain"`
	PrimaryType string `json:"primaryType"`
	Message CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage `json:"message"`
	Types map[string]interface{} `json:"types"`
}

// NewSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData instantiates a new SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData(domain SignatureDomain, primaryType string, message CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage, types map[string]interface{}) *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData {
	this := SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData{}
	this.Domain = domain
	this.PrimaryType = primaryType
	this.Message = message
	this.Types = types
	return &this
}

// NewSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoDataWithDefaults instantiates a new SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoDataWithDefaults() *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData {
	this := SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetDomain() SignatureDomain {
	if o == nil {
		var ret SignatureDomain
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetDomainOk() (*SignatureDomain, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) SetDomain(v SignatureDomain) {
	o.Domain = v
}

// GetPrimaryType returns the PrimaryType field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetPrimaryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryType
}

// GetPrimaryTypeOk returns a tuple with the PrimaryType field value
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetPrimaryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrimaryType, true
}

// SetPrimaryType sets field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) SetPrimaryType(v string) {
	o.PrimaryType = v
}

// GetMessage returns the Message field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetMessage() CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage {
	if o == nil {
		var ret CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetMessageOk() (*CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) SetMessage(v CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) {
	o.Message = v
}

// GetTypes returns the Types field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetTypes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) GetTypesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) SetTypes(v map[string]interface{}) {
	o.Types = v
}

func (o SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["primaryType"] = o.PrimaryType
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData struct {
	value *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData
	isSet bool
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) Get() *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData {
	return v.value
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) Set(val *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData(val *SdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) *NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData {
	return &NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData{value: val, isSet: true}
}

func (v NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

