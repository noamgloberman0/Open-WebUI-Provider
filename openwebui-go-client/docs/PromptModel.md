# PromptModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**UserId** | **string** |  | 
**Title** | **string** |  | 
**Content** | **string** |  | 
**Timestamp** | **int32** |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPromptModel

`func NewPromptModel(command string, userId string, title string, content string, timestamp int32, ) *PromptModel`

NewPromptModel instantiates a new PromptModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptModelWithDefaults

`func NewPromptModelWithDefaults() *PromptModel`

NewPromptModelWithDefaults instantiates a new PromptModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PromptModel) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PromptModel) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PromptModel) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetUserId

`func (o *PromptModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PromptModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PromptModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTitle

`func (o *PromptModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PromptModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PromptModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *PromptModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PromptModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PromptModel) SetContent(v string)`

SetContent sets Content field to given value.


### GetTimestamp

`func (o *PromptModel) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PromptModel) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PromptModel) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetAccessControl

`func (o *PromptModel) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *PromptModel) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *PromptModel) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *PromptModel) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *PromptModel) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *PromptModel) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


