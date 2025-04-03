# \GroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewGroupApiV1GroupsCreatePost**](GroupsAPI.md#CreateNewGroupApiV1GroupsCreatePost) | **Post** /api/v1/groups/create | Create New Group
[**DeleteGroupByIdApiV1GroupsIdIdDeleteDelete**](GroupsAPI.md#DeleteGroupByIdApiV1GroupsIdIdDeleteDelete) | **Delete** /api/v1/groups/id/{id}/delete | Delete Group By Id
[**GetGroupByIdApiV1GroupsIdIdGet**](GroupsAPI.md#GetGroupByIdApiV1GroupsIdIdGet) | **Get** /api/v1/groups/id/{id} | Get Group By Id
[**GetGroupsApiV1GroupsGet**](GroupsAPI.md#GetGroupsApiV1GroupsGet) | **Get** /api/v1/groups/ | Get Groups
[**UpdateGroupByIdApiV1GroupsIdIdUpdatePost**](GroupsAPI.md#UpdateGroupByIdApiV1GroupsIdIdUpdatePost) | **Post** /api/v1/groups/id/{id}/update | Update Group By Id



## CreateNewGroupApiV1GroupsCreatePost

> GroupResponse CreateNewGroupApiV1GroupsCreatePost(ctx).GroupForm(groupForm).Execute()

Create New Group

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
	groupForm := *openapiclient.NewGroupForm("Name_example", "Description_example") // GroupForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateNewGroupApiV1GroupsCreatePost(context.Background()).GroupForm(groupForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateNewGroupApiV1GroupsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewGroupApiV1GroupsCreatePost`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateNewGroupApiV1GroupsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewGroupApiV1GroupsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupForm** | [**GroupForm**](GroupForm.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupByIdApiV1GroupsIdIdDeleteDelete

> bool DeleteGroupByIdApiV1GroupsIdIdDeleteDelete(ctx, id).Execute()

Delete Group By Id

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
	resp, r, err := apiClient.GroupsAPI.DeleteGroupByIdApiV1GroupsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupByIdApiV1GroupsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupByIdApiV1GroupsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupByIdApiV1GroupsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupByIdApiV1GroupsIdIdDeleteDeleteRequest struct via the builder pattern


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


## GetGroupByIdApiV1GroupsIdIdGet

> GroupResponse GetGroupByIdApiV1GroupsIdIdGet(ctx, id).Execute()

Get Group By Id

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
	resp, r, err := apiClient.GroupsAPI.GetGroupByIdApiV1GroupsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupByIdApiV1GroupsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupByIdApiV1GroupsIdIdGet`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupByIdApiV1GroupsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByIdApiV1GroupsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsApiV1GroupsGet

> []GroupResponse GetGroupsApiV1GroupsGet(ctx).Execute()

Get Groups

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
	resp, r, err := apiClient.GroupsAPI.GetGroupsApiV1GroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsApiV1GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsApiV1GroupsGet`: []GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsApiV1GroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsApiV1GroupsGetRequest struct via the builder pattern


### Return type

[**[]GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupByIdApiV1GroupsIdIdUpdatePost

> GroupResponse UpdateGroupByIdApiV1GroupsIdIdUpdatePost(ctx, id).GroupUpdateForm(groupUpdateForm).Execute()

Update Group By Id

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
	groupUpdateForm := *openapiclient.NewGroupUpdateForm("Name_example", "Description_example") // GroupUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupByIdApiV1GroupsIdIdUpdatePost(context.Background(), id).GroupUpdateForm(groupUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupByIdApiV1GroupsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupByIdApiV1GroupsIdIdUpdatePost`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupByIdApiV1GroupsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupByIdApiV1GroupsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateForm** | [**GroupUpdateForm**](GroupUpdateForm.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

