# GetItemById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Item**](Item.md) |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetItemById200Response

`func NewGetItemById200Response(data Item, meta map[string]interface{}, ) *GetItemById200Response`

NewGetItemById200Response instantiates a new GetItemById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetItemById200ResponseWithDefaults

`func NewGetItemById200ResponseWithDefaults() *GetItemById200Response`

NewGetItemById200ResponseWithDefaults instantiates a new GetItemById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetItemById200Response) GetData() Item`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetItemById200Response) GetDataOk() (*Item, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetItemById200Response) SetData(v Item)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetItemById200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetItemById200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetItemById200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


