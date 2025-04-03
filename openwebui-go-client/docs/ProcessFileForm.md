# ProcessFileForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** |  | 
**Content** | Pointer to **NullableString** |  | [optional] 
**CollectionName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProcessFileForm

`func NewProcessFileForm(fileId string, ) *ProcessFileForm`

NewProcessFileForm instantiates a new ProcessFileForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessFileFormWithDefaults

`func NewProcessFileFormWithDefaults() *ProcessFileForm`

NewProcessFileFormWithDefaults instantiates a new ProcessFileForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ProcessFileForm) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ProcessFileForm) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ProcessFileForm) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetContent

`func (o *ProcessFileForm) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProcessFileForm) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProcessFileForm) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ProcessFileForm) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ProcessFileForm) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ProcessFileForm) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCollectionName

`func (o *ProcessFileForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ProcessFileForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ProcessFileForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *ProcessFileForm) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *ProcessFileForm) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *ProcessFileForm) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


