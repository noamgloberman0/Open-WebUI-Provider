# \EvaluationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedbackApiV1EvaluationsFeedbackPost**](EvaluationsAPI.md#CreateFeedbackApiV1EvaluationsFeedbackPost) | **Post** /api/v1/evaluations/feedback | Create Feedback
[**DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete**](EvaluationsAPI.md#DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete) | **Delete** /api/v1/evaluations/feedbacks/all | Delete All Feedbacks
[**DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete**](EvaluationsAPI.md#DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete) | **Delete** /api/v1/evaluations/feedback/{id} | Delete Feedback By Id
[**DeleteFeedbacksApiV1EvaluationsFeedbacksDelete**](EvaluationsAPI.md#DeleteFeedbacksApiV1EvaluationsFeedbacksDelete) | **Delete** /api/v1/evaluations/feedbacks | Delete Feedbacks
[**GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet**](EvaluationsAPI.md#GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet) | **Get** /api/v1/evaluations/feedbacks/all/export | Get All Feedbacks
[**GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet**](EvaluationsAPI.md#GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet) | **Get** /api/v1/evaluations/feedbacks/all | Get All Feedbacks
[**GetConfigApiV1EvaluationsConfigGet**](EvaluationsAPI.md#GetConfigApiV1EvaluationsConfigGet) | **Get** /api/v1/evaluations/config | Get Config
[**GetFeedbackByIdApiV1EvaluationsFeedbackIdGet**](EvaluationsAPI.md#GetFeedbackByIdApiV1EvaluationsFeedbackIdGet) | **Get** /api/v1/evaluations/feedback/{id} | Get Feedback By Id
[**GetFeedbacksApiV1EvaluationsFeedbacksUserGet**](EvaluationsAPI.md#GetFeedbacksApiV1EvaluationsFeedbacksUserGet) | **Get** /api/v1/evaluations/feedbacks/user | Get Feedbacks
[**UpdateConfigApiV1EvaluationsConfigPost**](EvaluationsAPI.md#UpdateConfigApiV1EvaluationsConfigPost) | **Post** /api/v1/evaluations/config | Update Config
[**UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost**](EvaluationsAPI.md#UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost) | **Post** /api/v1/evaluations/feedback/{id} | Update Feedback By Id



## CreateFeedbackApiV1EvaluationsFeedbackPost

> FeedbackModel CreateFeedbackApiV1EvaluationsFeedbackPost(ctx).FeedbackForm(feedbackForm).Execute()

Create Feedback

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
	feedbackForm := *openapiclient.NewFeedbackForm("Type_example") // FeedbackForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.CreateFeedbackApiV1EvaluationsFeedbackPost(context.Background()).FeedbackForm(feedbackForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.CreateFeedbackApiV1EvaluationsFeedbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedbackApiV1EvaluationsFeedbackPost`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.CreateFeedbackApiV1EvaluationsFeedbackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackApiV1EvaluationsFeedbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackForm** | [**FeedbackForm**](FeedbackForm.md) |  | 

### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete

> interface{} DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete(ctx).Execute()

Delete All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllFeedbacksApiV1EvaluationsFeedbacksAllDeleteRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete

> interface{} DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete(ctx, id).Execute()

Delete Feedback By Id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteFeedbackByIdApiV1EvaluationsFeedbackIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedbackByIdApiV1EvaluationsFeedbackIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeedbacksApiV1EvaluationsFeedbacksDelete

> bool DeleteFeedbacksApiV1EvaluationsFeedbacksDelete(ctx).Execute()

Delete Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.DeleteFeedbacksApiV1EvaluationsFeedbacksDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteFeedbacksApiV1EvaluationsFeedbacksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedbacksApiV1EvaluationsFeedbacksDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteFeedbacksApiV1EvaluationsFeedbacksDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedbacksApiV1EvaluationsFeedbacksDeleteRequest struct via the builder pattern


### Return type

**bool**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet

> []FeedbackModel GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet(ctx).Execute()

Get All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet`: []FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeedbacksApiV1EvaluationsFeedbacksAllExportGetRequest struct via the builder pattern


### Return type

[**[]FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet

> []FeedbackUserResponse GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet(ctx).Execute()

Get All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet`: []FeedbackUserResponse
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetAllFeedbacksApiV1EvaluationsFeedbacksAllGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeedbacksApiV1EvaluationsFeedbacksAllGetRequest struct via the builder pattern


### Return type

[**[]FeedbackUserResponse**](FeedbackUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigApiV1EvaluationsConfigGet

> interface{} GetConfigApiV1EvaluationsConfigGet(ctx).Execute()

Get Config

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
	resp, r, err := apiClient.EvaluationsAPI.GetConfigApiV1EvaluationsConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetConfigApiV1EvaluationsConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigApiV1EvaluationsConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetConfigApiV1EvaluationsConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigApiV1EvaluationsConfigGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedbackByIdApiV1EvaluationsFeedbackIdGet

> FeedbackModel GetFeedbackByIdApiV1EvaluationsFeedbackIdGet(ctx, id).Execute()

Get Feedback By Id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.GetFeedbackByIdApiV1EvaluationsFeedbackIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetFeedbackByIdApiV1EvaluationsFeedbackIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedbackByIdApiV1EvaluationsFeedbackIdGet`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetFeedbackByIdApiV1EvaluationsFeedbackIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackByIdApiV1EvaluationsFeedbackIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedbacksApiV1EvaluationsFeedbacksUserGet

> []FeedbackUserResponse GetFeedbacksApiV1EvaluationsFeedbacksUserGet(ctx).Execute()

Get Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetFeedbacksApiV1EvaluationsFeedbacksUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetFeedbacksApiV1EvaluationsFeedbacksUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedbacksApiV1EvaluationsFeedbacksUserGet`: []FeedbackUserResponse
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetFeedbacksApiV1EvaluationsFeedbacksUserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbacksApiV1EvaluationsFeedbacksUserGetRequest struct via the builder pattern


### Return type

[**[]FeedbackUserResponse**](FeedbackUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigApiV1EvaluationsConfigPost

> interface{} UpdateConfigApiV1EvaluationsConfigPost(ctx).UpdateConfigForm(updateConfigForm).Execute()

Update Config

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
	updateConfigForm := *openapiclient.NewUpdateConfigForm() // UpdateConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.UpdateConfigApiV1EvaluationsConfigPost(context.Background()).UpdateConfigForm(updateConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.UpdateConfigApiV1EvaluationsConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigApiV1EvaluationsConfigPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.UpdateConfigApiV1EvaluationsConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigApiV1EvaluationsConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigForm** | [**UpdateConfigForm**](UpdateConfigForm.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost

> FeedbackModel UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost(ctx, id).FeedbackForm(feedbackForm).Execute()

Update Feedback By Id

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
	id := "id_example" // string | 
	feedbackForm := *openapiclient.NewFeedbackForm("Type_example") // FeedbackForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost(context.Background(), id).FeedbackForm(feedbackForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.UpdateFeedbackByIdApiV1EvaluationsFeedbackIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedbackByIdApiV1EvaluationsFeedbackIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackForm** | [**FeedbackForm**](FeedbackForm.md) |  | 

### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

