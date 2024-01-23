# \AnalyticsAPI

All URIs are relative to *https://example.saviyntcloud.com/ECM/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchControlAttributes**](AnalyticsAPI.md#FetchControlAttributes) | **Post** /fetchControlAttributes | Fetch Control Attributes
[**FetchControlDetails**](AnalyticsAPI.md#FetchControlDetails) | **Post** /fetchControlDetails | Fetch Analytics Details
[**FetchControlDetailsES**](AnalyticsAPI.md#FetchControlDetailsES) | **Post** /fetchControlDetailsES | Fetch Analytics Details ES
[**FetchControlList**](AnalyticsAPI.md#FetchControlList) | **Post** /fetchControlList | Fetch List of Analytics
[**FetchControlListES**](AnalyticsAPI.md#FetchControlListES) | **Post** /fetchControlListES | Fetch List of Analytics ES
[**FetchRuntimeControlsData**](AnalyticsAPI.md#FetchRuntimeControlsData) | **Post** /fetchRuntimeControlsData | Fetch Runtime Controls Data
[**FetchRuntimeControlsDataV2**](AnalyticsAPI.md#FetchRuntimeControlsDataV2) | **Post** /fetchRuntimeControlsDataV2 | Fetch Runtime Controls Data V2
[**RunAnalyticsControls**](AnalyticsAPI.md#RunAnalyticsControls) | **Post** /runAnalyticsControls | Run Analytics Controls



## FetchControlAttributes

> FetchControlAttributes(ctx).FetchControlAttributesRequest(fetchControlAttributesRequest).Execute()

Fetch Control Attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fetchControlAttributesRequest := *openapiclient.NewFetchControlAttributesRequest("Analyticstype_example") // FetchControlAttributesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchControlAttributes(context.Background()).FetchControlAttributesRequest(fetchControlAttributesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchControlAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchControlAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchControlAttributesRequest** | [**FetchControlAttributesRequest**](FetchControlAttributesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchControlDetails

> FetchControlDetails(ctx).ControlId(controlId).Max(max).Offset(offset).Execute()

Fetch Analytics Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	controlId := "controlId_example" // string | 
	max := "max_example" // string |  (optional)
	offset := "offset_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchControlDetails(context.Background()).ControlId(controlId).Max(max).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchControlDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchControlDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlId** | **string** |  | 
 **max** | **string** |  | 
 **offset** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchControlDetailsES

> FetchControlDetailsES(ctx).Execute()

Fetch Analytics Details ES



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchControlDetailsES(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchControlDetailsES``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchControlDetailsESRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchControlList

> FetchControlListResponse FetchControlList(ctx).FetchControlListRequest(fetchControlListRequest).Execute()

Fetch List of Analytics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fetchControlListRequest := *openapiclient.NewFetchControlListRequest() // FetchControlListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.FetchControlList(context.Background()).FetchControlListRequest(fetchControlListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchControlList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchControlList`: FetchControlListResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.FetchControlList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchControlListRequest** | [**FetchControlListRequest**](FetchControlListRequest.md) |  | 

### Return type

[**FetchControlListResponse**](FetchControlListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchControlListES

> FetchControlListES(ctx).Execute()

Fetch List of Analytics ES



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchControlListES(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchControlListES``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchControlListESRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRuntimeControlsData

> FetchRuntimeControlsData(ctx).FetchRuntimeControlsDataRequest(fetchRuntimeControlsDataRequest).Execute()

Fetch Runtime Controls Data

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fetchRuntimeControlsDataRequest := *openapiclient.NewFetchRuntimeControlsDataRequest() // FetchRuntimeControlsDataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchRuntimeControlsData(context.Background()).FetchRuntimeControlsDataRequest(fetchRuntimeControlsDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchRuntimeControlsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRuntimeControlsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchRuntimeControlsDataRequest** | [**FetchRuntimeControlsDataRequest**](FetchRuntimeControlsDataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRuntimeControlsDataV2

> FetchRuntimeControlsDataV2(ctx).FetchRuntimeControlsDataV2Request(fetchRuntimeControlsDataV2Request).Execute()

Fetch Runtime Controls Data V2

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fetchRuntimeControlsDataV2Request := *openapiclient.NewFetchRuntimeControlsDataV2Request() // FetchRuntimeControlsDataV2Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchRuntimeControlsDataV2(context.Background()).FetchRuntimeControlsDataV2Request(fetchRuntimeControlsDataV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FetchRuntimeControlsDataV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRuntimeControlsDataV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchRuntimeControlsDataV2Request** | [**FetchRuntimeControlsDataV2Request**](FetchRuntimeControlsDataV2Request.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAnalyticsControls

> RunAnalyticsControls(ctx).Jobgroup(jobgroup).Jobname(jobname).AnalyticsCategories(analyticsCategories).AnalyticsApplications(analyticsApplications).Analyticssubcategories(analyticssubcategories).Execute()

Run Analytics Controls

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobgroup := "jobgroup_example" // string | 
	jobname := "jobname_example" // string | 
	analyticsCategories := "analyticsCategories_example" // string | 
	analyticsApplications := "analyticsApplications_example" // string |  (optional)
	analyticssubcategories := "analyticssubcategories_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.RunAnalyticsControls(context.Background()).Jobgroup(jobgroup).Jobname(jobname).AnalyticsCategories(analyticsCategories).AnalyticsApplications(analyticsApplications).Analyticssubcategories(analyticssubcategories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.RunAnalyticsControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunAnalyticsControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobgroup** | **string** |  | 
 **jobname** | **string** |  | 
 **analyticsCategories** | **string** |  | 
 **analyticsApplications** | **string** |  | 
 **analyticssubcategories** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

