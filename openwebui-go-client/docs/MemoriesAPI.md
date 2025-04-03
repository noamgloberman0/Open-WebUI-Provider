# \MemoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemoryApiV1MemoriesAddPost**](MemoriesAPI.md#AddMemoryApiV1MemoriesAddPost) | **Post** /api/v1/memories/add | Add Memory
[**DeleteMemoryByIdApiV1MemoriesMemoryIdDelete**](MemoriesAPI.md#DeleteMemoryByIdApiV1MemoriesMemoryIdDelete) | **Delete** /api/v1/memories/{memory_id} | Delete Memory By Id
[**DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete**](MemoriesAPI.md#DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete) | **Delete** /api/v1/memories/delete/user | Delete Memory By User Id
[**GetEmbeddingsApiV1MemoriesEfGet**](MemoriesAPI.md#GetEmbeddingsApiV1MemoriesEfGet) | **Get** /api/v1/memories/ef | Get Embeddings
[**GetMemoriesApiV1MemoriesGet**](MemoriesAPI.md#GetMemoriesApiV1MemoriesGet) | **Get** /api/v1/memories/ | Get Memories
[**QueryMemoryApiV1MemoriesQueryPost**](MemoriesAPI.md#QueryMemoryApiV1MemoriesQueryPost) | **Post** /api/v1/memories/query | Query Memory
[**ResetMemoryFromVectorDbApiV1MemoriesResetPost**](MemoriesAPI.md#ResetMemoryFromVectorDbApiV1MemoriesResetPost) | **Post** /api/v1/memories/reset | Reset Memory From Vector Db
[**UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost**](MemoriesAPI.md#UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost) | **Post** /api/v1/memories/{memory_id}/update | Update Memory By Id



## AddMemoryApiV1MemoriesAddPost

> MemoryModel AddMemoryApiV1MemoriesAddPost(ctx).AddMemoryForm(addMemoryForm).Execute()

Add Memory

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
	addMemoryForm := *openapiclient.NewAddMemoryForm("Content_example") // AddMemoryForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.AddMemoryApiV1MemoriesAddPost(context.Background()).AddMemoryForm(addMemoryForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.AddMemoryApiV1MemoriesAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMemoryApiV1MemoriesAddPost`: MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.AddMemoryApiV1MemoriesAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMemoryApiV1MemoriesAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addMemoryForm** | [**AddMemoryForm**](AddMemoryForm.md) |  | 

### Return type

[**MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemoryByIdApiV1MemoriesMemoryIdDelete

> bool DeleteMemoryByIdApiV1MemoriesMemoryIdDelete(ctx, memoryId).Execute()

Delete Memory By Id

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
	memoryId := "memoryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.DeleteMemoryByIdApiV1MemoriesMemoryIdDelete(context.Background(), memoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.DeleteMemoryByIdApiV1MemoriesMemoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemoryByIdApiV1MemoriesMemoryIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.DeleteMemoryByIdApiV1MemoriesMemoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemoryByIdApiV1MemoriesMemoryIdDeleteRequest struct via the builder pattern


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


## DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete

> bool DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete(ctx).Execute()

Delete Memory By User Id

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
	resp, r, err := apiClient.MemoriesAPI.DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemoryByUserIdApiV1MemoriesDeleteUserDeleteRequest struct via the builder pattern


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


## GetEmbeddingsApiV1MemoriesEfGet

> interface{} GetEmbeddingsApiV1MemoriesEfGet(ctx).Execute()

Get Embeddings

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
	resp, r, err := apiClient.MemoriesAPI.GetEmbeddingsApiV1MemoriesEfGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.GetEmbeddingsApiV1MemoriesEfGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingsApiV1MemoriesEfGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.GetEmbeddingsApiV1MemoriesEfGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingsApiV1MemoriesEfGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemoriesApiV1MemoriesGet

> []MemoryModel GetMemoriesApiV1MemoriesGet(ctx).Execute()

Get Memories

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
	resp, r, err := apiClient.MemoriesAPI.GetMemoriesApiV1MemoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.GetMemoriesApiV1MemoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMemoriesApiV1MemoriesGet`: []MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.GetMemoriesApiV1MemoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoriesApiV1MemoriesGetRequest struct via the builder pattern


### Return type

[**[]MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMemoryApiV1MemoriesQueryPost

> interface{} QueryMemoryApiV1MemoriesQueryPost(ctx).QueryMemoryForm(queryMemoryForm).Execute()

Query Memory

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
	queryMemoryForm := *openapiclient.NewQueryMemoryForm("Content_example") // QueryMemoryForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.QueryMemoryApiV1MemoriesQueryPost(context.Background()).QueryMemoryForm(queryMemoryForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.QueryMemoryApiV1MemoriesQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMemoryApiV1MemoriesQueryPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.QueryMemoryApiV1MemoriesQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMemoryApiV1MemoriesQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMemoryForm** | [**QueryMemoryForm**](QueryMemoryForm.md) |  | 

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


## ResetMemoryFromVectorDbApiV1MemoriesResetPost

> bool ResetMemoryFromVectorDbApiV1MemoriesResetPost(ctx).Execute()

Reset Memory From Vector Db

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
	resp, r, err := apiClient.MemoriesAPI.ResetMemoryFromVectorDbApiV1MemoriesResetPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.ResetMemoryFromVectorDbApiV1MemoriesResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetMemoryFromVectorDbApiV1MemoriesResetPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.ResetMemoryFromVectorDbApiV1MemoriesResetPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetMemoryFromVectorDbApiV1MemoriesResetPostRequest struct via the builder pattern


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


## UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost

> MemoryModel UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost(ctx, memoryId).MemoryUpdateModel(memoryUpdateModel).Execute()

Update Memory By Id

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
	memoryId := "memoryId_example" // string | 
	memoryUpdateModel := *openapiclient.NewMemoryUpdateModel() // MemoryUpdateModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost(context.Background(), memoryId).MemoryUpdateModel(memoryUpdateModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost`: MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryUpdateModel** | [**MemoryUpdateModel**](MemoryUpdateModel.md) |  | 

### Return type

[**MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

