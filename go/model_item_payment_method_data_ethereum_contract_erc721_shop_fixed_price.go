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

// ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice struct for ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice
type ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice struct {
	PaymentMethod string `json:"paymentMethod"`
	ContractDataERC721Shop ContractDataERC721Shop `json:"contractDataERC721Shop"`
}

// NewItemPaymentMethodDataEthereumContractERC721ShopFixedPrice instantiates a new ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPaymentMethodDataEthereumContractERC721ShopFixedPrice(paymentMethod string, contractDataERC721Shop ContractDataERC721Shop) *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice {
	this := ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice{}
	this.PaymentMethod = paymentMethod
	this.ContractDataERC721Shop = contractDataERC721Shop
	return &this
}

// NewItemPaymentMethodDataEthereumContractERC721ShopFixedPriceWithDefaults instantiates a new ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPaymentMethodDataEthereumContractERC721ShopFixedPriceWithDefaults() *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice {
	this := ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetContractDataERC721Shop returns the ContractDataERC721Shop field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) GetContractDataERC721Shop() ContractDataERC721Shop {
	if o == nil {
		var ret ContractDataERC721Shop
		return ret
	}

	return o.ContractDataERC721Shop
}

// GetContractDataERC721ShopOk returns a tuple with the ContractDataERC721Shop field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) GetContractDataERC721ShopOk() (*ContractDataERC721Shop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractDataERC721Shop, true
}

// SetContractDataERC721Shop sets field value
func (o *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) SetContractDataERC721Shop(v ContractDataERC721Shop) {
	o.ContractDataERC721Shop = v
}

func (o ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if true {
		toSerialize["contractDataERC721Shop"] = o.ContractDataERC721Shop
	}
	return json.Marshal(toSerialize)
}

type NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice struct {
	value *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice
	isSet bool
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) Get() *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice {
	return v.value
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) Set(val *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice(val *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) *NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice {
	return &NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice{value: val, isSet: true}
}

func (v NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


