# ToolUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Meta** | [**ToolMeta**](ToolMeta.md) |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**User** | Pointer to [**NullableOpenWebuiModelsUsersUserResponse**](OpenWebuiModelsUsersUserResponse.md) |  | [optional] 

## Methods

### NewToolUserResponse

`func NewToolUserResponse(id string, userId string, name string, meta ToolMeta, updatedAt int32, createdAt int32, ) *ToolUserResponse`

NewToolUserResponse instantiates a new ToolUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolUserResponseWithDefaults

`func NewToolUserResponseWithDefaults() *ToolUserResponse`

NewToolUserResponseWithDefaults instantiates a new ToolUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ToolUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ToolUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ToolUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *ToolUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *ToolUserResponse) GetMeta() ToolMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ToolUserResponse) GetMetaOk() (*ToolMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ToolUserResponse) SetMeta(v ToolMeta)`

SetMeta sets Meta field to given value.


### GetAccessControl

`func (o *ToolUserResponse) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ToolUserResponse) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ToolUserResponse) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ToolUserResponse) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *ToolUserResponse) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *ToolUserResponse) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetUpdatedAt

`func (o *ToolUserResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ToolUserResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ToolUserResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ToolUserResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolUserResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolUserResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUser

`func (o *ToolUserResponse) GetUser() OpenWebuiModelsUsersUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ToolUserResponse) GetUserOk() (*OpenWebuiModelsUsersUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ToolUserResponse) SetUser(v OpenWebuiModelsUsersUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ToolUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ToolUserResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ToolUserResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


