# \RetrievalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEntriesFromCollectionApiV1RetrievalDeletePost**](RetrievalAPI.md#DeleteEntriesFromCollectionApiV1RetrievalDeletePost) | **Post** /api/v1/retrieval/delete | Delete Entries From Collection
[**GetEmbeddingConfigApiV1RetrievalEmbeddingGet**](RetrievalAPI.md#GetEmbeddingConfigApiV1RetrievalEmbeddingGet) | **Get** /api/v1/retrieval/embedding | Get Embedding Config
[**GetEmbeddingsApiV1RetrievalEfTextGet**](RetrievalAPI.md#GetEmbeddingsApiV1RetrievalEfTextGet) | **Get** /api/v1/retrieval/ef/{text} | Get Embeddings
[**GetQuerySettingsApiV1RetrievalQuerySettingsGet**](RetrievalAPI.md#GetQuerySettingsApiV1RetrievalQuerySettingsGet) | **Get** /api/v1/retrieval/query/settings | Get Query Settings
[**GetRagConfigApiV1RetrievalConfigGet**](RetrievalAPI.md#GetRagConfigApiV1RetrievalConfigGet) | **Get** /api/v1/retrieval/config | Get Rag Config
[**GetRagTemplateApiV1RetrievalTemplateGet**](RetrievalAPI.md#GetRagTemplateApiV1RetrievalTemplateGet) | **Get** /api/v1/retrieval/template | Get Rag Template
[**GetReraankingConfigApiV1RetrievalRerankingGet**](RetrievalAPI.md#GetReraankingConfigApiV1RetrievalRerankingGet) | **Get** /api/v1/retrieval/reranking | Get Reraanking Config
[**GetStatusApiV1RetrievalGet**](RetrievalAPI.md#GetStatusApiV1RetrievalGet) | **Get** /api/v1/retrieval/ | Get Status
[**ProcessFileApiV1RetrievalProcessFilePost**](RetrievalAPI.md#ProcessFileApiV1RetrievalProcessFilePost) | **Post** /api/v1/retrieval/process/file | Process File
[**ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost**](RetrievalAPI.md#ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost) | **Post** /api/v1/retrieval/process/files/batch | Process Files Batch
[**ProcessTextApiV1RetrievalProcessTextPost**](RetrievalAPI.md#ProcessTextApiV1RetrievalProcessTextPost) | **Post** /api/v1/retrieval/process/text | Process Text
[**ProcessWebApiV1RetrievalProcessWebPost**](RetrievalAPI.md#ProcessWebApiV1RetrievalProcessWebPost) | **Post** /api/v1/retrieval/process/web | Process Web
[**ProcessWebSearchApiV1RetrievalProcessWebSearchPost**](RetrievalAPI.md#ProcessWebSearchApiV1RetrievalProcessWebSearchPost) | **Post** /api/v1/retrieval/process/web/search | Process Web Search
[**ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost**](RetrievalAPI.md#ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost) | **Post** /api/v1/retrieval/process/youtube | Process Youtube Video
[**QueryCollectionHandlerApiV1RetrievalQueryCollectionPost**](RetrievalAPI.md#QueryCollectionHandlerApiV1RetrievalQueryCollectionPost) | **Post** /api/v1/retrieval/query/collection | Query Collection Handler
[**QueryDocHandlerApiV1RetrievalQueryDocPost**](RetrievalAPI.md#QueryDocHandlerApiV1RetrievalQueryDocPost) | **Post** /api/v1/retrieval/query/doc | Query Doc Handler
[**ResetUploadDirApiV1RetrievalResetUploadsPost**](RetrievalAPI.md#ResetUploadDirApiV1RetrievalResetUploadsPost) | **Post** /api/v1/retrieval/reset/uploads | Reset Upload Dir
[**ResetVectorDbApiV1RetrievalResetDbPost**](RetrievalAPI.md#ResetVectorDbApiV1RetrievalResetDbPost) | **Post** /api/v1/retrieval/reset/db | Reset Vector Db
[**UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost**](RetrievalAPI.md#UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost) | **Post** /api/v1/retrieval/embedding/update | Update Embedding Config
[**UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost**](RetrievalAPI.md#UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost) | **Post** /api/v1/retrieval/query/settings/update | Update Query Settings
[**UpdateRagConfigApiV1RetrievalConfigUpdatePost**](RetrievalAPI.md#UpdateRagConfigApiV1RetrievalConfigUpdatePost) | **Post** /api/v1/retrieval/config/update | Update Rag Config
[**UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost**](RetrievalAPI.md#UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost) | **Post** /api/v1/retrieval/reranking/update | Update Reranking Config



## DeleteEntriesFromCollectionApiV1RetrievalDeletePost

> interface{} DeleteEntriesFromCollectionApiV1RetrievalDeletePost(ctx).DeleteForm(deleteForm).Execute()

Delete Entries From Collection

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
	deleteForm := *openapiclient.NewDeleteForm("CollectionName_example", "FileId_example") // DeleteForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.DeleteEntriesFromCollectionApiV1RetrievalDeletePost(context.Background()).DeleteForm(deleteForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.DeleteEntriesFromCollectionApiV1RetrievalDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEntriesFromCollectionApiV1RetrievalDeletePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.DeleteEntriesFromCollectionApiV1RetrievalDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntriesFromCollectionApiV1RetrievalDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteForm** | [**DeleteForm**](DeleteForm.md) |  | 

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


## GetEmbeddingConfigApiV1RetrievalEmbeddingGet

> interface{} GetEmbeddingConfigApiV1RetrievalEmbeddingGet(ctx).Execute()

Get Embedding Config

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
	resp, r, err := apiClient.RetrievalAPI.GetEmbeddingConfigApiV1RetrievalEmbeddingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetEmbeddingConfigApiV1RetrievalEmbeddingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingConfigApiV1RetrievalEmbeddingGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetEmbeddingConfigApiV1RetrievalEmbeddingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingConfigApiV1RetrievalEmbeddingGetRequest struct via the builder pattern


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


## GetEmbeddingsApiV1RetrievalEfTextGet

> interface{} GetEmbeddingsApiV1RetrievalEfTextGet(ctx, text).Execute()

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
	text := "text_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.GetEmbeddingsApiV1RetrievalEfTextGet(context.Background(), text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetEmbeddingsApiV1RetrievalEfTextGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingsApiV1RetrievalEfTextGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetEmbeddingsApiV1RetrievalEfTextGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**text** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingsApiV1RetrievalEfTextGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetQuerySettingsApiV1RetrievalQuerySettingsGet

> interface{} GetQuerySettingsApiV1RetrievalQuerySettingsGet(ctx).Execute()

Get Query Settings

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
	resp, r, err := apiClient.RetrievalAPI.GetQuerySettingsApiV1RetrievalQuerySettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetQuerySettingsApiV1RetrievalQuerySettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuerySettingsApiV1RetrievalQuerySettingsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetQuerySettingsApiV1RetrievalQuerySettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuerySettingsApiV1RetrievalQuerySettingsGetRequest struct via the builder pattern


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


## GetRagConfigApiV1RetrievalConfigGet

> interface{} GetRagConfigApiV1RetrievalConfigGet(ctx).Execute()

Get Rag Config

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
	resp, r, err := apiClient.RetrievalAPI.GetRagConfigApiV1RetrievalConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetRagConfigApiV1RetrievalConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRagConfigApiV1RetrievalConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetRagConfigApiV1RetrievalConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRagConfigApiV1RetrievalConfigGetRequest struct via the builder pattern


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


## GetRagTemplateApiV1RetrievalTemplateGet

> interface{} GetRagTemplateApiV1RetrievalTemplateGet(ctx).Execute()

Get Rag Template

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
	resp, r, err := apiClient.RetrievalAPI.GetRagTemplateApiV1RetrievalTemplateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetRagTemplateApiV1RetrievalTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRagTemplateApiV1RetrievalTemplateGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetRagTemplateApiV1RetrievalTemplateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRagTemplateApiV1RetrievalTemplateGetRequest struct via the builder pattern


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


## GetReraankingConfigApiV1RetrievalRerankingGet

> interface{} GetReraankingConfigApiV1RetrievalRerankingGet(ctx).Execute()

Get Reraanking Config

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
	resp, r, err := apiClient.RetrievalAPI.GetReraankingConfigApiV1RetrievalRerankingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetReraankingConfigApiV1RetrievalRerankingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReraankingConfigApiV1RetrievalRerankingGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetReraankingConfigApiV1RetrievalRerankingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReraankingConfigApiV1RetrievalRerankingGetRequest struct via the builder pattern


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


## GetStatusApiV1RetrievalGet

> interface{} GetStatusApiV1RetrievalGet(ctx).Execute()

Get Status

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
	resp, r, err := apiClient.RetrievalAPI.GetStatusApiV1RetrievalGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.GetStatusApiV1RetrievalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusApiV1RetrievalGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.GetStatusApiV1RetrievalGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusApiV1RetrievalGetRequest struct via the builder pattern


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


## ProcessFileApiV1RetrievalProcessFilePost

> interface{} ProcessFileApiV1RetrievalProcessFilePost(ctx).ProcessFileForm(processFileForm).Execute()

Process File

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
	processFileForm := *openapiclient.NewProcessFileForm("FileId_example") // ProcessFileForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessFileApiV1RetrievalProcessFilePost(context.Background()).ProcessFileForm(processFileForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessFileApiV1RetrievalProcessFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessFileApiV1RetrievalProcessFilePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessFileApiV1RetrievalProcessFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessFileApiV1RetrievalProcessFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processFileForm** | [**ProcessFileForm**](ProcessFileForm.md) |  | 

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


## ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost

> BatchProcessFilesResponse ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost(ctx).BatchProcessFilesForm(batchProcessFilesForm).Execute()

Process Files Batch



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
	batchProcessFilesForm := *openapiclient.NewBatchProcessFilesForm([]openapiclient.FileModel{*openapiclient.NewFileModel("Id_example", "UserId_example", "Filename_example", NullableInt32(123), NullableInt32(123))}, "CollectionName_example") // BatchProcessFilesForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost(context.Background()).BatchProcessFilesForm(batchProcessFilesForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost`: BatchProcessFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessFilesBatchApiV1RetrievalProcessFilesBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessFilesBatchApiV1RetrievalProcessFilesBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchProcessFilesForm** | [**BatchProcessFilesForm**](BatchProcessFilesForm.md) |  | 

### Return type

[**BatchProcessFilesResponse**](BatchProcessFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessTextApiV1RetrievalProcessTextPost

> interface{} ProcessTextApiV1RetrievalProcessTextPost(ctx).ProcessTextForm(processTextForm).Execute()

Process Text

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
	processTextForm := *openapiclient.NewProcessTextForm("Name_example", "Content_example") // ProcessTextForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessTextApiV1RetrievalProcessTextPost(context.Background()).ProcessTextForm(processTextForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessTextApiV1RetrievalProcessTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessTextApiV1RetrievalProcessTextPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessTextApiV1RetrievalProcessTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessTextApiV1RetrievalProcessTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processTextForm** | [**ProcessTextForm**](ProcessTextForm.md) |  | 

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


## ProcessWebApiV1RetrievalProcessWebPost

> interface{} ProcessWebApiV1RetrievalProcessWebPost(ctx).ProcessUrlForm(processUrlForm).Execute()

Process Web

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
	processUrlForm := *openapiclient.NewProcessUrlForm("Url_example") // ProcessUrlForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessWebApiV1RetrievalProcessWebPost(context.Background()).ProcessUrlForm(processUrlForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessWebApiV1RetrievalProcessWebPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessWebApiV1RetrievalProcessWebPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessWebApiV1RetrievalProcessWebPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessWebApiV1RetrievalProcessWebPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processUrlForm** | [**ProcessUrlForm**](ProcessUrlForm.md) |  | 

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


## ProcessWebSearchApiV1RetrievalProcessWebSearchPost

> interface{} ProcessWebSearchApiV1RetrievalProcessWebSearchPost(ctx).SearchForm(searchForm).Execute()

Process Web Search

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
	searchForm := *openapiclient.NewSearchForm("Query_example") // SearchForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessWebSearchApiV1RetrievalProcessWebSearchPost(context.Background()).SearchForm(searchForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessWebSearchApiV1RetrievalProcessWebSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessWebSearchApiV1RetrievalProcessWebSearchPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessWebSearchApiV1RetrievalProcessWebSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessWebSearchApiV1RetrievalProcessWebSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchForm** | [**SearchForm**](SearchForm.md) |  | 

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


## ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost

> interface{} ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost(ctx).ProcessUrlForm(processUrlForm).Execute()

Process Youtube Video

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
	processUrlForm := *openapiclient.NewProcessUrlForm("Url_example") // ProcessUrlForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost(context.Background()).ProcessUrlForm(processUrlForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ProcessYoutubeVideoApiV1RetrievalProcessYoutubePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessYoutubeVideoApiV1RetrievalProcessYoutubePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processUrlForm** | [**ProcessUrlForm**](ProcessUrlForm.md) |  | 

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


## QueryCollectionHandlerApiV1RetrievalQueryCollectionPost

> interface{} QueryCollectionHandlerApiV1RetrievalQueryCollectionPost(ctx).QueryCollectionsForm(queryCollectionsForm).Execute()

Query Collection Handler

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
	queryCollectionsForm := *openapiclient.NewQueryCollectionsForm([]string{"CollectionNames_example"}, "Query_example") // QueryCollectionsForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.QueryCollectionHandlerApiV1RetrievalQueryCollectionPost(context.Background()).QueryCollectionsForm(queryCollectionsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.QueryCollectionHandlerApiV1RetrievalQueryCollectionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCollectionHandlerApiV1RetrievalQueryCollectionPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.QueryCollectionHandlerApiV1RetrievalQueryCollectionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCollectionHandlerApiV1RetrievalQueryCollectionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCollectionsForm** | [**QueryCollectionsForm**](QueryCollectionsForm.md) |  | 

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


## QueryDocHandlerApiV1RetrievalQueryDocPost

> interface{} QueryDocHandlerApiV1RetrievalQueryDocPost(ctx).QueryDocForm(queryDocForm).Execute()

Query Doc Handler

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
	queryDocForm := *openapiclient.NewQueryDocForm("CollectionName_example", "Query_example") // QueryDocForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.QueryDocHandlerApiV1RetrievalQueryDocPost(context.Background()).QueryDocForm(queryDocForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.QueryDocHandlerApiV1RetrievalQueryDocPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryDocHandlerApiV1RetrievalQueryDocPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.QueryDocHandlerApiV1RetrievalQueryDocPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDocHandlerApiV1RetrievalQueryDocPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryDocForm** | [**QueryDocForm**](QueryDocForm.md) |  | 

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


## ResetUploadDirApiV1RetrievalResetUploadsPost

> bool ResetUploadDirApiV1RetrievalResetUploadsPost(ctx).Execute()

Reset Upload Dir

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
	resp, r, err := apiClient.RetrievalAPI.ResetUploadDirApiV1RetrievalResetUploadsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ResetUploadDirApiV1RetrievalResetUploadsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUploadDirApiV1RetrievalResetUploadsPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ResetUploadDirApiV1RetrievalResetUploadsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetUploadDirApiV1RetrievalResetUploadsPostRequest struct via the builder pattern


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


## ResetVectorDbApiV1RetrievalResetDbPost

> interface{} ResetVectorDbApiV1RetrievalResetDbPost(ctx).Execute()

Reset Vector Db

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
	resp, r, err := apiClient.RetrievalAPI.ResetVectorDbApiV1RetrievalResetDbPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.ResetVectorDbApiV1RetrievalResetDbPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetVectorDbApiV1RetrievalResetDbPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.ResetVectorDbApiV1RetrievalResetDbPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetVectorDbApiV1RetrievalResetDbPostRequest struct via the builder pattern


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


## UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost

> interface{} UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost(ctx).EmbeddingModelUpdateForm(embeddingModelUpdateForm).Execute()

Update Embedding Config

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
	embeddingModelUpdateForm := *openapiclient.NewEmbeddingModelUpdateForm("EmbeddingEngine_example", "EmbeddingModel_example") // EmbeddingModelUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost(context.Background()).EmbeddingModelUpdateForm(embeddingModelUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.UpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmbeddingConfigApiV1RetrievalEmbeddingUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingModelUpdateForm** | [**EmbeddingModelUpdateForm**](EmbeddingModelUpdateForm.md) |  | 

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


## UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost

> interface{} UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost(ctx).QuerySettingsForm(querySettingsForm).Execute()

Update Query Settings

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
	querySettingsForm := *openapiclient.NewQuerySettingsForm() // QuerySettingsForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost(context.Background()).QuerySettingsForm(querySettingsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.UpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuerySettingsApiV1RetrievalQuerySettingsUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querySettingsForm** | [**QuerySettingsForm**](QuerySettingsForm.md) |  | 

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


## UpdateRagConfigApiV1RetrievalConfigUpdatePost

> interface{} UpdateRagConfigApiV1RetrievalConfigUpdatePost(ctx).ConfigUpdateForm(configUpdateForm).Execute()

Update Rag Config

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
	configUpdateForm := *openapiclient.NewConfigUpdateForm() // ConfigUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.UpdateRagConfigApiV1RetrievalConfigUpdatePost(context.Background()).ConfigUpdateForm(configUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.UpdateRagConfigApiV1RetrievalConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRagConfigApiV1RetrievalConfigUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.UpdateRagConfigApiV1RetrievalConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRagConfigApiV1RetrievalConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configUpdateForm** | [**ConfigUpdateForm**](ConfigUpdateForm.md) |  | 

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


## UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost

> interface{} UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost(ctx).RerankingModelUpdateForm(rerankingModelUpdateForm).Execute()

Update Reranking Config

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
	rerankingModelUpdateForm := *openapiclient.NewRerankingModelUpdateForm("RerankingModel_example") // RerankingModelUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrievalAPI.UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost(context.Background()).RerankingModelUpdateForm(rerankingModelUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrievalAPI.UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RetrievalAPI.UpdateRerankingConfigApiV1RetrievalRerankingUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRerankingConfigApiV1RetrievalRerankingUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rerankingModelUpdateForm** | [**RerankingModelUpdateForm**](RerankingModelUpdateForm.md) |  | 

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

