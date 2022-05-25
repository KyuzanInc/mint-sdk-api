# ProductERC721

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**Status** | **string** |  | 
**TokenStandardType** | [**TokenStandardType**](TokenStandardType.md) |  | [default to ERC721]
**ContractERC721Id** | **string** |  | 
**TokenId** | **float32** |  | 
**TokenURI** | **NullableString** |  | 
**CreatorAddress** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**ItemStockDocumentId** | **NullableString** |  | 
**ProductBlueprintId** | **NullableString** |  | 
**ProductGroupId** | **NullableString** |  | 

## Methods

### NewProductERC721

`func NewProductERC721(id string, createAt time.Time, updateAt time.Time, status string, tokenStandardType TokenStandardType, contractERC721Id string, tokenId float32, tokenURI NullableString, creatorAddress string, metadata map[string]interface{}, itemStockDocumentId NullableString, productBlueprintId NullableString, productGroupId NullableString, ) *ProductERC721`

NewProductERC721 instantiates a new ProductERC721 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductERC721WithDefaults

`func NewProductERC721WithDefaults() *ProductERC721`

NewProductERC721WithDefaults instantiates a new ProductERC721 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductERC721) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductERC721) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductERC721) SetId(v string)`

SetId sets Id field to given value.


### GetCreateAt

`func (o *ProductERC721) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *ProductERC721) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *ProductERC721) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *ProductERC721) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *ProductERC721) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *ProductERC721) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetStatus

`func (o *ProductERC721) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductERC721) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductERC721) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenStandardType

`func (o *ProductERC721) GetTokenStandardType() TokenStandardType`

GetTokenStandardType returns the TokenStandardType field if non-nil, zero value otherwise.

### GetTokenStandardTypeOk

`func (o *ProductERC721) GetTokenStandardTypeOk() (*TokenStandardType, bool)`

GetTokenStandardTypeOk returns a tuple with the TokenStandardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandardType

`func (o *ProductERC721) SetTokenStandardType(v TokenStandardType)`

SetTokenStandardType sets TokenStandardType field to given value.


### GetContractERC721Id

`func (o *ProductERC721) GetContractERC721Id() string`

GetContractERC721Id returns the ContractERC721Id field if non-nil, zero value otherwise.

### GetContractERC721IdOk

`func (o *ProductERC721) GetContractERC721IdOk() (*string, bool)`

GetContractERC721IdOk returns a tuple with the ContractERC721Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractERC721Id

`func (o *ProductERC721) SetContractERC721Id(v string)`

SetContractERC721Id sets ContractERC721Id field to given value.


### GetTokenId

`func (o *ProductERC721) GetTokenId() float32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ProductERC721) GetTokenIdOk() (*float32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ProductERC721) SetTokenId(v float32)`

SetTokenId sets TokenId field to given value.


### GetTokenURI

`func (o *ProductERC721) GetTokenURI() string`

GetTokenURI returns the TokenURI field if non-nil, zero value otherwise.

### GetTokenURIOk

`func (o *ProductERC721) GetTokenURIOk() (*string, bool)`

GetTokenURIOk returns a tuple with the TokenURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURI

`func (o *ProductERC721) SetTokenURI(v string)`

SetTokenURI sets TokenURI field to given value.


### SetTokenURINil

`func (o *ProductERC721) SetTokenURINil(b bool)`

 SetTokenURINil sets the value for TokenURI to be an explicit nil

### UnsetTokenURI
`func (o *ProductERC721) UnsetTokenURI()`

UnsetTokenURI ensures that no value is present for TokenURI, not even an explicit nil
### GetCreatorAddress

`func (o *ProductERC721) GetCreatorAddress() string`

GetCreatorAddress returns the CreatorAddress field if non-nil, zero value otherwise.

### GetCreatorAddressOk

`func (o *ProductERC721) GetCreatorAddressOk() (*string, bool)`

GetCreatorAddressOk returns a tuple with the CreatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorAddress

`func (o *ProductERC721) SetCreatorAddress(v string)`

SetCreatorAddress sets CreatorAddress field to given value.


### GetMetadata

`func (o *ProductERC721) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductERC721) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductERC721) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *ProductERC721) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ProductERC721) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetItemStockDocumentId

`func (o *ProductERC721) GetItemStockDocumentId() string`

GetItemStockDocumentId returns the ItemStockDocumentId field if non-nil, zero value otherwise.

### GetItemStockDocumentIdOk

`func (o *ProductERC721) GetItemStockDocumentIdOk() (*string, bool)`

GetItemStockDocumentIdOk returns a tuple with the ItemStockDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemStockDocumentId

`func (o *ProductERC721) SetItemStockDocumentId(v string)`

SetItemStockDocumentId sets ItemStockDocumentId field to given value.


### SetItemStockDocumentIdNil

`func (o *ProductERC721) SetItemStockDocumentIdNil(b bool)`

 SetItemStockDocumentIdNil sets the value for ItemStockDocumentId to be an explicit nil

### UnsetItemStockDocumentId
`func (o *ProductERC721) UnsetItemStockDocumentId()`

UnsetItemStockDocumentId ensures that no value is present for ItemStockDocumentId, not even an explicit nil
### GetProductBlueprintId

`func (o *ProductERC721) GetProductBlueprintId() string`

GetProductBlueprintId returns the ProductBlueprintId field if non-nil, zero value otherwise.

### GetProductBlueprintIdOk

`func (o *ProductERC721) GetProductBlueprintIdOk() (*string, bool)`

GetProductBlueprintIdOk returns a tuple with the ProductBlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBlueprintId

`func (o *ProductERC721) SetProductBlueprintId(v string)`

SetProductBlueprintId sets ProductBlueprintId field to given value.


### SetProductBlueprintIdNil

`func (o *ProductERC721) SetProductBlueprintIdNil(b bool)`

 SetProductBlueprintIdNil sets the value for ProductBlueprintId to be an explicit nil

### UnsetProductBlueprintId
`func (o *ProductERC721) UnsetProductBlueprintId()`

UnsetProductBlueprintId ensures that no value is present for ProductBlueprintId, not even an explicit nil
### GetProductGroupId

`func (o *ProductERC721) GetProductGroupId() string`

GetProductGroupId returns the ProductGroupId field if non-nil, zero value otherwise.

### GetProductGroupIdOk

`func (o *ProductERC721) GetProductGroupIdOk() (*string, bool)`

GetProductGroupIdOk returns a tuple with the ProductGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroupId

`func (o *ProductERC721) SetProductGroupId(v string)`

SetProductGroupId sets ProductGroupId field to given value.


### SetProductGroupIdNil

`func (o *ProductERC721) SetProductGroupIdNil(b bool)`

 SetProductGroupIdNil sets the value for ProductGroupId to be an explicit nil

### UnsetProductGroupId
`func (o *ProductERC721) UnsetProductGroupId()`

UnsetProductGroupId ensures that no value is present for ProductGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


