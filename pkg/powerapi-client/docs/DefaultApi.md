# \DefaultApi

All URIs are relative to *https://}/power/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregationServiceActionsAggregationServiceResetPost**](DefaultApi.md#AggregationServiceActionsAggregationServiceResetPost) | **Post** /AggregationService/Actions/AggregationService.Reset | Request aggregate system reset
[**ComputerSystemsGet**](DefaultApi.md#ComputerSystemsGet) | **Get** /ComputerSystems | Get computer systems
[**ComputerSystemsNameActionsComputerSystemResetPost**](DefaultApi.md#ComputerSystemsNameActionsComputerSystemResetPost) | **Post** /ComputerSystems/{name}/Actions/ComputerSystem.Reset | Request system reset
[**ComputerSystemsNameGet**](DefaultApi.md#ComputerSystemsNameGet) | **Get** /ComputerSystems/{name} | Get a specific computer system state



## AggregationServiceActionsAggregationServiceResetPost

> AggregationResetBody AggregationServiceActionsAggregationServiceResetPost(ctx).AggregationResetBody(aggregationResetBody).Execute()

Request aggregate system reset

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
    aggregationResetBody := *openapiclient.NewAggregationResetBody() // AggregationResetBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AggregationServiceActionsAggregationServiceResetPost(context.Background()).AggregationResetBody(aggregationResetBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AggregationServiceActionsAggregationServiceResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregationServiceActionsAggregationServiceResetPost`: AggregationResetBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AggregationServiceActionsAggregationServiceResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregationServiceActionsAggregationServiceResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregationResetBody** | [**AggregationResetBody**](AggregationResetBody.md) |  | 

### Return type

[**AggregationResetBody**](AggregationResetBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerSystemsGet

> ComputerSystemCollection ComputerSystemsGet(ctx).Execute()

Get computer systems

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ComputerSystemsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ComputerSystemsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputerSystemsGet`: ComputerSystemCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ComputerSystemsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiComputerSystemsGetRequest struct via the builder pattern


### Return type

[**ComputerSystemCollection**](ComputerSystemCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerSystemsNameActionsComputerSystemResetPost

> ResetRequestBody ComputerSystemsNameActionsComputerSystemResetPost(ctx, name).ResetRequestBody(resetRequestBody).Execute()

Request system reset

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
    name := "name_example" // string | Unique name of the computer system
    resetRequestBody := *openapiclient.NewResetRequestBody(openapiclient.ResetType("On")) // ResetRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ComputerSystemsNameActionsComputerSystemResetPost(context.Background(), name).ResetRequestBody(resetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ComputerSystemsNameActionsComputerSystemResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputerSystemsNameActionsComputerSystemResetPost`: ResetRequestBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ComputerSystemsNameActionsComputerSystemResetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name of the computer system | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerSystemsNameActionsComputerSystemResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetRequestBody** | [**ResetRequestBody**](ResetRequestBody.md) |  | 

### Return type

[**ResetRequestBody**](ResetRequestBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerSystemsNameGet

> ComputerSystem ComputerSystemsNameGet(ctx, name).Execute()

Get a specific computer system state

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
    name := "name_example" // string | Unique name of the computer system

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ComputerSystemsNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ComputerSystemsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputerSystemsNameGet`: ComputerSystem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ComputerSystemsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name of the computer system | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerSystemsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerSystem**](ComputerSystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

