# KnowledgeResponse

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
**Files** | Pointer to [**[]KnowledgeResponseFilesInner**](KnowledgeResponseFilesInner.md) |  | [optional] 

## Methods

### NewKnowledgeResponse

`func NewKnowledgeResponse(id string, userId string, name string, description string, createdAt int32, updatedAt int32, ) *KnowledgeResponse`

NewKnowledgeResponse instantiates a new KnowledgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeResponseWithDefaults

`func NewKnowledgeResponseWithDefaults() *KnowledgeResponse`

NewKnowledgeResponseWithDefaults instantiates a new KnowledgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *KnowledgeResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KnowledgeResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KnowledgeResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *KnowledgeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KnowledgeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KnowledgeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KnowledgeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetData

`func (o *KnowledgeResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KnowledgeResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KnowledgeResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *KnowledgeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *KnowledgeResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KnowledgeResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *KnowledgeResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *KnowledgeResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *KnowledgeResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *KnowledgeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *KnowledgeResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *KnowledgeResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAccessControl

`func (o *KnowledgeResponse) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *KnowledgeResponse) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *KnowledgeResponse) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *KnowledgeResponse) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *KnowledgeResponse) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *KnowledgeResponse) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetCreatedAt

`func (o *KnowledgeResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KnowledgeResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KnowledgeResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KnowledgeResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KnowledgeResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KnowledgeResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFiles

`func (o *KnowledgeResponse) GetFiles() []KnowledgeResponseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeResponse) GetFilesOk() (*[]KnowledgeResponseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeResponse) SetFiles(v []KnowledgeResponseFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *KnowledgeResponse) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *KnowledgeResponse) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


