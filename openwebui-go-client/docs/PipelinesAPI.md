# \PipelinesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPipelineApiV1PipelinesAddPost**](PipelinesAPI.md#AddPipelineApiV1PipelinesAddPost) | **Post** /api/v1/pipelines/add | Add Pipeline
[**DeletePipelineApiV1PipelinesDeleteDelete**](PipelinesAPI.md#DeletePipelineApiV1PipelinesDeleteDelete) | **Delete** /api/v1/pipelines/delete | Delete Pipeline
[**GetPipelineValvesApiV1PipelinesPipelineIdValvesGet**](PipelinesAPI.md#GetPipelineValvesApiV1PipelinesPipelineIdValvesGet) | **Get** /api/v1/pipelines/{pipeline_id}/valves | Get Pipeline Valves
[**GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet**](PipelinesAPI.md#GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet) | **Get** /api/v1/pipelines/{pipeline_id}/valves/spec | Get Pipeline Valves Spec
[**GetPipelinesApiV1PipelinesGet**](PipelinesAPI.md#GetPipelinesApiV1PipelinesGet) | **Get** /api/v1/pipelines/ | Get Pipelines
[**GetPipelinesListApiV1PipelinesListGet**](PipelinesAPI.md#GetPipelinesListApiV1PipelinesListGet) | **Get** /api/v1/pipelines/list | Get Pipelines List
[**UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost**](PipelinesAPI.md#UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost) | **Post** /api/v1/pipelines/{pipeline_id}/valves/update | Update Pipeline Valves
[**UploadPipelineApiV1PipelinesUploadPost**](PipelinesAPI.md#UploadPipelineApiV1PipelinesUploadPost) | **Post** /api/v1/pipelines/upload | Upload Pipeline



## AddPipelineApiV1PipelinesAddPost

> interface{} AddPipelineApiV1PipelinesAddPost(ctx).AddPipelineForm(addPipelineForm).Execute()

Add Pipeline

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
	addPipelineForm := *openapiclient.NewAddPipelineForm("Url_example", int32(123)) // AddPipelineForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.AddPipelineApiV1PipelinesAddPost(context.Background()).AddPipelineForm(addPipelineForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.AddPipelineApiV1PipelinesAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPipelineApiV1PipelinesAddPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.AddPipelineApiV1PipelinesAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPipelineApiV1PipelinesAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPipelineForm** | [**AddPipelineForm**](AddPipelineForm.md) |  | 

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


## DeletePipelineApiV1PipelinesDeleteDelete

> interface{} DeletePipelineApiV1PipelinesDeleteDelete(ctx).DeletePipelineForm(deletePipelineForm).Execute()

Delete Pipeline

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
	deletePipelineForm := *openapiclient.NewDeletePipelineForm("Id_example", int32(123)) // DeletePipelineForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.DeletePipelineApiV1PipelinesDeleteDelete(context.Background()).DeletePipelineForm(deletePipelineForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.DeletePipelineApiV1PipelinesDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePipelineApiV1PipelinesDeleteDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.DeletePipelineApiV1PipelinesDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineApiV1PipelinesDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePipelineForm** | [**DeletePipelineForm**](DeletePipelineForm.md) |  | 

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


## GetPipelineValvesApiV1PipelinesPipelineIdValvesGet

> interface{} GetPipelineValvesApiV1PipelinesPipelineIdValvesGet(ctx, pipelineId).UrlIdx(urlIdx).Execute()

Get Pipeline Valves

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
	pipelineId := "pipelineId_example" // string | 
	urlIdx := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.GetPipelineValvesApiV1PipelinesPipelineIdValvesGet(context.Background(), pipelineId).UrlIdx(urlIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetPipelineValvesApiV1PipelinesPipelineIdValvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineValvesApiV1PipelinesPipelineIdValvesGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetPipelineValvesApiV1PipelinesPipelineIdValvesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineValvesApiV1PipelinesPipelineIdValvesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlIdx** | **int32** |  | 

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


## GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet

> interface{} GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet(ctx, pipelineId).UrlIdx(urlIdx).Execute()

Get Pipeline Valves Spec

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
	pipelineId := "pipelineId_example" // string | 
	urlIdx := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet(context.Background(), pipelineId).UrlIdx(urlIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineValvesSpecApiV1PipelinesPipelineIdValvesSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlIdx** | **int32** |  | 

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


## GetPipelinesApiV1PipelinesGet

> interface{} GetPipelinesApiV1PipelinesGet(ctx).UrlIdx(urlIdx).Execute()

Get Pipelines

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
	urlIdx := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.GetPipelinesApiV1PipelinesGet(context.Background()).UrlIdx(urlIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetPipelinesApiV1PipelinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesApiV1PipelinesGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetPipelinesApiV1PipelinesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesApiV1PipelinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlIdx** | **int32** |  | 

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


## GetPipelinesListApiV1PipelinesListGet

> interface{} GetPipelinesListApiV1PipelinesListGet(ctx).Execute()

Get Pipelines List

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
	resp, r, err := apiClient.PipelinesAPI.GetPipelinesListApiV1PipelinesListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetPipelinesListApiV1PipelinesListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelinesListApiV1PipelinesListGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetPipelinesListApiV1PipelinesListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesListApiV1PipelinesListGetRequest struct via the builder pattern


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


## UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost

> interface{} UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost(ctx, pipelineId).UrlIdx(urlIdx).Body(body).Execute()

Update Pipeline Valves

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
	pipelineId := "pipelineId_example" // string | 
	urlIdx := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost(context.Background(), pipelineId).UrlIdx(urlIdx).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.UpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineValvesApiV1PipelinesPipelineIdValvesUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlIdx** | **int32** |  | 
 **body** | **map[string]interface{}** |  | 

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


## UploadPipelineApiV1PipelinesUploadPost

> interface{} UploadPipelineApiV1PipelinesUploadPost(ctx).UrlIdx(urlIdx).File(file).Execute()

Upload Pipeline

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
	urlIdx := int32(56) // int32 | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.UploadPipelineApiV1PipelinesUploadPost(context.Background()).UrlIdx(urlIdx).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.UploadPipelineApiV1PipelinesUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPipelineApiV1PipelinesUploadPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.UploadPipelineApiV1PipelinesUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPipelineApiV1PipelinesUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlIdx** | **int32** |  | 
 **file** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

