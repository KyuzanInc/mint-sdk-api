# Bid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ChainType** | [**ChainType**](ChainType.md) |  | 
**NetworkId** | [**NetworkId**](NetworkId.md) |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**ItemDocumentId** | **string** |  | 
**Bidder** | **string** |  | 
**BidPrice** | **float32** |  | 
**CryptoCurrencyRate** | [**CryptoCurrencyRate**](CryptoCurrencyRate.md) |  | 
**TransactionAt** | **time.Time** |  | 
**TransactionHash** | **string** |  | 

## Methods

### NewBid

`func NewBid(id string, chainType ChainType, networkId NetworkId, createAt time.Time, updateAt time.Time, itemDocumentId string, bidder string, bidPrice float32, cryptoCurrencyRate CryptoCurrencyRate, transactionAt time.Time, transactionHash string, ) *Bid`

NewBid instantiates a new Bid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidWithDefaults

`func NewBidWithDefaults() *Bid`

NewBidWithDefaults instantiates a new Bid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bid) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bid) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bid) SetId(v string)`

SetId sets Id field to given value.


### GetChainType

`func (o *Bid) GetChainType() ChainType`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *Bid) GetChainTypeOk() (*ChainType, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *Bid) SetChainType(v ChainType)`

SetChainType sets ChainType field to given value.


### GetNetworkId

`func (o *Bid) GetNetworkId() NetworkId`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *Bid) GetNetworkIdOk() (*NetworkId, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *Bid) SetNetworkId(v NetworkId)`

SetNetworkId sets NetworkId field to given value.


### GetCreateAt

`func (o *Bid) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *Bid) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *Bid) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *Bid) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *Bid) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *Bid) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetItemDocumentId

`func (o *Bid) GetItemDocumentId() string`

GetItemDocumentId returns the ItemDocumentId field if non-nil, zero value otherwise.

### GetItemDocumentIdOk

`func (o *Bid) GetItemDocumentIdOk() (*string, bool)`

GetItemDocumentIdOk returns a tuple with the ItemDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDocumentId

`func (o *Bid) SetItemDocumentId(v string)`

SetItemDocumentId sets ItemDocumentId field to given value.


### GetBidder

`func (o *Bid) GetBidder() string`

GetBidder returns the Bidder field if non-nil, zero value otherwise.

### GetBidderOk

`func (o *Bid) GetBidderOk() (*string, bool)`

GetBidderOk returns a tuple with the Bidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidder

`func (o *Bid) SetBidder(v string)`

SetBidder sets Bidder field to given value.


### GetBidPrice

`func (o *Bid) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Bid) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Bid) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.


### GetCryptoCurrencyRate

`func (o *Bid) GetCryptoCurrencyRate() CryptoCurrencyRate`

GetCryptoCurrencyRate returns the CryptoCurrencyRate field if non-nil, zero value otherwise.

### GetCryptoCurrencyRateOk

`func (o *Bid) GetCryptoCurrencyRateOk() (*CryptoCurrencyRate, bool)`

GetCryptoCurrencyRateOk returns a tuple with the CryptoCurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyRate

`func (o *Bid) SetCryptoCurrencyRate(v CryptoCurrencyRate)`

SetCryptoCurrencyRate sets CryptoCurrencyRate field to given value.


### GetTransactionAt

`func (o *Bid) GetTransactionAt() time.Time`

GetTransactionAt returns the TransactionAt field if non-nil, zero value otherwise.

### GetTransactionAtOk

`func (o *Bid) GetTransactionAtOk() (*time.Time, bool)`

GetTransactionAtOk returns a tuple with the TransactionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAt

`func (o *Bid) SetTransactionAt(v time.Time)`

SetTransactionAt sets TransactionAt field to given value.


### GetTransactionHash

`func (o *Bid) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Bid) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Bid) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


