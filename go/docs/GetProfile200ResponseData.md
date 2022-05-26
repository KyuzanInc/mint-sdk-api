# GetProfile200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | [**WalletAddressProfile**](WalletAddressProfile.md) |  | 
**AvatarImageUrl** | **string** |  | 

## Methods

### NewGetProfile200ResponseData

`func NewGetProfile200ResponseData(profile WalletAddressProfile, avatarImageUrl string, ) *GetProfile200ResponseData`

NewGetProfile200ResponseData instantiates a new GetProfile200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfile200ResponseDataWithDefaults

`func NewGetProfile200ResponseDataWithDefaults() *GetProfile200ResponseData`

NewGetProfile200ResponseDataWithDefaults instantiates a new GetProfile200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *GetProfile200ResponseData) GetProfile() WalletAddressProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetProfile200ResponseData) GetProfileOk() (*WalletAddressProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetProfile200ResponseData) SetProfile(v WalletAddressProfile)`

SetProfile sets Profile field to given value.


### GetAvatarImageUrl

`func (o *GetProfile200ResponseData) GetAvatarImageUrl() string`

GetAvatarImageUrl returns the AvatarImageUrl field if non-nil, zero value otherwise.

### GetAvatarImageUrlOk

`func (o *GetProfile200ResponseData) GetAvatarImageUrlOk() (*string, bool)`

GetAvatarImageUrlOk returns a tuple with the AvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarImageUrl

`func (o *GetProfile200ResponseData) SetAvatarImageUrl(v string)`

SetAvatarImageUrl sets AvatarImageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


