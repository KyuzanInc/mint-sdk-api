# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Previews** | [**[]PreviewMedia**](PreviewMedia.md) | Itemのプレビュー用URL | 
**Type** | [**ItemType**](ItemType.md) |  | 
**StartAt** | **time.Time** |  | 
**EndAt** | **time.Time** |  | 
**Price** | **float32** |  | 
**CryptoCurrencyRate** | [**CryptoCurrencyRate**](CryptoCurrencyRate.md) |  | 
**FeeRatePermill** | **float32** | Mintに支払われる取引手数料 | 
**Tags** | **[]string** | 任意のTag | 
**PaymentMethodData** | [**OneOfItemPaymentMethodDataEthereumContractERC721ShopFixedPriceItemPaymentMethodDataEthereumContractERC721ShopAuctionItemPaymentMethodDataCreditCardStripeFixedPrice**](oneOf&lt;ItemPaymentMethodDataEthereumContractERC721ShopFixedPrice,ItemPaymentMethodDataEthereumContractERC721ShopAuction,ItemPaymentMethodDataCreditCardStripeFixedPrice&gt;.md) | paymentMethodによって異なるデータ | 
**ItemStockIds** | **[]string** |  | 
**AvailableStockNum** | **float32** |  | 
**ProductERC721Ids** | **[]string** |  | 
**Bids** | [**[]Bid**](Bid.md) | オークションItem以外は空配列が入る | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewItem

`func NewItem(id string, createAt time.Time, updateAt time.Time, name string, description string, previews []PreviewMedia, type_ ItemType, startAt time.Time, endAt time.Time, price float32, cryptoCurrencyRate CryptoCurrencyRate, feeRatePermill float32, tags []string, paymentMethodData OneOfItemPaymentMethodDataEthereumContractERC721ShopFixedPriceItemPaymentMethodDataEthereumContractERC721ShopAuctionItemPaymentMethodDataCreditCardStripeFixedPrice, itemStockIds []string, availableStockNum float32, productERC721Ids []string, bids []Bid, metadata map[string]interface{}, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.


### GetCreateAt

`func (o *Item) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *Item) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *Item) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *Item) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *Item) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *Item) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPreviews

`func (o *Item) GetPreviews() []PreviewMedia`

GetPreviews returns the Previews field if non-nil, zero value otherwise.

### GetPreviewsOk

`func (o *Item) GetPreviewsOk() (*[]PreviewMedia, bool)`

GetPreviewsOk returns a tuple with the Previews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviews

`func (o *Item) SetPreviews(v []PreviewMedia)`

SetPreviews sets Previews field to given value.


### GetType

`func (o *Item) GetType() ItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Item) GetTypeOk() (*ItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Item) SetType(v ItemType)`

SetType sets Type field to given value.


### GetStartAt

`func (o *Item) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Item) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Item) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *Item) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Item) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Item) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### GetPrice

`func (o *Item) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Item) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Item) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCryptoCurrencyRate

`func (o *Item) GetCryptoCurrencyRate() CryptoCurrencyRate`

GetCryptoCurrencyRate returns the CryptoCurrencyRate field if non-nil, zero value otherwise.

### GetCryptoCurrencyRateOk

`func (o *Item) GetCryptoCurrencyRateOk() (*CryptoCurrencyRate, bool)`

GetCryptoCurrencyRateOk returns a tuple with the CryptoCurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyRate

`func (o *Item) SetCryptoCurrencyRate(v CryptoCurrencyRate)`

SetCryptoCurrencyRate sets CryptoCurrencyRate field to given value.


### GetFeeRatePermill

`func (o *Item) GetFeeRatePermill() float32`

GetFeeRatePermill returns the FeeRatePermill field if non-nil, zero value otherwise.

### GetFeeRatePermillOk

`func (o *Item) GetFeeRatePermillOk() (*float32, bool)`

GetFeeRatePermillOk returns a tuple with the FeeRatePermill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRatePermill

`func (o *Item) SetFeeRatePermill(v float32)`

SetFeeRatePermill sets FeeRatePermill field to given value.


### GetTags

`func (o *Item) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Item) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Item) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetPaymentMethodData

`func (o *Item) GetPaymentMethodData() OneOfItemPaymentMethodDataEthereumContractERC721ShopFixedPriceItemPaymentMethodDataEthereumContractERC721ShopAuctionItemPaymentMethodDataCreditCardStripeFixedPrice`

GetPaymentMethodData returns the PaymentMethodData field if non-nil, zero value otherwise.

### GetPaymentMethodDataOk

`func (o *Item) GetPaymentMethodDataOk() (*OneOfItemPaymentMethodDataEthereumContractERC721ShopFixedPriceItemPaymentMethodDataEthereumContractERC721ShopAuctionItemPaymentMethodDataCreditCardStripeFixedPrice, bool)`

GetPaymentMethodDataOk returns a tuple with the PaymentMethodData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodData

`func (o *Item) SetPaymentMethodData(v OneOfItemPaymentMethodDataEthereumContractERC721ShopFixedPriceItemPaymentMethodDataEthereumContractERC721ShopAuctionItemPaymentMethodDataCreditCardStripeFixedPrice)`

SetPaymentMethodData sets PaymentMethodData field to given value.


### GetItemStockIds

`func (o *Item) GetItemStockIds() []string`

GetItemStockIds returns the ItemStockIds field if non-nil, zero value otherwise.

### GetItemStockIdsOk

`func (o *Item) GetItemStockIdsOk() (*[]string, bool)`

GetItemStockIdsOk returns a tuple with the ItemStockIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemStockIds

`func (o *Item) SetItemStockIds(v []string)`

SetItemStockIds sets ItemStockIds field to given value.


### GetAvailableStockNum

`func (o *Item) GetAvailableStockNum() float32`

GetAvailableStockNum returns the AvailableStockNum field if non-nil, zero value otherwise.

### GetAvailableStockNumOk

`func (o *Item) GetAvailableStockNumOk() (*float32, bool)`

GetAvailableStockNumOk returns a tuple with the AvailableStockNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStockNum

`func (o *Item) SetAvailableStockNum(v float32)`

SetAvailableStockNum sets AvailableStockNum field to given value.


### GetProductERC721Ids

`func (o *Item) GetProductERC721Ids() []string`

GetProductERC721Ids returns the ProductERC721Ids field if non-nil, zero value otherwise.

### GetProductERC721IdsOk

`func (o *Item) GetProductERC721IdsOk() (*[]string, bool)`

GetProductERC721IdsOk returns a tuple with the ProductERC721Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductERC721Ids

`func (o *Item) SetProductERC721Ids(v []string)`

SetProductERC721Ids sets ProductERC721Ids field to given value.


### GetBids

`func (o *Item) GetBids() []Bid`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *Item) GetBidsOk() (*[]Bid, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *Item) SetBids(v []Bid)`

SetBids sets Bids field to given value.


### GetMetadata

`func (o *Item) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Item) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Item) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *Item) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Item) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


