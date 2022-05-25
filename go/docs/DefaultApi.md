# \DefaultApi

All URIs are relative to *http://localhost:5500/mint-v2-development/asia-northeast1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateItemStockPhysicalShippingInfo**](DefaultApi.md#CreateOrUpdateItemStockPhysicalShippingInfo) | **Post** /sdk_v4/itemStockPhysicalShippingInfos/createOrUpdateItemStockPhysicalShippingInfo | API for creating or updating item stock physical shipping info for given item stock id
[**CreateStripePaymentIntent**](DefaultApi.md#CreateStripePaymentIntent) | **Post** /sdk_v4/stripePayment/createStripePaymentIntent | クレジットカード(Stripe)で指定のアイテムを購入するためのPyamentIntentを作成し、対応するSecretを返す
[**GetAvatar**](DefaultApi.md#GetAvatar) | **Get** /sdk_v4/avatar | アバター画像の署名付きURLの取得
[**GetBiddedItemStocksByWalletAddress**](DefaultApi.md#GetBiddedItemStocksByWalletAddress) | **Get** /sdk_v4/itemStocks/getBiddedItemStocksByWalletAddress | 指定したwalletAddressでBidしたItemStockを取得する
[**GetBoughtItemStocksByWalletAddress**](DefaultApi.md#GetBoughtItemStocksByWalletAddress) | **Get** /sdk_v4/itemStocks/getBoughtItemStocksByWalletAddress | 指定したwalletAddressで購入または落札したItemStockを取得する
[**GetContractERC721ById**](DefaultApi.md#GetContractERC721ById) | **Get** /sdk_v4/contracts/getContractERC721ById | ContractERC721をId指定で取得する
[**GetItemById**](DefaultApi.md#GetItemById) | **Get** /sdk_v4/items/{itemId} | プロジェクトのItemをId指定で取得する
[**GetItemByTokenERC721**](DefaultApi.md#GetItemByTokenERC721) | **Get** /sdk_v4/items/getItemByTokenERC721 | get item by tokenERC721 id
[**GetItemStockById**](DefaultApi.md#GetItemStockById) | **Get** /sdk_v4/itemStocks/getById | ItemStockをId指定で取得する
[**GetItemStockPhysicalShippingInfoByItemStockId**](DefaultApi.md#GetItemStockPhysicalShippingInfoByItemStockId) | **Post** /sdk_v4/itemStockPhysicalShippingInfos/getItemStockPhysicalShippingInfoByItemStockId | API for getting item stock physical shipping info by item stock id
[**GetItemStockPhysicalShippingInfoStatusByItemStockId**](DefaultApi.md#GetItemStockPhysicalShippingInfoStatusByItemStockId) | **Get** /sdk_v4/itemStockPhysicalShippingInfos/getItemStockPhysicalShippingInfoStatusByItemStockId | API for getting item stock physical shipping info status by item stock id
[**GetItems**](DefaultApi.md#GetItems) | **Get** /sdk_v4/items | プロジェクトのItemを全て取得する
[**GetPaymentIntentById**](DefaultApi.md#GetPaymentIntentById) | **Get** /sdk_v4/paymentIntents/{paymentIntentId} | This API is responsible to get payment intent by its id
[**GetProductERC721ById**](DefaultApi.md#GetProductERC721ById) | **Get** /sdk_v4/products/getProductERC721ById | Itemにパックされていて、ItemのstatusがpublishなProductERC721を取得
[**GetProfile**](DefaultApi.md#GetProfile) | **Get** /sdk_v4/profile | ウォレットに紐づくプロフィールの取得
[**GetSellableItemStockERC721Id**](DefaultApi.md#GetSellableItemStockERC721Id) | **Get** /sdk_v4/items/getSellableItemStockERC721Id | スマコンで販売している&#x60;Item&#x60;の販売可能な(まだ売れていない)&#x60;ItemStockId&#x60;を取得する
[**GetSignByItemStockId**](DefaultApi.md#GetSignByItemStockId) | **Get** /sdk_v4/itemStocks/sign | Item購入に関してスマコンの操作に必要なSignを返す
[**GetTokenERC721ById**](DefaultApi.md#GetTokenERC721ById) | **Get** /sdk_v4/tokens/{tokenId} | get TokenERC721 by Id
[**GetTokenERC721sByWalletAddress**](DefaultApi.md#GetTokenERC721sByWalletAddress) | **Get** /sdk_v4/tokens/getTokenERC721sByWalletAddress | walletAddressに紐づくTokenERC721を全て取得する
[**GetTokentERC721sByWalletAddressFromAnyContract**](DefaultApi.md#GetTokentERC721sByWalletAddressFromAnyContract) | **Get** /sdk_v4/tokens/getTokentERC721sByWalletAddressFromAnyContract | get TokenERC721s by specifying wallet address and contract address
[**HasNft**](DefaultApi.md#HasNft) | **Get** /sdk_v4/tokens/hasNft | has token id of nft or not in contract
[**HasNfts**](DefaultApi.md#HasNfts) | **Get** /sdk_v4/tokens/hasNfts | has nfts or not in contract
[**UpdateProfile**](DefaultApi.md#UpdateProfile) | **Post** /sdk_v4/profile | ウォレットに紐づくプロフィールの作成



## CreateOrUpdateItemStockPhysicalShippingInfo

> InlineResponse2009 CreateOrUpdateItemStockPhysicalShippingInfo(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).InlineObject1(inlineObject1).Execute()

API for creating or updating item stock physical shipping info for given item stock id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemStockId := "itemStockId_example" // string | 
    inlineObject1 := *openapiclient.NewInlineObject1(*openapiclient.NewSdkV4ItemStockPhysicalShippingInfosCreateOrUpdateItemStockPhysicalShippingInfoData(*openapiclient.NewSignatureDomain(openapiclient.NetworkIdString("1"), "Name_example", "Version_example"), "PrimaryType_example", *openapiclient.NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage("FirstName_example", "LastName_example", "Country_example", "Email_example", "PostalCode_example", "City_example", "State_example", "Address1_example", "PhoneNumber_example", "Address2_example", "Address3_example", "RequestTimestamp_example"), map[string]interface{}(123)), "Signature_example") // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOrUpdateItemStockPhysicalShippingInfo(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateItemStockPhysicalShippingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateItemStockPhysicalShippingInfo`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrUpdateItemStockPhysicalShippingInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateItemStockPhysicalShippingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemStockId** | **string** |  | 
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStripePaymentIntent

> InlineResponse20019 CreateStripePaymentIntent(ctx).MintAccessToken(mintAccessToken).InlineObject3(inlineObject3).Execute()

クレジットカード(Stripe)で指定のアイテムを購入するためのPyamentIntentを作成し、対応するSecretを返す

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    inlineObject3 := *openapiclient.NewInlineObject3("ItemId_example", "ToAddress_example", openapiclient.UserResidence("jp")) // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateStripePaymentIntent(context.Background()).MintAccessToken(mintAccessToken).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStripePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStripePaymentIntent`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStripePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStripePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatar

> InlineResponse20017 GetAvatar(ctx).MintAccessToken(mintAccessToken).Execute()

アバター画像の署名付きURLの取得

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAvatar(context.Background()).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvatar`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBiddedItemStocksByWalletAddress

> InlineResponse2005 GetBiddedItemStocksByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).OnlyBeforeEnd(onlyBeforeEnd).SortBy(sortBy).SortDirection(sortDirection).Execute()

指定したwalletAddressでBidしたItemStockを取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    page := "page_example" // string | 
    perPage := "perPage_example" // string | 
    onlyBeforeEnd := "onlyBeforeEnd_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBiddedItemStocksByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).OnlyBeforeEnd(onlyBeforeEnd).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBiddedItemStocksByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBiddedItemStocksByWalletAddress`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBiddedItemStocksByWalletAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBiddedItemStocksByWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **page** | **string** |  | 
 **perPage** | **string** |  | 
 **onlyBeforeEnd** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoughtItemStocksByWalletAddress

> InlineResponse2006 GetBoughtItemStocksByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()

指定したwalletAddressで購入または落札したItemStockを取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    page := "page_example" // string | 
    perPage := "perPage_example" // string | 
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBoughtItemStocksByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBoughtItemStocksByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoughtItemStocksByWalletAddress`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBoughtItemStocksByWalletAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBoughtItemStocksByWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **page** | **string** |  | 
 **perPage** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractERC721ById

> InlineResponse20018 GetContractERC721ById(ctx).MintAccessToken(mintAccessToken).ContractId(contractId).Execute()

ContractERC721をId指定で取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    contractId := "contractId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetContractERC721ById(context.Background()).MintAccessToken(mintAccessToken).ContractId(contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetContractERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractERC721ById`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetContractERC721ById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractERC721ByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **contractId** | **string** |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemById

> InlineResponse2001 GetItemById(ctx, itemId).MintAccessToken(mintAccessToken).Execute()

プロジェクトのItemをId指定で取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItemById(context.Background(), itemId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemById`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemByTokenERC721

> InlineResponse2001 GetItemByTokenERC721(ctx).MintAccessToken(mintAccessToken).TokenId(tokenId).Execute()

get item by tokenERC721 id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItemByTokenERC721(context.Background()).MintAccessToken(mintAccessToken).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemByTokenERC721``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemByTokenERC721`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItemByTokenERC721`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByTokenERC721Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **tokenId** | **string** |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockById

> InlineResponse2004 GetItemStockById(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()

ItemStockをId指定で取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemStockId := "itemStockId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItemStockById(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockById`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItemStockById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemStockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemStockId** | **string** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockPhysicalShippingInfoByItemStockId

> InlineResponse2008 GetItemStockPhysicalShippingInfoByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).InlineObject(inlineObject).Execute()

API for getting item stock physical shipping info by item stock id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemStockId := "itemStockId_example" // string | 
    inlineObject := *openapiclient.NewInlineObject(*openapiclient.NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdData(*openapiclient.NewSignatureDomain(openapiclient.NetworkIdString("1"), "Name_example", "Version_example"), "PrimaryType_example", *openapiclient.NewSdkV4ItemStockPhysicalShippingInfosGetItemStockPhysicalShippingInfoByItemStockIdDataMessage(), map[string]interface{}(123)), "Signature_example") // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItemStockPhysicalShippingInfoByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockPhysicalShippingInfoByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockPhysicalShippingInfoByItemStockId`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItemStockPhysicalShippingInfoByItemStockId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemStockPhysicalShippingInfoByItemStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemStockId** | **string** |  | 
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockPhysicalShippingInfoStatusByItemStockId

> InlineResponse2007 GetItemStockPhysicalShippingInfoStatusByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()

API for getting item stock physical shipping info status by item stock id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemStockId := "itemStockId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItemStockPhysicalShippingInfoStatusByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockPhysicalShippingInfoStatusByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockPhysicalShippingInfoStatusByItemStockId`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItemStockPhysicalShippingInfoStatusByItemStockId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemStockPhysicalShippingInfoStatusByItemStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemStockId** | **string** |  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> InlineResponse200 GetItems(ctx).MintAccessToken(mintAccessToken).Page(page).PerPage(perPage).SaleStatus(saleStatus).OnlyAvailableStock(onlyAvailableStock).PaymentMethod(paymentMethod).Tags(tags).SortBy(sortBy).SortDirection(sortDirection).Execute()

プロジェクトのItemを全て取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    page := "page_example" // string | 
    perPage := "perPage_example" // string | 
    saleStatus := "saleStatus_example" // string |  (optional)
    onlyAvailableStock := "onlyAvailableStock_example" // string |  (optional)
    paymentMethod := "paymentMethod_example" // string |  (optional)
    tags := "tags_example" // string | , 区切りで指定 (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetItems(context.Background()).MintAccessToken(mintAccessToken).Page(page).PerPage(perPage).SaleStatus(saleStatus).OnlyAvailableStock(onlyAvailableStock).PaymentMethod(paymentMethod).Tags(tags).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **page** | **string** |  | 
 **perPage** | **string** |  | 
 **saleStatus** | **string** |  | 
 **onlyAvailableStock** | **string** |  | 
 **paymentMethod** | **string** |  | 
 **tags** | **string** | , 区切りで指定 | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentIntentById

> InlineResponse20011 GetPaymentIntentById(ctx, paymentIntentId).MintAccessToken(mintAccessToken).Execute()

This API is responsible to get payment intent by its id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    paymentIntentId := "paymentIntentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetPaymentIntentById(context.Background(), paymentIntentId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPaymentIntentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentIntentById`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPaymentIntentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentIntentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentIntentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 


### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductERC721ById

> InlineResponse20010 GetProductERC721ById(ctx).MintAccessToken(mintAccessToken).Id(id).Execute()

Itemにパックされていて、ItemのstatusがpublishなProductERC721を取得

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetProductERC721ById(context.Background()).MintAccessToken(mintAccessToken).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProductERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductERC721ById`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProductERC721ById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductERC721ByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **id** | **string** |  | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> InlineResponse20015 GetProfile(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Execute()

ウォレットに紐づくプロフィールの取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetProfile(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellableItemStockERC721Id

> InlineResponse2002 GetSellableItemStockERC721Id(ctx).MintAccessToken(mintAccessToken).ItemId(itemId).Execute()

スマコンで販売している`Item`の販売可能な(まだ売れていない)`ItemStockId`を取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSellableItemStockERC721Id(context.Background()).MintAccessToken(mintAccessToken).ItemId(itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSellableItemStockERC721Id``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellableItemStockERC721Id`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSellableItemStockERC721Id`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellableItemStockERC721IdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemId** | **string** |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignByItemStockId

> InlineResponse2003 GetSignByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).SignatureType(signatureType).WalletAddress(walletAddress).Residence(residence).Execute()

Item購入に関してスマコンの操作に必要なSignを返す

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    itemStockId := "itemStockId_example" // string | 
    signatureType := openapiclient.SignatureType("ethereum-contract-erc721-shop-fixed-price") // SignatureType | 
    walletAddress := "walletAddress_example" // string | 購入時のみ必須 (optional)
    residence := "residence_example" // string | 購入時のみ必須 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSignByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).SignatureType(signatureType).WalletAddress(walletAddress).Residence(residence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSignByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignByItemStockId`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSignByItemStockId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignByItemStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **itemStockId** | **string** |  | 
 **signatureType** | [**SignatureType**](SignatureType.md) |  | 
 **walletAddress** | **string** | 購入時のみ必須 | 
 **residence** | **string** | 購入時のみ必須 | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenERC721ById

> InlineResponse20014 GetTokenERC721ById(ctx, tokenId).MintAccessToken(mintAccessToken).Execute()

get TokenERC721 by Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTokenERC721ById(context.Background(), tokenId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokenERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenERC721ById`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTokenERC721ById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenERC721ByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 


### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenERC721sByWalletAddress

> InlineResponse20012 GetTokenERC721sByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).Execute()

walletAddressに紐づくTokenERC721を全て取得する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    page := "page_example" // string | 
    perPage := "perPage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTokenERC721sByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokenERC721sByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenERC721sByWalletAddress`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTokenERC721sByWalletAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenERC721sByWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **page** | **string** |  | 
 **perPage** | **string** |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokentERC721sByWalletAddressFromAnyContract

> InlineResponse20012 GetTokentERC721sByWalletAddressFromAnyContract(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()

get TokenERC721s by specifying wallet address and contract address

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    contractAddress := "contractAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTokentERC721sByWalletAddressFromAnyContract(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokentERC721sByWalletAddressFromAnyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokentERC721sByWalletAddressFromAnyContract`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTokentERC721sByWalletAddressFromAnyContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokentERC721sByWalletAddressFromAnyContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **contractAddress** | **string** |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasNft

> InlineResponse20013 HasNft(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).TokenId(tokenId).Execute()

has token id of nft or not in contract

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    contractAddress := "contractAddress_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HasNft(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HasNft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasNft`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HasNft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHasNftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **contractAddress** | **string** |  | 
 **tokenId** | **string** |  | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasNfts

> InlineResponse20013 HasNfts(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()

has nfts or not in contract

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    walletAddress := "walletAddress_example" // string | 
    contractAddress := "contractAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HasNfts(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HasNfts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasNfts`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HasNfts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHasNftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **walletAddress** | **string** |  | 
 **contractAddress** | **string** |  | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> InlineResponse20016 UpdateProfile(ctx).MintAccessToken(mintAccessToken).InlineObject2(inlineObject2).Execute()

ウォレットに紐づくプロフィールの作成

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mintAccessToken := "mintAccessToken_example" // string | 
    inlineObject2 := *openapiclient.NewInlineObject2(*openapiclient.NewWalletAddressProfile("WalletAddress_example", "AvatarImageId_example", "DisplayName_example", "Bio_example", "TwitterAccountName_example", "InstagramAccountName_example", "HomepageUrl_example"), "Signature_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateProfile(context.Background()).MintAccessToken(mintAccessToken).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

