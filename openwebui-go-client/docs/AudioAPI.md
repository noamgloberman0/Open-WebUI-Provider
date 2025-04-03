# \AudioAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioConfigApiV1AudioConfigGet**](AudioAPI.md#GetAudioConfigApiV1AudioConfigGet) | **Get** /api/v1/audio/config | Get Audio Config
[**GetModelsApiV1AudioModelsGet**](AudioAPI.md#GetModelsApiV1AudioModelsGet) | **Get** /api/v1/audio/models | Get Models
[**GetVoicesApiV1AudioVoicesGet**](AudioAPI.md#GetVoicesApiV1AudioVoicesGet) | **Get** /api/v1/audio/voices | Get Voices
[**SpeechApiV1AudioSpeechPost**](AudioAPI.md#SpeechApiV1AudioSpeechPost) | **Post** /api/v1/audio/speech | Speech
[**TranscriptionApiV1AudioTranscriptionsPost**](AudioAPI.md#TranscriptionApiV1AudioTranscriptionsPost) | **Post** /api/v1/audio/transcriptions | Transcription
[**UpdateAudioConfigApiV1AudioConfigUpdatePost**](AudioAPI.md#UpdateAudioConfigApiV1AudioConfigUpdatePost) | **Post** /api/v1/audio/config/update | Update Audio Config



## GetAudioConfigApiV1AudioConfigGet

> interface{} GetAudioConfigApiV1AudioConfigGet(ctx).Execute()

Get Audio Config

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
	resp, r, err := apiClient.AudioAPI.GetAudioConfigApiV1AudioConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.GetAudioConfigApiV1AudioConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudioConfigApiV1AudioConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.GetAudioConfigApiV1AudioConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioConfigApiV1AudioConfigGetRequest struct via the builder pattern


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


## GetModelsApiV1AudioModelsGet

> interface{} GetModelsApiV1AudioModelsGet(ctx).Execute()

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
	resp, r, err := apiClient.AudioAPI.GetModelsApiV1AudioModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.GetModelsApiV1AudioModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsApiV1AudioModelsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.GetModelsApiV1AudioModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsApiV1AudioModelsGetRequest struct via the builder pattern


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


## GetVoicesApiV1AudioVoicesGet

> interface{} GetVoicesApiV1AudioVoicesGet(ctx).Execute()

Get Voices

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
	resp, r, err := apiClient.AudioAPI.GetVoicesApiV1AudioVoicesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.GetVoicesApiV1AudioVoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoicesApiV1AudioVoicesGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.GetVoicesApiV1AudioVoicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicesApiV1AudioVoicesGetRequest struct via the builder pattern


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


## SpeechApiV1AudioSpeechPost

> interface{} SpeechApiV1AudioSpeechPost(ctx).Execute()

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
	resp, r, err := apiClient.AudioAPI.SpeechApiV1AudioSpeechPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.SpeechApiV1AudioSpeechPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeechApiV1AudioSpeechPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.SpeechApiV1AudioSpeechPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpeechApiV1AudioSpeechPostRequest struct via the builder pattern


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


## TranscriptionApiV1AudioTranscriptionsPost

> interface{} TranscriptionApiV1AudioTranscriptionsPost(ctx).File(file).Execute()

Transcription

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.TranscriptionApiV1AudioTranscriptionsPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.TranscriptionApiV1AudioTranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranscriptionApiV1AudioTranscriptionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.TranscriptionApiV1AudioTranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTranscriptionApiV1AudioTranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudioConfigApiV1AudioConfigUpdatePost

> interface{} UpdateAudioConfigApiV1AudioConfigUpdatePost(ctx).AudioConfigUpdateForm(audioConfigUpdateForm).Execute()

Update Audio Config

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
	audioConfigUpdateForm := *openapiclient.NewAudioConfigUpdateForm(*openapiclient.NewTTSConfigForm("OPENAI_API_BASE_URL_example", "OPENAI_API_KEY_example", "API_KEY_example", "ENGINE_example", "MODEL_example", "VOICE_example", "SPLIT_ON_example", "AZURE_SPEECH_REGION_example", "AZURE_SPEECH_OUTPUT_FORMAT_example"), *openapiclient.NewSTTConfigForm("OPENAI_API_BASE_URL_example", "OPENAI_API_KEY_example", "ENGINE_example", "MODEL_example", "WHISPER_MODEL_example", "DEEPGRAM_API_KEY_example")) // AudioConfigUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.UpdateAudioConfigApiV1AudioConfigUpdatePost(context.Background()).AudioConfigUpdateForm(audioConfigUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.UpdateAudioConfigApiV1AudioConfigUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAudioConfigApiV1AudioConfigUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.UpdateAudioConfigApiV1AudioConfigUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudioConfigApiV1AudioConfigUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audioConfigUpdateForm** | [**AudioConfigUpdateForm**](AudioConfigUpdateForm.md) |  | 

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

