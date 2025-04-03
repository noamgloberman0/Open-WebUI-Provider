# GenerateEmbedForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | [**Input**](Input.md) |  | 
**Truncate** | Pointer to **NullableBool** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**KeepAlive** | Pointer to [**NullableKeepAlive**](KeepAlive.md) |  | [optional] 

## Methods

### NewGenerateEmbedForm

`func NewGenerateEmbedForm(model string, input Input, ) *GenerateEmbedForm`

NewGenerateEmbedForm instantiates a new GenerateEmbedForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateEmbedFormWithDefaults

`func NewGenerateEmbedFormWithDefaults() *GenerateEmbedForm`

NewGenerateEmbedFormWithDefaults instantiates a new GenerateEmbedForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GenerateEmbedForm) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GenerateEmbedForm) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GenerateEmbedForm) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *GenerateEmbedForm) GetInput() Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GenerateEmbedForm) GetInputOk() (*Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GenerateEmbedForm) SetInput(v Input)`

SetInput sets Input field to given value.


### GetTruncate

`func (o *GenerateEmbedForm) GetTruncate() bool`

GetTruncate returns the Truncate field if non-nil, zero value otherwise.

### GetTruncateOk

`func (o *GenerateEmbedForm) GetTruncateOk() (*bool, bool)`

GetTruncateOk returns a tuple with the Truncate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncate

`func (o *GenerateEmbedForm) SetTruncate(v bool)`

SetTruncate sets Truncate field to given value.

### HasTruncate

`func (o *GenerateEmbedForm) HasTruncate() bool`

HasTruncate returns a boolean if a field has been set.

### SetTruncateNil

`func (o *GenerateEmbedForm) SetTruncateNil(b bool)`

 SetTruncateNil sets the value for Truncate to be an explicit nil

### UnsetTruncate
`func (o *GenerateEmbedForm) UnsetTruncate()`

UnsetTruncate ensures that no value is present for Truncate, not even an explicit nil
### GetOptions

`func (o *GenerateEmbedForm) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GenerateEmbedForm) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GenerateEmbedForm) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GenerateEmbedForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GenerateEmbedForm) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GenerateEmbedForm) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetKeepAlive

`func (o *GenerateEmbedForm) GetKeepAlive() KeepAlive`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *GenerateEmbedForm) GetKeepAliveOk() (*KeepAlive, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *GenerateEmbedForm) SetKeepAlive(v KeepAlive)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *GenerateEmbedForm) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### SetKeepAliveNil

`func (o *GenerateEmbedForm) SetKeepAliveNil(b bool)`

 SetKeepAliveNil sets the value for KeepAlive to be an explicit nil

### UnsetKeepAlive
`func (o *GenerateEmbedForm) UnsetKeepAlive()`

UnsetKeepAlive ensures that no value is present for KeepAlive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


