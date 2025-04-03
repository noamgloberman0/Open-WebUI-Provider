# WebConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | [**WebSearchConfig**](WebSearchConfig.md) |  | 
**ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION** | Pointer to **NullableBool** |  | [optional] 
**BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWebConfig

`func NewWebConfig(search WebSearchConfig, ) *WebConfig`

NewWebConfig instantiates a new WebConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebConfigWithDefaults

`func NewWebConfigWithDefaults() *WebConfig`

NewWebConfigWithDefaults instantiates a new WebConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *WebConfig) GetSearch() WebSearchConfig`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *WebConfig) GetSearchOk() (*WebSearchConfig, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *WebConfig) SetSearch(v WebSearchConfig)`

SetSearch sets Search field to given value.


### GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION

`func (o *WebConfig) GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION() bool`

GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION returns the ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field if non-nil, zero value otherwise.

### GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONOk

`func (o *WebConfig) GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONOk() (*bool, bool)`

GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONOk returns a tuple with the ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION

`func (o *WebConfig) SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION(v bool)`

SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION sets ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field to given value.

### HasENABLE_RAG_WEB_LOADER_SSL_VERIFICATION

`func (o *WebConfig) HasENABLE_RAG_WEB_LOADER_SSL_VERIFICATION() bool`

HasENABLE_RAG_WEB_LOADER_SSL_VERIFICATION returns a boolean if a field has been set.

### SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONNil

`func (o *WebConfig) SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONNil(b bool)`

 SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONNil sets the value for ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION to be an explicit nil

### UnsetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION
`func (o *WebConfig) UnsetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION()`

UnsetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION ensures that no value is present for ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION, not even an explicit nil
### GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL

`func (o *WebConfig) GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL() bool`

GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL returns the BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field if non-nil, zero value otherwise.

### GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALOk

`func (o *WebConfig) GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALOk() (*bool, bool)`

GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALOk returns a tuple with the BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL

`func (o *WebConfig) SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL(v bool)`

SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL sets BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field to given value.

### HasBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL

`func (o *WebConfig) HasBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL() bool`

HasBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL returns a boolean if a field has been set.

### SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALNil

`func (o *WebConfig) SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALNil(b bool)`

 SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALNil sets the value for BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL to be an explicit nil

### UnsetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL
`func (o *WebConfig) UnsetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL()`

UnsetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL ensures that no value is present for BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


