# ConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Engine** | **string** |  | 
**PromptGeneration** | **bool** |  | 
**Openai** | [**OpenWebuiRoutersImagesOpenAIConfigForm**](OpenWebuiRoutersImagesOpenAIConfigForm.md) |  | 
**Automatic1111** | [**Automatic1111ConfigForm**](Automatic1111ConfigForm.md) |  | 
**Comfyui** | [**ComfyUIConfigForm**](ComfyUIConfigForm.md) |  | 
**Gemini** | [**GeminiConfigForm**](GeminiConfigForm.md) |  | 

## Methods

### NewConfigForm

`func NewConfigForm(enabled bool, engine string, promptGeneration bool, openai OpenWebuiRoutersImagesOpenAIConfigForm, automatic1111 Automatic1111ConfigForm, comfyui ComfyUIConfigForm, gemini GeminiConfigForm, ) *ConfigForm`

NewConfigForm instantiates a new ConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFormWithDefaults

`func NewConfigFormWithDefaults() *ConfigForm`

NewConfigFormWithDefaults instantiates a new ConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ConfigForm) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigForm) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigForm) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEngine

`func (o *ConfigForm) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ConfigForm) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ConfigForm) SetEngine(v string)`

SetEngine sets Engine field to given value.


### GetPromptGeneration

`func (o *ConfigForm) GetPromptGeneration() bool`

GetPromptGeneration returns the PromptGeneration field if non-nil, zero value otherwise.

### GetPromptGenerationOk

`func (o *ConfigForm) GetPromptGenerationOk() (*bool, bool)`

GetPromptGenerationOk returns a tuple with the PromptGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptGeneration

`func (o *ConfigForm) SetPromptGeneration(v bool)`

SetPromptGeneration sets PromptGeneration field to given value.


### GetOpenai

`func (o *ConfigForm) GetOpenai() OpenWebuiRoutersImagesOpenAIConfigForm`

GetOpenai returns the Openai field if non-nil, zero value otherwise.

### GetOpenaiOk

`func (o *ConfigForm) GetOpenaiOk() (*OpenWebuiRoutersImagesOpenAIConfigForm, bool)`

GetOpenaiOk returns a tuple with the Openai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenai

`func (o *ConfigForm) SetOpenai(v OpenWebuiRoutersImagesOpenAIConfigForm)`

SetOpenai sets Openai field to given value.


### GetAutomatic1111

`func (o *ConfigForm) GetAutomatic1111() Automatic1111ConfigForm`

GetAutomatic1111 returns the Automatic1111 field if non-nil, zero value otherwise.

### GetAutomatic1111Ok

`func (o *ConfigForm) GetAutomatic1111Ok() (*Automatic1111ConfigForm, bool)`

GetAutomatic1111Ok returns a tuple with the Automatic1111 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic1111

`func (o *ConfigForm) SetAutomatic1111(v Automatic1111ConfigForm)`

SetAutomatic1111 sets Automatic1111 field to given value.


### GetComfyui

`func (o *ConfigForm) GetComfyui() ComfyUIConfigForm`

GetComfyui returns the Comfyui field if non-nil, zero value otherwise.

### GetComfyuiOk

`func (o *ConfigForm) GetComfyuiOk() (*ComfyUIConfigForm, bool)`

GetComfyuiOk returns a tuple with the Comfyui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComfyui

`func (o *ConfigForm) SetComfyui(v ComfyUIConfigForm)`

SetComfyui sets Comfyui field to given value.


### GetGemini

`func (o *ConfigForm) GetGemini() GeminiConfigForm`

GetGemini returns the Gemini field if non-nil, zero value otherwise.

### GetGeminiOk

`func (o *ConfigForm) GetGeminiOk() (*GeminiConfigForm, bool)`

GetGeminiOk returns a tuple with the Gemini field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemini

`func (o *ConfigForm) SetGemini(v GeminiConfigForm)`

SetGemini sets Gemini field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


