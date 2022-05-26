# CreateStripePaymentIntent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublishableKey** | **string** | クライアント側でloadStripeに対して渡す公開可能なAPI-Key | 
**Secret** | **string** | StripeのPaymentIntentのClientSecret | 

## Methods

### NewCreateStripePaymentIntent200Response

`func NewCreateStripePaymentIntent200Response(publishableKey string, secret string, ) *CreateStripePaymentIntent200Response`

NewCreateStripePaymentIntent200Response instantiates a new CreateStripePaymentIntent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripePaymentIntent200ResponseWithDefaults

`func NewCreateStripePaymentIntent200ResponseWithDefaults() *CreateStripePaymentIntent200Response`

NewCreateStripePaymentIntent200ResponseWithDefaults instantiates a new CreateStripePaymentIntent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishableKey

`func (o *CreateStripePaymentIntent200Response) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *CreateStripePaymentIntent200Response) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *CreateStripePaymentIntent200Response) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.


### GetSecret

`func (o *CreateStripePaymentIntent200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateStripePaymentIntent200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateStripePaymentIntent200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


