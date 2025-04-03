# ModelUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**BaseModelId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Params** | **map[string]interface{}** |  | 
**Meta** | [**ModelMeta**](ModelMeta.md) |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**IsActive** | **bool** |  | 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**User** | Pointer to [**NullableOpenWebuiModelsUsersUserResponse**](OpenWebuiModelsUsersUserResponse.md) |  | [optional] 

## Methods

### NewModelUserResponse

`func NewModelUserResponse(id string, userId string, name string, params map[string]interface{}, meta ModelMeta, isActive bool, updatedAt int32, createdAt int32, ) *ModelUserResponse`

NewModelUserResponse instantiates a new ModelUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserResponseWithDefaults

`func NewModelUserResponseWithDefaults() *ModelUserResponse`

NewModelUserResponseWithDefaults instantiates a new ModelUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ModelUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBaseModelId

`func (o *ModelUserResponse) GetBaseModelId() string`

GetBaseModelId returns the BaseModelId field if non-nil, zero value otherwise.

### GetBaseModelIdOk

`func (o *ModelUserResponse) GetBaseModelIdOk() (*string, bool)`

GetBaseModelIdOk returns a tuple with the BaseModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModelId

`func (o *ModelUserResponse) SetBaseModelId(v string)`

SetBaseModelId sets BaseModelId field to given value.

### HasBaseModelId

`func (o *ModelUserResponse) HasBaseModelId() bool`

HasBaseModelId returns a boolean if a field has been set.

### SetBaseModelIdNil

`func (o *ModelUserResponse) SetBaseModelIdNil(b bool)`

 SetBaseModelIdNil sets the value for BaseModelId to be an explicit nil

### UnsetBaseModelId
`func (o *ModelUserResponse) UnsetBaseModelId()`

UnsetBaseModelId ensures that no value is present for BaseModelId, not even an explicit nil
### GetName

`func (o *ModelUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *ModelUserResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelUserResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelUserResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetMeta

`func (o *ModelUserResponse) GetMeta() ModelMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ModelUserResponse) GetMetaOk() (*ModelMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ModelUserResponse) SetMeta(v ModelMeta)`

SetMeta sets Meta field to given value.


### GetAccessControl

`func (o *ModelUserResponse) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ModelUserResponse) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ModelUserResponse) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ModelUserResponse) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *ModelUserResponse) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *ModelUserResponse) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetIsActive

`func (o *ModelUserResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelUserResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelUserResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUpdatedAt

`func (o *ModelUserResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelUserResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelUserResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ModelUserResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelUserResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelUserResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUser

`func (o *ModelUserResponse) GetUser() OpenWebuiModelsUsersUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelUserResponse) GetUserOk() (*OpenWebuiModelsUsersUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelUserResponse) SetUser(v OpenWebuiModelsUsersUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ModelUserResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ModelUserResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


