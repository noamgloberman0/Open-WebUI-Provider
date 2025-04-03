# ProcessTextForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Content** | **string** |  | 
**CollectionName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProcessTextForm

`func NewProcessTextForm(name string, content string, ) *ProcessTextForm`

NewProcessTextForm instantiates a new ProcessTextForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessTextFormWithDefaults

`func NewProcessTextFormWithDefaults() *ProcessTextForm`

NewProcessTextFormWithDefaults instantiates a new ProcessTextForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProcessTextForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessTextForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessTextForm) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *ProcessTextForm) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProcessTextForm) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProcessTextForm) SetContent(v string)`

SetContent sets Content field to given value.


### GetCollectionName

`func (o *ProcessTextForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ProcessTextForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ProcessTextForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *ProcessTextForm) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *ProcessTextForm) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *ProcessTextForm) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


