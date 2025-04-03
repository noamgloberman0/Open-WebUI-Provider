# ProcessUrlForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewProcessUrlForm

`func NewProcessUrlForm(url string, ) *ProcessUrlForm`

NewProcessUrlForm instantiates a new ProcessUrlForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessUrlFormWithDefaults

`func NewProcessUrlFormWithDefaults() *ProcessUrlForm`

NewProcessUrlFormWithDefaults instantiates a new ProcessUrlForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *ProcessUrlForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ProcessUrlForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ProcessUrlForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *ProcessUrlForm) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *ProcessUrlForm) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *ProcessUrlForm) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil
### GetUrl

`func (o *ProcessUrlForm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProcessUrlForm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProcessUrlForm) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


