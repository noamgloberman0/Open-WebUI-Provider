# \PromptsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewPromptApiV1PromptsCreatePost**](PromptsAPI.md#CreateNewPromptApiV1PromptsCreatePost) | **Post** /api/v1/prompts/create | Create New Prompt
[**DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete**](PromptsAPI.md#DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete) | **Delete** /api/v1/prompts/command/{command}/delete | Delete Prompt By Command
[**GetPromptByCommandApiV1PromptsCommandCommandGet**](PromptsAPI.md#GetPromptByCommandApiV1PromptsCommandCommandGet) | **Get** /api/v1/prompts/command/{command} | Get Prompt By Command
[**GetPromptListApiV1PromptsListGet**](PromptsAPI.md#GetPromptListApiV1PromptsListGet) | **Get** /api/v1/prompts/list | Get Prompt List
[**GetPromptsApiV1PromptsGet**](PromptsAPI.md#GetPromptsApiV1PromptsGet) | **Get** /api/v1/prompts/ | Get Prompts
[**UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost**](PromptsAPI.md#UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost) | **Post** /api/v1/prompts/command/{command}/update | Update Prompt By Command



## CreateNewPromptApiV1PromptsCreatePost

> PromptModel CreateNewPromptApiV1PromptsCreatePost(ctx).PromptForm(promptForm).Execute()

Create New Prompt

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
	promptForm := *openapiclient.NewPromptForm("Command_example", "Title_example", "Content_example") // PromptForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.CreateNewPromptApiV1PromptsCreatePost(context.Background()).PromptForm(promptForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CreateNewPromptApiV1PromptsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewPromptApiV1PromptsCreatePost`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CreateNewPromptApiV1PromptsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewPromptApiV1PromptsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptForm** | [**PromptForm**](PromptForm.md) |  | 

### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete

> bool DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete(ctx, command).Execute()

Delete Prompt By Command

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
	command := "command_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete(context.Background(), command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.DeletePromptByCommandApiV1PromptsCommandCommandDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromptByCommandApiV1PromptsCommandCommandDeleteDeleteRequest struct via the builder pattern


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


## GetPromptByCommandApiV1PromptsCommandCommandGet

> PromptModel GetPromptByCommandApiV1PromptsCommandCommandGet(ctx, command).Execute()

Get Prompt By Command

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
	command := "command_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.GetPromptByCommandApiV1PromptsCommandCommandGet(context.Background(), command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptByCommandApiV1PromptsCommandCommandGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptByCommandApiV1PromptsCommandCommandGet`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptByCommandApiV1PromptsCommandCommandGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptByCommandApiV1PromptsCommandCommandGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptListApiV1PromptsListGet

> []PromptUserResponse GetPromptListApiV1PromptsListGet(ctx).Execute()

Get Prompt List

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
	resp, r, err := apiClient.PromptsAPI.GetPromptListApiV1PromptsListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptListApiV1PromptsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptListApiV1PromptsListGet`: []PromptUserResponse
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptListApiV1PromptsListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptListApiV1PromptsListGetRequest struct via the builder pattern


### Return type

[**[]PromptUserResponse**](PromptUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsApiV1PromptsGet

> []PromptModel GetPromptsApiV1PromptsGet(ctx).Execute()

Get Prompts

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
	resp, r, err := apiClient.PromptsAPI.GetPromptsApiV1PromptsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptsApiV1PromptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptsApiV1PromptsGet`: []PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptsApiV1PromptsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsApiV1PromptsGetRequest struct via the builder pattern


### Return type

[**[]PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost

> PromptModel UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost(ctx, command).PromptForm(promptForm).Execute()

Update Prompt By Command

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
	command := "command_example" // string | 
	promptForm := *openapiclient.NewPromptForm("Command_example", "Title_example", "Content_example") // PromptForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost(context.Background(), command).PromptForm(promptForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.UpdatePromptByCommandApiV1PromptsCommandCommandUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePromptByCommandApiV1PromptsCommandCommandUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promptForm** | [**PromptForm**](PromptForm.md) |  | 

### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

