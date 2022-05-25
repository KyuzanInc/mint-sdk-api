# ItemStock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**Status** | [**ItemStockStatus**](ItemStockStatus.md) |  | 
**ProductsData** | [**[]ProductERC721**](ProductERC721.md) |  | 
**Item** | [**Item**](Item.md) |  | 

## Methods

### NewItemStock

`func NewItemStock(id string, createAt time.Time, updateAt time.Time, status ItemStockStatus, productsData []ProductERC721, item Item, ) *ItemStock`

NewItemStock instantiates a new ItemStock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStockWithDefaults

`func NewItemStockWithDefaults() *ItemStock`

NewItemStockWithDefaults instantiates a new ItemStock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemStock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemStock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemStock) SetId(v string)`

SetId sets Id field to given value.


### GetCreateAt

`func (o *ItemStock) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *ItemStock) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *ItemStock) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *ItemStock) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *ItemStock) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *ItemStock) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetStatus

`func (o *ItemStock) GetStatus() ItemStockStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemStock) GetStatusOk() (*ItemStockStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemStock) SetStatus(v ItemStockStatus)`

SetStatus sets Status field to given value.


### GetProductsData

`func (o *ItemStock) GetProductsData() []ProductERC721`

GetProductsData returns the ProductsData field if non-nil, zero value otherwise.

### GetProductsDataOk

`func (o *ItemStock) GetProductsDataOk() (*[]ProductERC721, bool)`

GetProductsDataOk returns a tuple with the ProductsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsData

`func (o *ItemStock) SetProductsData(v []ProductERC721)`

SetProductsData sets ProductsData field to given value.


### GetItem

`func (o *ItemStock) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemStock) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemStock) SetItem(v Item)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


