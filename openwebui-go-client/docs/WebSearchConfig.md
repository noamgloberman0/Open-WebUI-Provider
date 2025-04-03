# WebSearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Engine** | Pointer to **NullableString** |  | [optional] 
**SearxngQueryUrl** | Pointer to **NullableString** |  | [optional] 
**GooglePseApiKey** | Pointer to **NullableString** |  | [optional] 
**GooglePseEngineId** | Pointer to **NullableString** |  | [optional] 
**BraveSearchApiKey** | Pointer to **NullableString** |  | [optional] 
**KagiSearchApiKey** | Pointer to **NullableString** |  | [optional] 
**MojeekSearchApiKey** | Pointer to **NullableString** |  | [optional] 
**BochaSearchApiKey** | Pointer to **NullableString** |  | [optional] 
**SerpstackApiKey** | Pointer to **NullableString** |  | [optional] 
**SerpstackHttps** | Pointer to **NullableBool** |  | [optional] 
**SerperApiKey** | Pointer to **NullableString** |  | [optional] 
**SerplyApiKey** | Pointer to **NullableString** |  | [optional] 
**TavilyApiKey** | Pointer to **NullableString** |  | [optional] 
**SearchapiApiKey** | Pointer to **NullableString** |  | [optional] 
**SearchapiEngine** | Pointer to **NullableString** |  | [optional] 
**SerpapiApiKey** | Pointer to **NullableString** |  | [optional] 
**SerpapiEngine** | Pointer to **NullableString** |  | [optional] 
**JinaApiKey** | Pointer to **NullableString** |  | [optional] 
**BingSearchV7Endpoint** | Pointer to **NullableString** |  | [optional] 
**BingSearchV7SubscriptionKey** | Pointer to **NullableString** |  | [optional] 
**ExaApiKey** | Pointer to **NullableString** |  | [optional] 
**PerplexityApiKey** | Pointer to **NullableString** |  | [optional] 
**ResultCount** | Pointer to **NullableInt32** |  | [optional] 
**ConcurrentRequests** | Pointer to **NullableInt32** |  | [optional] 
**TrustEnv** | Pointer to **NullableBool** |  | [optional] 
**DomainFilterList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebSearchConfig

`func NewWebSearchConfig(enabled bool, ) *WebSearchConfig`

NewWebSearchConfig instantiates a new WebSearchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSearchConfigWithDefaults

`func NewWebSearchConfigWithDefaults() *WebSearchConfig`

NewWebSearchConfigWithDefaults instantiates a new WebSearchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebSearchConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebSearchConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebSearchConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEngine

