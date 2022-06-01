# CreateOrUpdateItemStockPhysicalShippingInfoRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | [**SignatureDomain**](SignatureDomain.md) |  | 
**PrimaryType** | **string** |  | 
**Message** | [**CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage**](CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage.md) |  | 
**Types** | **map[string]interface{}** |  | 

## Methods

### NewCreateOrUpdateItemStockPhysicalShippingInfoRequestData

`func NewCreateOrUpdateItemStockPhysicalShippingInfoRequestData(domain SignatureDomain, primaryType string, message CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage, types map[string]interface{}, ) *CreateOrUpdateItemStockPhysicalShippingInfoRequestData`

NewCreateOrUpdateItemStockPhysicalShippingInfoRequestData instantiates a new CreateOrUpdateItemStockPhysicalShippingInfoRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateItemStockPhysicalShippingInfoRequestDataWithDefaults

`func NewCreateOrUpdateItemStockPhysicalShippingInfoRequestDataWithDefaults() *CreateOrUpdateItemStockPhysicalShippingInfoRequestData`

NewCreateOrUpdateItemStockPhysicalShippingInfoRequestDataWithDefaults instantiates a new CreateOrUpdateItemStockPhysicalShippingInfoRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetDomain() SignatureDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetDomainOk() (*SignatureDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) SetDomain(v SignatureDomain)`

SetDomain sets Domain field to given value.


### GetPrimaryType

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.


### GetMessage

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetMessage() CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetMessageOk() (*CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) SetMessage(v CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage)`

SetMessage sets Message field to given value.


### GetTypes

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetTypes() map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) GetTypesOk() (*map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestData) SetTypes(v map[string]interface{})`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


