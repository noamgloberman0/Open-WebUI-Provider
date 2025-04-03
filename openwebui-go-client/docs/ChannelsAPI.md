# \ChannelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost**](ChannelsAPI.md#AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost) | **Post** /api/v1/channels/{id}/messages/{message_id}/reactions/add | Add Reaction To Message
[**CreateNewChannelApiV1ChannelsCreatePost**](ChannelsAPI.md#CreateNewChannelApiV1ChannelsCreatePost) | **Post** /api/v1/channels/create | Create New Channel
[**DeleteChannelByIdApiV1ChannelsIdDeleteDelete**](ChannelsAPI.md#DeleteChannelByIdApiV1ChannelsIdDeleteDelete) | **Delete** /api/v1/channels/{id}/delete | Delete Channel By Id
[**DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete**](ChannelsAPI.md#DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete) | **Delete** /api/v1/channels/{id}/messages/{message_id}/delete | Delete Message By Id
[**GetChannelByIdApiV1ChannelsIdGet**](ChannelsAPI.md#GetChannelByIdApiV1ChannelsIdGet) | **Get** /api/v1/channels/{id} | Get Channel By Id
[**GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet**](ChannelsAPI.md#GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet) | **Get** /api/v1/channels/{id}/messages/{message_id} | Get Channel Message
[**GetChannelMessagesApiV1ChannelsIdMessagesGet**](ChannelsAPI.md#GetChannelMessagesApiV1ChannelsIdMessagesGet) | **Get** /api/v1/channels/{id}/messages | Get Channel Messages
[**GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet**](ChannelsAPI.md#GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet) | **Get** /api/v1/channels/{id}/messages/{message_id}/thread | Get Channel Thread Messages
[**GetChannelsApiV1ChannelsGet**](ChannelsAPI.md#GetChannelsApiV1ChannelsGet) | **Get** /api/v1/channels/ | Get Channels
[**PostNewMessageApiV1ChannelsIdMessagesPostPost**](ChannelsAPI.md#PostNewMessageApiV1ChannelsIdMessagesPostPost) | **Post** /api/v1/channels/{id}/messages/post | Post New Message
[**RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost**](ChannelsAPI.md#RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost) | **Post** /api/v1/channels/{id}/messages/{message_id}/reactions/remove | Remove Reaction By Id And User Id And Name
[**UpdateChannelByIdApiV1ChannelsIdUpdatePost**](ChannelsAPI.md#UpdateChannelByIdApiV1ChannelsIdUpdatePost) | **Post** /api/v1/channels/{id}/update | Update Channel By Id
[**UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost**](ChannelsAPI.md#UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost) | **Post** /api/v1/channels/{id}/messages/{message_id}/update | Update Message By Id



## AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost

> bool AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost(ctx, id, messageId).ReactionForm(reactionForm).Execute()

Add Reaction To Message

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
	messageId := "messageId_example" // string | 
	reactionForm := *openapiclient.NewReactionForm("Name_example") // ReactionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost(context.Background(), id, messageId).ReactionForm(reactionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.AddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReactionToMessageApiV1ChannelsIdMessagesMessageIdReactionsAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reactionForm** | [**ReactionForm**](ReactionForm.md) |  | 

### Return type

**bool**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewChannelApiV1ChannelsCreatePost

> ChannelModel CreateNewChannelApiV1ChannelsCreatePost(ctx).ChannelForm(channelForm).Execute()

Create New Channel

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
	channelForm := *openapiclient.NewChannelForm("Name_example") // ChannelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CreateNewChannelApiV1ChannelsCreatePost(context.Background()).ChannelForm(channelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CreateNewChannelApiV1ChannelsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewChannelApiV1ChannelsCreatePost`: ChannelModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CreateNewChannelApiV1ChannelsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewChannelApiV1ChannelsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelForm** | [**ChannelForm**](ChannelForm.md) |  | 

### Return type

[**ChannelModel**](ChannelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelByIdApiV1ChannelsIdDeleteDelete

> bool DeleteChannelByIdApiV1ChannelsIdDeleteDelete(ctx, id).Execute()

Delete Channel By Id

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
	resp, r, err := apiClient.ChannelsAPI.DeleteChannelByIdApiV1ChannelsIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteChannelByIdApiV1ChannelsIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChannelByIdApiV1ChannelsIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.DeleteChannelByIdApiV1ChannelsIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelByIdApiV1ChannelsIdDeleteDeleteRequest struct via the builder pattern


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


## DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete

> bool DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete(ctx, id, messageId).Execute()

Delete Message By Id

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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete(context.Background(), id, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.DeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageByIdApiV1ChannelsIdMessagesMessageIdDeleteDeleteRequest struct via the builder pattern


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


## GetChannelByIdApiV1ChannelsIdGet

> ChannelModel GetChannelByIdApiV1ChannelsIdGet(ctx, id).Execute()

Get Channel By Id

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
	resp, r, err := apiClient.ChannelsAPI.GetChannelByIdApiV1ChannelsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelByIdApiV1ChannelsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelByIdApiV1ChannelsIdGet`: ChannelModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelByIdApiV1ChannelsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelByIdApiV1ChannelsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelModel**](ChannelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet

> MessageUserResponse GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet(ctx, id, messageId).Execute()

Get Channel Message

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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet(context.Background(), id, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet`: MessageUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelMessageApiV1ChannelsIdMessagesMessageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelMessageApiV1ChannelsIdMessagesMessageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageUserResponse**](MessageUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelMessagesApiV1ChannelsIdMessagesGet

> []MessageUserResponse GetChannelMessagesApiV1ChannelsIdMessagesGet(ctx, id).Skip(skip).Limit(limit).Execute()

Get Channel Messages

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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelMessagesApiV1ChannelsIdMessagesGet(context.Background(), id).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelMessagesApiV1ChannelsIdMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelMessagesApiV1ChannelsIdMessagesGet`: []MessageUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelMessagesApiV1ChannelsIdMessagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelMessagesApiV1ChannelsIdMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]MessageUserResponse**](MessageUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet

> []MessageUserResponse GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet(ctx, id, messageId).Skip(skip).Limit(limit).Execute()

Get Channel Thread Messages

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
	messageId := "messageId_example" // string | 
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet(context.Background(), id, messageId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet`: []MessageUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelThreadMessagesApiV1ChannelsIdMessagesMessageIdThreadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]MessageUserResponse**](MessageUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsApiV1ChannelsGet

> []ChannelModel GetChannelsApiV1ChannelsGet(ctx).Execute()

Get Channels

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
	resp, r, err := apiClient.ChannelsAPI.GetChannelsApiV1ChannelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelsApiV1ChannelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsApiV1ChannelsGet`: []ChannelModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelsApiV1ChannelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsApiV1ChannelsGetRequest struct via the builder pattern


### Return type

[**[]ChannelModel**](ChannelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNewMessageApiV1ChannelsIdMessagesPostPost

> MessageModel PostNewMessageApiV1ChannelsIdMessagesPostPost(ctx, id).MessageForm(messageForm).Execute()

Post New Message

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
	messageForm := *openapiclient.NewMessageForm("Content_example") // MessageForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.PostNewMessageApiV1ChannelsIdMessagesPostPost(context.Background(), id).MessageForm(messageForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PostNewMessageApiV1ChannelsIdMessagesPostPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNewMessageApiV1ChannelsIdMessagesPostPost`: MessageModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PostNewMessageApiV1ChannelsIdMessagesPostPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNewMessageApiV1ChannelsIdMessagesPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageForm** | [**MessageForm**](MessageForm.md) |  | 

### Return type

[**MessageModel**](MessageModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost

> bool RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost(ctx, id, messageId).ReactionForm(reactionForm).Execute()

Remove Reaction By Id And User Id And Name

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
	messageId := "messageId_example" // string | 
	reactionForm := *openapiclient.NewReactionForm("Name_example") // ReactionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost(context.Background(), id, messageId).ReactionForm(reactionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.RemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReactionByIdAndUserIdAndNameApiV1ChannelsIdMessagesMessageIdReactionsRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reactionForm** | [**ReactionForm**](ReactionForm.md) |  | 

### Return type

**bool**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelByIdApiV1ChannelsIdUpdatePost

> ChannelModel UpdateChannelByIdApiV1ChannelsIdUpdatePost(ctx, id).ChannelForm(channelForm).Execute()

Update Channel By Id

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
	channelForm := *openapiclient.NewChannelForm("Name_example") // ChannelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.UpdateChannelByIdApiV1ChannelsIdUpdatePost(context.Background(), id).ChannelForm(channelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateChannelByIdApiV1ChannelsIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannelByIdApiV1ChannelsIdUpdatePost`: ChannelModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateChannelByIdApiV1ChannelsIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelByIdApiV1ChannelsIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelForm** | [**ChannelForm**](ChannelForm.md) |  | 

### Return type

[**ChannelModel**](ChannelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost

> MessageModel UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost(ctx, id, messageId).MessageForm(messageForm).Execute()

Update Message By Id

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
	messageId := "messageId_example" // string | 
	messageForm := *openapiclient.NewMessageForm("Content_example") // MessageForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost(context.Background(), id, messageId).MessageForm(messageForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost`: MessageModel
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageByIdApiV1ChannelsIdMessagesMessageIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messageForm** | [**MessageForm**](MessageForm.md) |  | 

### Return type

[**MessageModel**](MessageModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

