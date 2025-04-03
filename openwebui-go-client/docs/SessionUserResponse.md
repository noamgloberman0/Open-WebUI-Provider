# SessionUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Role** | **string** |  | 
**ProfileImageUrl** | **string** |  | 
**Token** | **string** |  | 
**TokenType** | **string** |  | 
**ExpiresAt** | Pointer to **NullableInt32** |  | [optional] 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSessionUserResponse

`func NewSessionUserResponse(id string, email string, name string, role string, profileImageUrl string, token string, tokenType string, ) *SessionUserResponse`

NewSessionUserResponse instantiates a new SessionUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionUserResponseWithDefaults

`func NewSessionUserResponseWithDefaults() *SessionUserResponse`

NewSessionUserResponseWithDefaults instantiates a new SessionUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *SessionUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SessionUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SessionUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *SessionUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *SessionUserResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SessionUserResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SessionUserResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetProfileImageUrl

`func (o *SessionUserResponse) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *SessionUserResponse) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *SessionUserResponse) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.


### GetToken

`func (o *SessionUserResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SessionUserResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SessionUserResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTokenType

`func (o *SessionUserResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *SessionUserResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *SessionUserResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresAt

`func (o *SessionUserResponse) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SessionUserResponse) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SessionUserResponse) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SessionUserResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *SessionUserResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SessionUserResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetPermissions

`func (o *SessionUserResponse) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SessionUserResponse) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SessionUserResponse) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SessionUserResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SessionUserResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SessionUserResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


