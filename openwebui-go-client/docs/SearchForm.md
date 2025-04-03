# SearchForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **NullableString** |  | [optional] 
**Query** | **string** |  | 

## Methods

### NewSearchForm

`func NewSearchForm(query string, ) *SearchForm`

NewSearchForm instantiates a new SearchForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFormWithDefaults

`func NewSearchFormWithDefaults() *SearchForm`

NewSearchFormWithDefaults instantiates a new SearchForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *SearchForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *SearchForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *SearchForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *SearchForm) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *SearchForm) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *SearchForm) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil
### GetQuery

`func (o *SearchForm) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchForm) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchForm) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


