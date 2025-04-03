# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatActionApiChatActionsActionIdPost**](DefaultAPI.md#ChatActionApiChatActionsActionIdPost) | **Post** /api/chat/actions/{action_id} | Chat Action
[**ChatCompletedApiChatCompletedPost**](DefaultAPI.md#ChatCompletedApiChatCompletedPost) | **Post** /api/chat/completed | Chat Completed
[**ChatCompletionApiChatCompletionsPost**](DefaultAPI.md#ChatCompletionApiChatCompletionsPost) | **Post** /api/chat/completions | Chat Completion
[**GetAppChangelogApiChangelogGet**](DefaultAPI.md#GetAppChangelogApiChangelogGet) | **Get** /api/changelog | Get App Changelog
[**GetAppConfigApiConfigGet**](DefaultAPI.md#GetAppConfigApiConfigGet) | **Get** /api/config | Get App Config
[**GetAppLatestReleaseVersionApiVersionUpdatesGet**](DefaultAPI.md#GetAppLatestReleaseVersionApiVersionUpdatesGet) | **Get** /api/version/updates | Get App Latest Release Version
[**GetAppVersionApiVersionGet**](DefaultAPI.md#GetAppVersionApiVersionGet) | **Get** /api/version | Get App Version
[**GetBaseModelsApiModelsBaseGet**](DefaultAPI.md#GetBaseModelsApiModelsBaseGet) | **Get** /api/models/base | Get Base Models
[**GetManifestJsonManifestJsonGet**](DefaultAPI.md#GetManifestJsonManifestJsonGet) | **Get** /manifest.json | Get Manifest Json
[**GetModelsApiModelsGet**](DefaultAPI.md#GetModelsApiModelsGet) | **Get** /api/models | Get Models
[**GetOpensearchXmlOpensearchXmlGet**](DefaultAPI.md#GetOpensearchXmlOpensearchXmlGet) | **Get** /opensearch.xml | Get Opensearch Xml
[**GetWebhookUrlApiWebhookGet**](DefaultAPI.md#GetWebhookUrlApiWebhookGet) | **Get** /api/webhook | Get Webhook Url
[**HealthcheckHealthGet**](DefaultAPI.md#HealthcheckHealthGet) | **Get** /health | Healthcheck
[**HealthcheckWithDbHealthDbGet**](DefaultAPI.md#HealthcheckWithDbHealthDbGet) | **Get** /health/db | Healthcheck With Db
[**ListTasksEndpointApiTasksGet**](DefaultAPI.md#ListTasksEndpointApiTasksGet) | **Get** /api/tasks | List Tasks Endpoint
[**OauthCallbackOauthProviderCallbackGet**](DefaultAPI.md#OauthCallbackOauthProviderCallbackGet) | **Get** /oauth/{provider}/callback | Oauth Callback
[**OauthLoginOauthProviderLoginGet**](DefaultAPI.md#OauthLoginOauthProviderLoginGet) | **Get** /oauth/{provider}/login | Oauth Login
[**StopTaskEndpointApiTasksStopTaskIdPost**](DefaultAPI.md#StopTaskEndpointApiTasksStopTaskIdPost) | **Post** /api/tasks/stop/{task_id} | Stop Task Endpoint
[**UpdateWebhookUrlApiWebhookPost**](DefaultAPI.md#UpdateWebhookUrlApiWebhookPost) | **Post** /api/webhook | Update Webhook Url



## ChatActionApiChatActionsActionIdPost

> interface{} ChatActionApiChatActionsActionIdPost(ctx, actionId).Body(body).Execute()

Chat Action

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
	actionId := "actionId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ChatActionApiChatActionsActionIdPost(context.Background(), actionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ChatActionApiChatActionsActionIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatActionApiChatActionsActionIdPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ChatActionApiChatActionsActionIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatActionApiChatActionsActionIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## ChatCompletedApiChatCompletedPost

> interface{} ChatCompletedApiChatCompletedPost(ctx).Body(body).Execute()

Chat Completed

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ChatCompletedApiChatCompletedPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ChatCompletedApiChatCompletedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatCompletedApiChatCompletedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ChatCompletedApiChatCompletedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCompletedApiChatCompletedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## ChatCompletionApiChatCompletionsPost

> interface{} ChatCompletionApiChatCompletionsPost(ctx).Body(body).Execute()

Chat Completion

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ChatCompletionApiChatCompletionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ChatCompletionApiChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatCompletionApiChatCompletionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ChatCompletionApiChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCompletionApiChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## GetAppChangelogApiChangelogGet

> interface{} GetAppChangelogApiChangelogGet(ctx).Execute()

Get App Changelog

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
	resp, r, err := apiClient.DefaultAPI.GetAppChangelogApiChangelogGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAppChangelogApiChangelogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppChangelogApiChangelogGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAppChangelogApiChangelogGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppChangelogApiChangelogGetRequest struct via the builder pattern


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


## GetAppConfigApiConfigGet

> interface{} GetAppConfigApiConfigGet(ctx).Execute()

Get App Config

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
	resp, r, err := apiClient.DefaultAPI.GetAppConfigApiConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAppConfigApiConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppConfigApiConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAppConfigApiConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppConfigApiConfigGetRequest struct via the builder pattern


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


## GetAppLatestReleaseVersionApiVersionUpdatesGet

> interface{} GetAppLatestReleaseVersionApiVersionUpdatesGet(ctx).Execute()

Get App Latest Release Version

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
	resp, r, err := apiClient.DefaultAPI.GetAppLatestReleaseVersionApiVersionUpdatesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAppLatestReleaseVersionApiVersionUpdatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppLatestReleaseVersionApiVersionUpdatesGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAppLatestReleaseVersionApiVersionUpdatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppLatestReleaseVersionApiVersionUpdatesGetRequest struct via the builder pattern


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


## GetAppVersionApiVersionGet

> interface{} GetAppVersionApiVersionGet(ctx).Execute()

Get App Version

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
	resp, r, err := apiClient.DefaultAPI.GetAppVersionApiVersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAppVersionApiVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppVersionApiVersionGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAppVersionApiVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppVersionApiVersionGetRequest struct via the builder pattern


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


## GetBaseModelsApiModelsBaseGet

> interface{} GetBaseModelsApiModelsBaseGet(ctx).Execute()

Get Base Models

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
	resp, r, err := apiClient.DefaultAPI.GetBaseModelsApiModelsBaseGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBaseModelsApiModelsBaseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseModelsApiModelsBaseGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBaseModelsApiModelsBaseGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseModelsApiModelsBaseGetRequest struct via the builder pattern


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


## GetManifestJsonManifestJsonGet

> interface{} GetManifestJsonManifestJsonGet(ctx).Execute()

Get Manifest Json

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
	resp, r, err := apiClient.DefaultAPI.GetManifestJsonManifestJsonGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetManifestJsonManifestJsonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManifestJsonManifestJsonGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetManifestJsonManifestJsonGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManifestJsonManifestJsonGetRequest struct via the builder pattern


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


## GetModelsApiModelsGet

> interface{} GetModelsApiModelsGet(ctx).Execute()

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
	resp, r, err := apiClient.DefaultAPI.GetModelsApiModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetModelsApiModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsApiModelsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetModelsApiModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsApiModelsGetRequest struct via the builder pattern


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


## GetOpensearchXmlOpensearchXmlGet

> interface{} GetOpensearchXmlOpensearchXmlGet(ctx).Execute()

Get Opensearch Xml

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
	resp, r, err := apiClient.DefaultAPI.GetOpensearchXmlOpensearchXmlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOpensearchXmlOpensearchXmlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpensearchXmlOpensearchXmlGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOpensearchXmlOpensearchXmlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpensearchXmlOpensearchXmlGetRequest struct via the builder pattern


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


## GetWebhookUrlApiWebhookGet

> interface{} GetWebhookUrlApiWebhookGet(ctx).Execute()

Get Webhook Url

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
	resp, r, err := apiClient.DefaultAPI.GetWebhookUrlApiWebhookGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookUrlApiWebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookUrlApiWebhookGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookUrlApiWebhookGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookUrlApiWebhookGetRequest struct via the builder pattern


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


## HealthcheckHealthGet

> interface{} HealthcheckHealthGet(ctx).Execute()

Healthcheck

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
	resp, r, err := apiClient.DefaultAPI.HealthcheckHealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthcheckHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthcheckHealthGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthcheckHealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckHealthGetRequest struct via the builder pattern


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


## HealthcheckWithDbHealthDbGet

> interface{} HealthcheckWithDbHealthDbGet(ctx).Execute()

Healthcheck With Db

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
	resp, r, err := apiClient.DefaultAPI.HealthcheckWithDbHealthDbGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthcheckWithDbHealthDbGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthcheckWithDbHealthDbGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthcheckWithDbHealthDbGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckWithDbHealthDbGetRequest struct via the builder pattern


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


## ListTasksEndpointApiTasksGet

> interface{} ListTasksEndpointApiTasksGet(ctx).Execute()

List Tasks Endpoint

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
	resp, r, err := apiClient.DefaultAPI.ListTasksEndpointApiTasksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTasksEndpointApiTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasksEndpointApiTasksGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTasksEndpointApiTasksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksEndpointApiTasksGetRequest struct via the builder pattern


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


## OauthCallbackOauthProviderCallbackGet

> interface{} OauthCallbackOauthProviderCallbackGet(ctx, provider).Execute()

Oauth Callback

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OauthCallbackOauthProviderCallbackGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OauthCallbackOauthProviderCallbackGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthCallbackOauthProviderCallbackGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OauthCallbackOauthProviderCallbackGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthCallbackOauthProviderCallbackGetRequest struct via the builder pattern


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


## OauthLoginOauthProviderLoginGet

> interface{} OauthLoginOauthProviderLoginGet(ctx, provider).Execute()

Oauth Login

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OauthLoginOauthProviderLoginGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OauthLoginOauthProviderLoginGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthLoginOauthProviderLoginGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OauthLoginOauthProviderLoginGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthLoginOauthProviderLoginGetRequest struct via the builder pattern


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


## StopTaskEndpointApiTasksStopTaskIdPost

> interface{} StopTaskEndpointApiTasksStopTaskIdPost(ctx, taskId).Execute()

Stop Task Endpoint

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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StopTaskEndpointApiTasksStopTaskIdPost(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopTaskEndpointApiTasksStopTaskIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopTaskEndpointApiTasksStopTaskIdPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StopTaskEndpointApiTasksStopTaskIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTaskEndpointApiTasksStopTaskIdPostRequest struct via the builder pattern


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


## UpdateWebhookUrlApiWebhookPost

> interface{} UpdateWebhookUrlApiWebhookPost(ctx).UrlForm(urlForm).Execute()

Update Webhook Url

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
	urlForm := *openapiclient.NewUrlForm("Url_example") // UrlForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWebhookUrlApiWebhookPost(context.Background()).UrlForm(urlForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhookUrlApiWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookUrlApiWebhookPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhookUrlApiWebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookUrlApiWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlForm** | [**UrlForm**](UrlForm.md) |  | 

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

