# ContentExtractionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** |  | [optional] [default to ""]
**TikaServerUrl** | Pointer to **NullableString** |  | [optional] 
**DocumentIntelligenceConfig** | Pointer to [**NullableDocumentIntelligenceConfigForm**](DocumentIntelligenceConfigForm.md) |  | [optional] 

## Methods

### NewContentExtractionConfig

`func NewContentExtractionConfig() *ContentExtractionConfig`

NewContentExtractionConfig instantiates a new ContentExtractionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentExtractionConfigWithDefaults

`func NewContentExtractionConfigWithDefaults() *ContentExtractionConfig`

NewContentExtractionConfigWithDefaults instantiates a new ContentExtractionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *ContentExtractionConfig) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ContentExtractionConfig) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ContentExtractionConfig) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *ContentExtractionConfig) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetTikaServerUrl

`func (o *ContentExtractionConfig) GetTikaServerUrl() string`

GetTikaServerUrl returns the TikaServerUrl field if non-nil, zero value otherwise.

### GetTikaServerUrlOk

`func (o *ContentExtractionConfig) GetTikaServerUrlOk() (*string, bool)`

GetTikaServerUrlOk returns a tuple with the TikaServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikaServerUrl

`func (o *ContentExtractionConfig) SetTikaServerUrl(v string)`

SetTikaServerUrl sets TikaServerUrl field to given value.

### HasTikaServerUrl

`func (o *ContentExtractionConfig) HasTikaServerUrl() bool`

HasTikaServerUrl returns a boolean if a field has been set.

### SetTikaServerUrlNil

`func (o *ContentExtractionConfig) SetTikaServerUrlNil(b bool)`

 SetTikaServerUrlNil sets the value for TikaServerUrl to be an explicit nil

### UnsetTikaServerUrl
`func (o *ContentExtractionConfig) UnsetTikaServerUrl()`

UnsetTikaServerUrl ensures that no value is present for TikaServerUrl, not even an explicit nil
### GetDocumentIntelligenceConfig

`func (o *ContentExtractionConfig) GetDocumentIntelligenceConfig() DocumentIntelligenceConfigForm`

GetDocumentIntelligenceConfig returns the DocumentIntelligenceConfig field if non-nil, zero value otherwise.

### GetDocumentIntelligenceConfigOk

`func (o *ContentExtractionConfig) GetDocumentIntelligenceConfigOk() (*DocumentIntelligenceConfigForm, bool)`

GetDocumentIntelligenceConfigOk returns a tuple with the DocumentIntelligenceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIntelligenceConfig

`func (o *ContentExtractionConfig) SetDocumentIntelligenceConfig(v DocumentIntelligenceConfigForm)`

SetDocumentIntelligenceConfig sets DocumentIntelligenceConfig field to given value.

### HasDocumentIntelligenceConfig

`func (o *ContentExtractionConfig) HasDocumentIntelligenceConfig() bool`

HasDocumentIntelligenceConfig returns a boolean if a field has been set.

### SetDocumentIntelligenceConfigNil

`func (o *ContentExtractionConfig) SetDocumentIntelligenceConfigNil(b bool)`

 SetDocumentIntelligenceConfigNil sets the value for DocumentIntelligenceConfig to be an explicit nil

### UnsetDocumentIntelligenceConfig
`func (o *ContentExtractionConfig) UnsetDocumentIntelligenceConfig()`

UnsetDocumentIntelligenceConfig ensures that no value is present for DocumentIntelligenceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


