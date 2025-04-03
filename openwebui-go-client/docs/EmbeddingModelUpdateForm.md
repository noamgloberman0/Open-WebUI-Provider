# EmbeddingModelUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenaiConfig** | Pointer to [**NullableOpenWebuiRoutersRetrievalOpenAIConfigForm**](OpenWebuiRoutersRetrievalOpenAIConfigForm.md) |  | [optional] 
**OllamaConfig** | Pointer to [**NullableOpenWebuiRoutersRetrievalOllamaConfigForm**](OpenWebuiRoutersRetrievalOllamaConfigForm.md) |  | [optional] 
**EmbeddingEngine** | **string** |  | 
**EmbeddingModel** | **string** |  | 
**EmbeddingBatchSize** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEmbeddingModelUpdateForm

`func NewEmbeddingModelUpdateForm(embeddingEngine string, embeddingModel string, ) *EmbeddingModelUpdateForm`

NewEmbeddingModelUpdateForm instantiates a new EmbeddingModelUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingModelUpdateFormWithDefaults

`func NewEmbeddingModelUpdateFormWithDefaults() *EmbeddingModelUpdateForm`

NewEmbeddingModelUpdateFormWithDefaults instantiates a new EmbeddingModelUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenaiConfig

`func (o *EmbeddingModelUpdateForm) GetOpenaiConfig() OpenWebuiRoutersRetrievalOpenAIConfigForm`

GetOpenaiConfig returns the OpenaiConfig field if non-nil, zero value otherwise.

### GetOpenaiConfigOk

`func (o *EmbeddingModelUpdateForm) GetOpenaiConfigOk() (*OpenWebuiRoutersRetrievalOpenAIConfigForm, bool)`

GetOpenaiConfigOk returns a tuple with the OpenaiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiConfig

`func (o *EmbeddingModelUpdateForm) SetOpenaiConfig(v OpenWebuiRoutersRetrievalOpenAIConfigForm)`

SetOpenaiConfig sets OpenaiConfig field to given value.

### HasOpenaiConfig

`func (o *EmbeddingModelUpdateForm) HasOpenaiConfig() bool`

HasOpenaiConfig returns a boolean if a field has been set.

### SetOpenaiConfigNil

`func (o *EmbeddingModelUpdateForm) SetOpenaiConfigNil(b bool)`

 SetOpenaiConfigNil sets the value for OpenaiConfig to be an explicit nil

### UnsetOpenaiConfig
`func (o *EmbeddingModelUpdateForm) UnsetOpenaiConfig()`

UnsetOpenaiConfig ensures that no value is present for OpenaiConfig, not even an explicit nil
### GetOllamaConfig

`func (o *EmbeddingModelUpdateForm) GetOllamaConfig() OpenWebuiRoutersRetrievalOllamaConfigForm`

GetOllamaConfig returns the OllamaConfig field if non-nil, zero value otherwise.

### GetOllamaConfigOk

`func (o *EmbeddingModelUpdateForm) GetOllamaConfigOk() (*OpenWebuiRoutersRetrievalOllamaConfigForm, bool)`

GetOllamaConfigOk returns a tuple with the OllamaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOllamaConfig

`func (o *EmbeddingModelUpdateForm) SetOllamaConfig(v OpenWebuiRoutersRetrievalOllamaConfigForm)`

SetOllamaConfig sets OllamaConfig field to given value.

### HasOllamaConfig

`func (o *EmbeddingModelUpdateForm) HasOllamaConfig() bool`

HasOllamaConfig returns a boolean if a field has been set.

### SetOllamaConfigNil

`func (o *EmbeddingModelUpdateForm) SetOllamaConfigNil(b bool)`

 SetOllamaConfigNil sets the value for OllamaConfig to be an explicit nil

### UnsetOllamaConfig
`func (o *EmbeddingModelUpdateForm) UnsetOllamaConfig()`

UnsetOllamaConfig ensures that no value is present for OllamaConfig, not even an explicit nil
### GetEmbeddingEngine

`func (o *EmbeddingModelUpdateForm) GetEmbeddingEngine() string`

GetEmbeddingEngine returns the EmbeddingEngine field if non-nil, zero value otherwise.

### GetEmbeddingEngineOk

`func (o *EmbeddingModelUpdateForm) GetEmbeddingEngineOk() (*string, bool)`

GetEmbeddingEngineOk returns a tuple with the EmbeddingEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingEngine

`func (o *EmbeddingModelUpdateForm) SetEmbeddingEngine(v string)`

SetEmbeddingEngine sets EmbeddingEngine field to given value.


### GetEmbeddingModel

`func (o *EmbeddingModelUpdateForm) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *EmbeddingModelUpdateForm) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *EmbeddingModelUpdateForm) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### GetEmbeddingBatchSize

`func (o *EmbeddingModelUpdateForm) GetEmbeddingBatchSize() int32`

GetEmbeddingBatchSize returns the EmbeddingBatchSize field if non-nil, zero value otherwise.

### GetEmbeddingBatchSizeOk

`func (o *EmbeddingModelUpdateForm) GetEmbeddingBatchSizeOk() (*int32, bool)`

GetEmbeddingBatchSizeOk returns a tuple with the EmbeddingBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingBatchSize

`func (o *EmbeddingModelUpdateForm) SetEmbeddingBatchSize(v int32)`

SetEmbeddingBatchSize sets EmbeddingBatchSize field to given value.

### HasEmbeddingBatchSize

`func (o *EmbeddingModelUpdateForm) HasEmbeddingBatchSize() bool`

HasEmbeddingBatchSize returns a boolean if a field has been set.

### SetEmbeddingBatchSizeNil

`func (o *EmbeddingModelUpdateForm) SetEmbeddingBatchSizeNil(b bool)`

 SetEmbeddingBatchSizeNil sets the value for EmbeddingBatchSize to be an explicit nil

### UnsetEmbeddingBatchSize
`func (o *EmbeddingModelUpdateForm) UnsetEmbeddingBatchSize()`

UnsetEmbeddingBatchSize ensures that no value is present for EmbeddingBatchSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


