# \AuthsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserApiV1AuthsAddPost**](AuthsAPI.md#AddUserApiV1AuthsAddPost) | **Post** /api/v1/auths/add | Add User
[**DeleteApiKeyApiV1AuthsApiKeyDelete**](AuthsAPI.md#DeleteApiKeyApiV1AuthsApiKeyDelete) | **Delete** /api/v1/auths/api_key | Delete Api Key
[**GenerateApiKeyApiV1AuthsApiKeyPost**](AuthsAPI.md#GenerateApiKeyApiV1AuthsApiKeyPost) | **Post** /api/v1/auths/api_key | Generate Api Key
[**GetAdminConfigApiV1AuthsAdminConfigGet**](AuthsAPI.md#GetAdminConfigApiV1AuthsAdminConfigGet) | **Get** /api/v1/auths/admin/config | Get Admin Config
[**GetAdminDetailsApiV1AuthsAdminDetailsGet**](AuthsAPI.md#GetAdminDetailsApiV1AuthsAdminDetailsGet) | **Get** /api/v1/auths/admin/details | Get Admin Details
[**GetApiKeyApiV1AuthsApiKeyGet**](AuthsAPI.md#GetApiKeyApiV1AuthsApiKeyGet) | **Get** /api/v1/auths/api_key | Get Api Key
[**GetLdapConfigApiV1AuthsAdminConfigLdapGet**](AuthsAPI.md#GetLdapConfigApiV1AuthsAdminConfigLdapGet) | **Get** /api/v1/auths/admin/config/ldap | Get Ldap Config
[**GetLdapServerApiV1AuthsAdminConfigLdapServerGet**](AuthsAPI.md#GetLdapServerApiV1AuthsAdminConfigLdapServerGet) | **Get** /api/v1/auths/admin/config/ldap/server | Get Ldap Server
[**GetSessionUserApiV1AuthsGet**](AuthsAPI.md#GetSessionUserApiV1AuthsGet) | **Get** /api/v1/auths/ | Get Session User
[**LdapAuthApiV1AuthsLdapPost**](AuthsAPI.md#LdapAuthApiV1AuthsLdapPost) | **Post** /api/v1/auths/ldap | Ldap Auth
[**SigninApiV1AuthsSigninPost**](AuthsAPI.md#SigninApiV1AuthsSigninPost) | **Post** /api/v1/auths/signin | Signin
[**SignoutApiV1AuthsSignoutGet**](AuthsAPI.md#SignoutApiV1AuthsSignoutGet) | **Get** /api/v1/auths/signout | Signout
[**SignupApiV1AuthsSignupPost**](AuthsAPI.md#SignupApiV1AuthsSignupPost) | **Post** /api/v1/auths/signup | Signup
[**UpdateAdminConfigApiV1AuthsAdminConfigPost**](AuthsAPI.md#UpdateAdminConfigApiV1AuthsAdminConfigPost) | **Post** /api/v1/auths/admin/config | Update Admin Config
[**UpdateLdapConfigApiV1AuthsAdminConfigLdapPost**](AuthsAPI.md#UpdateLdapConfigApiV1AuthsAdminConfigLdapPost) | **Post** /api/v1/auths/admin/config/ldap | Update Ldap Config
[**UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost**](AuthsAPI.md#UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost) | **Post** /api/v1/auths/admin/config/ldap/server | Update Ldap Server
[**UpdatePasswordApiV1AuthsUpdatePasswordPost**](AuthsAPI.md#UpdatePasswordApiV1AuthsUpdatePasswordPost) | **Post** /api/v1/auths/update/password | Update Password
[**UpdateProfileApiV1AuthsUpdateProfilePost**](AuthsAPI.md#UpdateProfileApiV1AuthsUpdateProfilePost) | **Post** /api/v1/auths/update/profile | Update Profile



## AddUserApiV1AuthsAddPost

> SigninResponse AddUserApiV1AuthsAddPost(ctx).AddUserForm(addUserForm).Execute()

Add User

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
	addUserForm := *openapiclient.NewAddUserForm("Name_example", "Email_example", "Password_example") // AddUserForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.AddUserApiV1AuthsAddPost(context.Background()).AddUserForm(addUserForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.AddUserApiV1AuthsAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserApiV1AuthsAddPost`: SigninResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.AddUserApiV1AuthsAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserApiV1AuthsAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserForm** | [**AddUserForm**](AddUserForm.md) |  | 

### Return type

[**SigninResponse**](SigninResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeyApiV1AuthsApiKeyDelete

> bool DeleteApiKeyApiV1AuthsApiKeyDelete(ctx).Execute()

Delete Api Key

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
	resp, r, err := apiClient.AuthsAPI.DeleteApiKeyApiV1AuthsApiKeyDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.DeleteApiKeyApiV1AuthsApiKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKeyApiV1AuthsApiKeyDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.DeleteApiKeyApiV1AuthsApiKeyDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyApiV1AuthsApiKeyDeleteRequest struct via the builder pattern


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


## GenerateApiKeyApiV1AuthsApiKeyPost

> ApiKey GenerateApiKeyApiV1AuthsApiKeyPost(ctx).Execute()

Generate Api Key

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
	resp, r, err := apiClient.AuthsAPI.GenerateApiKeyApiV1AuthsApiKeyPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GenerateApiKeyApiV1AuthsApiKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateApiKeyApiV1AuthsApiKeyPost`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GenerateApiKeyApiV1AuthsApiKeyPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateApiKeyApiV1AuthsApiKeyPostRequest struct via the builder pattern


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminConfigApiV1AuthsAdminConfigGet

> interface{} GetAdminConfigApiV1AuthsAdminConfigGet(ctx).Execute()

Get Admin Config

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
	resp, r, err := apiClient.AuthsAPI.GetAdminConfigApiV1AuthsAdminConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetAdminConfigApiV1AuthsAdminConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminConfigApiV1AuthsAdminConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetAdminConfigApiV1AuthsAdminConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminConfigApiV1AuthsAdminConfigGetRequest struct via the builder pattern


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


## GetAdminDetailsApiV1AuthsAdminDetailsGet

> interface{} GetAdminDetailsApiV1AuthsAdminDetailsGet(ctx).Execute()

Get Admin Details

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
	resp, r, err := apiClient.AuthsAPI.GetAdminDetailsApiV1AuthsAdminDetailsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetAdminDetailsApiV1AuthsAdminDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminDetailsApiV1AuthsAdminDetailsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetAdminDetailsApiV1AuthsAdminDetailsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminDetailsApiV1AuthsAdminDetailsGetRequest struct via the builder pattern


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


## GetApiKeyApiV1AuthsApiKeyGet

> ApiKey GetApiKeyApiV1AuthsApiKeyGet(ctx).Execute()

Get Api Key

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
	resp, r, err := apiClient.AuthsAPI.GetApiKeyApiV1AuthsApiKeyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetApiKeyApiV1AuthsApiKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyApiV1AuthsApiKeyGet`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetApiKeyApiV1AuthsApiKeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyApiV1AuthsApiKeyGetRequest struct via the builder pattern


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConfigApiV1AuthsAdminConfigLdapGet

> interface{} GetLdapConfigApiV1AuthsAdminConfigLdapGet(ctx).Execute()

Get Ldap Config

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
	resp, r, err := apiClient.AuthsAPI.GetLdapConfigApiV1AuthsAdminConfigLdapGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetLdapConfigApiV1AuthsAdminConfigLdapGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapConfigApiV1AuthsAdminConfigLdapGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetLdapConfigApiV1AuthsAdminConfigLdapGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConfigApiV1AuthsAdminConfigLdapGetRequest struct via the builder pattern


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


## GetLdapServerApiV1AuthsAdminConfigLdapServerGet

> LdapServerConfig GetLdapServerApiV1AuthsAdminConfigLdapServerGet(ctx).Execute()

Get Ldap Server

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
	resp, r, err := apiClient.AuthsAPI.GetLdapServerApiV1AuthsAdminConfigLdapServerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetLdapServerApiV1AuthsAdminConfigLdapServerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapServerApiV1AuthsAdminConfigLdapServerGet`: LdapServerConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetLdapServerApiV1AuthsAdminConfigLdapServerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServerApiV1AuthsAdminConfigLdapServerGetRequest struct via the builder pattern


### Return type

[**LdapServerConfig**](LdapServerConfig.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionUserApiV1AuthsGet

> SessionUserResponse GetSessionUserApiV1AuthsGet(ctx).Execute()

Get Session User

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
	resp, r, err := apiClient.AuthsAPI.GetSessionUserApiV1AuthsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetSessionUserApiV1AuthsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionUserApiV1AuthsGet`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetSessionUserApiV1AuthsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionUserApiV1AuthsGetRequest struct via the builder pattern


### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapAuthApiV1AuthsLdapPost

> SessionUserResponse LdapAuthApiV1AuthsLdapPost(ctx).LdapForm(ldapForm).Execute()

Ldap Auth

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
	ldapForm := *openapiclient.NewLdapForm("User_example", "Password_example") // LdapForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.LdapAuthApiV1AuthsLdapPost(context.Background()).LdapForm(ldapForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.LdapAuthApiV1AuthsLdapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LdapAuthApiV1AuthsLdapPost`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.LdapAuthApiV1AuthsLdapPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLdapAuthApiV1AuthsLdapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapForm** | [**LdapForm**](LdapForm.md) |  | 

### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninApiV1AuthsSigninPost

> SessionUserResponse SigninApiV1AuthsSigninPost(ctx).SigninForm(signinForm).Execute()

Signin

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
	signinForm := *openapiclient.NewSigninForm("Email_example", "Password_example") // SigninForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.SigninApiV1AuthsSigninPost(context.Background()).SigninForm(signinForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SigninApiV1AuthsSigninPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninApiV1AuthsSigninPost`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SigninApiV1AuthsSigninPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiV1AuthsSigninPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signinForm** | [**SigninForm**](SigninForm.md) |  | 

### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignoutApiV1AuthsSignoutGet

> interface{} SignoutApiV1AuthsSignoutGet(ctx).Execute()

Signout

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
	resp, r, err := apiClient.AuthsAPI.SignoutApiV1AuthsSignoutGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SignoutApiV1AuthsSignoutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignoutApiV1AuthsSignoutGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SignoutApiV1AuthsSignoutGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSignoutApiV1AuthsSignoutGetRequest struct via the builder pattern


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


## SignupApiV1AuthsSignupPost

> SessionUserResponse SignupApiV1AuthsSignupPost(ctx).SignupForm(signupForm).Execute()

Signup

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
	signupForm := *openapiclient.NewSignupForm("Name_example", "Email_example", "Password_example") // SignupForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.SignupApiV1AuthsSignupPost(context.Background()).SignupForm(signupForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SignupApiV1AuthsSignupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignupApiV1AuthsSignupPost`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SignupApiV1AuthsSignupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiV1AuthsSignupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupForm** | [**SignupForm**](SignupForm.md) |  | 

### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminConfigApiV1AuthsAdminConfigPost

> interface{} UpdateAdminConfigApiV1AuthsAdminConfigPost(ctx).AdminConfig(adminConfig).Execute()

Update Admin Config

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
	adminConfig := *openapiclient.NewAdminConfig(false, "WEBUI_URL_example", false, false, false, "API_KEY_ALLOWED_ENDPOINTS_example", false, "DEFAULT_USER_ROLE_example", "JWT_EXPIRES_IN_example", false, false) // AdminConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateAdminConfigApiV1AuthsAdminConfigPost(context.Background()).AdminConfig(adminConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateAdminConfigApiV1AuthsAdminConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminConfigApiV1AuthsAdminConfigPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateAdminConfigApiV1AuthsAdminConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminConfigApiV1AuthsAdminConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminConfig** | [**AdminConfig**](AdminConfig.md) |  | 

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


## UpdateLdapConfigApiV1AuthsAdminConfigLdapPost

> interface{} UpdateLdapConfigApiV1AuthsAdminConfigLdapPost(ctx).LdapConfigForm(ldapConfigForm).Execute()

Update Ldap Config

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
	ldapConfigForm := *openapiclient.NewLdapConfigForm() // LdapConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateLdapConfigApiV1AuthsAdminConfigLdapPost(context.Background()).LdapConfigForm(ldapConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateLdapConfigApiV1AuthsAdminConfigLdapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLdapConfigApiV1AuthsAdminConfigLdapPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateLdapConfigApiV1AuthsAdminConfigLdapPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapConfigApiV1AuthsAdminConfigLdapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigForm** | [**LdapConfigForm**](LdapConfigForm.md) |  | 

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


## UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost

> interface{} UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost(ctx).LdapServerConfig(ldapServerConfig).Execute()

Update Ldap Server

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
	ldapServerConfig := *openapiclient.NewLdapServerConfig("Label_example", "Host_example", "AppDn_example", "AppDnPassword_example", "SearchBase_example") // LdapServerConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost(context.Background()).LdapServerConfig(ldapServerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateLdapServerApiV1AuthsAdminConfigLdapServerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapServerApiV1AuthsAdminConfigLdapServerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapServerConfig** | [**LdapServerConfig**](LdapServerConfig.md) |  | 

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


## UpdatePasswordApiV1AuthsUpdatePasswordPost

> bool UpdatePasswordApiV1AuthsUpdatePasswordPost(ctx).UpdatePasswordForm(updatePasswordForm).Execute()

Update Password

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
	updatePasswordForm := *openapiclient.NewUpdatePasswordForm("Password_example", "NewPassword_example") // UpdatePasswordForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdatePasswordApiV1AuthsUpdatePasswordPost(context.Background()).UpdatePasswordForm(updatePasswordForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdatePasswordApiV1AuthsUpdatePasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePasswordApiV1AuthsUpdatePasswordPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdatePasswordApiV1AuthsUpdatePasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordApiV1AuthsUpdatePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordForm** | [**UpdatePasswordForm**](UpdatePasswordForm.md) |  | 

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


## UpdateProfileApiV1AuthsUpdateProfilePost

> OpenWebuiModelsAuthsUserResponse UpdateProfileApiV1AuthsUpdateProfilePost(ctx).UpdateProfileForm(updateProfileForm).Execute()

Update Profile

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
	updateProfileForm := *openapiclient.NewUpdateProfileForm("ProfileImageUrl_example", "Name_example") // UpdateProfileForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateProfileApiV1AuthsUpdateProfilePost(context.Background()).UpdateProfileForm(updateProfileForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateProfileApiV1AuthsUpdateProfilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfileApiV1AuthsUpdateProfilePost`: OpenWebuiModelsAuthsUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateProfileApiV1AuthsUpdateProfilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileApiV1AuthsUpdateProfilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileForm** | [**UpdateProfileForm**](UpdateProfileForm.md) |  | 

### Return type

[**OpenWebuiModelsAuthsUserResponse**](OpenWebuiModelsAuthsUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

