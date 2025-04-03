# AddUserForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Password** | **string** |  | 
**ProfileImageUrl** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddUserForm

`func NewAddUserForm(name string, email string, password string, ) *AddUserForm`

NewAddUserForm instantiates a new AddUserForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserFormWithDefaults

`func NewAddUserFormWithDefaults() *AddUserForm`

NewAddUserFormWithDefaults instantiates a new AddUserForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddUserForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUserForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUserForm) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *AddUserForm) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUserForm) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUserForm) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AddUserForm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUserForm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUserForm) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProfileImageUrl

`func (o *AddUserForm) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *AddUserForm) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *AddUserForm) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.

### HasProfileImageUrl

`func (o *AddUserForm) HasProfileImageUrl() bool`

HasProfileImageUrl returns a boolean if a field has been set.

### SetProfileImageUrlNil

`func (o *AddUserForm) SetProfileImageUrlNil(b bool)`

 SetProfileImageUrlNil sets the value for ProfileImageUrl to be an explicit nil

### UnsetProfileImageUrl
`func (o *AddUserForm) UnsetProfileImageUrl()`

UnsetProfileImageUrl ensures that no value is present for ProfileImageUrl, not even an explicit nil
### GetRole

`func (o *AddUserForm) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddUserForm) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddUserForm) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddUserForm) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AddUserForm) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AddUserForm) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


