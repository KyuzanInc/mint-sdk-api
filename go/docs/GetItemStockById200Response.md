# GetItemStockById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ItemStock**](ItemStock.md) |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetItemStockById200Response

`func NewGetItemStockById200Response(data ItemStock, meta map[string]interface{}, ) *GetItemStockById200Response`

NewGetItemStockById200Response instantiates a new GetItemStockById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetItemStockById200ResponseWithDefaults

`func NewGetItemStockById200ResponseWithDefaults() *GetItemStockById200Response`

NewGetItemStockById200ResponseWithDefaults instantiates a new GetItemStockById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetItemStockById200Response) GetData() ItemStock`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetItemStockById200Response) GetDataOk() (*ItemStock, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetItemStockById200Response) SetData(v ItemStock)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetItemStockById200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetItemStockById200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetItemStockById200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


