# UserUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**ProfileImageUrl** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserUpdateForm

`func NewUserUpdateForm(name string, email string, profileImageUrl string, ) *UserUpdateForm`

NewUserUpdateForm instantiates a new UserUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateFormWithDefaults

`func NewUserUpdateFormWithDefaults() *UserUpdateForm`

NewUserUpdateFormWithDefaults instantiates a new UserUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserUpdateForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdateForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdateForm) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UserUpdateForm) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdateForm) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdateForm) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProfileImageUrl

`func (o *UserUpdateForm) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *UserUpdateForm) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *UserUpdateForm) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.


### GetPassword

`func (o *UserUpdateForm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserUpdateForm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserUpdateForm) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserUpdateForm) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UserUpdateForm) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserUpdateForm) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


