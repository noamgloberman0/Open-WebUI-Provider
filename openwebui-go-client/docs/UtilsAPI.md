# \UtilsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadChatAsPdfApiV1UtilsPdfPost**](UtilsAPI.md#DownloadChatAsPdfApiV1UtilsPdfPost) | **Post** /api/v1/utils/pdf | Download Chat As Pdf
[**DownloadDbApiV1UtilsDbDownloadGet**](UtilsAPI.md#DownloadDbApiV1UtilsDbDownloadGet) | **Get** /api/v1/utils/db/download | Download Db
[**DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet**](UtilsAPI.md#DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet) | **Get** /api/v1/utils/litellm/config | Download Litellm Config Yaml
[**ExecuteCodeApiV1UtilsCodeExecutePost**](UtilsAPI.md#ExecuteCodeApiV1UtilsCodeExecutePost) | **Post** /api/v1/utils/code/execute | Execute Code
[**FormatCodeApiV1UtilsCodeFormatPost**](UtilsAPI.md#FormatCodeApiV1UtilsCodeFormatPost) | **Post** /api/v1/utils/code/format | Format Code
[**GetGravatarApiV1UtilsGravatarGet**](UtilsAPI.md#GetGravatarApiV1UtilsGravatarGet) | **Get** /api/v1/utils/gravatar | Get Gravatar
[**GetHtmlFromMarkdownApiV1UtilsMarkdownPost**](UtilsAPI.md#GetHtmlFromMarkdownApiV1UtilsMarkdownPost) | **Post** /api/v1/utils/markdown | Get Html From Markdown



## DownloadChatAsPdfApiV1UtilsPdfPost

> interface{} DownloadChatAsPdfApiV1UtilsPdfPost(ctx).ChatTitleMessagesForm(chatTitleMessagesForm).Execute()

Download Chat As Pdf

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
	chatTitleMessagesForm := *openapiclient.NewChatTitleMessagesForm("Title_example", []map[string]interface{}{map[string]interface{}(123)}) // ChatTitleMessagesForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.DownloadChatAsPdfApiV1UtilsPdfPost(context.Background()).ChatTitleMessagesForm(chatTitleMessagesForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadChatAsPdfApiV1UtilsPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadChatAsPdfApiV1UtilsPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadChatAsPdfApiV1UtilsPdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadChatAsPdfApiV1UtilsPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatTitleMessagesForm** | [**ChatTitleMessagesForm**](ChatTitleMessagesForm.md) |  | 

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


## DownloadDbApiV1UtilsDbDownloadGet

> interface{} DownloadDbApiV1UtilsDbDownloadGet(ctx).Execute()

Download Db

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
	resp, r, err := apiClient.UtilsAPI.DownloadDbApiV1UtilsDbDownloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadDbApiV1UtilsDbDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDbApiV1UtilsDbDownloadGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadDbApiV1UtilsDbDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDbApiV1UtilsDbDownloadGetRequest struct via the builder pattern


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


## DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet

> interface{} DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet(ctx).Execute()

Download Litellm Config Yaml

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
	resp, r, err := apiClient.UtilsAPI.DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadLitellmConfigYamlApiV1UtilsLitellmConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLitellmConfigYamlApiV1UtilsLitellmConfigGetRequest struct via the builder pattern


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


## ExecuteCodeApiV1UtilsCodeExecutePost

> interface{} ExecuteCodeApiV1UtilsCodeExecutePost(ctx).CodeForm(codeForm).Execute()

Execute Code

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
	codeForm := *openapiclient.NewCodeForm("Code_example") // CodeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.ExecuteCodeApiV1UtilsCodeExecutePost(context.Background()).CodeForm(codeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.ExecuteCodeApiV1UtilsCodeExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteCodeApiV1UtilsCodeExecutePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.ExecuteCodeApiV1UtilsCodeExecutePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteCodeApiV1UtilsCodeExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeForm** | [**CodeForm**](CodeForm.md) |  | 

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


## FormatCodeApiV1UtilsCodeFormatPost

> interface{} FormatCodeApiV1UtilsCodeFormatPost(ctx).CodeForm(codeForm).Execute()

Format Code

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
	codeForm := *openapiclient.NewCodeForm("Code_example") // CodeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.FormatCodeApiV1UtilsCodeFormatPost(context.Background()).CodeForm(codeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.FormatCodeApiV1UtilsCodeFormatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormatCodeApiV1UtilsCodeFormatPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.FormatCodeApiV1UtilsCodeFormatPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFormatCodeApiV1UtilsCodeFormatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeForm** | [**CodeForm**](CodeForm.md) |  | 

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


## GetGravatarApiV1UtilsGravatarGet

> interface{} GetGravatarApiV1UtilsGravatarGet(ctx).Email(email).Execute()

Get Gravatar

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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.GetGravatarApiV1UtilsGravatarGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.GetGravatarApiV1UtilsGravatarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGravatarApiV1UtilsGravatarGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.GetGravatarApiV1UtilsGravatarGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGravatarApiV1UtilsGravatarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

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


## GetHtmlFromMarkdownApiV1UtilsMarkdownPost

> interface{} GetHtmlFromMarkdownApiV1UtilsMarkdownPost(ctx).MarkdownForm(markdownForm).Execute()

Get Html From Markdown

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
	markdownForm := *openapiclient.NewMarkdownForm("Md_example") // MarkdownForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.GetHtmlFromMarkdownApiV1UtilsMarkdownPost(context.Background()).MarkdownForm(markdownForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.GetHtmlFromMarkdownApiV1UtilsMarkdownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHtmlFromMarkdownApiV1UtilsMarkdownPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.GetHtmlFromMarkdownApiV1UtilsMarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHtmlFromMarkdownApiV1UtilsMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markdownForm** | [**MarkdownForm**](MarkdownForm.md) |  | 

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

