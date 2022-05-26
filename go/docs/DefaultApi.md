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

> CreateOrUpdateItemStockPhysicalShippingInfo200Response CreateOrUpdateItemStockPhysicalShippingInfo(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).CreateOrUpdateItemStockPhysicalShippingInfoRequest(createOrUpdateItemStockPhysicalShippingInfoRequest).Execute()

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
    createOrUpdateItemStockPhysicalShippingInfoRequest := *openapiclient.NewCreateOrUpdateItemStockPhysicalShippingInfoRequest(*openapiclient.NewCreateOrUpdateItemStockPhysicalShippingInfoRequestData(*openapiclient.NewSignatureDomain(openapiclient.NetworkIdString("1"), "Name_example", "Version_example"), "PrimaryType_example", *openapiclient.NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage("FirstName_example", "LastName_example", "Country_example", "Email_example", "PostalCode_example", "City_example", "State_example", "Address1_example", "PhoneNumber_example", "Address2_example", "Address3_example", "RequestTimestamp_example"), map[string]interface{}(123)), "Signature_example") // CreateOrUpdateItemStockPhysicalShippingInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOrUpdateItemStockPhysicalShippingInfo(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).CreateOrUpdateItemStockPhysicalShippingInfoRequest(createOrUpdateItemStockPhysicalShippingInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateItemStockPhysicalShippingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateItemStockPhysicalShippingInfo`: CreateOrUpdateItemStockPhysicalShippingInfo200Response
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
 **createOrUpdateItemStockPhysicalShippingInfoRequest** | [**CreateOrUpdateItemStockPhysicalShippingInfoRequest**](CreateOrUpdateItemStockPhysicalShippingInfoRequest.md) |  | 

### Return type

