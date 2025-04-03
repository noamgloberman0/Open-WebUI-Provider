# ModelForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BaseModelId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Meta** | [**ModelMeta**](ModelMeta.md) |  | 
**Params** | **map[string]interface{}** |  | 
**AccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewModelForm

`func NewModelForm(id string, name string, meta ModelMeta, params map[string]interface{}, ) *ModelForm`

NewModelForm instantiates a new ModelForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFormWithDefaults

`func NewModelFormWithDefaults() *ModelForm`

NewModelFormWithDefaults instantiates a new ModelForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelForm) SetId(v string)`

SetId sets Id field to given value.


### GetBaseModelId

`func (o *ModelForm) GetBaseModelId() string`

GetBaseModelId returns the BaseModelId field if non-nil, zero value otherwise.

### GetBaseModelIdOk

`func (o *ModelForm) GetBaseModelIdOk() (*string, bool)`

GetBaseModelIdOk returns a tuple with the BaseModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModelId

`func (o *ModelForm) SetBaseModelId(v string)`

SetBaseModelId sets BaseModelId field to given value.

### HasBaseModelId

`func (o *ModelForm) HasBaseModelId() bool`

HasBaseModelId returns a boolean if a field has been set.

### SetBaseModelIdNil

`func (o *ModelForm) SetBaseModelIdNil(b bool)`

 SetBaseModelIdNil sets the value for BaseModelId to be an explicit nil

### UnsetBaseModelId
`func (o *ModelForm) UnsetBaseModelId()`

UnsetBaseModelId ensures that no value is present for BaseModelId, not even an explicit nil
### GetName

`func (o *ModelForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelForm) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *ModelForm) GetMeta() ModelMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ModelForm) GetMetaOk() (*ModelMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ModelForm) SetMeta(v ModelMeta)`

SetMeta sets Meta field to given value.


### GetParams

`func (o *ModelForm) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelForm) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelForm) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetAccessControl

`func (o *ModelForm) GetAccessControl() map[string]interface{}`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ModelForm) GetAccessControlOk() (*map[string]interface{}, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ModelForm) SetAccessControl(v map[string]interface{})`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ModelForm) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### SetAccessControlNil

`func (o *ModelForm) SetAccessControlNil(b bool)`

 SetAccessControlNil sets the value for AccessControl to be an explicit nil

### UnsetAccessControl
`func (o *ModelForm) UnsetAccessControl()`

UnsetAccessControl ensures that no value is present for AccessControl, not even an explicit nil
### GetIsActive

`func (o *ModelForm) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelForm) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelForm) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelForm) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


