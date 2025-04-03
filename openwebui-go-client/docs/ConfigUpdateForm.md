# ConfigUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RAG_FULL_CONTEXT** | Pointer to **NullableBool** |  | [optional] 
**BYPASS_EMBEDDING_AND_RETRIEVAL** | Pointer to **NullableBool** |  | [optional] 
**PdfExtractImages** | Pointer to **NullableBool** |  | [optional] 
**EnableGoogleDriveIntegration** | Pointer to **NullableBool** |  | [optional] 
**EnableOnedriveIntegration** | Pointer to **NullableBool** |  | [optional] 
**File** | Pointer to [**NullableFileConfig**](FileConfig.md) |  | [optional] 
**ContentExtraction** | Pointer to [**NullableContentExtractionConfig**](ContentExtractionConfig.md) |  | [optional] 
**Chunk** | Pointer to [**NullableChunkParamUpdateForm**](ChunkParamUpdateForm.md) |  | [optional] 
**Youtube** | Pointer to [**NullableYoutubeLoaderConfig**](YoutubeLoaderConfig.md) |  | [optional] 
**Web** | Pointer to [**NullableWebConfig**](WebConfig.md) |  | [optional] 

## Methods

### NewConfigUpdateForm

`func NewConfigUpdateForm() *ConfigUpdateForm`

NewConfigUpdateForm instantiates a new ConfigUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigUpdateFormWithDefaults

`func NewConfigUpdateFormWithDefaults() *ConfigUpdateForm`

NewConfigUpdateFormWithDefaults instantiates a new ConfigUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRAG_FULL_CONTEXT

`func (o *ConfigUpdateForm) GetRAG_FULL_CONTEXT() bool`

GetRAG_FULL_CONTEXT returns the RAG_FULL_CONTEXT field if non-nil, zero value otherwise.

### GetRAG_FULL_CONTEXTOk

`func (o *ConfigUpdateForm) GetRAG_FULL_CONTEXTOk() (*bool, bool)`

GetRAG_FULL_CONTEXTOk returns a tuple with the RAG_FULL_CONTEXT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAG_FULL_CONTEXT

`func (o *ConfigUpdateForm) SetRAG_FULL_CONTEXT(v bool)`

SetRAG_FULL_CONTEXT sets RAG_FULL_CONTEXT field to given value.

### HasRAG_FULL_CONTEXT

`func (o *ConfigUpdateForm) HasRAG_FULL_CONTEXT() bool`

HasRAG_FULL_CONTEXT returns a boolean if a field has been set.

### SetRAG_FULL_CONTEXTNil

`func (o *ConfigUpdateForm) SetRAG_FULL_CONTEXTNil(b bool)`

 SetRAG_FULL_CONTEXTNil sets the value for RAG_FULL_CONTEXT to be an explicit nil

### UnsetRAG_FULL_CONTEXT
`func (o *ConfigUpdateForm) UnsetRAG_FULL_CONTEXT()`

UnsetRAG_FULL_CONTEXT ensures that no value is present for RAG_FULL_CONTEXT, not even an explicit nil
### GetBYPASS_EMBEDDING_AND_RETRIEVAL

`func (o *ConfigUpdateForm) GetBYPASS_EMBEDDING_AND_RETRIEVAL() bool`

GetBYPASS_EMBEDDING_AND_RETRIEVAL returns the BYPASS_EMBEDDING_AND_RETRIEVAL field if non-nil, zero value otherwise.

### GetBYPASS_EMBEDDING_AND_RETRIEVALOk

`func (o *ConfigUpdateForm) GetBYPASS_EMBEDDING_AND_RETRIEVALOk() (*bool, bool)`

GetBYPASS_EMBEDDING_AND_RETRIEVALOk returns a tuple with the BYPASS_EMBEDDING_AND_RETRIEVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBYPASS_EMBEDDING_AND_RETRIEVAL

`func (o *ConfigUpdateForm) SetBYPASS_EMBEDDING_AND_RETRIEVAL(v bool)`

SetBYPASS_EMBEDDING_AND_RETRIEVAL sets BYPASS_EMBEDDING_AND_RETRIEVAL field to given value.

### HasBYPASS_EMBEDDING_AND_RETRIEVAL

`func (o *ConfigUpdateForm) HasBYPASS_EMBEDDING_AND_RETRIEVAL() bool`

HasBYPASS_EMBEDDING_AND_RETRIEVAL returns a boolean if a field has been set.

