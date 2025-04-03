# \ModelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewModelApiV1ModelsCreatePost**](ModelsAPI.md#CreateNewModelApiV1ModelsCreatePost) | **Post** /api/v1/models/create | Create New Model
[**DeleteAllModelsApiV1ModelsDeleteAllDelete**](ModelsAPI.md#DeleteAllModelsApiV1ModelsDeleteAllDelete) | **Delete** /api/v1/models/delete/all | Delete All Models
[**DeleteModelByIdApiV1ModelsModelDeleteDelete**](ModelsAPI.md#DeleteModelByIdApiV1ModelsModelDeleteDelete) | **Delete** /api/v1/models/model/delete | Delete Model By Id
[**GetBaseModelsApiV1ModelsBaseGet**](ModelsAPI.md#GetBaseModelsApiV1ModelsBaseGet) | **Get** /api/v1/models/base | Get Base Models
[**GetModelByIdApiV1ModelsModelGet**](ModelsAPI.md#GetModelByIdApiV1ModelsModelGet) | **Get** /api/v1/models/model | Get Model By Id
[**GetModelsApiV1ModelsGet**](ModelsAPI.md#GetModelsApiV1ModelsGet) | **Get** /api/v1/models/ | Get Models
[**ToggleModelByIdApiV1ModelsModelTogglePost**](ModelsAPI.md#ToggleModelByIdApiV1ModelsModelTogglePost) | **Post** /api/v1/models/model/toggle | Toggle Model By Id
[**UpdateModelByIdApiV1ModelsModelUpdatePost**](ModelsAPI.md#UpdateModelByIdApiV1ModelsModelUpdatePost) | **Post** /api/v1/models/model/update | Update Model By Id



## CreateNewModelApiV1ModelsCreatePost

> ModelModel CreateNewModelApiV1ModelsCreatePost(ctx).ModelForm(modelForm).Execute()

Create New Model

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
	modelForm := *openapiclient.NewModelForm("Id_example", "Name_example", *openapiclient.NewModelMeta(), map[string]interface{}{"key": interface{}(123)}) // ModelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.CreateNewModelApiV1ModelsCreatePost(context.Background()).ModelForm(modelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.CreateNewModelApiV1ModelsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewModelApiV1ModelsCreatePost`: ModelModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.CreateNewModelApiV1ModelsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewModelApiV1ModelsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelForm** | [**ModelForm**](ModelForm.md) |  | 

### Return type

[**ModelModel**](ModelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllModelsApiV1ModelsDeleteAllDelete

> bool DeleteAllModelsApiV1ModelsDeleteAllDelete(ctx).Execute()

Delete All Models

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
	resp, r, err := apiClient.ModelsAPI.DeleteAllModelsApiV1ModelsDeleteAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.DeleteAllModelsApiV1ModelsDeleteAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllModelsApiV1ModelsDeleteAllDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.DeleteAllModelsApiV1ModelsDeleteAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllModelsApiV1ModelsDeleteAllDeleteRequest struct via the builder pattern


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


## DeleteModelByIdApiV1ModelsModelDeleteDelete

> bool DeleteModelByIdApiV1ModelsModelDeleteDelete(ctx).Id(id).Execute()

Delete Model By Id

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
	resp, r, err := apiClient.ModelsAPI.DeleteModelByIdApiV1ModelsModelDeleteDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.DeleteModelByIdApiV1ModelsModelDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModelByIdApiV1ModelsModelDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.DeleteModelByIdApiV1ModelsModelDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModelByIdApiV1ModelsModelDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

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


## GetBaseModelsApiV1ModelsBaseGet

> []ModelResponse GetBaseModelsApiV1ModelsBaseGet(ctx).Execute()

Get Base Models

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
	resp, r, err := apiClient.ModelsAPI.GetBaseModelsApiV1ModelsBaseGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetBaseModelsApiV1ModelsBaseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseModelsApiV1ModelsBaseGet`: []ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetBaseModelsApiV1ModelsBaseGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseModelsApiV1ModelsBaseGetRequest struct via the builder pattern


### Return type

[**[]ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelByIdApiV1ModelsModelGet

> ModelResponse GetModelByIdApiV1ModelsModelGet(ctx).Id(id).Execute()

Get Model By Id

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
	resp, r, err := apiClient.ModelsAPI.GetModelByIdApiV1ModelsModelGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelByIdApiV1ModelsModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelByIdApiV1ModelsModelGet`: ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetModelByIdApiV1ModelsModelGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelByIdApiV1ModelsModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsApiV1ModelsGet

> []ModelUserResponse GetModelsApiV1ModelsGet(ctx).Id(id).Execute()

Get Models

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
	id := "id_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.GetModelsApiV1ModelsGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelsApiV1ModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsApiV1ModelsGet`: []ModelUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetModelsApiV1ModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsApiV1ModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**[]ModelUserResponse**](ModelUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleModelByIdApiV1ModelsModelTogglePost

> ModelResponse ToggleModelByIdApiV1ModelsModelTogglePost(ctx).Id(id).Execute()

Toggle Model By Id

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
	resp, r, err := apiClient.ModelsAPI.ToggleModelByIdApiV1ModelsModelTogglePost(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ToggleModelByIdApiV1ModelsModelTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleModelByIdApiV1ModelsModelTogglePost`: ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ToggleModelByIdApiV1ModelsModelTogglePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleModelByIdApiV1ModelsModelTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelByIdApiV1ModelsModelUpdatePost

> ModelModel UpdateModelByIdApiV1ModelsModelUpdatePost(ctx).Id(id).ModelForm(modelForm).Execute()

Update Model By Id

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
	modelForm := *openapiclient.NewModelForm("Id_example", "Name_example", *openapiclient.NewModelMeta(), map[string]interface{}{"key": interface{}(123)}) // ModelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.UpdateModelByIdApiV1ModelsModelUpdatePost(context.Background()).Id(id).ModelForm(modelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.UpdateModelByIdApiV1ModelsModelUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateModelByIdApiV1ModelsModelUpdatePost`: ModelModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.UpdateModelByIdApiV1ModelsModelUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModelByIdApiV1ModelsModelUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **modelForm** | [**ModelForm**](ModelForm.md) |  | 

### Return type

[**ModelModel**](ModelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

