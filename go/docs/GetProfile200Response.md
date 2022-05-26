# GetProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NullableGetProfile200ResponseData**](GetProfile200ResponseData.md) |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetProfile200Response

`func NewGetProfile200Response(data NullableGetProfile200ResponseData, meta map[string]interface{}, ) *GetProfile200Response`

NewGetProfile200Response instantiates a new GetProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfile200ResponseWithDefaults

`func NewGetProfile200ResponseWithDefaults() *GetProfile200Response`

NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetProfile200Response) GetData() GetProfile200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProfile200Response) GetDataOk() (*GetProfile200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProfile200Response) SetData(v GetProfile200ResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *GetProfile200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetProfile200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *GetProfile200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProfile200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProfile200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


