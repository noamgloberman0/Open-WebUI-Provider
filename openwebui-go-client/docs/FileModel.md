# FileModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Hash** | Pointer to **NullableString** |  | [optional] 
**Filename** | **string** |  | 
**Path** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **NullableInt32** |  | 
**UpdatedAt** | **NullableInt32** |  | 

## Methods

### NewFileModel

`func NewFileModel(id string, userId string, filename string, createdAt NullableInt32, updatedAt NullableInt32, ) *FileModel`

NewFileModel instantiates a new FileModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileModelWithDefaults

`func NewFileModelWithDefaults() *FileModel`

NewFileModelWithDefaults instantiates a new FileModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileModel) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *FileModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHash

`func (o *FileModel) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *FileModel) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *FileModel) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *FileModel) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *FileModel) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *FileModel) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFilename

`func (o *FileModel) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileModel) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileModel) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPath

`func (o *FileModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *FileModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *FileModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetData

`func (o *FileModel) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FileModel) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FileModel) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FileModel) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *FileModel) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *FileModel) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *FileModel) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileModel) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileModel) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileModel) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *FileModel) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *FileModel) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAccessControl

`func (o *FileModel) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *FileModel) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *FileModel) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *FileModel) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *FileModel) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *FileModel) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetCreatedAt

`func (o *FileModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *FileModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *FileModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *FileModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FileModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FileModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


