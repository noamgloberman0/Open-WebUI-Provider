# \FunctionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewFunctionApiV1FunctionsCreatePost**](FunctionsAPI.md#CreateNewFunctionApiV1FunctionsCreatePost) | **Post** /api/v1/functions/create | Create New Function
[**DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete**](FunctionsAPI.md#DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete) | **Delete** /api/v1/functions/id/{id}/delete | Delete Function By Id
[**GetFunctionByIdApiV1FunctionsIdIdGet**](FunctionsAPI.md#GetFunctionByIdApiV1FunctionsIdIdGet) | **Get** /api/v1/functions/id/{id} | Get Function By Id
[**GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet**](FunctionsAPI.md#GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet) | **Get** /api/v1/functions/id/{id}/valves/user | Get Function User Valves By Id
[**GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet**](FunctionsAPI.md#GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet) | **Get** /api/v1/functions/id/{id}/valves/user/spec | Get Function User Valves Spec By Id
[**GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet**](FunctionsAPI.md#GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet) | **Get** /api/v1/functions/id/{id}/valves | Get Function Valves By Id
[**GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet**](FunctionsAPI.md#GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet) | **Get** /api/v1/functions/id/{id}/valves/spec | Get Function Valves Spec By Id
[**GetFunctionsApiV1FunctionsExportGet**](FunctionsAPI.md#GetFunctionsApiV1FunctionsExportGet) | **Get** /api/v1/functions/export | Get Functions
[**GetFunctionsApiV1FunctionsGet**](FunctionsAPI.md#GetFunctionsApiV1FunctionsGet) | **Get** /api/v1/functions/ | Get Functions
[**ToggleFunctionByIdApiV1FunctionsIdIdTogglePost**](FunctionsAPI.md#ToggleFunctionByIdApiV1FunctionsIdIdTogglePost) | **Post** /api/v1/functions/id/{id}/toggle | Toggle Function By Id
[**ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost**](FunctionsAPI.md#ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost) | **Post** /api/v1/functions/id/{id}/toggle/global | Toggle Global By Id
[**UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost**](FunctionsAPI.md#UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost) | **Post** /api/v1/functions/id/{id}/update | Update Function By Id
[**UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost**](FunctionsAPI.md#UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost) | **Post** /api/v1/functions/id/{id}/valves/user/update | Update Function User Valves By Id
[**UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost**](FunctionsAPI.md#UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost) | **Post** /api/v1/functions/id/{id}/valves/update | Update Function Valves By Id



## CreateNewFunctionApiV1FunctionsCreatePost

> FunctionResponse CreateNewFunctionApiV1FunctionsCreatePost(ctx).FunctionForm(functionForm).Execute()

Create New Function

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
	functionForm := *openapiclient.NewFunctionForm("Id_example", "Name_example", "Content_example", *openapiclient.NewFunctionMeta()) // FunctionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.CreateNewFunctionApiV1FunctionsCreatePost(context.Background()).FunctionForm(functionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateNewFunctionApiV1FunctionsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewFunctionApiV1FunctionsCreatePost`: FunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateNewFunctionApiV1FunctionsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewFunctionApiV1FunctionsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionForm** | [**FunctionForm**](FunctionForm.md) |  | 

### Return type

[**FunctionResponse**](FunctionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete

> bool DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete(ctx, id).Execute()

Delete Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteFunctionByIdApiV1FunctionsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionByIdApiV1FunctionsIdIdDeleteDeleteRequest struct via the builder pattern


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


## GetFunctionByIdApiV1FunctionsIdIdGet

> FunctionModel GetFunctionByIdApiV1FunctionsIdIdGet(ctx, id).Execute()

Get Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionByIdApiV1FunctionsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionByIdApiV1FunctionsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionByIdApiV1FunctionsIdIdGet`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionByIdApiV1FunctionsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionByIdApiV1FunctionsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet

> map[string]interface{} GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet(ctx, id).Execute()

Get Function User Valves By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet

> map[string]interface{} GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet(ctx, id).Execute()

Get Function User Valves Spec By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionUserValvesSpecByIdApiV1FunctionsIdIdValvesUserSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet

> map[string]interface{} GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet(ctx, id).Execute()

Get Function Valves By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionValvesByIdApiV1FunctionsIdIdValvesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionValvesByIdApiV1FunctionsIdIdValvesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet

> map[string]interface{} GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet(ctx, id).Execute()

Get Function Valves Spec By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionValvesSpecByIdApiV1FunctionsIdIdValvesSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsApiV1FunctionsExportGet

> []FunctionModel GetFunctionsApiV1FunctionsExportGet(ctx).Execute()

Get Functions

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsApiV1FunctionsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsApiV1FunctionsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsApiV1FunctionsExportGet`: []FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsApiV1FunctionsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsApiV1FunctionsExportGetRequest struct via the builder pattern


### Return type

[**[]FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsApiV1FunctionsGet

> []FunctionResponse GetFunctionsApiV1FunctionsGet(ctx).Execute()

Get Functions

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsApiV1FunctionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsApiV1FunctionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsApiV1FunctionsGet`: []FunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsApiV1FunctionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsApiV1FunctionsGetRequest struct via the builder pattern


### Return type

[**[]FunctionResponse**](FunctionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFunctionByIdApiV1FunctionsIdIdTogglePost

> FunctionModel ToggleFunctionByIdApiV1FunctionsIdIdTogglePost(ctx, id).Execute()

Toggle Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.ToggleFunctionByIdApiV1FunctionsIdIdTogglePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ToggleFunctionByIdApiV1FunctionsIdIdTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFunctionByIdApiV1FunctionsIdIdTogglePost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ToggleFunctionByIdApiV1FunctionsIdIdTogglePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFunctionByIdApiV1FunctionsIdIdTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost

> FunctionModel ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost(ctx, id).Execute()

Toggle Global By Id

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
	resp, r, err := apiClient.FunctionsAPI.ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleGlobalByIdApiV1FunctionsIdIdToggleGlobalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost

> FunctionModel UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost(ctx, id).FunctionForm(functionForm).Execute()

Update Function By Id

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
	functionForm := *openapiclient.NewFunctionForm("Id_example", "Name_example", "Content_example", *openapiclient.NewFunctionMeta()) // FunctionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost(context.Background(), id).FunctionForm(functionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionByIdApiV1FunctionsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionByIdApiV1FunctionsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionForm** | [**FunctionForm**](FunctionForm.md) |  | 

### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost

> map[string]interface{} UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost(ctx, id).Body(body).Execute()

Update Function User Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionUserValvesByIdApiV1FunctionsIdIdValvesUserUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost

> map[string]interface{} UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost(ctx, id).Body(body).Execute()

Update Function Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionValvesByIdApiV1FunctionsIdIdValvesUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

