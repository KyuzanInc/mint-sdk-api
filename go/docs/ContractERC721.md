# ContractERC721

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreateAt** | **time.Time** |  | 
**UpdateAt** | **time.Time** |  | 
**Name** | **string** |  | 
**ChainType** | [**ChainType**](ChainType.md) |  | 
**NetworkId** | [**NetworkId**](NetworkId.md) |  | 
**Address** | **string** |  | 
**InitialDeployBlockNumber** | **float32** |  | 
**TokenStandardType** | [**TokenStandardType**](TokenStandardType.md) |  | [default to ERC721]

## Methods

### NewContractERC721

`func NewContractERC721(id string, createAt time.Time, updateAt time.Time, name string, chainType ChainType, networkId NetworkId, address string, initialDeployBlockNumber float32, tokenStandardType TokenStandardType, ) *ContractERC721`

NewContractERC721 instantiates a new ContractERC721 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractERC721WithDefaults

`func NewContractERC721WithDefaults() *ContractERC721`

NewContractERC721WithDefaults instantiates a new ContractERC721 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractERC721) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractERC721) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractERC721) SetId(v string)`

SetId sets Id field to given value.


### GetCreateAt

`func (o *ContractERC721) GetCreateAt() time.Time`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *ContractERC721) GetCreateAtOk() (*time.Time, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *ContractERC721) SetCreateAt(v time.Time)`

SetCreateAt sets CreateAt field to given value.


### GetUpdateAt

`func (o *ContractERC721) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *ContractERC721) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *ContractERC721) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.


### GetName

`func (o *ContractERC721) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractERC721) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractERC721) SetName(v string)`

SetName sets Name field to given value.


### GetChainType

`func (o *ContractERC721) GetChainType() ChainType`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ContractERC721) GetChainTypeOk() (*ChainType, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ContractERC721) SetChainType(v ChainType)`

SetChainType sets ChainType field to given value.


### GetNetworkId

`func (o *ContractERC721) GetNetworkId() NetworkId`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ContractERC721) GetNetworkIdOk() (*NetworkId, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ContractERC721) SetNetworkId(v NetworkId)`

SetNetworkId sets NetworkId field to given value.


### GetAddress

`func (o *ContractERC721) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractERC721) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractERC721) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetInitialDeployBlockNumber

`func (o *ContractERC721) GetInitialDeployBlockNumber() float32`

GetInitialDeployBlockNumber returns the InitialDeployBlockNumber field if non-nil, zero value otherwise.

### GetInitialDeployBlockNumberOk

`func (o *ContractERC721) GetInitialDeployBlockNumberOk() (*float32, bool)`

GetInitialDeployBlockNumberOk returns a tuple with the InitialDeployBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDeployBlockNumber

`func (o *ContractERC721) SetInitialDeployBlockNumber(v float32)`

SetInitialDeployBlockNumber sets InitialDeployBlockNumber field to given value.


### GetTokenStandardType

`func (o *ContractERC721) GetTokenStandardType() TokenStandardType`

GetTokenStandardType returns the TokenStandardType field if non-nil, zero value otherwise.

### GetTokenStandardTypeOk

`func (o *ContractERC721) GetTokenStandardTypeOk() (*TokenStandardType, bool)`

GetTokenStandardTypeOk returns a tuple with the TokenStandardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandardType

`func (o *ContractERC721) SetTokenStandardType(v TokenStandardType)`

SetTokenStandardType sets TokenStandardType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


