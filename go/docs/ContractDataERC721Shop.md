# ContractDataERC721Shop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainType** | [**ChainType**](ChainType.md) |  | 
**NetworkId** | [**NetworkId**](NetworkId.md) |  | 
**ContractAddress** | **string** |  | 
**Abi** | **string** | JSON.stringifyしたもの | 

## Methods

### NewContractDataERC721Shop

`func NewContractDataERC721Shop(chainType ChainType, networkId NetworkId, contractAddress string, abi string, ) *ContractDataERC721Shop`

NewContractDataERC721Shop instantiates a new ContractDataERC721Shop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractDataERC721ShopWithDefaults

`func NewContractDataERC721ShopWithDefaults() *ContractDataERC721Shop`

NewContractDataERC721ShopWithDefaults instantiates a new ContractDataERC721Shop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainType

`func (o *ContractDataERC721Shop) GetChainType() ChainType`

GetChainType returns the ChainType field if non-nil, zero value otherwise.

### GetChainTypeOk

`func (o *ContractDataERC721Shop) GetChainTypeOk() (*ChainType, bool)`

GetChainTypeOk returns a tuple with the ChainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainType

`func (o *ContractDataERC721Shop) SetChainType(v ChainType)`

SetChainType sets ChainType field to given value.


### GetNetworkId

`func (o *ContractDataERC721Shop) GetNetworkId() NetworkId`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ContractDataERC721Shop) GetNetworkIdOk() (*NetworkId, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ContractDataERC721Shop) SetNetworkId(v NetworkId)`

SetNetworkId sets NetworkId field to given value.


### GetContractAddress

`func (o *ContractDataERC721Shop) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ContractDataERC721Shop) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ContractDataERC721Shop) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetAbi

`func (o *ContractDataERC721Shop) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *ContractDataERC721Shop) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *ContractDataERC721Shop) SetAbi(v string)`

SetAbi sets Abi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


