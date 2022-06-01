# CreateStripePaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | Item id that will be purchased | 
**ToAddress** | **string** | Buyer wallet address | 
**UserResidence** | [**UserResidence**](UserResidence.md) |  | 

## Methods

### NewCreateStripePaymentIntentRequest

`func NewCreateStripePaymentIntentRequest(itemId string, toAddress string, userResidence UserResidence, ) *CreateStripePaymentIntentRequest`

NewCreateStripePaymentIntentRequest instantiates a new CreateStripePaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripePaymentIntentRequestWithDefaults

`func NewCreateStripePaymentIntentRequestWithDefaults() *CreateStripePaymentIntentRequest`

NewCreateStripePaymentIntentRequestWithDefaults instantiates a new CreateStripePaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *CreateStripePaymentIntentRequest) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CreateStripePaymentIntentRequest) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CreateStripePaymentIntentRequest) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetToAddress

`func (o *CreateStripePaymentIntentRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateStripePaymentIntentRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateStripePaymentIntentRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetUserResidence

`func (o *CreateStripePaymentIntentRequest) GetUserResidence() UserResidence`

GetUserResidence returns the UserResidence field if non-nil, zero value otherwise.

### GetUserResidenceOk

`func (o *CreateStripePaymentIntentRequest) GetUserResidenceOk() (*UserResidence, bool)`

GetUserResidenceOk returns a tuple with the UserResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResidence

`func (o *CreateStripePaymentIntentRequest) SetUserResidence(v UserResidence)`

SetUserResidence sets UserResidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


