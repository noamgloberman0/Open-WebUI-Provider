# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigApiV1ImagesConfigGet**](ImagesAPI.md#GetConfigApiV1ImagesConfigGet) | **Get** /api/v1/images/config | Get Config
[**GetImageConfigApiV1ImagesImageConfigGet**](ImagesAPI.md#GetImageConfigApiV1ImagesImageConfigGet) | **Get** /api/v1/images/image/config | Get Image Config
[**GetModelsApiV1ImagesModelsGet**](ImagesAPI.md#GetModelsApiV1ImagesModelsGet) | **Get** /api/v1/images/models | Get Models
[**ImageGenerationsApiV1ImagesGenerationsPost**](ImagesAPI.md#ImageGenerationsApiV1ImagesGenerationsPost) | **Post** /api/v1/images/generations | Image Generations
[**UpdateConfigApiV1ImagesConfigUpdatePost**](ImagesAPI.md#UpdateConfigApiV1ImagesConfigUpdatePost) | **Post** /api/v1/images/config/update | Update Config
[**UpdateImageConfigApiV1ImagesImageConfigUpdatePost**](ImagesAPI.md#UpdateImageConfigApiV1ImagesImageConfigUpdatePost) | **Post** /api/v1/images/image/config/update | Update Image Config
[**VerifyUrlApiV1ImagesConfigUrlVerifyGet**](ImagesAPI.md#VerifyUrlApiV1ImagesConfigUrlVerifyGet) | **Get** /api/v1/images/config/url/verify | Verify Url



## GetConfigApiV1ImagesConfigGet

> interface{} GetConfigApiV1ImagesConfigGet(ctx).Execute()

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
	resp, r, err := apiClient.ImagesAPI.GetConfigApiV1ImagesConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetConfigApiV1ImagesConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigApiV1ImagesConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetConfigApiV1ImagesConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigApiV1ImagesConfigGetRequest struct via the builder pattern


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


## GetImageConfigApiV1ImagesImageConfigGet

> interface{} GetImageConfigApiV1ImagesImageConfigGet(ctx).Execute()

Get Image Config

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
	resp, r, err := apiClient.ImagesAPI.GetImageConfigApiV1ImagesImageConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageConfigApiV1ImagesImageConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageConfigApiV1ImagesImageConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageConfigApiV1ImagesImageConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageConfigApiV1ImagesImageConfigGetRequest struct via the builder pattern


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


## GetModelsApiV1ImagesModelsGet

> interface{} GetModelsApiV1ImagesModelsGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetModelsApiV1ImagesModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetModelsApiV1ImagesModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsApiV1ImagesModelsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetModelsApiV1ImagesModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsApiV1ImagesModelsGetRequest struct via the builder pattern


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


## ImageGenerationsApiV1ImagesGenerationsPost

> interface{} ImageGenerationsApiV1ImagesGenerationsPost(ctx).GenerateImageForm(generateImageForm).Execute()

Image Generations

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
	generateImageForm := *openapiclient.NewGenerateImageForm("Prompt_example") // GenerateImageForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImageGenerationsApiV1ImagesGenerationsPost(context.Background()).GenerateImageForm(generateImageForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImageGenerationsApiV1ImagesGenerationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageGenerationsApiV1ImagesGenerationsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImageGenerationsApiV1ImagesGenerationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageGenerationsApiV1ImagesGenerationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateImageForm** | [**GenerateImageForm**](GenerateImageForm.md) |  | 

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


## UpdateConfigApiV1ImagesConfigUpdatePost

> interface{} UpdateConfigApiV1ImagesConfigUpdatePost(ctx).ConfigForm(configForm).Execute()

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
	configForm := *openapiclient.NewConfigForm(false, "Engine_example", false, *openapiclient.NewOpenWebuiRoutersImagesOpenAIConfigForm("OPENAI_API_BASE_URL_example", "OPENAI_API_KEY_example"), *openapiclient.NewAutomatic1111ConfigForm("AUTOMATIC1111BASEURL_example", "AUTOMATIC1111APIAUTH_example", "TODO", "AUTOMATIC1111SAMPLER_example", "AUTOMATIC1111SCHEDULER_example"), *openapiclient.NewComfyUIConfigForm("COMFYUI_BASE_URL_example", "COMFYUI_API_KEY_example", "COMFYUI_WORKFLOW_example", []map[string]interface{}{map[string]interface{}(123)}), *openapiclient.NewGeminiConfigForm("GEMINI_API_BASE_URL_example", "GEMINI_API_KEY_example")) // ConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.UpdateConfigApiV1ImagesConfigUpdatePost(context.Background()).ConfigForm(configForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UpdateConfigApiV1ImagesConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigApiV1ImagesConfigUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UpdateConfigApiV1ImagesConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigApiV1ImagesConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configForm** | [**ConfigForm**](ConfigForm.md) |  | 

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


## UpdateImageConfigApiV1ImagesImageConfigUpdatePost

> interface{} UpdateImageConfigApiV1ImagesImageConfigUpdatePost(ctx).ImageConfigForm(imageConfigForm).Execute()

Update Image Config

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
	imageConfigForm := *openapiclient.NewImageConfigForm("MODEL_example", "IMAGE_SIZE_example", int32(123)) // ImageConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.UpdateImageConfigApiV1ImagesImageConfigUpdatePost(context.Background()).ImageConfigForm(imageConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UpdateImageConfigApiV1ImagesImageConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImageConfigApiV1ImagesImageConfigUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UpdateImageConfigApiV1ImagesImageConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageConfigApiV1ImagesImageConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageConfigForm** | [**ImageConfigForm**](ImageConfigForm.md) |  | 

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


## VerifyUrlApiV1ImagesConfigUrlVerifyGet

> interface{} VerifyUrlApiV1ImagesConfigUrlVerifyGet(ctx).Execute()

Verify Url

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
	resp, r, err := apiClient.ImagesAPI.VerifyUrlApiV1ImagesConfigUrlVerifyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.VerifyUrlApiV1ImagesConfigUrlVerifyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyUrlApiV1ImagesConfigUrlVerifyGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.VerifyUrlApiV1ImagesConfigUrlVerifyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUrlApiV1ImagesConfigUrlVerifyGetRequest struct via the builder pattern


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

