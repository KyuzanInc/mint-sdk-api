# ItemPaymentMethodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | **string** |  | 
**ContractDataERC721Shop** | [**ContractDataERC721Shop**](ContractDataERC721Shop.md) |  | 
**MinBidPercentage** | **float32** |  | 
**DefaultEndAt** | **time.Time** |  | 
**Currency** | [**CreditCardStripeCurrencyType**](CreditCardStripeCurrencyType.md) |  | 

## Methods

### NewItemPaymentMethodData

`func NewItemPaymentMethodData(paymentMethod string, contractDataERC721Shop ContractDataERC721Shop, minBidPercentage float32, defaultEndAt time.Time, currency CreditCardStripeCurrencyType, ) *ItemPaymentMethodData`

NewItemPaymentMethodData instantiates a new ItemPaymentMethodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPaymentMethodDataWithDefaults

`func NewItemPaymentMethodDataWithDefaults() *ItemPaymentMethodData`

NewItemPaymentMethodDataWithDefaults instantiates a new ItemPaymentMethodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *ItemPaymentMethodData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ItemPaymentMethodData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ItemPaymentMethodData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetContractDataERC721Shop

`func (o *ItemPaymentMethodData) GetContractDataERC721Shop() ContractDataERC721Shop`

GetContractDataERC721Shop returns the ContractDataERC721Shop field if non-nil, zero value otherwise.

### GetContractDataERC721ShopOk

`func (o *ItemPaymentMethodData) GetContractDataERC721ShopOk() (*ContractDataERC721Shop, bool)`

GetContractDataERC721ShopOk returns a tuple with the ContractDataERC721Shop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDataERC721Shop

`func (o *ItemPaymentMethodData) SetContractDataERC721Shop(v ContractDataERC721Shop)`

SetContractDataERC721Shop sets ContractDataERC721Shop field to given value.


### GetMinBidPercentage

`func (o *ItemPaymentMethodData) GetMinBidPercentage() float32`

GetMinBidPercentage returns the MinBidPercentage field if non-nil, zero value otherwise.

### GetMinBidPercentageOk

`func (o *ItemPaymentMethodData) GetMinBidPercentageOk() (*float32, bool)`

GetMinBidPercentageOk returns a tuple with the MinBidPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBidPercentage

`func (o *ItemPaymentMethodData) SetMinBidPercentage(v float32)`

SetMinBidPercentage sets MinBidPercentage field to given value.


### GetDefaultEndAt

`func (o *ItemPaymentMethodData) GetDefaultEndAt() time.Time`

GetDefaultEndAt returns the DefaultEndAt field if non-nil, zero value otherwise.

### GetDefaultEndAtOk

`func (o *ItemPaymentMethodData) GetDefaultEndAtOk() (*time.Time, bool)`

GetDefaultEndAtOk returns a tuple with the DefaultEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEndAt

`func (o *ItemPaymentMethodData) SetDefaultEndAt(v time.Time)`

SetDefaultEndAt sets DefaultEndAt field to given value.


### GetCurrency

`func (o *ItemPaymentMethodData) GetCurrency() CreditCardStripeCurrencyType`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ItemPaymentMethodData) GetCurrencyOk() (*CreditCardStripeCurrencyType, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ItemPaymentMethodData) SetCurrency(v CreditCardStripeCurrencyType)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


