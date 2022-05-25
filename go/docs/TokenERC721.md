# TokenERC721

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**TokenStandardType** | [**TokenStandardType**](TokenStandardType.md) |  | [default to ERC721]
**ContractERC721Id** | **string** |  | 
**TokenId** | **float32** |  | 
**TokenURI** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**MintTransactionHash** | **string** |  | 
**InitialOwnerAddress** | **string** |  | 
**CurrentOwnerAddress** | **string** |  | 
**TransferHistory** | [**[]TransferData**](TransferData.md) |  | 

## Methods

### NewTokenERC721

`func NewTokenERC721(id string, createAt time.Time, updateAt time.Time, tokenStandardType TokenStandardType, contractERC721Id string, tokenId float32, tokenURI string, metadata map[string]interface{}, mintTransactionHash string, initialOwnerAddress string, currentOwnerAddress string, transferHistory []TransferData, ) *TokenERC721`

NewTokenERC721 instantiates a new TokenERC721 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenERC721WithDefaults

`func NewTokenERC721WithDefaults() *TokenERC721`

NewTokenERC721WithDefaults instantiates a new TokenERC721 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenERC721) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenERC721) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenERC721) SetId(v string)`

SetId sets Id field to given value.


### GetCreateAt

`func (o *TokenERC721) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *TokenERC721) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *TokenERC721) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *TokenERC721) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *TokenERC721) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *TokenERC721) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetTokenStandardType

`func (o *TokenERC721) GetTokenStandardType() TokenStandardType`

GetTokenStandardType returns the TokenStandardType field if non-nil, zero value otherwise.

### GetTokenStandardTypeOk

`func (o *TokenERC721) GetTokenStandardTypeOk() (*TokenStandardType, bool)`

GetTokenStandardTypeOk returns a tuple with the TokenStandardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandardType

`func (o *TokenERC721) SetTokenStandardType(v TokenStandardType)`

SetTokenStandardType sets TokenStandardType field to given value.


### GetContractERC721Id

`func (o *TokenERC721) GetContractERC721Id() string`

GetContractERC721Id returns the ContractERC721Id field if non-nil, zero value otherwise.

### GetContractERC721IdOk

`func (o *TokenERC721) GetContractERC721IdOk() (*string, bool)`

GetContractERC721IdOk returns a tuple with the ContractERC721Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractERC721Id

`func (o *TokenERC721) SetContractERC721Id(v string)`

SetContractERC721Id sets ContractERC721Id field to given value.


### GetTokenId

`func (o *TokenERC721) GetTokenId() float32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenERC721) GetTokenIdOk() (*float32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenERC721) SetTokenId(v float32)`

SetTokenId sets TokenId field to given value.


### GetTokenURI

`func (o *TokenERC721) GetTokenURI() string`

GetTokenURI returns the TokenURI field if non-nil, zero value otherwise.

### GetTokenURIOk

`func (o *TokenERC721) GetTokenURIOk() (*string, bool)`

GetTokenURIOk returns a tuple with the TokenURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURI

`func (o *TokenERC721) SetTokenURI(v string)`

SetTokenURI sets TokenURI field to given value.


### GetMetadata

`func (o *TokenERC721) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenERC721) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenERC721) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMintTransactionHash

`func (o *TokenERC721) GetMintTransactionHash() string`

GetMintTransactionHash returns the MintTransactionHash field if non-nil, zero value otherwise.

### GetMintTransactionHashOk

`func (o *TokenERC721) GetMintTransactionHashOk() (*string, bool)`

GetMintTransactionHashOk returns a tuple with the MintTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintTransactionHash

`func (o *TokenERC721) SetMintTransactionHash(v string)`

SetMintTransactionHash sets MintTransactionHash field to given value.


### GetInitialOwnerAddress

`func (o *TokenERC721) GetInitialOwnerAddress() string`

GetInitialOwnerAddress returns the InitialOwnerAddress field if non-nil, zero value otherwise.

### GetInitialOwnerAddressOk

`func (o *TokenERC721) GetInitialOwnerAddressOk() (*string, bool)`

GetInitialOwnerAddressOk returns a tuple with the InitialOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialOwnerAddress

`func (o *TokenERC721) SetInitialOwnerAddress(v string)`

SetInitialOwnerAddress sets InitialOwnerAddress field to given value.


### GetCurrentOwnerAddress

`func (o *TokenERC721) GetCurrentOwnerAddress() string`

GetCurrentOwnerAddress returns the CurrentOwnerAddress field if non-nil, zero value otherwise.

### GetCurrentOwnerAddressOk

`func (o *TokenERC721) GetCurrentOwnerAddressOk() (*string, bool)`

GetCurrentOwnerAddressOk returns a tuple with the CurrentOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOwnerAddress

`func (o *TokenERC721) SetCurrentOwnerAddress(v string)`

SetCurrentOwnerAddress sets CurrentOwnerAddress field to given value.


### GetTransferHistory

`func (o *TokenERC721) GetTransferHistory() []TransferData`

GetTransferHistory returns the TransferHistory field if non-nil, zero value otherwise.

### GetTransferHistoryOk

`func (o *TokenERC721) GetTransferHistoryOk() (*[]TransferData, bool)`

GetTransferHistoryOk returns a tuple with the TransferHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferHistory

`func (o *TokenERC721) SetTransferHistory(v []TransferData)`

SetTransferHistory sets TransferHistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


