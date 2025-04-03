# \FilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllFilesApiV1FilesAllDelete**](FilesAPI.md#DeleteAllFilesApiV1FilesAllDelete) | **Delete** /api/v1/files/all | Delete All Files
[**DeleteFileByIdApiV1FilesIdDelete**](FilesAPI.md#DeleteFileByIdApiV1FilesIdDelete) | **Delete** /api/v1/files/{id} | Delete File By Id
[**GetFileByIdApiV1FilesIdGet**](FilesAPI.md#GetFileByIdApiV1FilesIdGet) | **Get** /api/v1/files/{id} | Get File By Id
[**GetFileContentByIdApiV1FilesIdContentFileNameGet**](FilesAPI.md#GetFileContentByIdApiV1FilesIdContentFileNameGet) | **Get** /api/v1/files/{id}/content/{file_name} | Get File Content By Id
[**GetFileContentByIdApiV1FilesIdContentGet**](FilesAPI.md#GetFileContentByIdApiV1FilesIdContentGet) | **Get** /api/v1/files/{id}/content | Get File Content By Id
[**GetFileDataContentByIdApiV1FilesIdDataContentGet**](FilesAPI.md#GetFileDataContentByIdApiV1FilesIdDataContentGet) | **Get** /api/v1/files/{id}/data/content | Get File Data Content By Id
[**GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet**](FilesAPI.md#GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet) | **Get** /api/v1/files/{id}/content/html | Get Html File Content By Id
[**ListFilesApiV1FilesGet**](FilesAPI.md#ListFilesApiV1FilesGet) | **Get** /api/v1/files/ | List Files
[**UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost**](FilesAPI.md#UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost) | **Post** /api/v1/files/{id}/data/content/update | Update File Data Content By Id
[**UploadFileApiV1FilesPost**](FilesAPI.md#UploadFileApiV1FilesPost) | **Post** /api/v1/files/ | Upload File



## DeleteAllFilesApiV1FilesAllDelete

> interface{} DeleteAllFilesApiV1FilesAllDelete(ctx).Execute()

Delete All Files

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
	resp, r, err := apiClient.FilesAPI.DeleteAllFilesApiV1FilesAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteAllFilesApiV1FilesAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllFilesApiV1FilesAllDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteAllFilesApiV1FilesAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllFilesApiV1FilesAllDeleteRequest struct via the builder pattern


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


## DeleteFileByIdApiV1FilesIdDelete

> interface{} DeleteFileByIdApiV1FilesIdDelete(ctx, id).Execute()

Delete File By Id

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
	resp, r, err := apiClient.FilesAPI.DeleteFileByIdApiV1FilesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFileByIdApiV1FilesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileByIdApiV1FilesIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFileByIdApiV1FilesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileByIdApiV1FilesIdDeleteRequest struct via the builder pattern


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


## GetFileByIdApiV1FilesIdGet

> FileModel GetFileByIdApiV1FilesIdGet(ctx, id).Execute()

Get File By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileByIdApiV1FilesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileByIdApiV1FilesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileByIdApiV1FilesIdGet`: FileModel
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileByIdApiV1FilesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileByIdApiV1FilesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileModel**](FileModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileContentByIdApiV1FilesIdContentFileNameGet

> interface{} GetFileContentByIdApiV1FilesIdContentFileNameGet(ctx, id).Execute()

Get File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileContentByIdApiV1FilesIdContentFileNameGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileContentByIdApiV1FilesIdContentFileNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileContentByIdApiV1FilesIdContentFileNameGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileContentByIdApiV1FilesIdContentFileNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileContentByIdApiV1FilesIdContentFileNameGetRequest struct via the builder pattern


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


## GetFileContentByIdApiV1FilesIdContentGet

> interface{} GetFileContentByIdApiV1FilesIdContentGet(ctx, id).Execute()

Get File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileContentByIdApiV1FilesIdContentGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileContentByIdApiV1FilesIdContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileContentByIdApiV1FilesIdContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileContentByIdApiV1FilesIdContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileContentByIdApiV1FilesIdContentGetRequest struct via the builder pattern


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


## GetFileDataContentByIdApiV1FilesIdDataContentGet

> interface{} GetFileDataContentByIdApiV1FilesIdDataContentGet(ctx, id).Execute()

Get File Data Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileDataContentByIdApiV1FilesIdDataContentGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileDataContentByIdApiV1FilesIdDataContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileDataContentByIdApiV1FilesIdDataContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileDataContentByIdApiV1FilesIdDataContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDataContentByIdApiV1FilesIdDataContentGetRequest struct via the builder pattern


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


## GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet

> interface{} GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet(ctx, id).Execute()

Get Html File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetHtmlFileContentByIdApiV1FilesIdContentHtmlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHtmlFileContentByIdApiV1FilesIdContentHtmlGetRequest struct via the builder pattern


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


## ListFilesApiV1FilesGet

> []FileModelResponse ListFilesApiV1FilesGet(ctx).Execute()

List Files

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
	resp, r, err := apiClient.FilesAPI.ListFilesApiV1FilesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ListFilesApiV1FilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFilesApiV1FilesGet`: []FileModelResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ListFilesApiV1FilesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesApiV1FilesGetRequest struct via the builder pattern


### Return type

[**[]FileModelResponse**](FileModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost

> interface{} UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost(ctx, id).ContentForm(contentForm).Execute()

Update File Data Content By Id

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
	contentForm := *openapiclient.NewContentForm("Content_example") // ContentForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost(context.Background(), id).ContentForm(contentForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileDataContentByIdApiV1FilesIdDataContentUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentForm** | [**ContentForm**](ContentForm.md) |  | 

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


## UploadFileApiV1FilesPost

> FileModelResponse UploadFileApiV1FilesPost(ctx).File(file).FileMetadata(fileMetadata).Execute()

Upload File

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
	file := os.NewFile(1234, "some_file") // *os.File | 
	fileMetadata := map[string]interface{}{ ... } // map[string]interface{} |  (optional) (default to {})

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFileApiV1FilesPost(context.Background()).File(file).FileMetadata(fileMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFileApiV1FilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileApiV1FilesPost`: FileModelResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFileApiV1FilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileApiV1FilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **fileMetadata** | [**map[string]interface{}**](map[string]interface{}.md) |  | [default to {}]

### Return type

[**FileModelResponse**](FileModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