### SetBYPASS_EMBEDDING_AND_RETRIEVALNil

`func (o *ConfigUpdateForm) SetBYPASS_EMBEDDING_AND_RETRIEVALNil(b bool)`

 SetBYPASS_EMBEDDING_AND_RETRIEVALNil sets the value for BYPASS_EMBEDDING_AND_RETRIEVAL to be an explicit nil

### UnsetBYPASS_EMBEDDING_AND_RETRIEVAL
`func (o *ConfigUpdateForm) UnsetBYPASS_EMBEDDING_AND_RETRIEVAL()`

UnsetBYPASS_EMBEDDING_AND_RETRIEVAL ensures that no value is present for BYPASS_EMBEDDING_AND_RETRIEVAL, not even an explicit nil
### GetPdfExtractImages

`func (o *ConfigUpdateForm) GetPdfExtractImages() bool`

GetPdfExtractImages returns the PdfExtractImages field if non-nil, zero value otherwise.

### GetPdfExtractImagesOk

`func (o *ConfigUpdateForm) GetPdfExtractImagesOk() (*bool, bool)`

GetPdfExtractImagesOk returns a tuple with the PdfExtractImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfExtractImages

`func (o *ConfigUpdateForm) SetPdfExtractImages(v bool)`

SetPdfExtractImages sets PdfExtractImages field to given value.

### HasPdfExtractImages

`func (o *ConfigUpdateForm) HasPdfExtractImages() bool`

HasPdfExtractImages returns a boolean if a field has been set.

### SetPdfExtractImagesNil

`func (o *ConfigUpdateForm) SetPdfExtractImagesNil(b bool)`

 SetPdfExtractImagesNil sets the value for PdfExtractImages to be an explicit nil

### UnsetPdfExtractImages
`func (o *ConfigUpdateForm) UnsetPdfExtractImages()`

UnsetPdfExtractImages ensures that no value is present for PdfExtractImages, not even an explicit nil
### GetEnableGoogleDriveIntegration

`func (o *ConfigUpdateForm) GetEnableGoogleDriveIntegration() bool`

GetEnableGoogleDriveIntegration returns the EnableGoogleDriveIntegration field if non-nil, zero value otherwise.

### GetEnableGoogleDriveIntegrationOk

`func (o *ConfigUpdateForm) GetEnableGoogleDriveIntegrationOk() (*bool, bool)`

GetEnableGoogleDriveIntegrationOk returns a tuple with the EnableGoogleDriveIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGoogleDriveIntegration

`func (o *ConfigUpdateForm) SetEnableGoogleDriveIntegration(v bool)`

SetEnableGoogleDriveIntegration sets EnableGoogleDriveIntegration field to given value.

### HasEnableGoogleDriveIntegration

`func (o *ConfigUpdateForm) HasEnableGoogleDriveIntegration() bool`

HasEnableGoogleDriveIntegration returns a boolean if a field has been set.

### SetEnableGoogleDriveIntegrationNil

`func (o *ConfigUpdateForm) SetEnableGoogleDriveIntegrationNil(b bool)`

 SetEnableGoogleDriveIntegrationNil sets the value for EnableGoogleDriveIntegration to be an explicit nil

### UnsetEnableGoogleDriveIntegration
`func (o *ConfigUpdateForm) UnsetEnableGoogleDriveIntegration()`

UnsetEnableGoogleDriveIntegration ensures that no value is present for EnableGoogleDriveIntegration, not even an explicit nil
### GetEnableOnedriveIntegration

`func (o *ConfigUpdateForm) GetEnableOnedriveIntegration() bool`

GetEnableOnedriveIntegration returns the EnableOnedriveIntegration field if non-nil, zero value otherwise.

### GetEnableOnedriveIntegrationOk

`func (o *ConfigUpdateForm) GetEnableOnedriveIntegrationOk() (*bool, bool)`

GetEnableOnedriveIntegrationOk returns a tuple with the EnableOnedriveIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnedriveIntegration

`func (o *ConfigUpdateForm) SetEnableOnedriveIntegration(v bool)`

SetEnableOnedriveIntegration sets EnableOnedriveIntegration field to given value.

### HasEnableOnedriveIntegration

`func (o *ConfigUpdateForm) HasEnableOnedriveIntegration() bool`

HasEnableOnedriveIntegration returns a boolean if a field has been set.

### SetEnableOnedriveIntegrationNil

