# KnowledgeUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 
**User** | Pointer to [**NullableOpenWebuiModelsUsersUserResponse**](OpenWebuiModelsUsersUserResponse.md) |  | [optional] 
**Files** | Pointer to [**[]KnowledgeResponseFilesInner**](KnowledgeResponseFilesInner.md) |  | [optional] 

## Methods

### NewKnowledgeUserResponse

`func NewKnowledgeUserResponse(id string, userId string, name string, description string, createdAt int32, updatedAt int32, ) *KnowledgeUserResponse`

NewKnowledgeUserResponse instantiates a new KnowledgeUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeUserResponseWithDefaults

`func NewKnowledgeUserResponseWithDefaults() *KnowledgeUserResponse`

NewKnowledgeUserResponseWithDefaults instantiates a new KnowledgeUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *KnowledgeUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KnowledgeUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KnowledgeUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *KnowledgeUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KnowledgeUserResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KnowledgeUserResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KnowledgeUserResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetData

`func (o *KnowledgeUserResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KnowledgeUserResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KnowledgeUserResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *KnowledgeUserResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *KnowledgeUserResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KnowledgeUserResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *KnowledgeUserResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *KnowledgeUserResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *KnowledgeUserResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *KnowledgeUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *KnowledgeUserResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *KnowledgeUserResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAccessControl

`func (o *KnowledgeUserResponse) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *KnowledgeUserResponse) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *KnowledgeUserResponse) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *KnowledgeUserResponse) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *KnowledgeUserResponse) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *KnowledgeUserResponse) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetCreatedAt

`func (o *KnowledgeUserResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KnowledgeUserResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KnowledgeUserResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KnowledgeUserResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KnowledgeUserResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KnowledgeUserResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *KnowledgeUserResponse) GetUser() OpenWebuiModelsUsersUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KnowledgeUserResponse) GetUserOk() (*OpenWebuiModelsUsersUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KnowledgeUserResponse) SetUser(v OpenWebuiModelsUsersUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *KnowledgeUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *KnowledgeUserResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *KnowledgeUserResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetFiles

`func (o *KnowledgeUserResponse) GetFiles() []KnowledgeResponseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeUserResponse) GetFilesOk() (*[]KnowledgeResponseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeUserResponse) SetFiles(v []KnowledgeResponseFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeUserResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *KnowledgeUserResponse) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *KnowledgeUserResponse) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


