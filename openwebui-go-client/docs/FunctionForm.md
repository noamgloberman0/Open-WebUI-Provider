# FunctionForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Content** | **string** |  | 
**Meta** | [**FunctionMeta**](FunctionMeta.md) |  | 

## Methods

### NewFunctionForm

`func NewFunctionForm(id string, name string, content string, meta FunctionMeta, ) *FunctionForm`

NewFunctionForm instantiates a new FunctionForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionFormWithDefaults

`func NewFunctionFormWithDefaults() *FunctionForm`

NewFunctionFormWithDefaults instantiates a new FunctionForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionForm) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FunctionForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionForm) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *FunctionForm) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FunctionForm) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FunctionForm) SetContent(v string)`

SetContent sets Content field to given value.


### GetMeta

`func (o *FunctionForm) GetMeta() FunctionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FunctionForm) GetMetaOk() (*FunctionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FunctionForm) SetMeta(v FunctionMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


