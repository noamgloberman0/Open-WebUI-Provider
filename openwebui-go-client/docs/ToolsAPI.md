# \ToolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewToolsApiV1ToolsCreatePost**](ToolsAPI.md#CreateNewToolsApiV1ToolsCreatePost) | **Post** /api/v1/tools/create | Create New Tools
[**DeleteToolsByIdApiV1ToolsIdIdDeleteDelete**](ToolsAPI.md#DeleteToolsByIdApiV1ToolsIdIdDeleteDelete) | **Delete** /api/v1/tools/id/{id}/delete | Delete Tools By Id
[**ExportToolsApiV1ToolsExportGet**](ToolsAPI.md#ExportToolsApiV1ToolsExportGet) | **Get** /api/v1/tools/export | Export Tools
[**GetToolListApiV1ToolsListGet**](ToolsAPI.md#GetToolListApiV1ToolsListGet) | **Get** /api/v1/tools/list | Get Tool List
[**GetToolsApiV1ToolsGet**](ToolsAPI.md#GetToolsApiV1ToolsGet) | **Get** /api/v1/tools/ | Get Tools
[**GetToolsByIdApiV1ToolsIdIdGet**](ToolsAPI.md#GetToolsByIdApiV1ToolsIdIdGet) | **Get** /api/v1/tools/id/{id} | Get Tools By Id
[**GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet**](ToolsAPI.md#GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet) | **Get** /api/v1/tools/id/{id}/valves/user | Get Tools User Valves By Id
[**GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet**](ToolsAPI.md#GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet) | **Get** /api/v1/tools/id/{id}/valves/user/spec | Get Tools User Valves Spec By Id
[**GetToolsValvesByIdApiV1ToolsIdIdValvesGet**](ToolsAPI.md#GetToolsValvesByIdApiV1ToolsIdIdValvesGet) | **Get** /api/v1/tools/id/{id}/valves | Get Tools Valves By Id
[**GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet**](ToolsAPI.md#GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet) | **Get** /api/v1/tools/id/{id}/valves/spec | Get Tools Valves Spec By Id
[**UpdateToolsByIdApiV1ToolsIdIdUpdatePost**](ToolsAPI.md#UpdateToolsByIdApiV1ToolsIdIdUpdatePost) | **Post** /api/v1/tools/id/{id}/update | Update Tools By Id
[**UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost**](ToolsAPI.md#UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost) | **Post** /api/v1/tools/id/{id}/valves/user/update | Update Tools User Valves By Id
[**UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost**](ToolsAPI.md#UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost) | **Post** /api/v1/tools/id/{id}/valves/update | Update Tools Valves By Id



## CreateNewToolsApiV1ToolsCreatePost

> ToolResponse CreateNewToolsApiV1ToolsCreatePost(ctx).ToolForm(toolForm).Execute()

Create New Tools

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
	toolForm := *openapiclient.NewToolForm("Id_example", "Name_example", "Content_example", *openapiclient.NewToolMeta()) // ToolForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CreateNewToolsApiV1ToolsCreatePost(context.Background()).ToolForm(toolForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CreateNewToolsApiV1ToolsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewToolsApiV1ToolsCreatePost`: ToolResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CreateNewToolsApiV1ToolsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewToolsApiV1ToolsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolForm** | [**ToolForm**](ToolForm.md) |  | 

### Return type

[**ToolResponse**](ToolResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToolsByIdApiV1ToolsIdIdDeleteDelete

> bool DeleteToolsByIdApiV1ToolsIdIdDeleteDelete(ctx, id).Execute()

Delete Tools By Id

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
	resp, r, err := apiClient.ToolsAPI.DeleteToolsByIdApiV1ToolsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.DeleteToolsByIdApiV1ToolsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolsByIdApiV1ToolsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.DeleteToolsByIdApiV1ToolsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolsByIdApiV1ToolsIdIdDeleteDeleteRequest struct via the builder pattern


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


## ExportToolsApiV1ToolsExportGet

> []ToolModel ExportToolsApiV1ToolsExportGet(ctx).Execute()

Export Tools

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
	resp, r, err := apiClient.ToolsAPI.ExportToolsApiV1ToolsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ExportToolsApiV1ToolsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportToolsApiV1ToolsExportGet`: []ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ExportToolsApiV1ToolsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportToolsApiV1ToolsExportGetRequest struct via the builder pattern


### Return type

[**[]ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolListApiV1ToolsListGet

> []ToolUserResponse GetToolListApiV1ToolsListGet(ctx).Execute()

Get Tool List

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
	resp, r, err := apiClient.ToolsAPI.GetToolListApiV1ToolsListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolListApiV1ToolsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolListApiV1ToolsListGet`: []ToolUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolListApiV1ToolsListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolListApiV1ToolsListGetRequest struct via the builder pattern


### Return type

[**[]ToolUserResponse**](ToolUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsApiV1ToolsGet

> []ToolUserResponse GetToolsApiV1ToolsGet(ctx).Execute()

Get Tools

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
	resp, r, err := apiClient.ToolsAPI.GetToolsApiV1ToolsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsApiV1ToolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsApiV1ToolsGet`: []ToolUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsApiV1ToolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsApiV1ToolsGetRequest struct via the builder pattern


### Return type

[**[]ToolUserResponse**](ToolUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsByIdApiV1ToolsIdIdGet

> ToolModel GetToolsByIdApiV1ToolsIdIdGet(ctx, id).Execute()

Get Tools By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsByIdApiV1ToolsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsByIdApiV1ToolsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsByIdApiV1ToolsIdIdGet`: ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsByIdApiV1ToolsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsByIdApiV1ToolsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet

> map[string]interface{} GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet(ctx, id).Execute()

Get Tools User Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGetRequest struct via the builder pattern


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


## GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet

> map[string]interface{} GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet(ctx, id).Execute()

Get Tools User Valves Spec By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGetRequest struct via the builder pattern


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


## GetToolsValvesByIdApiV1ToolsIdIdValvesGet

> map[string]interface{} GetToolsValvesByIdApiV1ToolsIdIdValvesGet(ctx, id).Execute()

Get Tools Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsValvesByIdApiV1ToolsIdIdValvesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsValvesByIdApiV1ToolsIdIdValvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsValvesByIdApiV1ToolsIdIdValvesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsValvesByIdApiV1ToolsIdIdValvesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsValvesByIdApiV1ToolsIdIdValvesGetRequest struct via the builder pattern


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


## GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet

> map[string]interface{} GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet(ctx, id).Execute()

Get Tools Valves Spec By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGetRequest struct via the builder pattern


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


## UpdateToolsByIdApiV1ToolsIdIdUpdatePost

> ToolModel UpdateToolsByIdApiV1ToolsIdIdUpdatePost(ctx, id).ToolForm(toolForm).Execute()

Update Tools By Id

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
	toolForm := *openapiclient.NewToolForm("Id_example", "Name_example", "Content_example", *openapiclient.NewToolMeta()) // ToolForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.UpdateToolsByIdApiV1ToolsIdIdUpdatePost(context.Background(), id).ToolForm(toolForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsByIdApiV1ToolsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsByIdApiV1ToolsIdIdUpdatePost`: ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsByIdApiV1ToolsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsByIdApiV1ToolsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolForm** | [**ToolForm**](ToolForm.md) |  | 

### Return type

[**ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost

> map[string]interface{} UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost(ctx, id).Body(body).Execute()

Update Tools User Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePostRequest struct via the builder pattern


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


## UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost

> map[string]interface{} UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost(ctx, id).Body(body).Execute()

Update Tools Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePostRequest struct via the builder pattern


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

