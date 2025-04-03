# \OpenaiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateChatCompletionOpenaiChatCompletionsPost**](OpenaiAPI.md#GenerateChatCompletionOpenaiChatCompletionsPost) | **Post** /openai/chat/completions | Generate Chat Completion
[**GetConfigOpenaiConfigGet**](OpenaiAPI.md#GetConfigOpenaiConfigGet) | **Get** /openai/config | Get Config
[**GetModelsOpenaiModelsGet**](OpenaiAPI.md#GetModelsOpenaiModelsGet) | **Get** /openai/models | Get Models
[**GetModelsOpenaiModelsUrlIdxGet**](OpenaiAPI.md#GetModelsOpenaiModelsUrlIdxGet) | **Get** /openai/models/{url_idx} | Get Models
[**ProxyOpenaiPathPut**](OpenaiAPI.md#ProxyOpenaiPathPut) | **Get** /openai/{path} | Proxy
[**ProxyOpenaiPathPut_0**](OpenaiAPI.md#ProxyOpenaiPathPut_0) | **Put** /openai/{path} | Proxy
[**ProxyOpenaiPathPut_1**](OpenaiAPI.md#ProxyOpenaiPathPut_1) | **Post** /openai/{path} | Proxy
[**ProxyOpenaiPathPut_2**](OpenaiAPI.md#ProxyOpenaiPathPut_2) | **Delete** /openai/{path} | Proxy
[**SpeechOpenaiAudioSpeechPost**](OpenaiAPI.md#SpeechOpenaiAudioSpeechPost) | **Post** /openai/audio/speech | Speech
[**UpdateConfigOpenaiConfigUpdatePost**](OpenaiAPI.md#UpdateConfigOpenaiConfigUpdatePost) | **Post** /openai/config/update | Update Config
[**VerifyConnectionOpenaiVerifyPost**](OpenaiAPI.md#VerifyConnectionOpenaiVerifyPost) | **Post** /openai/verify | Verify Connection



## GenerateChatCompletionOpenaiChatCompletionsPost

> interface{} GenerateChatCompletionOpenaiChatCompletionsPost(ctx).Body(body).BypassFilter(bypassFilter).Execute()

Generate Chat Completion

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	bypassFilter := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.GenerateChatCompletionOpenaiChatCompletionsPost(context.Background()).Body(body).BypassFilter(bypassFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.GenerateChatCompletionOpenaiChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateChatCompletionOpenaiChatCompletionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.GenerateChatCompletionOpenaiChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateChatCompletionOpenaiChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **bypassFilter** | **bool** |  | 

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


## GetConfigOpenaiConfigGet

> interface{} GetConfigOpenaiConfigGet(ctx).Execute()

Get Config

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
	resp, r, err := apiClient.OpenaiAPI.GetConfigOpenaiConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.GetConfigOpenaiConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigOpenaiConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.GetConfigOpenaiConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigOpenaiConfigGetRequest struct via the builder pattern


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


## GetModelsOpenaiModelsGet

> interface{} GetModelsOpenaiModelsGet(ctx).UrlIdx(urlIdx).Execute()

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
	urlIdx := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.GetModelsOpenaiModelsGet(context.Background()).UrlIdx(urlIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.GetModelsOpenaiModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsOpenaiModelsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.GetModelsOpenaiModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsOpenaiModelsGetRequest struct via the builder pattern


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


## GetModelsOpenaiModelsUrlIdxGet

> interface{} GetModelsOpenaiModelsUrlIdxGet(ctx, urlIdx).Execute()

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
	urlIdx := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.GetModelsOpenaiModelsUrlIdxGet(context.Background(), urlIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.GetModelsOpenaiModelsUrlIdxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsOpenaiModelsUrlIdxGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.GetModelsOpenaiModelsUrlIdxGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**urlIdx** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsOpenaiModelsUrlIdxGetRequest struct via the builder pattern


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


## ProxyOpenaiPathPut

> interface{} ProxyOpenaiPathPut(ctx, path).Execute()

Proxy



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.ProxyOpenaiPathPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenaiPathPut`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.ProxyOpenaiPathPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenaiPathPutRequest struct via the builder pattern


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


## ProxyOpenaiPathPut_0

> interface{} ProxyOpenaiPathPut_0(ctx, path).Execute()

Proxy



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_0(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.ProxyOpenaiPathPut_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenaiPathPut_0`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.ProxyOpenaiPathPut_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenaiPathPut_1Request struct via the builder pattern


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


## ProxyOpenaiPathPut_1

> interface{} ProxyOpenaiPathPut_1(ctx, path).Execute()

Proxy



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_1(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.ProxyOpenaiPathPut_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenaiPathPut_1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.ProxyOpenaiPathPut_1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenaiPathPut_2Request struct via the builder pattern


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


## ProxyOpenaiPathPut_2

> interface{} ProxyOpenaiPathPut_2(ctx, path).Execute()

Proxy



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_2(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.ProxyOpenaiPathPut_2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenaiPathPut_2`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.ProxyOpenaiPathPut_2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenaiPathPut_3Request struct via the builder pattern


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


## SpeechOpenaiAudioSpeechPost

> interface{} SpeechOpenaiAudioSpeechPost(ctx).Execute()

Speech

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
	resp, r, err := apiClient.OpenaiAPI.SpeechOpenaiAudioSpeechPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.SpeechOpenaiAudioSpeechPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeechOpenaiAudioSpeechPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.SpeechOpenaiAudioSpeechPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpeechOpenaiAudioSpeechPostRequest struct via the builder pattern


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


## UpdateConfigOpenaiConfigUpdatePost

> interface{} UpdateConfigOpenaiConfigUpdatePost(ctx).OpenWebuiRoutersOpenaiOpenAIConfigForm(openWebuiRoutersOpenaiOpenAIConfigForm).Execute()

Update Config

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
	openWebuiRoutersOpenaiOpenAIConfigForm := *openapiclient.NewOpenWebuiRoutersOpenaiOpenAIConfigForm([]string{"OPENAI_API_BASE_URLS_example"}, []string{"OPENAI_API_KEYS_example"}, map[string]interface{}(123)) // OpenWebuiRoutersOpenaiOpenAIConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.UpdateConfigOpenaiConfigUpdatePost(context.Background()).OpenWebuiRoutersOpenaiOpenAIConfigForm(openWebuiRoutersOpenaiOpenAIConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.UpdateConfigOpenaiConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigOpenaiConfigUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.UpdateConfigOpenaiConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigOpenaiConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openWebuiRoutersOpenaiOpenAIConfigForm** | [**OpenWebuiRoutersOpenaiOpenAIConfigForm**](OpenWebuiRoutersOpenaiOpenAIConfigForm.md) |  | 

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


## VerifyConnectionOpenaiVerifyPost

> interface{} VerifyConnectionOpenaiVerifyPost(ctx).OpenWebuiRoutersOpenaiConnectionVerificationForm(openWebuiRoutersOpenaiConnectionVerificationForm).Execute()

Verify Connection

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
	openWebuiRoutersOpenaiConnectionVerificationForm := *openapiclient.NewOpenWebuiRoutersOpenaiConnectionVerificationForm("Url_example", "Key_example") // OpenWebuiRoutersOpenaiConnectionVerificationForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenaiAPI.VerifyConnectionOpenaiVerifyPost(context.Background()).OpenWebuiRoutersOpenaiConnectionVerificationForm(openWebuiRoutersOpenaiConnectionVerificationForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenaiAPI.VerifyConnectionOpenaiVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyConnectionOpenaiVerifyPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenaiAPI.VerifyConnectionOpenaiVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnectionOpenaiVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openWebuiRoutersOpenaiConnectionVerificationForm** | [**OpenWebuiRoutersOpenaiConnectionVerificationForm**](OpenWebuiRoutersOpenaiConnectionVerificationForm.md) |  | 

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

