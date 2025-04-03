# SignupForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Password** | **string** |  | 
**ProfileImageUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignupForm

`func NewSignupForm(name string, email string, password string, ) *SignupForm`

NewSignupForm instantiates a new SignupForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupFormWithDefaults

`func NewSignupFormWithDefaults() *SignupForm`

NewSignupFormWithDefaults instantiates a new SignupForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignupForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignupForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignupForm) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *SignupForm) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignupForm) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignupForm) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SignupForm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignupForm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignupForm) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProfileImageUrl

`func (o *SignupForm) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *SignupForm) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *SignupForm) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.

### HasProfileImageUrl

`func (o *SignupForm) HasProfileImageUrl() bool`

HasProfileImageUrl returns a boolean if a field has been set.

### SetProfileImageUrlNil

`func (o *SignupForm) SetProfileImageUrlNil(b bool)`

 SetProfileImageUrlNil sets the value for ProfileImageUrl to be an explicit nil

### UnsetProfileImageUrl
`func (o *SignupForm) UnsetProfileImageUrl()`

UnsetProfileImageUrl ensures that no value is present for ProfileImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


