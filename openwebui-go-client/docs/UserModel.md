# UserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "pending"]
**ProfileImageUrl** | **string** |  | 
**LastActiveAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Settings** | Pointer to [**NullableUserSettings**](UserSettings.md) |  | [optional] 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 
**OauthSub** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserModel

`func NewUserModel(id string, name string, email string, profileImageUrl string, lastActiveAt int32, updatedAt int32, createdAt int32, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserModel) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *UserModel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserModel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserModel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserModel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetProfileImageUrl

`func (o *UserModel) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *UserModel) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *UserModel) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.


### GetLastActiveAt

`func (o *UserModel) GetLastActiveAt() int32`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *UserModel) GetLastActiveAtOk() (*int32, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *UserModel) SetLastActiveAt(v int32)`

SetLastActiveAt sets LastActiveAt field to given value.


### GetUpdatedAt

`func (o *UserModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *UserModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetApiKey

`func (o *UserModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *UserModel) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *UserModel) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetSettings

`func (o *UserModel) GetSettings() UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserModel) GetSettingsOk() (*UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserModel) SetSettings(v UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserModel) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *UserModel) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *UserModel) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetInfo

`func (o *UserModel) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UserModel) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UserModel) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UserModel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *UserModel) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *UserModel) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetOauthSub

`func (o *UserModel) GetOauthSub() string`

GetOauthSub returns the OauthSub field if non-nil, zero value otherwise.

### GetOauthSubOk

`func (o *UserModel) GetOauthSubOk() (*string, bool)`

GetOauthSubOk returns a tuple with the OauthSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthSub

`func (o *UserModel) SetOauthSub(v string)`

SetOauthSub sets OauthSub field to given value.

### HasOauthSub

`func (o *UserModel) HasOauthSub() bool`

HasOauthSub returns a boolean if a field has been set.

### SetOauthSubNil

`func (o *UserModel) SetOauthSubNil(b bool)`

 SetOauthSubNil sets the value for OauthSub to be an explicit nil

### UnsetOauthSub
`func (o *UserModel) UnsetOauthSub()`

UnsetOauthSub ensures that no value is present for OauthSub, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