[**CreateOrUpdateItemStockPhysicalShippingInfo200Response**](CreateOrUpdateItemStockPhysicalShippingInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStripePaymentIntent

> CreateStripePaymentIntent200Response CreateStripePaymentIntent(ctx).MintAccessToken(mintAccessToken).CreateStripePaymentIntentRequest(createStripePaymentIntentRequest).Execute()

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
    createStripePaymentIntentRequest := *openapiclient.NewCreateStripePaymentIntentRequest("ItemId_example", "ToAddress_example", openapiclient.UserResidence("jp")) // CreateStripePaymentIntentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStripePaymentIntent(context.Background()).MintAccessToken(mintAccessToken).CreateStripePaymentIntentRequest(createStripePaymentIntentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStripePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStripePaymentIntent`: CreateStripePaymentIntent200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStripePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStripePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **createStripePaymentIntentRequest** | [**CreateStripePaymentIntentRequest**](CreateStripePaymentIntentRequest.md) |  | 

### Return type

[**CreateStripePaymentIntent200Response**](CreateStripePaymentIntent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatar

> GetAvatar200Response GetAvatar(ctx).MintAccessToken(mintAccessToken).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAvatar(context.Background()).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvatar`: GetAvatar200Response
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

[**GetAvatar200Response**](GetAvatar200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBiddedItemStocksByWalletAddress

> GetBiddedItemStocksByWalletAddress200Response GetBiddedItemStocksByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).OnlyBeforeEnd(onlyBeforeEnd).SortBy(sortBy).SortDirection(sortDirection).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBiddedItemStocksByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).OnlyBeforeEnd(onlyBeforeEnd).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBiddedItemStocksByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBiddedItemStocksByWalletAddress`: GetBiddedItemStocksByWalletAddress200Response
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

[**GetBiddedItemStocksByWalletAddress200Response**](GetBiddedItemStocksByWalletAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoughtItemStocksByWalletAddress

> GetBoughtItemStocksByWalletAddress200Response GetBoughtItemStocksByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBoughtItemStocksByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBoughtItemStocksByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoughtItemStocksByWalletAddress`: GetBoughtItemStocksByWalletAddress200Response
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

[**GetBoughtItemStocksByWalletAddress200Response**](GetBoughtItemStocksByWalletAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractERC721ById

> GetContractERC721ById200Response GetContractERC721ById(ctx).MintAccessToken(mintAccessToken).ContractId(contractId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetContractERC721ById(context.Background()).MintAccessToken(mintAccessToken).ContractId(contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetContractERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractERC721ById`: GetContractERC721ById200Response
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

[**GetContractERC721ById200Response**](GetContractERC721ById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemById

> GetItemById200Response GetItemById(ctx, itemId).MintAccessToken(mintAccessToken).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItemById(context.Background(), itemId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemById`: GetItemById200Response
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

[**GetItemById200Response**](GetItemById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemByTokenERC721

> GetItemById200Response GetItemByTokenERC721(ctx).MintAccessToken(mintAccessToken).TokenId(tokenId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItemByTokenERC721(context.Background()).MintAccessToken(mintAccessToken).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemByTokenERC721``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemByTokenERC721`: GetItemById200Response
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

[**GetItemById200Response**](GetItemById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockById

> GetItemStockById200Response GetItemStockById(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItemStockById(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockById`: GetItemStockById200Response
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

[**GetItemStockById200Response**](GetItemStockById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockPhysicalShippingInfoByItemStockId

> GetItemStockPhysicalShippingInfoByItemStockId200Response GetItemStockPhysicalShippingInfoByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).GetItemStockPhysicalShippingInfoByItemStockIdRequest(getItemStockPhysicalShippingInfoByItemStockIdRequest).Execute()

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
    getItemStockPhysicalShippingInfoByItemStockIdRequest := *openapiclient.NewGetItemStockPhysicalShippingInfoByItemStockIdRequest(*openapiclient.NewGetItemStockPhysicalShippingInfoByItemStockIdRequestData(*openapiclient.NewSignatureDomain(openapiclient.NetworkIdString("1"), "Name_example", "Version_example"), "PrimaryType_example", *openapiclient.NewGetItemStockPhysicalShippingInfoByItemStockIdRequestDataMessage(), map[string]interface{}(123)), "Signature_example") // GetItemStockPhysicalShippingInfoByItemStockIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItemStockPhysicalShippingInfoByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).GetItemStockPhysicalShippingInfoByItemStockIdRequest(getItemStockPhysicalShippingInfoByItemStockIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockPhysicalShippingInfoByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockPhysicalShippingInfoByItemStockId`: GetItemStockPhysicalShippingInfoByItemStockId200Response
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
 **getItemStockPhysicalShippingInfoByItemStockIdRequest** | [**GetItemStockPhysicalShippingInfoByItemStockIdRequest**](GetItemStockPhysicalShippingInfoByItemStockIdRequest.md) |  | 

### Return type

[**GetItemStockPhysicalShippingInfoByItemStockId200Response**](GetItemStockPhysicalShippingInfoByItemStockId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemStockPhysicalShippingInfoStatusByItemStockId

> GetItemStockPhysicalShippingInfoStatusByItemStockId200Response GetItemStockPhysicalShippingInfoStatusByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItemStockPhysicalShippingInfoStatusByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItemStockPhysicalShippingInfoStatusByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemStockPhysicalShippingInfoStatusByItemStockId`: GetItemStockPhysicalShippingInfoStatusByItemStockId200Response
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

[**GetItemStockPhysicalShippingInfoStatusByItemStockId200Response**](GetItemStockPhysicalShippingInfoStatusByItemStockId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> GetItems200Response GetItems(ctx).MintAccessToken(mintAccessToken).Page(page).PerPage(perPage).SaleStatus(saleStatus).OnlyAvailableStock(onlyAvailableStock).PaymentMethod(paymentMethod).Tags(tags).SortBy(sortBy).SortDirection(sortDirection).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetItems(context.Background()).MintAccessToken(mintAccessToken).Page(page).PerPage(perPage).SaleStatus(saleStatus).OnlyAvailableStock(onlyAvailableStock).PaymentMethod(paymentMethod).Tags(tags).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItems`: GetItems200Response
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

[**GetItems200Response**](GetItems200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentIntentById

> GetPaymentIntentById200Response GetPaymentIntentById(ctx, paymentIntentId).MintAccessToken(mintAccessToken).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPaymentIntentById(context.Background(), paymentIntentId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPaymentIntentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentIntentById`: GetPaymentIntentById200Response
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

[**GetPaymentIntentById200Response**](GetPaymentIntentById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductERC721ById

> GetProductERC721ById200Response GetProductERC721ById(ctx).MintAccessToken(mintAccessToken).Id(id).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProductERC721ById(context.Background()).MintAccessToken(mintAccessToken).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProductERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductERC721ById`: GetProductERC721ById200Response
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

[**GetProductERC721ById200Response**](GetProductERC721ById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> GetProfile200Response GetProfile(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProfile(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: GetProfile200Response
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

[**GetProfile200Response**](GetProfile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellableItemStockERC721Id

> GetSellableItemStockERC721Id200Response GetSellableItemStockERC721Id(ctx).MintAccessToken(mintAccessToken).ItemId(itemId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSellableItemStockERC721Id(context.Background()).MintAccessToken(mintAccessToken).ItemId(itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSellableItemStockERC721Id``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellableItemStockERC721Id`: GetSellableItemStockERC721Id200Response
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

[**GetSellableItemStockERC721Id200Response**](GetSellableItemStockERC721Id200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignByItemStockId

> GetSignByItemStockId200Response GetSignByItemStockId(ctx).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).SignatureType(signatureType).WalletAddress(walletAddress).Residence(residence).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSignByItemStockId(context.Background()).MintAccessToken(mintAccessToken).ItemStockId(itemStockId).SignatureType(signatureType).WalletAddress(walletAddress).Residence(residence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSignByItemStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignByItemStockId`: GetSignByItemStockId200Response
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

[**GetSignByItemStockId200Response**](GetSignByItemStockId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenERC721ById

> GetTokenERC721ById200Response GetTokenERC721ById(ctx, tokenId).MintAccessToken(mintAccessToken).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTokenERC721ById(context.Background(), tokenId).MintAccessToken(mintAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokenERC721ById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenERC721ById`: GetTokenERC721ById200Response
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

[**GetTokenERC721ById200Response**](GetTokenERC721ById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenERC721sByWalletAddress

> GetTokenERC721sByWalletAddress200Response GetTokenERC721sByWalletAddress(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTokenERC721sByWalletAddress(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokenERC721sByWalletAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenERC721sByWalletAddress`: GetTokenERC721sByWalletAddress200Response
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

[**GetTokenERC721sByWalletAddress200Response**](GetTokenERC721sByWalletAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokentERC721sByWalletAddressFromAnyContract

> GetTokenERC721sByWalletAddress200Response GetTokentERC721sByWalletAddressFromAnyContract(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTokentERC721sByWalletAddressFromAnyContract(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokentERC721sByWalletAddressFromAnyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokentERC721sByWalletAddressFromAnyContract`: GetTokenERC721sByWalletAddress200Response
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

[**GetTokenERC721sByWalletAddress200Response**](GetTokenERC721sByWalletAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasNft

> HasNfts200Response HasNft(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).TokenId(tokenId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.HasNft(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HasNft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasNft`: HasNfts200Response
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

[**HasNfts200Response**](HasNfts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasNfts

> HasNfts200Response HasNfts(ctx).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.HasNfts(context.Background()).MintAccessToken(mintAccessToken).WalletAddress(walletAddress).ContractAddress(contractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HasNfts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasNfts`: HasNfts200Response
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

[**HasNfts200Response**](HasNfts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile200Response UpdateProfile(ctx).MintAccessToken(mintAccessToken).UpdateProfileRequest(updateProfileRequest).Execute()

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
    updateProfileRequest := *openapiclient.NewUpdateProfileRequest(*openapiclient.NewWalletAddressProfile("WalletAddress_example", "AvatarImageId_example", "DisplayName_example", "Bio_example", "TwitterAccountName_example", "InstagramAccountName_example", "HomepageUrl_example"), "Signature_example") // UpdateProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateProfile(context.Background()).MintAccessToken(mintAccessToken).UpdateProfileRequest(updateProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: UpdateProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintAccessToken** | **string** |  | 
 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

[**UpdateProfile200Response**](UpdateProfile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

