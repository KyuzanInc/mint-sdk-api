# CryptoCurrencyRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**Currency** | [**CryptoCurrencyType**](CryptoCurrencyType.md) |  | 
**Jpy** | **float32** |  | 
**Eur** | **float32** |  | 
**Usd** | **float32** |  | 

## Methods

### NewCryptoCurrencyRate

`func NewCryptoCurrencyRate(createAt time.Time, updateAt time.Time, currency CryptoCurrencyType, jpy float32, eur float32, usd float32, ) *CryptoCurrencyRate`

NewCryptoCurrencyRate instantiates a new CryptoCurrencyRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoCurrencyRateWithDefaults

`func NewCryptoCurrencyRateWithDefaults() *CryptoCurrencyRate`

NewCryptoCurrencyRateWithDefaults instantiates a new CryptoCurrencyRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateAt

`func (o *CryptoCurrencyRate) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *CryptoCurrencyRate) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *CryptoCurrencyRate) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *CryptoCurrencyRate) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *CryptoCurrencyRate) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *CryptoCurrencyRate) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetCurrency

`func (o *CryptoCurrencyRate) GetCurrency() CryptoCurrencyType`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CryptoCurrencyRate) GetCurrencyOk() (*CryptoCurrencyType, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CryptoCurrencyRate) SetCurrency(v CryptoCurrencyType)`

SetCurrency sets Currency field to given value.


### GetJpy

`func (o *CryptoCurrencyRate) GetJpy() float32`

GetJpy returns the Jpy field if non-nil, zero value otherwise.

### GetJpyOk

`func (o *CryptoCurrencyRate) GetJpyOk() (*float32, bool)`

GetJpyOk returns a tuple with the Jpy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpy

`func (o *CryptoCurrencyRate) SetJpy(v float32)`

SetJpy sets Jpy field to given value.


### GetEur

`func (o *CryptoCurrencyRate) GetEur() float32`

GetEur returns the Eur field if non-nil, zero value otherwise.

### GetEurOk

`func (o *CryptoCurrencyRate) GetEurOk() (*float32, bool)`

GetEurOk returns a tuple with the Eur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEur

`func (o *CryptoCurrencyRate) SetEur(v float32)`

SetEur sets Eur field to given value.


### GetUsd

`func (o *CryptoCurrencyRate) GetUsd() float32`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *CryptoCurrencyRate) GetUsdOk() (*float32, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *CryptoCurrencyRate) SetUsd(v float32)`

SetUsd sets Usd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


