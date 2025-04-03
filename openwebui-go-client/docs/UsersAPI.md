# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserByIdApiV1UsersUserIdDelete**](UsersAPI.md#DeleteUserByIdApiV1UsersUserIdDelete) | **Delete** /api/v1/users/{user_id} | Delete User By Id
[**GetUserByIdApiV1UsersUserIdGet**](UsersAPI.md#GetUserByIdApiV1UsersUserIdGet) | **Get** /api/v1/users/{user_id} | Get User By Id
[**GetUserGroupsApiV1UsersGroupsGet**](UsersAPI.md#GetUserGroupsApiV1UsersGroupsGet) | **Get** /api/v1/users/groups | Get User Groups
[**GetUserInfoBySessionUserApiV1UsersUserInfoGet**](UsersAPI.md#GetUserInfoBySessionUserApiV1UsersUserInfoGet) | **Get** /api/v1/users/user/info | Get User Info By Session User
[**GetUserPermissionsApiV1UsersDefaultPermissionsGet**](UsersAPI.md#GetUserPermissionsApiV1UsersDefaultPermissionsGet) | **Get** /api/v1/users/default/permissions | Get User Permissions
[**GetUserPermissisionsApiV1UsersPermissionsGet**](UsersAPI.md#GetUserPermissisionsApiV1UsersPermissionsGet) | **Get** /api/v1/users/permissions | Get User Permissisions
[**GetUserSettingsBySessionUserApiV1UsersUserSettingsGet**](UsersAPI.md#GetUserSettingsBySessionUserApiV1UsersUserSettingsGet) | **Get** /api/v1/users/user/settings | Get User Settings By Session User
[**GetUsersApiV1UsersGet**](UsersAPI.md#GetUsersApiV1UsersGet) | **Get** /api/v1/users/ | Get Users
[**UpdateUserByIdApiV1UsersUserIdUpdatePost**](UsersAPI.md#UpdateUserByIdApiV1UsersUserIdUpdatePost) | **Post** /api/v1/users/{user_id}/update | Update User By Id
[**UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost**](UsersAPI.md#UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost) | **Post** /api/v1/users/user/info/update | Update User Info By Session User
[**UpdateUserPermissionsApiV1UsersDefaultPermissionsPost**](UsersAPI.md#UpdateUserPermissionsApiV1UsersDefaultPermissionsPost) | **Post** /api/v1/users/default/permissions | Update User Permissions
[**UpdateUserRoleApiV1UsersUpdateRolePost**](UsersAPI.md#UpdateUserRoleApiV1UsersUpdateRolePost) | **Post** /api/v1/users/update/role | Update User Role
[**UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost**](UsersAPI.md#UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost) | **Post** /api/v1/users/user/settings/update | Update User Settings By Session User



## DeleteUserByIdApiV1UsersUserIdDelete

> bool DeleteUserByIdApiV1UsersUserIdDelete(ctx, userId).Execute()

Delete User By Id

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserByIdApiV1UsersUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserByIdApiV1UsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserByIdApiV1UsersUserIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserByIdApiV1UsersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByIdApiV1UsersUserIdDeleteRequest struct via the builder pattern


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


## GetUserByIdApiV1UsersUserIdGet

> OpenWebuiRoutersUsersUserResponse GetUserByIdApiV1UsersUserIdGet(ctx, userId).Execute()

Get User By Id

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserByIdApiV1UsersUserIdGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserByIdApiV1UsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByIdApiV1UsersUserIdGet`: OpenWebuiRoutersUsersUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserByIdApiV1UsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdApiV1UsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenWebuiRoutersUsersUserResponse**](OpenWebuiRoutersUsersUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupsApiV1UsersGroupsGet

> interface{} GetUserGroupsApiV1UsersGroupsGet(ctx).Execute()

Get User Groups

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
	resp, r, err := apiClient.UsersAPI.GetUserGroupsApiV1UsersGroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupsApiV1UsersGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupsApiV1UsersGroupsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupsApiV1UsersGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsApiV1UsersGroupsGetRequest struct via the builder pattern


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


## GetUserInfoBySessionUserApiV1UsersUserInfoGet

> map[string]interface{} GetUserInfoBySessionUserApiV1UsersUserInfoGet(ctx).Execute()

Get User Info By Session User

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
	resp, r, err := apiClient.UsersAPI.GetUserInfoBySessionUserApiV1UsersUserInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserInfoBySessionUserApiV1UsersUserInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfoBySessionUserApiV1UsersUserInfoGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserInfoBySessionUserApiV1UsersUserInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoBySessionUserApiV1UsersUserInfoGetRequest struct via the builder pattern


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


## GetUserPermissionsApiV1UsersDefaultPermissionsGet

> UserPermissions GetUserPermissionsApiV1UsersDefaultPermissionsGet(ctx).Execute()

Get User Permissions

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
	resp, r, err := apiClient.UsersAPI.GetUserPermissionsApiV1UsersDefaultPermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissionsApiV1UsersDefaultPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissionsApiV1UsersDefaultPermissionsGet`: UserPermissions
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissionsApiV1UsersDefaultPermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsApiV1UsersDefaultPermissionsGetRequest struct via the builder pattern


### Return type

[**UserPermissions**](UserPermissions.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPermissisionsApiV1UsersPermissionsGet

> interface{} GetUserPermissisionsApiV1UsersPermissionsGet(ctx).Execute()

Get User Permissisions

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
	resp, r, err := apiClient.UsersAPI.GetUserPermissisionsApiV1UsersPermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissisionsApiV1UsersPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissisionsApiV1UsersPermissionsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissisionsApiV1UsersPermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissisionsApiV1UsersPermissionsGetRequest struct via the builder pattern


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


## GetUserSettingsBySessionUserApiV1UsersUserSettingsGet

> UserSettings GetUserSettingsBySessionUserApiV1UsersUserSettingsGet(ctx).Execute()

Get User Settings By Session User

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
	resp, r, err := apiClient.UsersAPI.GetUserSettingsBySessionUserApiV1UsersUserSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSettingsBySessionUserApiV1UsersUserSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSettingsBySessionUserApiV1UsersUserSettingsGet`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSettingsBySessionUserApiV1UsersUserSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsBySessionUserApiV1UsersUserSettingsGetRequest struct via the builder pattern


### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersApiV1UsersGet

> []UserModel GetUsersApiV1UsersGet(ctx).Skip(skip).Limit(limit).Execute()

Get Users

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
	skip := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersApiV1UsersGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersApiV1UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersApiV1UsersGet`: []UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersApiV1UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersApiV1UsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**[]UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserByIdApiV1UsersUserIdUpdatePost

> UserModel UpdateUserByIdApiV1UsersUserIdUpdatePost(ctx, userId).UserUpdateForm(userUpdateForm).Execute()

Update User By Id

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
	userId := "userId_example" // string | 
	userUpdateForm := *openapiclient.NewUserUpdateForm("Name_example", "Email_example", "ProfileImageUrl_example") // UserUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserByIdApiV1UsersUserIdUpdatePost(context.Background(), userId).UserUpdateForm(userUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserByIdApiV1UsersUserIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserByIdApiV1UsersUserIdUpdatePost`: UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserByIdApiV1UsersUserIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserByIdApiV1UsersUserIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateForm** | [**UserUpdateForm**](UserUpdateForm.md) |  | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost

> map[string]interface{} UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost(ctx).Body(body).Execute()

Update User Info By Session User

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
	resp, r, err := apiClient.UsersAPI.UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserInfoBySessionUserApiV1UsersUserInfoUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## UpdateUserPermissionsApiV1UsersDefaultPermissionsPost

> interface{} UpdateUserPermissionsApiV1UsersDefaultPermissionsPost(ctx).UserPermissions(userPermissions).Execute()

Update User Permissions

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
	userPermissions := *openapiclient.NewUserPermissions(*openapiclient.NewWorkspacePermissions(), *openapiclient.NewChatPermissions(), *openapiclient.NewFeaturesPermissions()) // UserPermissions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserPermissionsApiV1UsersDefaultPermissionsPost(context.Background()).UserPermissions(userPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserPermissionsApiV1UsersDefaultPermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPermissionsApiV1UsersDefaultPermissionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserPermissionsApiV1UsersDefaultPermissionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionsApiV1UsersDefaultPermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPermissions** | [**UserPermissions**](UserPermissions.md) |  | 

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


## UpdateUserRoleApiV1UsersUpdateRolePost

> UserModel UpdateUserRoleApiV1UsersUpdateRolePost(ctx).UserRoleUpdateForm(userRoleUpdateForm).Execute()

Update User Role

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
	userRoleUpdateForm := *openapiclient.NewUserRoleUpdateForm("Id_example", "Role_example") // UserRoleUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserRoleApiV1UsersUpdateRolePost(context.Background()).UserRoleUpdateForm(userRoleUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserRoleApiV1UsersUpdateRolePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserRoleApiV1UsersUpdateRolePost`: UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserRoleApiV1UsersUpdateRolePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRoleApiV1UsersUpdateRolePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRoleUpdateForm** | [**UserRoleUpdateForm**](UserRoleUpdateForm.md) |  | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost

> UserSettings UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost(ctx).UserSettings(userSettings).Execute()

Update User Settings By Session User

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
	userSettings := *openapiclient.NewUserSettings() // UserSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost(context.Background()).UserSettings(userSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsBySessionUserApiV1UsersUserSettingsUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSettings** | [**UserSettings**](UserSettings.md) |  | 

### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

