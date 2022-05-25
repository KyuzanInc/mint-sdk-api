# SignatureDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | [**NetworkIdString**](NetworkIdString.md) |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewSignatureDomain

`func NewSignatureDomain(chainId NetworkIdString, name string, version string, ) *SignatureDomain`

NewSignatureDomain instantiates a new SignatureDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureDomainWithDefaults

`func NewSignatureDomainWithDefaults() *SignatureDomain`

NewSignatureDomainWithDefaults instantiates a new SignatureDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *SignatureDomain) GetChainId() NetworkIdString`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SignatureDomain) GetChainIdOk() (*NetworkIdString, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SignatureDomain) SetChainId(v NetworkIdString)`

SetChainId sets ChainId field to given value.


### GetName

`func (o *SignatureDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignatureDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignatureDomain) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *SignatureDomain) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignatureDomain) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignatureDomain) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


