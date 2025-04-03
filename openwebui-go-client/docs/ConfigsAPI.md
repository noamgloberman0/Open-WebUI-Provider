# \ConfigsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportConfigApiV1ConfigsExportGet**](ConfigsAPI.md#ExportConfigApiV1ConfigsExportGet) | **Get** /api/v1/configs/export | Export Config
[**GetBannersApiV1ConfigsBannersGet**](ConfigsAPI.md#GetBannersApiV1ConfigsBannersGet) | **Get** /api/v1/configs/banners | Get Banners
[**GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet**](ConfigsAPI.md#GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet) | **Get** /api/v1/configs/code_execution | Get Code Execution Config
[**GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet**](ConfigsAPI.md#GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet) | **Get** /api/v1/configs/direct_connections | Get Direct Connections Config
[**GetModelsConfigApiV1ConfigsModelsGet**](ConfigsAPI.md#GetModelsConfigApiV1ConfigsModelsGet) | **Get** /api/v1/configs/models | Get Models Config
[**ImportConfigApiV1ConfigsImportPost**](ConfigsAPI.md#ImportConfigApiV1ConfigsImportPost) | **Post** /api/v1/configs/import | Import Config
[**SetBannersApiV1ConfigsBannersPost**](ConfigsAPI.md#SetBannersApiV1ConfigsBannersPost) | **Post** /api/v1/configs/banners | Set Banners
[**SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost**](ConfigsAPI.md#SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost) | **Post** /api/v1/configs/code_execution | Set Code Execution Config
[**SetDefaultSuggestionsApiV1ConfigsSuggestionsPost**](ConfigsAPI.md#SetDefaultSuggestionsApiV1ConfigsSuggestionsPost) | **Post** /api/v1/configs/suggestions | Set Default Suggestions
[**SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost**](ConfigsAPI.md#SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost) | **Post** /api/v1/configs/direct_connections | Set Direct Connections Config
[**SetModelsConfigApiV1ConfigsModelsPost**](ConfigsAPI.md#SetModelsConfigApiV1ConfigsModelsPost) | **Post** /api/v1/configs/models | Set Models Config



## ExportConfigApiV1ConfigsExportGet

> map[string]interface{} ExportConfigApiV1ConfigsExportGet(ctx).Execute()

Export Config

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
	resp, r, err := apiClient.ConfigsAPI.ExportConfigApiV1ConfigsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ExportConfigApiV1ConfigsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportConfigApiV1ConfigsExportGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ExportConfigApiV1ConfigsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigApiV1ConfigsExportGetRequest struct via the builder pattern


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


## GetBannersApiV1ConfigsBannersGet

> []BannerModel GetBannersApiV1ConfigsBannersGet(ctx).Execute()

Get Banners

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
	resp, r, err := apiClient.ConfigsAPI.GetBannersApiV1ConfigsBannersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetBannersApiV1ConfigsBannersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBannersApiV1ConfigsBannersGet`: []BannerModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetBannersApiV1ConfigsBannersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannersApiV1ConfigsBannersGetRequest struct via the builder pattern


### Return type

[**[]BannerModel**](BannerModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet

> CodeInterpreterConfigForm GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet(ctx).Execute()

Get Code Execution Config

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
	resp, r, err := apiClient.ConfigsAPI.GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet`: CodeInterpreterConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetCodeExecutionConfigApiV1ConfigsCodeExecutionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeExecutionConfigApiV1ConfigsCodeExecutionGetRequest struct via the builder pattern


### Return type

[**CodeInterpreterConfigForm**](CodeInterpreterConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet

> DirectConnectionsConfigForm GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet(ctx).Execute()

Get Direct Connections Config

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
	resp, r, err := apiClient.ConfigsAPI.GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet`: DirectConnectionsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectConnectionsConfigApiV1ConfigsDirectConnectionsGetRequest struct via the builder pattern


### Return type

[**DirectConnectionsConfigForm**](DirectConnectionsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsConfigApiV1ConfigsModelsGet

> ModelsConfigForm GetModelsConfigApiV1ConfigsModelsGet(ctx).Execute()

Get Models Config

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
	resp, r, err := apiClient.ConfigsAPI.GetModelsConfigApiV1ConfigsModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetModelsConfigApiV1ConfigsModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsConfigApiV1ConfigsModelsGet`: ModelsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetModelsConfigApiV1ConfigsModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsConfigApiV1ConfigsModelsGetRequest struct via the builder pattern


### Return type

[**ModelsConfigForm**](ModelsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportConfigApiV1ConfigsImportPost

> map[string]interface{} ImportConfigApiV1ConfigsImportPost(ctx).ImportConfigForm(importConfigForm).Execute()

Import Config

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
	importConfigForm := *openapiclient.NewImportConfigForm(map[string]interface{}(123)) // ImportConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.ImportConfigApiV1ConfigsImportPost(context.Background()).ImportConfigForm(importConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ImportConfigApiV1ConfigsImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportConfigApiV1ConfigsImportPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ImportConfigApiV1ConfigsImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportConfigApiV1ConfigsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importConfigForm** | [**ImportConfigForm**](ImportConfigForm.md) |  | 

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


## SetBannersApiV1ConfigsBannersPost

> []BannerModel SetBannersApiV1ConfigsBannersPost(ctx).SetBannersForm(setBannersForm).Execute()

Set Banners

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
	setBannersForm := *openapiclient.NewSetBannersForm([]openapiclient.BannerModel{*openapiclient.NewBannerModel("Id_example", "Type_example", "Content_example", false, int32(123))}) // SetBannersForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetBannersApiV1ConfigsBannersPost(context.Background()).SetBannersForm(setBannersForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetBannersApiV1ConfigsBannersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBannersApiV1ConfigsBannersPost`: []BannerModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetBannersApiV1ConfigsBannersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBannersApiV1ConfigsBannersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setBannersForm** | [**SetBannersForm**](SetBannersForm.md) |  | 

### Return type

[**[]BannerModel**](BannerModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost

> CodeInterpreterConfigForm SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost(ctx).CodeInterpreterConfigForm(codeInterpreterConfigForm).Execute()

Set Code Execution Config

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
	codeInterpreterConfigForm := *openapiclient.NewCodeInterpreterConfigForm(false, "CODE_EXECUTION_ENGINE_example", "CODE_EXECUTION_JUPYTER_URL_example", "CODE_EXECUTION_JUPYTER_AUTH_example", "CODE_EXECUTION_JUPYTER_AUTH_TOKEN_example", "CODE_EXECUTION_JUPYTER_AUTH_PASSWORD_example", NullableInt32(123), false, "CODE_INTERPRETER_ENGINE_example", "CODE_INTERPRETER_PROMPT_TEMPLATE_example", "CODE_INTERPRETER_JUPYTER_URL_example", "CODE_INTERPRETER_JUPYTER_AUTH_example", "CODE_INTERPRETER_JUPYTER_AUTH_TOKEN_example", "CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD_example", NullableInt32(123)) // CodeInterpreterConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost(context.Background()).CodeInterpreterConfigForm(codeInterpreterConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost`: CodeInterpreterConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetCodeExecutionConfigApiV1ConfigsCodeExecutionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetCodeExecutionConfigApiV1ConfigsCodeExecutionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeInterpreterConfigForm** | [**CodeInterpreterConfigForm**](CodeInterpreterConfigForm.md) |  | 

### Return type

[**CodeInterpreterConfigForm**](CodeInterpreterConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultSuggestionsApiV1ConfigsSuggestionsPost

> []PromptSuggestion SetDefaultSuggestionsApiV1ConfigsSuggestionsPost(ctx).SetDefaultSuggestionsForm(setDefaultSuggestionsForm).Execute()

Set Default Suggestions

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
	setDefaultSuggestionsForm := *openapiclient.NewSetDefaultSuggestionsForm([]openapiclient.PromptSuggestion{*openapiclient.NewPromptSuggestion([]string{"Title_example"}, "Content_example")}) // SetDefaultSuggestionsForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetDefaultSuggestionsApiV1ConfigsSuggestionsPost(context.Background()).SetDefaultSuggestionsForm(setDefaultSuggestionsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetDefaultSuggestionsApiV1ConfigsSuggestionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultSuggestionsApiV1ConfigsSuggestionsPost`: []PromptSuggestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetDefaultSuggestionsApiV1ConfigsSuggestionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultSuggestionsApiV1ConfigsSuggestionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultSuggestionsForm** | [**SetDefaultSuggestionsForm**](SetDefaultSuggestionsForm.md) |  | 

### Return type

[**[]PromptSuggestion**](PromptSuggestion.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost

> DirectConnectionsConfigForm SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost(ctx).DirectConnectionsConfigForm(directConnectionsConfigForm).Execute()

Set Direct Connections Config

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
	directConnectionsConfigForm := *openapiclient.NewDirectConnectionsConfigForm(false) // DirectConnectionsConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost(context.Background()).DirectConnectionsConfigForm(directConnectionsConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost`: DirectConnectionsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDirectConnectionsConfigApiV1ConfigsDirectConnectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directConnectionsConfigForm** | [**DirectConnectionsConfigForm**](DirectConnectionsConfigForm.md) |  | 

### Return type

[**DirectConnectionsConfigForm**](DirectConnectionsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetModelsConfigApiV1ConfigsModelsPost

> ModelsConfigForm SetModelsConfigApiV1ConfigsModelsPost(ctx).ModelsConfigForm(modelsConfigForm).Execute()

Set Models Config

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
	modelsConfigForm := *openapiclient.NewModelsConfigForm("DEFAULT_MODELS_example", []string{"MODEL_ORDER_LIST_example"}) // ModelsConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetModelsConfigApiV1ConfigsModelsPost(context.Background()).ModelsConfigForm(modelsConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetModelsConfigApiV1ConfigsModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetModelsConfigApiV1ConfigsModelsPost`: ModelsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetModelsConfigApiV1ConfigsModelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetModelsConfigApiV1ConfigsModelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelsConfigForm** | [**ModelsConfigForm**](ModelsConfigForm.md) |  | 

### Return type

[**ModelsConfigForm**](ModelsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

