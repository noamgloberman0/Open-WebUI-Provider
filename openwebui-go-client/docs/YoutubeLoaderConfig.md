# YoutubeLoaderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **[]string** |  | 
**Translation** | Pointer to **NullableString** |  | [optional] 
**ProxyUrl** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewYoutubeLoaderConfig

`func NewYoutubeLoaderConfig(language []string, ) *YoutubeLoaderConfig`

NewYoutubeLoaderConfig instantiates a new YoutubeLoaderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYoutubeLoaderConfigWithDefaults

`func NewYoutubeLoaderConfigWithDefaults() *YoutubeLoaderConfig`

NewYoutubeLoaderConfigWithDefaults instantiates a new YoutubeLoaderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *YoutubeLoaderConfig) GetLanguage() []string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *YoutubeLoaderConfig) GetLanguageOk() (*[]string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *YoutubeLoaderConfig) SetLanguage(v []string)`

SetLanguage sets Language field to given value.


### GetTranslation

`func (o *YoutubeLoaderConfig) GetTranslation() string`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *YoutubeLoaderConfig) GetTranslationOk() (*string, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *YoutubeLoaderConfig) SetTranslation(v string)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *YoutubeLoaderConfig) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### SetTranslationNil

`func (o *YoutubeLoaderConfig) SetTranslationNil(b bool)`

 SetTranslationNil sets the value for Translation to be an explicit nil

### UnsetTranslation
`func (o *YoutubeLoaderConfig) UnsetTranslation()`

UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
### GetProxyUrl

`func (o *YoutubeLoaderConfig) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *YoutubeLoaderConfig) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *YoutubeLoaderConfig) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *YoutubeLoaderConfig) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


