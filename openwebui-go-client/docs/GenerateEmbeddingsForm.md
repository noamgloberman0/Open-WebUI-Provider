# GenerateEmbeddingsForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Prompt** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**KeepAlive** | Pointer to [**NullableKeepAlive**](KeepAlive.md) |  | [optional] 

## Methods

### NewGenerateEmbeddingsForm

`func NewGenerateEmbeddingsForm(model string, prompt string, ) *GenerateEmbeddingsForm`

NewGenerateEmbeddingsForm instantiates a new GenerateEmbeddingsForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateEmbeddingsFormWithDefaults

`func NewGenerateEmbeddingsFormWithDefaults() *GenerateEmbeddingsForm`

NewGenerateEmbeddingsFormWithDefaults instantiates a new GenerateEmbeddingsForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GenerateEmbeddingsForm) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GenerateEmbeddingsForm) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GenerateEmbeddingsForm) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *GenerateEmbeddingsForm) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GenerateEmbeddingsForm) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GenerateEmbeddingsForm) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetOptions

`func (o *GenerateEmbeddingsForm) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GenerateEmbeddingsForm) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GenerateEmbeddingsForm) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GenerateEmbeddingsForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GenerateEmbeddingsForm) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GenerateEmbeddingsForm) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetKeepAlive

`func (o *GenerateEmbeddingsForm) GetKeepAlive() KeepAlive`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *GenerateEmbeddingsForm) GetKeepAliveOk() (*KeepAlive, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *GenerateEmbeddingsForm) SetKeepAlive(v KeepAlive)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *GenerateEmbeddingsForm) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### SetKeepAliveNil

`func (o *GenerateEmbeddingsForm) SetKeepAliveNil(b bool)`

 SetKeepAliveNil sets the value for KeepAlive to be an explicit nil

### UnsetKeepAlive
`func (o *GenerateEmbeddingsForm) UnsetKeepAlive()`

UnsetKeepAlive ensures that no value is present for KeepAlive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


