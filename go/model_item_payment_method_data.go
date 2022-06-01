/*
sdk_api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ItemPaymentMethodData - paymentMethodによって異なるデータ
type ItemPaymentMethodData struct {
	ItemPaymentMethodDataCreditCardStripeFixedPrice *ItemPaymentMethodDataCreditCardStripeFixedPrice
	ItemPaymentMethodDataEthereumContractERC721ShopAuction *ItemPaymentMethodDataEthereumContractERC721ShopAuction
	ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice
}

// ItemPaymentMethodDataCreditCardStripeFixedPriceAsItemPaymentMethodData is a convenience function that returns ItemPaymentMethodDataCreditCardStripeFixedPrice wrapped in ItemPaymentMethodData
func ItemPaymentMethodDataCreditCardStripeFixedPriceAsItemPaymentMethodData(v *ItemPaymentMethodDataCreditCardStripeFixedPrice) ItemPaymentMethodData {
	return ItemPaymentMethodData{
		ItemPaymentMethodDataCreditCardStripeFixedPrice: v,
	}
}

// ItemPaymentMethodDataEthereumContractERC721ShopAuctionAsItemPaymentMethodData is a convenience function that returns ItemPaymentMethodDataEthereumContractERC721ShopAuction wrapped in ItemPaymentMethodData
func ItemPaymentMethodDataEthereumContractERC721ShopAuctionAsItemPaymentMethodData(v *ItemPaymentMethodDataEthereumContractERC721ShopAuction) ItemPaymentMethodData {
	return ItemPaymentMethodData{
		ItemPaymentMethodDataEthereumContractERC721ShopAuction: v,
	}
}

// ItemPaymentMethodDataEthereumContractERC721ShopFixedPriceAsItemPaymentMethodData is a convenience function that returns ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice wrapped in ItemPaymentMethodData
func ItemPaymentMethodDataEthereumContractERC721ShopFixedPriceAsItemPaymentMethodData(v *ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) ItemPaymentMethodData {
	return ItemPaymentMethodData{
		ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ItemPaymentMethodData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ItemPaymentMethodDataCreditCardStripeFixedPrice
	err = newStrictDecoder(data).Decode(&dst.ItemPaymentMethodDataCreditCardStripeFixedPrice)
	if err == nil {
		jsonItemPaymentMethodDataCreditCardStripeFixedPrice, _ := json.Marshal(dst.ItemPaymentMethodDataCreditCardStripeFixedPrice)
		if string(jsonItemPaymentMethodDataCreditCardStripeFixedPrice) == "{}" { // empty struct
			dst.ItemPaymentMethodDataCreditCardStripeFixedPrice = nil
		} else {
			match++
		}
	} else {
		dst.ItemPaymentMethodDataCreditCardStripeFixedPrice = nil
	}

	// try to unmarshal data into ItemPaymentMethodDataEthereumContractERC721ShopAuction
	err = newStrictDecoder(data).Decode(&dst.ItemPaymentMethodDataEthereumContractERC721ShopAuction)
	if err == nil {
		jsonItemPaymentMethodDataEthereumContractERC721ShopAuction, _ := json.Marshal(dst.ItemPaymentMethodDataEthereumContractERC721ShopAuction)
		if string(jsonItemPaymentMethodDataEthereumContractERC721ShopAuction) == "{}" { // empty struct
			dst.ItemPaymentMethodDataEthereumContractERC721ShopAuction = nil
		} else {
			match++
		}
	} else {
		dst.ItemPaymentMethodDataEthereumContractERC721ShopAuction = nil
	}

	// try to unmarshal data into ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice
	err = newStrictDecoder(data).Decode(&dst.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice)
	if err == nil {
		jsonItemPaymentMethodDataEthereumContractERC721ShopFixedPrice, _ := json.Marshal(dst.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice)
		if string(jsonItemPaymentMethodDataEthereumContractERC721ShopFixedPrice) == "{}" { // empty struct
			dst.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice = nil
		} else {
			match++
		}
	} else {
		dst.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ItemPaymentMethodDataCreditCardStripeFixedPrice = nil
		dst.ItemPaymentMethodDataEthereumContractERC721ShopAuction = nil
		dst.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ItemPaymentMethodData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ItemPaymentMethodData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ItemPaymentMethodData) MarshalJSON() ([]byte, error) {
	if src.ItemPaymentMethodDataCreditCardStripeFixedPrice != nil {
		return json.Marshal(&src.ItemPaymentMethodDataCreditCardStripeFixedPrice)
	}

	if src.ItemPaymentMethodDataEthereumContractERC721ShopAuction != nil {
		return json.Marshal(&src.ItemPaymentMethodDataEthereumContractERC721ShopAuction)
	}

	if src.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice != nil {
		return json.Marshal(&src.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ItemPaymentMethodData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ItemPaymentMethodDataCreditCardStripeFixedPrice != nil {
		return obj.ItemPaymentMethodDataCreditCardStripeFixedPrice
	}

	if obj.ItemPaymentMethodDataEthereumContractERC721ShopAuction != nil {
		return obj.ItemPaymentMethodDataEthereumContractERC721ShopAuction
	}

	if obj.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice != nil {
		return obj.ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice
	}

	// all schemas are nil
	return nil
}

type NullableItemPaymentMethodData struct {
	value *ItemPaymentMethodData
	isSet bool
}

func (v NullableItemPaymentMethodData) Get() *ItemPaymentMethodData {
	return v.value
}

func (v *NullableItemPaymentMethodData) Set(val *ItemPaymentMethodData) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPaymentMethodData) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPaymentMethodData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPaymentMethodData(val *ItemPaymentMethodData) *NullableItemPaymentMethodData {
	return &NullableItemPaymentMethodData{value: val, isSet: true}
}

func (v NullableItemPaymentMethodData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPaymentMethodData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


