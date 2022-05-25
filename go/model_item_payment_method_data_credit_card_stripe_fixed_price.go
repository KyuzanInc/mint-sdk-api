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

// ItemPaymentMethodDataCreditCardStripeFixedPrice struct for ItemPaymentMethodDataCreditCardStripeFixedPrice
type ItemPaymentMethodDataCreditCardStripeFixedPrice struct {
	PaymentMethod string `json:"paymentMethod"`
	Currency CreditCardStripeCurrencyType `json:"currency"`
}

// NewItemPaymentMethodDataCreditCardStripeFixedPrice instantiates a new ItemPaymentMethodDataCreditCardStripeFixedPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPaymentMethodDataCreditCardStripeFixedPrice(paymentMethod string, currency CreditCardStripeCurrencyType) *ItemPaymentMethodDataCreditCardStripeFixedPrice {
	this := ItemPaymentMethodDataCreditCardStripeFixedPrice{}
	this.PaymentMethod = paymentMethod
	this.Currency = currency
	return &this
}

// NewItemPaymentMethodDataCreditCardStripeFixedPriceWithDefaults instantiates a new ItemPaymentMethodDataCreditCardStripeFixedPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPaymentMethodDataCreditCardStripeFixedPriceWithDefaults() *ItemPaymentMethodDataCreditCardStripeFixedPrice {
	this := ItemPaymentMethodDataCreditCardStripeFixedPrice{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) GetPaymentMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetCurrency returns the Currency field value
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) GetCurrency() CreditCardStripeCurrencyType {
	if o == nil {
		var ret CreditCardStripeCurrencyType
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) GetCurrencyOk() (*CreditCardStripeCurrencyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ItemPaymentMethodDataCreditCardStripeFixedPrice) SetCurrency(v CreditCardStripeCurrencyType) {
	o.Currency = v
}

func (o ItemPaymentMethodDataCreditCardStripeFixedPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableItemPaymentMethodDataCreditCardStripeFixedPrice struct {
	value *ItemPaymentMethodDataCreditCardStripeFixedPrice
	isSet bool
}

func (v NullableItemPaymentMethodDataCreditCardStripeFixedPrice) Get() *ItemPaymentMethodDataCreditCardStripeFixedPrice {
	return v.value
}

func (v *NullableItemPaymentMethodDataCreditCardStripeFixedPrice) Set(val *ItemPaymentMethodDataCreditCardStripeFixedPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPaymentMethodDataCreditCardStripeFixedPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPaymentMethodDataCreditCardStripeFixedPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPaymentMethodDataCreditCardStripeFixedPrice(val *ItemPaymentMethodDataCreditCardStripeFixedPrice) *NullableItemPaymentMethodDataCreditCardStripeFixedPrice {
	return &NullableItemPaymentMethodDataCreditCardStripeFixedPrice{value: val, isSet: true}
}

func (v NullableItemPaymentMethodDataCreditCardStripeFixedPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPaymentMethodDataCreditCardStripeFixedPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


