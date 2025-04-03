# PromptUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**UserId** | **string** |  | 
**Title** | **string** |  | 
**Content** | **string** |  | 
**Timestamp** | **int32** |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**User** | Pointer to [**NullableOpenWebuiModelsUsersUserResponse**](OpenWebuiModelsUsersUserResponse.md) |  | [optional] 

## Methods

### NewPromptUserResponse

`func NewPromptUserResponse(command string, userId string, title string, content string, timestamp int32, ) *PromptUserResponse`

NewPromptUserResponse instantiates a new PromptUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptUserResponseWithDefaults

`func NewPromptUserResponseWithDefaults() *PromptUserResponse`

NewPromptUserResponseWithDefaults instantiates a new PromptUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PromptUserResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PromptUserResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PromptUserResponse) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetUserId

`func (o *PromptUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PromptUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PromptUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTitle

`func (o *PromptUserResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PromptUserResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PromptUserResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *PromptUserResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PromptUserResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PromptUserResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetTimestamp

`func (o *PromptUserResponse) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PromptUserResponse) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PromptUserResponse) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetAccessControl

`func (o *PromptUserResponse) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *PromptUserResponse) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *PromptUserResponse) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *PromptUserResponse) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *PromptUserResponse) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *PromptUserResponse) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetUser

`func (o *PromptUserResponse) GetUser() OpenWebuiModelsUsersUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PromptUserResponse) GetUserOk() (*OpenWebuiModelsUsersUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PromptUserResponse) SetUser(v OpenWebuiModelsUsersUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *PromptUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PromptUserResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PromptUserResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


