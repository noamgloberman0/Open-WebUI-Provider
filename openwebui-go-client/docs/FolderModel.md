# FolderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Items** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**IsExpanded** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewFolderModel

`func NewFolderModel(id string, userId string, name string, createdAt int32, updatedAt int32, ) *FolderModel`

NewFolderModel instantiates a new FolderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderModelWithDefaults

`func NewFolderModelWithDefaults() *FolderModel`

NewFolderModelWithDefaults instantiates a new FolderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FolderModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FolderModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FolderModel) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *FolderModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FolderModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FolderModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FolderModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *FolderModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *FolderModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetUserId

`func (o *FolderModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FolderModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FolderModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *FolderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderModel) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *FolderModel) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FolderModel) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FolderModel) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *FolderModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *FolderModel) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *FolderModel) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetMeta

`func (o *FolderModel) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FolderModel) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FolderModel) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FolderModel) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *FolderModel) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *FolderModel) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetIsExpanded

`func (o *FolderModel) GetIsExpanded() bool`

GetIsExpanded returns the IsExpanded field if non-nil, zero value otherwise.

### GetIsExpandedOk

`func (o *FolderModel) GetIsExpandedOk() (*bool, bool)`

GetIsExpandedOk returns a tuple with the IsExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpanded

`func (o *FolderModel) SetIsExpanded(v bool)`

SetIsExpanded sets IsExpanded field to given value.

### HasIsExpanded

`func (o *FolderModel) HasIsExpanded() bool`

HasIsExpanded returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FolderModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FolderModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FolderModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FolderModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FolderModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FolderModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


