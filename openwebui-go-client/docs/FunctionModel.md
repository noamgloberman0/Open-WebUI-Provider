# FunctionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Content** | **string** |  | 
**Meta** | [**FunctionMeta**](FunctionMeta.md) |  | 
**IsActive** | Pointer to **bool** |  | [optional] [default to false]
**IsGlobal** | Pointer to **bool** |  | [optional] [default to false]
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 

## Methods

### NewFunctionModel

`func NewFunctionModel(id string, userId string, name string, type_ string, content string, meta FunctionMeta, updatedAt int32, createdAt int32, ) *FunctionModel`

NewFunctionModel instantiates a new FunctionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionModelWithDefaults

`func NewFunctionModelWithDefaults() *FunctionModel`

NewFunctionModelWithDefaults instantiates a new FunctionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionModel) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *FunctionModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FunctionModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FunctionModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *FunctionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FunctionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionModel) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *FunctionModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FunctionModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FunctionModel) SetContent(v string)`

SetContent sets Content field to given value.


### GetMeta

`func (o *FunctionModel) GetMeta() FunctionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FunctionModel) GetMetaOk() (*FunctionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FunctionModel) SetMeta(v FunctionMeta)`

SetMeta sets Meta field to given value.


### GetIsActive

`func (o *FunctionModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *FunctionModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *FunctionModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *FunctionModel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsGlobal

`func (o *FunctionModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *FunctionModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *FunctionModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *FunctionModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FunctionModel) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FunctionModel) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FunctionModel) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *FunctionModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


