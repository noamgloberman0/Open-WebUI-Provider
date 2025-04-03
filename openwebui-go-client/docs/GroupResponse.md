# GroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] [default to []]
**CreatedAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewGroupResponse

`func NewGroupResponse(id string, userId string, name string, description string, createdAt int32, updatedAt int32, ) *GroupResponse`

NewGroupResponse instantiates a new GroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResponseWithDefaults

`func NewGroupResponseWithDefaults() *GroupResponse`

NewGroupResponseWithDefaults instantiates a new GroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *GroupResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *GroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissions

`func (o *GroupResponse) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupResponse) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupResponse) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *GroupResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GroupResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetData

`func (o *GroupResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GroupResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GroupResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GroupResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *GroupResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroupResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroupResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroupResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *GroupResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *GroupResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetUserIds

`func (o *GroupResponse) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GroupResponse) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GroupResponse) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GroupResponse) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GroupResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


