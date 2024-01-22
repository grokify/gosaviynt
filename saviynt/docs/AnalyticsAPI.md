# \AnalyticsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRuntimeControlsDataV2**](AnalyticsAPI.md#FetchRuntimeControlsDataV2) | **Post** /fetchRuntimeControlsDataV2 | Fetch Runtime Controls Data V2



## FetchRuntimeControlsDataV2

> FetchRuntimeControlsDataV2(ctx).FetchRuntimeControlsDataV2(fetchRuntimeControlsDataV2).Execute()

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
	fetchRuntimeControlsDataV2 := *openapiclient.NewFetchRuntimeControlsDataV2() // FetchRuntimeControlsDataV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.FetchRuntimeControlsDataV2(context.Background()).FetchRuntimeControlsDataV2(fetchRuntimeControlsDataV2).Execute()
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
 **fetchRuntimeControlsDataV2** | [**FetchRuntimeControlsDataV2**](FetchRuntimeControlsDataV2.md) |  | 

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

