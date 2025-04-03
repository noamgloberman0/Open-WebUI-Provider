# FileModelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Hash** | Pointer to [**Hash**](Hash.md) |  | [optional] 
**Filename** | **string** |  | 
**Data** | Pointer to [**Data**](Data.md) |  | [optional] 
**Meta** | [**FileMeta**](FileMeta.md) |  | 
**CreatedAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewFileModelResponse

`func NewFileModelResponse(id string, userId string, filename string, meta FileMeta, createdAt int32, updatedAt int32, ) *FileModelResponse`

NewFileModelResponse instantiates a new FileModelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileModelResponseWithDefaults

`func NewFileModelResponseWithDefaults() *FileModelResponse`

NewFileModelResponseWithDefaults instantiates a new FileModelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileModelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileModelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileModelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *FileModelResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileModelResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileModelResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHash

`func (o *FileModelResponse) GetHash() Hash`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *FileModelResponse) GetHashOk() (*Hash, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *FileModelResponse) SetHash(v Hash)`

SetHash sets Hash field to given value.

### HasHash

`func (o *FileModelResponse) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetFilename

`func (o *FileModelResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileModelResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileModelResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetData

`func (o *FileModelResponse) GetData() Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FileModelResponse) GetDataOk() (*Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FileModelResponse) SetData(v Data)`

SetData sets Data field to given value.

### HasData

`func (o *FileModelResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *FileModelResponse) GetMeta() FileMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileModelResponse) GetMetaOk() (*FileMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileModelResponse) SetMeta(v FileMeta)`

SetMeta sets Meta field to given value.


### GetCreatedAt

`func (o *FileModelResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileModelResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileModelResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FileModelResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileModelResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileModelResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


