# GenerateCompletionForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Prompt** | **string** |  | 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**System** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Context** | Pointer to **[]int32** |  | [optional] 
**Stream** | Pointer to **NullableBool** |  | [optional] 
**Raw** | Pointer to **NullableBool** |  | [optional] 
**KeepAlive** | Pointer to [**NullableKeepAlive**](KeepAlive.md) |  | [optional] 

## Methods

### NewGenerateCompletionForm

`func NewGenerateCompletionForm(model string, prompt string, ) *GenerateCompletionForm`

NewGenerateCompletionForm instantiates a new GenerateCompletionForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCompletionFormWithDefaults

`func NewGenerateCompletionFormWithDefaults() *GenerateCompletionForm`

NewGenerateCompletionFormWithDefaults instantiates a new GenerateCompletionForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GenerateCompletionForm) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GenerateCompletionForm) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GenerateCompletionForm) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *GenerateCompletionForm) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GenerateCompletionForm) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GenerateCompletionForm) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSuffix

`func (o *GenerateCompletionForm) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *GenerateCompletionForm) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *GenerateCompletionForm) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *GenerateCompletionForm) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *GenerateCompletionForm) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *GenerateCompletionForm) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetImages

`func (o *GenerateCompletionForm) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GenerateCompletionForm) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GenerateCompletionForm) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *GenerateCompletionForm) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *GenerateCompletionForm) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *GenerateCompletionForm) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetFormat

`func (o *GenerateCompletionForm) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateCompletionForm) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateCompletionForm) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GenerateCompletionForm) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *GenerateCompletionForm) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *GenerateCompletionForm) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetOptions

`func (o *GenerateCompletionForm) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GenerateCompletionForm) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GenerateCompletionForm) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GenerateCompletionForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GenerateCompletionForm) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GenerateCompletionForm) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetSystem

`func (o *GenerateCompletionForm) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *GenerateCompletionForm) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *GenerateCompletionForm) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *GenerateCompletionForm) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *GenerateCompletionForm) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *GenerateCompletionForm) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTemplate

`func (o *GenerateCompletionForm) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateCompletionForm) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateCompletionForm) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GenerateCompletionForm) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *GenerateCompletionForm) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *GenerateCompletionForm) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetContext

`func (o *GenerateCompletionForm) GetContext() []int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GenerateCompletionForm) GetContextOk() (*[]int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GenerateCompletionForm) SetContext(v []int32)`

SetContext sets Context field to given value.

### HasContext

`func (o *GenerateCompletionForm) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *GenerateCompletionForm) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *GenerateCompletionForm) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetStream

`func (o *GenerateCompletionForm) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *GenerateCompletionForm) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *GenerateCompletionForm) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *GenerateCompletionForm) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *GenerateCompletionForm) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *GenerateCompletionForm) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil
### GetRaw

`func (o *GenerateCompletionForm) GetRaw() bool`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *GenerateCompletionForm) GetRawOk() (*bool, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *GenerateCompletionForm) SetRaw(v bool)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *GenerateCompletionForm) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *GenerateCompletionForm) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *GenerateCompletionForm) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetKeepAlive

`func (o *GenerateCompletionForm) GetKeepAlive() KeepAlive`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *GenerateCompletionForm) GetKeepAliveOk() (*KeepAlive, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *GenerateCompletionForm) SetKeepAlive(v KeepAlive)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *GenerateCompletionForm) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### SetKeepAliveNil

`func (o *GenerateCompletionForm) SetKeepAliveNil(b bool)`

 SetKeepAliveNil sets the value for KeepAlive to be an explicit nil

### UnsetKeepAlive
`func (o *GenerateCompletionForm) UnsetKeepAlive()`

UnsetKeepAlive ensures that no value is present for KeepAlive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


