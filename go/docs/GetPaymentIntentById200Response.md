# GetPaymentIntentById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetPaymentIntentByIdResponseBodyData**](GetPaymentIntentByIdResponseBodyData.md) |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetPaymentIntentById200Response

`func NewGetPaymentIntentById200Response(data GetPaymentIntentByIdResponseBodyData, meta map[string]interface{}, ) *GetPaymentIntentById200Response`

NewGetPaymentIntentById200Response instantiates a new GetPaymentIntentById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentIntentById200ResponseWithDefaults

`func NewGetPaymentIntentById200ResponseWithDefaults() *GetPaymentIntentById200Response`

NewGetPaymentIntentById200ResponseWithDefaults instantiates a new GetPaymentIntentById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPaymentIntentById200Response) GetData() GetPaymentIntentByIdResponseBodyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPaymentIntentById200Response) GetDataOk() (*GetPaymentIntentByIdResponseBodyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPaymentIntentById200Response) SetData(v GetPaymentIntentByIdResponseBodyData)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetPaymentIntentById200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetPaymentIntentById200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetPaymentIntentById200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


