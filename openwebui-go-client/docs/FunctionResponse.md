# FunctionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Meta** | [**FunctionMeta**](FunctionMeta.md) |  | 
**IsActive** | **bool** |  | 
**IsGlobal** | **bool** |  | 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 

## Methods

### NewFunctionResponse

`func NewFunctionResponse(id string, userId string, type_ string, name string, meta FunctionMeta, isActive bool, isGlobal bool, updatedAt int32, createdAt int32, ) *FunctionResponse`

NewFunctionResponse instantiates a new FunctionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionResponseWithDefaults

`func NewFunctionResponseWithDefaults() *FunctionResponse`

NewFunctionResponseWithDefaults instantiates a new FunctionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *FunctionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FunctionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FunctionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetType

`func (o *FunctionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *FunctionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *FunctionResponse) GetMeta() FunctionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FunctionResponse) GetMetaOk() (*FunctionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FunctionResponse) SetMeta(v FunctionMeta)`

SetMeta sets Meta field to given value.


### GetIsActive

`func (o *FunctionResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *FunctionResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *FunctionResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsGlobal

`func (o *FunctionResponse) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *FunctionResponse) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *FunctionResponse) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.


### GetUpdatedAt

`func (o *FunctionResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FunctionResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FunctionResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *FunctionResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


