# GenerateImageForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **NullableString** |  | [optional] 
**Prompt** | **string** |  | 
**Size** | Pointer to **NullableString** |  | [optional] 
**N** | Pointer to **int32** |  | [optional] [default to 1]
**NegativePrompt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGenerateImageForm

`func NewGenerateImageForm(prompt string, ) *GenerateImageForm`

NewGenerateImageForm instantiates a new GenerateImageForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateImageFormWithDefaults

`func NewGenerateImageFormWithDefaults() *GenerateImageForm`

NewGenerateImageFormWithDefaults instantiates a new GenerateImageForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GenerateImageForm) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GenerateImageForm) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GenerateImageForm) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GenerateImageForm) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *GenerateImageForm) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *GenerateImageForm) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetPrompt

`func (o *GenerateImageForm) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GenerateImageForm) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GenerateImageForm) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSize

`func (o *GenerateImageForm) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GenerateImageForm) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GenerateImageForm) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GenerateImageForm) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GenerateImageForm) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GenerateImageForm) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetN

`func (o *GenerateImageForm) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *GenerateImageForm) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *GenerateImageForm) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *GenerateImageForm) HasN() bool`

HasN returns a boolean if a field has been set.

### GetNegativePrompt

`func (o *GenerateImageForm) GetNegativePrompt() string`

GetNegativePrompt returns the NegativePrompt field if non-nil, zero value otherwise.

### GetNegativePromptOk

`func (o *GenerateImageForm) GetNegativePromptOk() (*string, bool)`

GetNegativePromptOk returns a tuple with the NegativePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativePrompt

`func (o *GenerateImageForm) SetNegativePrompt(v string)`

SetNegativePrompt sets NegativePrompt field to given value.

### HasNegativePrompt

`func (o *GenerateImageForm) HasNegativePrompt() bool`

HasNegativePrompt returns a boolean if a field has been set.

### SetNegativePromptNil

`func (o *GenerateImageForm) SetNegativePromptNil(b bool)`

 SetNegativePromptNil sets the value for NegativePrompt to be an explicit nil

### UnsetNegativePrompt
`func (o *GenerateImageForm) UnsetNegativePrompt()`

UnsetNegativePrompt ensures that no value is present for NegativePrompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


