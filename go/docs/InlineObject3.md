# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | Item id that will be purchased | 
**ToAddress** | **string** | Buyer wallet address | 
**UserResidence** | [**UserResidence**](UserResidence.md) |  | 

## Methods

### NewInlineObject3

`func NewInlineObject3(itemId string, toAddress string, userResidence UserResidence, ) *InlineObject3`

NewInlineObject3 instantiates a new InlineObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject3WithDefaults

`func NewInlineObject3WithDefaults() *InlineObject3`

NewInlineObject3WithDefaults instantiates a new InlineObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *InlineObject3) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InlineObject3) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InlineObject3) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetToAddress

`func (o *InlineObject3) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *InlineObject3) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *InlineObject3) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetUserResidence

`func (o *InlineObject3) GetUserResidence() UserResidence`

GetUserResidence returns the UserResidence field if non-nil, zero value otherwise.

### GetUserResidenceOk

`func (o *InlineObject3) GetUserResidenceOk() (*UserResidence, bool)`

GetUserResidenceOk returns a tuple with the UserResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResidence

`func (o *InlineObject3) SetUserResidence(v UserResidence)`

SetUserResidence sets UserResidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


