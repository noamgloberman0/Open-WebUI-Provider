# \KnowledgeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost**](KnowledgeAPI.md#AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost) | **Post** /api/v1/knowledge/{id}/file/add | Add File To Knowledge By Id
[**AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost**](KnowledgeAPI.md#AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost) | **Post** /api/v1/knowledge/{id}/files/batch/add | Add Files To Knowledge Batch
[**CreateNewKnowledgeApiV1KnowledgeCreatePost**](KnowledgeAPI.md#CreateNewKnowledgeApiV1KnowledgeCreatePost) | **Post** /api/v1/knowledge/create | Create New Knowledge
[**DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete**](KnowledgeAPI.md#DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete) | **Delete** /api/v1/knowledge/{id}/delete | Delete Knowledge By Id
[**GetKnowledgeApiV1KnowledgeGet**](KnowledgeAPI.md#GetKnowledgeApiV1KnowledgeGet) | **Get** /api/v1/knowledge/ | Get Knowledge
[**GetKnowledgeByIdApiV1KnowledgeIdGet**](KnowledgeAPI.md#GetKnowledgeByIdApiV1KnowledgeIdGet) | **Get** /api/v1/knowledge/{id} | Get Knowledge By Id
[**GetKnowledgeListApiV1KnowledgeListGet**](KnowledgeAPI.md#GetKnowledgeListApiV1KnowledgeListGet) | **Get** /api/v1/knowledge/list | Get Knowledge List
[**RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost**](KnowledgeAPI.md#RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost) | **Post** /api/v1/knowledge/{id}/file/remove | Remove File From Knowledge By Id
[**ResetKnowledgeByIdApiV1KnowledgeIdResetPost**](KnowledgeAPI.md#ResetKnowledgeByIdApiV1KnowledgeIdResetPost) | **Post** /api/v1/knowledge/{id}/reset | Reset Knowledge By Id
[**UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost**](KnowledgeAPI.md#UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost) | **Post** /api/v1/knowledge/{id}/file/update | Update File From Knowledge By Id
[**UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost**](KnowledgeAPI.md#UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost) | **Post** /api/v1/knowledge/{id}/update | Update Knowledge By Id



## AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost

> KnowledgeFilesResponse AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Add File To Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.AddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToKnowledgeByIdApiV1KnowledgeIdFileAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost

> KnowledgeFilesResponse AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Add Files To Knowledge Batch



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
	knowledgeFileIdForm := []openapiclient.KnowledgeFileIdForm{*openapiclient.NewKnowledgeFileIdForm("FileId_example")} // []KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.AddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFilesToKnowledgeBatchApiV1KnowledgeIdFilesBatchAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**[]KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewKnowledgeApiV1KnowledgeCreatePost

> KnowledgeResponse CreateNewKnowledgeApiV1KnowledgeCreatePost(ctx).KnowledgeForm(knowledgeForm).Execute()

Create New Knowledge

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
	knowledgeForm := *openapiclient.NewKnowledgeForm("Name_example", "Description_example") // KnowledgeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.CreateNewKnowledgeApiV1KnowledgeCreatePost(context.Background()).KnowledgeForm(knowledgeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.CreateNewKnowledgeApiV1KnowledgeCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewKnowledgeApiV1KnowledgeCreatePost`: KnowledgeResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.CreateNewKnowledgeApiV1KnowledgeCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewKnowledgeApiV1KnowledgeCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeForm** | [**KnowledgeForm**](KnowledgeForm.md) |  | 

### Return type

[**KnowledgeResponse**](KnowledgeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete

> bool DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete(ctx, id).Execute()

Delete Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.DeleteKnowledgeByIdApiV1KnowledgeIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgeByIdApiV1KnowledgeIdDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetKnowledgeApiV1KnowledgeGet

> []KnowledgeUserResponse GetKnowledgeApiV1KnowledgeGet(ctx).Execute()

Get Knowledge

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeApiV1KnowledgeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeApiV1KnowledgeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeApiV1KnowledgeGet`: []KnowledgeUserResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeApiV1KnowledgeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeApiV1KnowledgeGetRequest struct via the builder pattern


### Return type

[**[]KnowledgeUserResponse**](KnowledgeUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeByIdApiV1KnowledgeIdGet

> KnowledgeFilesResponse GetKnowledgeByIdApiV1KnowledgeIdGet(ctx, id).Execute()

Get Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeByIdApiV1KnowledgeIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeByIdApiV1KnowledgeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeByIdApiV1KnowledgeIdGet`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeByIdApiV1KnowledgeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeByIdApiV1KnowledgeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeListApiV1KnowledgeListGet

> []KnowledgeUserResponse GetKnowledgeListApiV1KnowledgeListGet(ctx).Execute()

Get Knowledge List

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeListApiV1KnowledgeListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeListApiV1KnowledgeListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeListApiV1KnowledgeListGet`: []KnowledgeUserResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeListApiV1KnowledgeListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeListApiV1KnowledgeListGetRequest struct via the builder pattern


### Return type

[**[]KnowledgeUserResponse**](KnowledgeUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost

> KnowledgeFilesResponse RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Remove File From Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.RemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFileFromKnowledgeByIdApiV1KnowledgeIdFileRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetKnowledgeByIdApiV1KnowledgeIdResetPost

> KnowledgeResponse ResetKnowledgeByIdApiV1KnowledgeIdResetPost(ctx, id).Execute()

Reset Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.ResetKnowledgeByIdApiV1KnowledgeIdResetPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.ResetKnowledgeByIdApiV1KnowledgeIdResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetKnowledgeByIdApiV1KnowledgeIdResetPost`: KnowledgeResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.ResetKnowledgeByIdApiV1KnowledgeIdResetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetKnowledgeByIdApiV1KnowledgeIdResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeResponse**](KnowledgeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost

> KnowledgeFilesResponse UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Update File From Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.UpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileFromKnowledgeByIdApiV1KnowledgeIdFileUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost

> KnowledgeFilesResponse UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost(ctx, id).KnowledgeForm(knowledgeForm).Execute()

Update Knowledge By Id

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
	knowledgeForm := *openapiclient.NewKnowledgeForm("Name_example", "Description_example") // KnowledgeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost(context.Background(), id).KnowledgeForm(knowledgeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.UpdateKnowledgeByIdApiV1KnowledgeIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgeByIdApiV1KnowledgeIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeForm** | [**KnowledgeForm**](KnowledgeForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