`func (o *WebSearchConfig) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *WebSearchConfig) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *WebSearchConfig) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *WebSearchConfig) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### SetEngineNil

`func (o *WebSearchConfig) SetEngineNil(b bool)`

 SetEngineNil sets the value for Engine to be an explicit nil

### UnsetEngine
`func (o *WebSearchConfig) UnsetEngine()`

UnsetEngine ensures that no value is present for Engine, not even an explicit nil
### GetSearxngQueryUrl

`func (o *WebSearchConfig) GetSearxngQueryUrl() string`

GetSearxngQueryUrl returns the SearxngQueryUrl field if non-nil, zero value otherwise.

### GetSearxngQueryUrlOk

`func (o *WebSearchConfig) GetSearxngQueryUrlOk() (*string, bool)`

GetSearxngQueryUrlOk returns a tuple with the SearxngQueryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearxngQueryUrl

`func (o *WebSearchConfig) SetSearxngQueryUrl(v string)`

SetSearxngQueryUrl sets SearxngQueryUrl field to given value.

### HasSearxngQueryUrl

`func (o *WebSearchConfig) HasSearxngQueryUrl() bool`

HasSearxngQueryUrl returns a boolean if a field has been set.

### SetSearxngQueryUrlNil

`func (o *WebSearchConfig) SetSearxngQueryUrlNil(b bool)`

 SetSearxngQueryUrlNil sets the value for SearxngQueryUrl to be an explicit nil

### UnsetSearxngQueryUrl
`func (o *WebSearchConfig) UnsetSearxngQueryUrl()`

UnsetSearxngQueryUrl ensures that no value is present for SearxngQueryUrl, not even an explicit nil
### GetGooglePseApiKey

`func (o *WebSearchConfig) GetGooglePseApiKey() string`

GetGooglePseApiKey returns the GooglePseApiKey field if non-nil, zero value otherwise.

### GetGooglePseApiKeyOk

`func (o *WebSearchConfig) GetGooglePseApiKeyOk() (*string, bool)`

GetGooglePseApiKeyOk returns a tuple with the GooglePseApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePseApiKey

`func (o *WebSearchConfig) SetGooglePseApiKey(v string)`

SetGooglePseApiKey sets GooglePseApiKey field to given value.

### HasGooglePseApiKey

`func (o *WebSearchConfig) HasGooglePseApiKey() bool`

HasGooglePseApiKey returns a boolean if a field has been set.

### SetGooglePseApiKeyNil

`func (o *WebSearchConfig) SetGooglePseApiKeyNil(b bool)`

 SetGooglePseApiKeyNil sets the value for GooglePseApiKey to be an explicit nil

### UnsetGooglePseApiKey
`func (o *WebSearchConfig) UnsetGooglePseApiKey()`

UnsetGooglePseApiKey ensures that no value is present for GooglePseApiKey, not even an explicit nil
### GetGooglePseEngineId

`func (o *WebSearchConfig) GetGooglePseEngineId() string`

GetGooglePseEngineId returns the GooglePseEngineId field if non-nil, zero value otherwise.

### GetGooglePseEngineIdOk

`func (o *WebSearchConfig) GetGooglePseEngineIdOk() (*string, bool)`

GetGooglePseEngineIdOk returns a tuple with the GooglePseEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePseEngineId

`func (o *WebSearchConfig) SetGooglePseEngineId(v string)`

SetGooglePseEngineId sets GooglePseEngineId field to given value.

### HasGooglePseEngineId

`func (o *WebSearchConfig) HasGooglePseEngineId() bool`

HasGooglePseEngineId returns a boolean if a field has been set.

### SetGooglePseEngineIdNil

`func (o *WebSearchConfig) SetGooglePseEngineIdNil(b bool)`

 SetGooglePseEngineIdNil sets the value for GooglePseEngineId to be an explicit nil

### UnsetGooglePseEngineId
`func (o *WebSearchConfig) UnsetGooglePseEngineId()`

UnsetGooglePseEngineId ensures that no value is present for GooglePseEngineId, not even an explicit nil
### GetBraveSearchApiKey

`func (o *WebSearchConfig) GetBraveSearchApiKey() string`

GetBraveSearchApiKey returns the BraveSearchApiKey field if non-nil, zero value otherwise.

### GetBraveSearchApiKeyOk

`func (o *WebSearchConfig) GetBraveSearchApiKeyOk() (*string, bool)`

GetBraveSearchApiKeyOk returns a tuple with the BraveSearchApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBraveSearchApiKey

`func (o *WebSearchConfig) SetBraveSearchApiKey(v string)`

SetBraveSearchApiKey sets BraveSearchApiKey field to given value.

### HasBraveSearchApiKey

`func (o *WebSearchConfig) HasBraveSearchApiKey() bool`

HasBraveSearchApiKey returns a boolean if a field has been set.

### SetBraveSearchApiKeyNil

`func (o *WebSearchConfig) SetBraveSearchApiKeyNil(b bool)`

 SetBraveSearchApiKeyNil sets the value for BraveSearchApiKey to be an explicit nil

### UnsetBraveSearchApiKey
`func (o *WebSearchConfig) UnsetBraveSearchApiKey()`

UnsetBraveSearchApiKey ensures that no value is present for BraveSearchApiKey, not even an explicit nil
### GetKagiSearchApiKey

`func (o *WebSearchConfig) GetKagiSearchApiKey() string`

GetKagiSearchApiKey returns the KagiSearchApiKey field if non-nil, zero value otherwise.

### GetKagiSearchApiKeyOk

`func (o *WebSearchConfig) GetKagiSearchApiKeyOk() (*string, bool)`

GetKagiSearchApiKeyOk returns a tuple with the KagiSearchApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKagiSearchApiKey

`func (o *WebSearchConfig) SetKagiSearchApiKey(v string)`

SetKagiSearchApiKey sets KagiSearchApiKey field to given value.

### HasKagiSearchApiKey

`func (o *WebSearchConfig) HasKagiSearchApiKey() bool`

HasKagiSearchApiKey returns a boolean if a field has been set.

### SetKagiSearchApiKeyNil

`func (o *WebSearchConfig) SetKagiSearchApiKeyNil(b bool)`

 SetKagiSearchApiKeyNil sets the value for KagiSearchApiKey to be an explicit nil

### UnsetKagiSearchApiKey
`func (o *WebSearchConfig) UnsetKagiSearchApiKey()`

UnsetKagiSearchApiKey ensures that no value is present for KagiSearchApiKey, not even an explicit nil
### GetMojeekSearchApiKey

`func (o *WebSearchConfig) GetMojeekSearchApiKey() string`

GetMojeekSearchApiKey returns the MojeekSearchApiKey field if non-nil, zero value otherwise.

### GetMojeekSearchApiKeyOk

`func (o *WebSearchConfig) GetMojeekSearchApiKeyOk() (*string, bool)`

GetMojeekSearchApiKeyOk returns a tuple with the MojeekSearchApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMojeekSearchApiKey

`func (o *WebSearchConfig) SetMojeekSearchApiKey(v string)`

SetMojeekSearchApiKey sets MojeekSearchApiKey field to given value.

### HasMojeekSearchApiKey

`func (o *WebSearchConfig) HasMojeekSearchApiKey() bool`

HasMojeekSearchApiKey returns a boolean if a field has been set.

### SetMojeekSearchApiKeyNil

`func (o *WebSearchConfig) SetMojeekSearchApiKeyNil(b bool)`

 SetMojeekSearchApiKeyNil sets the value for MojeekSearchApiKey to be an explicit nil

### UnsetMojeekSearchApiKey
`func (o *WebSearchConfig) UnsetMojeekSearchApiKey()`

UnsetMojeekSearchApiKey ensures that no value is present for MojeekSearchApiKey, not even an explicit nil
### GetBochaSearchApiKey

`func (o *WebSearchConfig) GetBochaSearchApiKey() string`

GetBochaSearchApiKey returns the BochaSearchApiKey field if non-nil, zero value otherwise.

### GetBochaSearchApiKeyOk

`func (o *WebSearchConfig) GetBochaSearchApiKeyOk() (*string, bool)`

GetBochaSearchApiKeyOk returns a tuple with the BochaSearchApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBochaSearchApiKey

`func (o *WebSearchConfig) SetBochaSearchApiKey(v string)`

SetBochaSearchApiKey sets BochaSearchApiKey field to given value.

### HasBochaSearchApiKey

`func (o *WebSearchConfig) HasBochaSearchApiKey() bool`

HasBochaSearchApiKey returns a boolean if a field has been set.

### SetBochaSearchApiKeyNil

`func (o *WebSearchConfig) SetBochaSearchApiKeyNil(b bool)`

 SetBochaSearchApiKeyNil sets the value for BochaSearchApiKey to be an explicit nil

### UnsetBochaSearchApiKey
`func (o *WebSearchConfig) UnsetBochaSearchApiKey()`

UnsetBochaSearchApiKey ensures that no value is present for BochaSearchApiKey, not even an explicit nil
### GetSerpstackApiKey

`func (o *WebSearchConfig) GetSerpstackApiKey() string`

GetSerpstackApiKey returns the SerpstackApiKey field if non-nil, zero value otherwise.

### GetSerpstackApiKeyOk

`func (o *WebSearchConfig) GetSerpstackApiKeyOk() (*string, bool)`

GetSerpstackApiKeyOk returns a tuple with the SerpstackApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerpstackApiKey

`func (o *WebSearchConfig) SetSerpstackApiKey(v string)`

SetSerpstackApiKey sets SerpstackApiKey field to given value.

### HasSerpstackApiKey

`func (o *WebSearchConfig) HasSerpstackApiKey() bool`

HasSerpstackApiKey returns a boolean if a field has been set.

### SetSerpstackApiKeyNil

`func (o *WebSearchConfig) SetSerpstackApiKeyNil(b bool)`

 SetSerpstackApiKeyNil sets the value for SerpstackApiKey to be an explicit nil

### UnsetSerpstackApiKey
`func (o *WebSearchConfig) UnsetSerpstackApiKey()`

UnsetSerpstackApiKey ensures that no value is present for SerpstackApiKey, not even an explicit nil
### GetSerpstackHttps

`func (o *WebSearchConfig) GetSerpstackHttps() bool`

GetSerpstackHttps returns the SerpstackHttps field if non-nil, zero value otherwise.

### GetSerpstackHttpsOk

`func (o *WebSearchConfig) GetSerpstackHttpsOk() (*bool, bool)`

GetSerpstackHttpsOk returns a tuple with the SerpstackHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerpstackHttps

`func (o *WebSearchConfig) SetSerpstackHttps(v bool)`

SetSerpstackHttps sets SerpstackHttps field to given value.

### HasSerpstackHttps

`func (o *WebSearchConfig) HasSerpstackHttps() bool`

HasSerpstackHttps returns a boolean if a field has been set.

### SetSerpstackHttpsNil

`func (o *WebSearchConfig) SetSerpstackHttpsNil(b bool)`

 SetSerpstackHttpsNil sets the value for SerpstackHttps to be an explicit nil

### UnsetSerpstackHttps
`func (o *WebSearchConfig) UnsetSerpstackHttps()`

UnsetSerpstackHttps ensures that no value is present for SerpstackHttps, not even an explicit nil
### GetSerperApiKey

`func (o *WebSearchConfig) GetSerperApiKey() string`

GetSerperApiKey returns the SerperApiKey field if non-nil, zero value otherwise.

### GetSerperApiKeyOk

`func (o *WebSearchConfig) GetSerperApiKeyOk() (*string, bool)`

GetSerperApiKeyOk returns a tuple with the SerperApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerperApiKey

`func (o *WebSearchConfig) SetSerperApiKey(v string)`

SetSerperApiKey sets SerperApiKey field to given value.

### HasSerperApiKey

`func (o *WebSearchConfig) HasSerperApiKey() bool`

HasSerperApiKey returns a boolean if a field has been set.

### SetSerperApiKeyNil

`func (o *WebSearchConfig) SetSerperApiKeyNil(b bool)`

 SetSerperApiKeyNil sets the value for SerperApiKey to be an explicit nil

### UnsetSerperApiKey
`func (o *WebSearchConfig) UnsetSerperApiKey()`

UnsetSerperApiKey ensures that no value is present for SerperApiKey, not even an explicit nil
### GetSerplyApiKey

`func (o *WebSearchConfig) GetSerplyApiKey() string`

GetSerplyApiKey returns the SerplyApiKey field if non-nil, zero value otherwise.

### GetSerplyApiKeyOk

`func (o *WebSearchConfig) GetSerplyApiKeyOk() (*string, bool)`

GetSerplyApiKeyOk returns a tuple with the SerplyApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerplyApiKey

`func (o *WebSearchConfig) SetSerplyApiKey(v string)`

SetSerplyApiKey sets SerplyApiKey field to given value.

### HasSerplyApiKey

`func (o *WebSearchConfig) HasSerplyApiKey() bool`

HasSerplyApiKey returns a boolean if a field has been set.

### SetSerplyApiKeyNil

`func (o *WebSearchConfig) SetSerplyApiKeyNil(b bool)`

 SetSerplyApiKeyNil sets the value for SerplyApiKey to be an explicit nil

### UnsetSerplyApiKey
`func (o *WebSearchConfig) UnsetSerplyApiKey()`

UnsetSerplyApiKey ensures that no value is present for SerplyApiKey, not even an explicit nil
### GetTavilyApiKey

`func (o *WebSearchConfig) GetTavilyApiKey() string`

GetTavilyApiKey returns the TavilyApiKey field if non-nil, zero value otherwise.

### GetTavilyApiKeyOk

`func (o *WebSearchConfig) GetTavilyApiKeyOk() (*string, bool)`

GetTavilyApiKeyOk returns a tuple with the TavilyApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTavilyApiKey

`func (o *WebSearchConfig) SetTavilyApiKey(v string)`

SetTavilyApiKey sets TavilyApiKey field to given value.

### HasTavilyApiKey

`func (o *WebSearchConfig) HasTavilyApiKey() bool`

HasTavilyApiKey returns a boolean if a field has been set.

### SetTavilyApiKeyNil

`func (o *WebSearchConfig) SetTavilyApiKeyNil(b bool)`

 SetTavilyApiKeyNil sets the value for TavilyApiKey to be an explicit nil

### UnsetTavilyApiKey
`func (o *WebSearchConfig) UnsetTavilyApiKey()`

UnsetTavilyApiKey ensures that no value is present for TavilyApiKey, not even an explicit nil
### GetSearchapiApiKey

`func (o *WebSearchConfig) GetSearchapiApiKey() string`

GetSearchapiApiKey returns the SearchapiApiKey field if non-nil, zero value otherwise.

### GetSearchapiApiKeyOk

`func (o *WebSearchConfig) GetSearchapiApiKeyOk() (*string, bool)`

GetSearchapiApiKeyOk returns a tuple with the SearchapiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchapiApiKey

`func (o *WebSearchConfig) SetSearchapiApiKey(v string)`

SetSearchapiApiKey sets SearchapiApiKey field to given value.

### HasSearchapiApiKey

`func (o *WebSearchConfig) HasSearchapiApiKey() bool`

HasSearchapiApiKey returns a boolean if a field has been set.

### SetSearchapiApiKeyNil

`func (o *WebSearchConfig) SetSearchapiApiKeyNil(b bool)`

 SetSearchapiApiKeyNil sets the value for SearchapiApiKey to be an explicit nil

### UnsetSearchapiApiKey
`func (o *WebSearchConfig) UnsetSearchapiApiKey()`

UnsetSearchapiApiKey ensures that no value is present for SearchapiApiKey, not even an explicit nil
### GetSearchapiEngine

`func (o *WebSearchConfig) GetSearchapiEngine() string`

GetSearchapiEngine returns the SearchapiEngine field if non-nil, zero value otherwise.

### GetSearchapiEngineOk

`func (o *WebSearchConfig) GetSearchapiEngineOk() (*string, bool)`

GetSearchapiEngineOk returns a tuple with the SearchapiEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchapiEngine

`func (o *WebSearchConfig) SetSearchapiEngine(v string)`

SetSearchapiEngine sets SearchapiEngine field to given value.

### HasSearchapiEngine

`func (o *WebSearchConfig) HasSearchapiEngine() bool`

HasSearchapiEngine returns a boolean if a field has been set.

### SetSearchapiEngineNil

`func (o *WebSearchConfig) SetSearchapiEngineNil(b bool)`

 SetSearchapiEngineNil sets the value for SearchapiEngine to be an explicit nil

### UnsetSearchapiEngine
`func (o *WebSearchConfig) UnsetSearchapiEngine()`

UnsetSearchapiEngine ensures that no value is present for SearchapiEngine, not even an explicit nil
### GetSerpapiApiKey

`func (o *WebSearchConfig) GetSerpapiApiKey() string`

GetSerpapiApiKey returns the SerpapiApiKey field if non-nil, zero value otherwise.

### GetSerpapiApiKeyOk

`func (o *WebSearchConfig) GetSerpapiApiKeyOk() (*string, bool)`

GetSerpapiApiKeyOk returns a tuple with the SerpapiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerpapiApiKey

`func (o *WebSearchConfig) SetSerpapiApiKey(v string)`

SetSerpapiApiKey sets SerpapiApiKey field to given value.

### HasSerpapiApiKey

`func (o *WebSearchConfig) HasSerpapiApiKey() bool`

HasSerpapiApiKey returns a boolean if a field has been set.

### SetSerpapiApiKeyNil

`func (o *WebSearchConfig) SetSerpapiApiKeyNil(b bool)`

 SetSerpapiApiKeyNil sets the value for SerpapiApiKey to be an explicit nil

### UnsetSerpapiApiKey
`func (o *WebSearchConfig) UnsetSerpapiApiKey()`

UnsetSerpapiApiKey ensures that no value is present for SerpapiApiKey, not even an explicit nil
### GetSerpapiEngine

`func (o *WebSearchConfig) GetSerpapiEngine() string`

GetSerpapiEngine returns the SerpapiEngine field if non-nil, zero value otherwise.

### GetSerpapiEngineOk

`func (o *WebSearchConfig) GetSerpapiEngineOk() (*string, bool)`

GetSerpapiEngineOk returns a tuple with the SerpapiEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerpapiEngine

`func (o *WebSearchConfig) SetSerpapiEngine(v string)`

SetSerpapiEngine sets SerpapiEngine field to given value.

### HasSerpapiEngine

`func (o *WebSearchConfig) HasSerpapiEngine() bool`

HasSerpapiEngine returns a boolean if a field has been set.

### SetSerpapiEngineNil

`func (o *WebSearchConfig) SetSerpapiEngineNil(b bool)`

 SetSerpapiEngineNil sets the value for SerpapiEngine to be an explicit nil

### UnsetSerpapiEngine
`func (o *WebSearchConfig) UnsetSerpapiEngine()`

UnsetSerpapiEngine ensures that no value is present for SerpapiEngine, not even an explicit nil
### GetJinaApiKey

`func (o *WebSearchConfig) GetJinaApiKey() string`

GetJinaApiKey returns the JinaApiKey field if non-nil, zero value otherwise.

### GetJinaApiKeyOk

`func (o *WebSearchConfig) GetJinaApiKeyOk() (*string, bool)`

GetJinaApiKeyOk returns a tuple with the JinaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJinaApiKey

`func (o *WebSearchConfig) SetJinaApiKey(v string)`

SetJinaApiKey sets JinaApiKey field to given value.

### HasJinaApiKey

`func (o *WebSearchConfig) HasJinaApiKey() bool`

HasJinaApiKey returns a boolean if a field has been set.

### SetJinaApiKeyNil

`func (o *WebSearchConfig) SetJinaApiKeyNil(b bool)`

 SetJinaApiKeyNil sets the value for JinaApiKey to be an explicit nil

### UnsetJinaApiKey
`func (o *WebSearchConfig) UnsetJinaApiKey()`

UnsetJinaApiKey ensures that no value is present for JinaApiKey, not even an explicit nil
### GetBingSearchV7Endpoint

`func (o *WebSearchConfig) GetBingSearchV7Endpoint() string`

GetBingSearchV7Endpoint returns the BingSearchV7Endpoint field if non-nil, zero value otherwise.

### GetBingSearchV7EndpointOk

`func (o *WebSearchConfig) GetBingSearchV7EndpointOk() (*string, bool)`

GetBingSearchV7EndpointOk returns a tuple with the BingSearchV7Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBingSearchV7Endpoint

`func (o *WebSearchConfig) SetBingSearchV7Endpoint(v string)`

SetBingSearchV7Endpoint sets BingSearchV7Endpoint field to given value.

### HasBingSearchV7Endpoint

`func (o *WebSearchConfig) HasBingSearchV7Endpoint() bool`

HasBingSearchV7Endpoint returns a boolean if a field has been set.

### SetBingSearchV7EndpointNil

`func (o *WebSearchConfig) SetBingSearchV7EndpointNil(b bool)`

 SetBingSearchV7EndpointNil sets the value for BingSearchV7Endpoint to be an explicit nil

### UnsetBingSearchV7Endpoint
`func (o *WebSearchConfig) UnsetBingSearchV7Endpoint()`

UnsetBingSearchV7Endpoint ensures that no value is present for BingSearchV7Endpoint, not even an explicit nil
### GetBingSearchV7SubscriptionKey

`func (o *WebSearchConfig) GetBingSearchV7SubscriptionKey() string`

GetBingSearchV7SubscriptionKey returns the BingSearchV7SubscriptionKey field if non-nil, zero value otherwise.

### GetBingSearchV7SubscriptionKeyOk

`func (o *WebSearchConfig) GetBingSearchV7SubscriptionKeyOk() (*string, bool)`

GetBingSearchV7SubscriptionKeyOk returns a tuple with the BingSearchV7SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBingSearchV7SubscriptionKey

`func (o *WebSearchConfig) SetBingSearchV7SubscriptionKey(v string)`

SetBingSearchV7SubscriptionKey sets BingSearchV7SubscriptionKey field to given value.

### HasBingSearchV7SubscriptionKey

`func (o *WebSearchConfig) HasBingSearchV7SubscriptionKey() bool`

HasBingSearchV7SubscriptionKey returns a boolean if a field has been set.

### SetBingSearchV7SubscriptionKeyNil

`func (o *WebSearchConfig) SetBingSearchV7SubscriptionKeyNil(b bool)`

 SetBingSearchV7SubscriptionKeyNil sets the value for BingSearchV7SubscriptionKey to be an explicit nil

### UnsetBingSearchV7SubscriptionKey
`func (o *WebSearchConfig) UnsetBingSearchV7SubscriptionKey()`

UnsetBingSearchV7SubscriptionKey ensures that no value is present for BingSearchV7SubscriptionKey, not even an explicit nil
### GetExaApiKey

`func (o *WebSearchConfig) GetExaApiKey() string`

GetExaApiKey returns the ExaApiKey field if non-nil, zero value otherwise.

### GetExaApiKeyOk

`func (o *WebSearchConfig) GetExaApiKeyOk() (*string, bool)`

GetExaApiKeyOk returns a tuple with the ExaApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExaApiKey

`func (o *WebSearchConfig) SetExaApiKey(v string)`

SetExaApiKey sets ExaApiKey field to given value.

### HasExaApiKey

`func (o *WebSearchConfig) HasExaApiKey() bool`

HasExaApiKey returns a boolean if a field has been set.

### SetExaApiKeyNil

`func (o *WebSearchConfig) SetExaApiKeyNil(b bool)`

 SetExaApiKeyNil sets the value for ExaApiKey to be an explicit nil

### UnsetExaApiKey
`func (o *WebSearchConfig) UnsetExaApiKey()`

UnsetExaApiKey ensures that no value is present for ExaApiKey, not even an explicit nil
### GetPerplexityApiKey

`func (o *WebSearchConfig) GetPerplexityApiKey() string`

GetPerplexityApiKey returns the PerplexityApiKey field if non-nil, zero value otherwise.

### GetPerplexityApiKeyOk

`func (o *WebSearchConfig) GetPerplexityApiKeyOk() (*string, bool)`

GetPerplexityApiKeyOk returns a tuple with the PerplexityApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerplexityApiKey

`func (o *WebSearchConfig) SetPerplexityApiKey(v string)`

SetPerplexityApiKey sets PerplexityApiKey field to given value.

### HasPerplexityApiKey

`func (o *WebSearchConfig) HasPerplexityApiKey() bool`

HasPerplexityApiKey returns a boolean if a field has been set.

### SetPerplexityApiKeyNil

`func (o *WebSearchConfig) SetPerplexityApiKeyNil(b bool)`

 SetPerplexityApiKeyNil sets the value for PerplexityApiKey to be an explicit nil

### UnsetPerplexityApiKey
`func (o *WebSearchConfig) UnsetPerplexityApiKey()`

UnsetPerplexityApiKey ensures that no value is present for PerplexityApiKey, not even an explicit nil
### GetResultCount

`func (o *WebSearchConfig) GetResultCount() int32`

GetResultCount returns the ResultCount field if non-nil, zero value otherwise.

### GetResultCountOk

`func (o *WebSearchConfig) GetResultCountOk() (*int32, bool)`

GetResultCountOk returns a tuple with the ResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCount

`func (o *WebSearchConfig) SetResultCount(v int32)`

SetResultCount sets ResultCount field to given value.

### HasResultCount

`func (o *WebSearchConfig) HasResultCount() bool`

HasResultCount returns a boolean if a field has been set.

### SetResultCountNil

`func (o *WebSearchConfig) SetResultCountNil(b bool)`

 SetResultCountNil sets the value for ResultCount to be an explicit nil

### UnsetResultCount
`func (o *WebSearchConfig) UnsetResultCount()`

UnsetResultCount ensures that no value is present for ResultCount, not even an explicit nil
### GetConcurrentRequests

`func (o *WebSearchConfig) GetConcurrentRequests() int32`

GetConcurrentRequests returns the ConcurrentRequests field if non-nil, zero value otherwise.

### GetConcurrentRequestsOk

`func (o *WebSearchConfig) GetConcurrentRequestsOk() (*int32, bool)`

GetConcurrentRequestsOk returns a tuple with the ConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentRequests

`func (o *WebSearchConfig) SetConcurrentRequests(v int32)`

SetConcurrentRequests sets ConcurrentRequests field to given value.

### HasConcurrentRequests

`func (o *WebSearchConfig) HasConcurrentRequests() bool`

HasConcurrentRequests returns a boolean if a field has been set.

### SetConcurrentRequestsNil

`func (o *WebSearchConfig) SetConcurrentRequestsNil(b bool)`

 SetConcurrentRequestsNil sets the value for ConcurrentRequests to be an explicit nil

### UnsetConcurrentRequests
`func (o *WebSearchConfig) UnsetConcurrentRequests()`

UnsetConcurrentRequests ensures that no value is present for ConcurrentRequests, not even an explicit nil
### GetTrustEnv

`func (o *WebSearchConfig) GetTrustEnv() bool`

GetTrustEnv returns the TrustEnv field if non-nil, zero value otherwise.

### GetTrustEnvOk

`func (o *WebSearchConfig) GetTrustEnvOk() (*bool, bool)`

GetTrustEnvOk returns a tuple with the TrustEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustEnv

`func (o *WebSearchConfig) SetTrustEnv(v bool)`

SetTrustEnv sets TrustEnv field to given value.

### HasTrustEnv

`func (o *WebSearchConfig) HasTrustEnv() bool`

HasTrustEnv returns a boolean if a field has been set.

### SetTrustEnvNil

`func (o *WebSearchConfig) SetTrustEnvNil(b bool)`

 SetTrustEnvNil sets the value for TrustEnv to be an explicit nil

### UnsetTrustEnv
`func (o *WebSearchConfig) UnsetTrustEnv()`

UnsetTrustEnv ensures that no value is present for TrustEnv, not even an explicit nil
### GetDomainFilterList

`func (o *WebSearchConfig) GetDomainFilterList() []string`

GetDomainFilterList returns the DomainFilterList field if non-nil, zero value otherwise.

### GetDomainFilterListOk

`func (o *WebSearchConfig) GetDomainFilterListOk() (*[]string, bool)`

GetDomainFilterListOk returns a tuple with the DomainFilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainFilterList

`func (o *WebSearchConfig) SetDomainFilterList(v []string)`

SetDomainFilterList sets DomainFilterList field to given value.

### HasDomainFilterList

`func (o *WebSearchConfig) HasDomainFilterList() bool`

HasDomainFilterList returns a boolean if a field has been set.

### SetDomainFilterListNil

`func (o *WebSearchConfig) SetDomainFilterListNil(b bool)`

 SetDomainFilterListNil sets the value for DomainFilterList to be an explicit nil

### UnsetDomainFilterList
`func (o *WebSearchConfig) UnsetDomainFilterList()`

UnsetDomainFilterList ensures that no value is present for DomainFilterList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