`func (o *ConfigUpdateForm) SetEnableOnedriveIntegrationNil(b bool)`

 SetEnableOnedriveIntegrationNil sets the value for EnableOnedriveIntegration to be an explicit nil

### UnsetEnableOnedriveIntegration
`func (o *ConfigUpdateForm) UnsetEnableOnedriveIntegration()`

UnsetEnableOnedriveIntegration ensures that no value is present for EnableOnedriveIntegration, not even an explicit nil
### GetFile

`func (o *ConfigUpdateForm) GetFile() FileConfig`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ConfigUpdateForm) GetFileOk() (*FileConfig, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ConfigUpdateForm) SetFile(v FileConfig)`

SetFile sets File field to given value.

### HasFile

`func (o *ConfigUpdateForm) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *ConfigUpdateForm) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *ConfigUpdateForm) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetContentExtraction

`func (o *ConfigUpdateForm) GetContentExtraction() ContentExtractionConfig`

GetContentExtraction returns the ContentExtraction field if non-nil, zero value otherwise.

### GetContentExtractionOk

`func (o *ConfigUpdateForm) GetContentExtractionOk() (*ContentExtractionConfig, bool)`

GetContentExtractionOk returns a tuple with the ContentExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtraction

`func (o *ConfigUpdateForm) SetContentExtraction(v ContentExtractionConfig)`

SetContentExtraction sets ContentExtraction field to given value.

### HasContentExtraction

`func (o *ConfigUpdateForm) HasContentExtraction() bool`

HasContentExtraction returns a boolean if a field has been set.

### SetContentExtractionNil

`func (o *ConfigUpdateForm) SetContentExtractionNil(b bool)`

 SetContentExtractionNil sets the value for ContentExtraction to be an explicit nil

### UnsetContentExtraction
`func (o *ConfigUpdateForm) UnsetContentExtraction()`

UnsetContentExtraction ensures that no value is present for ContentExtraction, not even an explicit nil
### GetChunk

`func (o *ConfigUpdateForm) GetChunk() ChunkParamUpdateForm`

GetChunk returns the Chunk field if non-nil, zero value otherwise.

### GetChunkOk

`func (o *ConfigUpdateForm) GetChunkOk() (*ChunkParamUpdateForm, bool)`

GetChunkOk returns a tuple with the Chunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunk

`func (o *ConfigUpdateForm) SetChunk(v ChunkParamUpdateForm)`

SetChunk sets Chunk field to given value.

### HasChunk

`func (o *ConfigUpdateForm) HasChunk() bool`

HasChunk returns a boolean if a field has been set.

### SetChunkNil

`func (o *ConfigUpdateForm) SetChunkNil(b bool)`

 SetChunkNil sets the value for Chunk to be an explicit nil

### UnsetChunk
`func (o *ConfigUpdateForm) UnsetChunk()`

UnsetChunk ensures that no value is present for Chunk, not even an explicit nil
### GetYoutube

`func (o *ConfigUpdateForm) GetYoutube() YoutubeLoaderConfig`

GetYoutube returns the Youtube field if non-nil, zero value otherwise.

### GetYoutubeOk

`func (o *ConfigUpdateForm) GetYoutubeOk() (*YoutubeLoaderConfig, bool)`

GetYoutubeOk returns a tuple with the Youtube field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutube

`func (o *ConfigUpdateForm) SetYoutube(v YoutubeLoaderConfig)`

SetYoutube sets Youtube field to given value.

### HasYoutube

`func (o *ConfigUpdateForm) HasYoutube() bool`

HasYoutube returns a boolean if a field has been set.

### SetYoutubeNil

`func (o *ConfigUpdateForm) SetYoutubeNil(b bool)`

 SetYoutubeNil sets the value for Youtube to be an explicit nil

### UnsetYoutube
`func (o *ConfigUpdateForm) UnsetYoutube()`

UnsetYoutube ensures that no value is present for Youtube, not even an explicit nil
### GetWeb

`func (o *ConfigUpdateForm) GetWeb() WebConfig`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *ConfigUpdateForm) GetWebOk() (*WebConfig, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *ConfigUpdateForm) SetWeb(v WebConfig)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *ConfigUpdateForm) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### SetWebNil

`func (o *ConfigUpdateForm) SetWebNil(b bool)`

 SetWebNil sets the value for Web to be an explicit nil

### UnsetWeb
`func (o *ConfigUpdateForm) UnsetWeb()`

UnsetWeb ensures that no value is present for Web, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


