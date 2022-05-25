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
	"time"
)

// ItemPaymentMethodDataEthereumContractERC721ShopAuction struct for ItemPaymentMethodDataEthereumContractERC721ShopAuction
type ItemPaymentMethodDataEthereumContractERC721ShopAuction struct {
	PaymentMethod string `json:"paymentMethod"`
	ContractDataERC721Shop ContractDataERC721Shop `json:"contractDataERC721Shop"`
	MinBidPercentage float32 `json:"minBidPercentage"`
	DefaultEndAt time.Time `json:"defaultEndAt"`
}

// NewItemPaymentMethodDataEthereumContractERC721ShopAuction instantiates a new ItemPaymentMethodDataEthereumContractERC721ShopAuction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPaymentMethodDataEthereumContractERC721ShopAuction(paymentMethod string, contractDataERC721Shop ContractDataERC721Shop, minBidPercentage float32, defaultEndAt time.Time) *ItemPaymentMethodDataEthereumContractERC721ShopAuction {
	this := ItemPaymentMethodDataEthereumContractERC721ShopAuction{}
	this.PaymentMethod = paymentMethod
	this.ContractDataERC721Shop = contractDataERC721Shop
	this.MinBidPercentage = minBidPercentage
	this.DefaultEndAt = defaultEndAt
	return &this
}

// NewItemPaymentMethodDataEthereumContractERC721ShopAuctionWithDefaults instantiates a new ItemPaymentMethodDataEthereumContractERC721ShopAuction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPaymentMethodDataEthereumContractERC721ShopAuctionWithDefaults() *ItemPaymentMethodDataEthereumContractERC721ShopAuction {
	this := ItemPaymentMethodDataEthereumContractERC721ShopAuction{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetPaymentMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetContractDataERC721Shop returns the ContractDataERC721Shop field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetContractDataERC721Shop() ContractDataERC721Shop {
	if o == nil {
		var ret ContractDataERC721Shop
		return ret
	}

	return o.ContractDataERC721Shop
}

// GetContractDataERC721ShopOk returns a tuple with the ContractDataERC721Shop field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetContractDataERC721ShopOk() (*ContractDataERC721Shop, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractDataERC721Shop, true
}

// SetContractDataERC721Shop sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) SetContractDataERC721Shop(v ContractDataERC721Shop) {
	o.ContractDataERC721Shop = v
}

// GetMinBidPercentage returns the MinBidPercentage field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetMinBidPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinBidPercentage
}

// GetMinBidPercentageOk returns a tuple with the MinBidPercentage field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetMinBidPercentageOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinBidPercentage, true
}

// SetMinBidPercentage sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) SetMinBidPercentage(v float32) {
	o.MinBidPercentage = v
}

// GetDefaultEndAt returns the DefaultEndAt field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetDefaultEndAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DefaultEndAt
}

// GetDefaultEndAtOk returns a tuple with the DefaultEndAt field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) GetDefaultEndAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultEndAt, true
}

// SetDefaultEndAt sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopAuction) SetDefaultEndAt(v time.Time) {
	o.DefaultEndAt = v
}

func (o ItemPaymentMethodDataEthereumContractERC721ShopAuction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if true {
		toSerialize["contractDataERC721Shop"] = o.ContractDataERC721Shop
	}
	if true {
		toSerialize["minBidPercentage"] = o.MinBidPercentage
	}
	if true {
		toSerialize["defaultEndAt"] = o.DefaultEndAt
	}
	return json.Marshal(toSerialize)
}

type NullableItemPaymentMethodDataEthereumContractERC721ShopAuction struct {
	value *ItemPaymentMethodDataEthereumContractERC721ShopAuction
	isSet bool
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) Get() *ItemPaymentMethodDataEthereumContractERC721ShopAuction {
	return v.value
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) Set(val *ItemPaymentMethodDataEthereumContractERC721ShopAuction) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPaymentMethodDataEthereumContractERC721ShopAuction(val *ItemPaymentMethodDataEthereumContractERC721ShopAuction) *NullableItemPaymentMethodDataEthereumContractERC721ShopAuction {
	return &NullableItemPaymentMethodDataEthereumContractERC721ShopAuction{value: val, isSet: true}
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopAuction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

