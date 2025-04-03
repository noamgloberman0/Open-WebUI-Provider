# \FoldersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolderApiV1FoldersPost**](FoldersAPI.md#CreateFolderApiV1FoldersPost) | **Post** /api/v1/folders/ | Create Folder
[**DeleteFolderByIdApiV1FoldersIdDelete**](FoldersAPI.md#DeleteFolderByIdApiV1FoldersIdDelete) | **Delete** /api/v1/folders/{id} | Delete Folder By Id
[**GetFolderByIdApiV1FoldersIdGet**](FoldersAPI.md#GetFolderByIdApiV1FoldersIdGet) | **Get** /api/v1/folders/{id} | Get Folder By Id
[**GetFoldersApiV1FoldersGet**](FoldersAPI.md#GetFoldersApiV1FoldersGet) | **Get** /api/v1/folders/ | Get Folders
[**UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost**](FoldersAPI.md#UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost) | **Post** /api/v1/folders/{id}/update/expanded | Update Folder Is Expanded By Id
[**UpdateFolderNameByIdApiV1FoldersIdUpdatePost**](FoldersAPI.md#UpdateFolderNameByIdApiV1FoldersIdUpdatePost) | **Post** /api/v1/folders/{id}/update | Update Folder Name By Id
[**UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost**](FoldersAPI.md#UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost) | **Post** /api/v1/folders/{id}/update/parent | Update Folder Parent Id By Id



## CreateFolderApiV1FoldersPost

> interface{} CreateFolderApiV1FoldersPost(ctx).FolderForm(folderForm).Execute()

Create Folder

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
	folderForm := *openapiclient.NewFolderForm("Name_example") // FolderForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.CreateFolderApiV1FoldersPost(context.Background()).FolderForm(folderForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.CreateFolderApiV1FoldersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolderApiV1FoldersPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.CreateFolderApiV1FoldersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderApiV1FoldersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderForm** | [**FolderForm**](FolderForm.md) |  | 

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


## DeleteFolderByIdApiV1FoldersIdDelete

> interface{} DeleteFolderByIdApiV1FoldersIdDelete(ctx, id).Execute()

Delete Folder By Id

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
	resp, r, err := apiClient.FoldersAPI.DeleteFolderByIdApiV1FoldersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFolderByIdApiV1FoldersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolderByIdApiV1FoldersIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.DeleteFolderByIdApiV1FoldersIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderByIdApiV1FoldersIdDeleteRequest struct via the builder pattern


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


## GetFolderByIdApiV1FoldersIdGet

> FolderModel GetFolderByIdApiV1FoldersIdGet(ctx, id).Execute()

Get Folder By Id

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
	resp, r, err := apiClient.FoldersAPI.GetFolderByIdApiV1FoldersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolderByIdApiV1FoldersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByIdApiV1FoldersIdGet`: FolderModel
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolderByIdApiV1FoldersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByIdApiV1FoldersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderModel**](FolderModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFoldersApiV1FoldersGet

> []FolderModel GetFoldersApiV1FoldersGet(ctx).Execute()

Get Folders

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
	resp, r, err := apiClient.FoldersAPI.GetFoldersApiV1FoldersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFoldersApiV1FoldersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFoldersApiV1FoldersGet`: []FolderModel
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFoldersApiV1FoldersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersApiV1FoldersGetRequest struct via the builder pattern


### Return type

[**[]FolderModel**](FolderModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost

> interface{} UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost(ctx, id).FolderIsExpandedForm(folderIsExpandedForm).Execute()

Update Folder Is Expanded By Id

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
	folderIsExpandedForm := *openapiclient.NewFolderIsExpandedForm(false) // FolderIsExpandedForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost(context.Background(), id).FolderIsExpandedForm(folderIsExpandedForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderIsExpandedByIdApiV1FoldersIdUpdateExpandedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderIsExpandedForm** | [**FolderIsExpandedForm**](FolderIsExpandedForm.md) |  | 

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


## UpdateFolderNameByIdApiV1FoldersIdUpdatePost

> interface{} UpdateFolderNameByIdApiV1FoldersIdUpdatePost(ctx, id).FolderForm(folderForm).Execute()

Update Folder Name By Id

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
	folderForm := *openapiclient.NewFolderForm("Name_example") // FolderForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderNameByIdApiV1FoldersIdUpdatePost(context.Background(), id).FolderForm(folderForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderNameByIdApiV1FoldersIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderNameByIdApiV1FoldersIdUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderNameByIdApiV1FoldersIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderNameByIdApiV1FoldersIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderForm** | [**FolderForm**](FolderForm.md) |  | 

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


## UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost

> interface{} UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost(ctx, id).FolderParentIdForm(folderParentIdForm).Execute()

Update Folder Parent Id By Id

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
	folderParentIdForm := *openapiclient.NewFolderParentIdForm() // FolderParentIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost(context.Background(), id).FolderParentIdForm(folderParentIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderParentIdByIdApiV1FoldersIdUpdateParentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderParentIdForm** | [**FolderParentIdForm**](FolderParentIdForm.md) |  | 

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

