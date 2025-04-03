# ModelModel

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

## Methods

### NewModelModel

`func NewModelModel(id string, userId string, name string, params map[string]interface{}, meta ModelMeta, isActive bool, updatedAt int32, createdAt int32, ) *ModelModel`

NewModelModel instantiates a new ModelModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelModelWithDefaults

`func NewModelModelWithDefaults() *ModelModel`

NewModelModelWithDefaults instantiates a new ModelModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelModel) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ModelModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBaseModelId

`func (o *ModelModel) GetBaseModelId() string`

GetBaseModelId returns the BaseModelId field if non-nil, zero value otherwise.

### GetBaseModelIdOk

`func (o *ModelModel) GetBaseModelIdOk() (*string, bool)`

GetBaseModelIdOk returns a tuple with the BaseModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModelId

`func (o *ModelModel) SetBaseModelId(v string)`

SetBaseModelId sets BaseModelId field to given value.

### HasBaseModelId

`func (o *ModelModel) HasBaseModelId() bool`

HasBaseModelId returns a boolean if a field has been set.

### SetBaseModelIdNil

`func (o *ModelModel) SetBaseModelIdNil(b bool)`

 SetBaseModelIdNil sets the value for BaseModelId to be an explicit nil

### UnsetBaseModelId
`func (o *ModelModel) UnsetBaseModelId()`

UnsetBaseModelId ensures that no value is present for BaseModelId, not even an explicit nil
### GetName

`func (o *ModelModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelModel) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *ModelModel) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelModel) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelModel) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetMeta

`func (o *ModelModel) GetMeta() ModelMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ModelModel) GetMetaOk() (*ModelMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ModelModel) SetMeta(v ModelMeta)`

SetMeta sets Meta field to given value.


### GetAccessControl

`func (o *ModelModel) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ModelModel) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ModelModel) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ModelModel) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *ModelModel) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *ModelModel) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetIsActive

`func (o *ModelModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUpdatedAt

`func (o *ModelModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ModelModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


