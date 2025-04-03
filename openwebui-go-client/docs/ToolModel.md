# ToolModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Content** | **string** |  | 
**Specs** | **[]map[string]interface{}** |  | 
**Meta** | [**ToolMeta**](ToolMeta.md) |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 

## Methods

### NewToolModel

`func NewToolModel(id string, userId string, name string, content string, specs []map[string]interface{}, meta ToolMeta, updatedAt int32, createdAt int32, ) *ToolModel`

NewToolModel instantiates a new ToolModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolModelWithDefaults

`func NewToolModelWithDefaults() *ToolModel`

NewToolModelWithDefaults instantiates a new ToolModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolModel) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ToolModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ToolModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ToolModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *ToolModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolModel) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *ToolModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ToolModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ToolModel) SetContent(v string)`

SetContent sets Content field to given value.


### GetSpecs

`func (o *ToolModel) GetSpecs() []map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *ToolModel) GetSpecsOk() (*[]map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *ToolModel) SetSpecs(v []map[string]interface{})`

SetSpecs sets Specs field to given value.


### GetMeta

`func (o *ToolModel) GetMeta() ToolMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ToolModel) GetMetaOk() (*ToolMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ToolModel) SetMeta(v ToolMeta)`

SetMeta sets Meta field to given value.


### GetAccessControl

`func (o *ToolModel) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ToolModel) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ToolModel) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ToolModel) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *ToolModel) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *ToolModel) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetUpdatedAt

`func (o *ToolModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ToolModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ToolModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ToolModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


