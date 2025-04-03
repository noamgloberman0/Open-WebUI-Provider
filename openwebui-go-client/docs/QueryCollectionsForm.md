# QueryCollectionsForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionNames** | **[]string** |  | 
**Query** | **string** |  | 
**K** | Pointer to **NullableInt32** |  | [optional] 
**R** | Pointer to **NullableFloat32** |  | [optional] 
**Hybrid** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewQueryCollectionsForm

`func NewQueryCollectionsForm(collectionNames []string, query string, ) *QueryCollectionsForm`

NewQueryCollectionsForm instantiates a new QueryCollectionsForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCollectionsFormWithDefaults

`func NewQueryCollectionsFormWithDefaults() *QueryCollectionsForm`

NewQueryCollectionsFormWithDefaults instantiates a new QueryCollectionsForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionNames

`func (o *QueryCollectionsForm) GetCollectionNames() []string`

GetCollectionNames returns the CollectionNames field if non-nil, zero value otherwise.

### GetCollectionNamesOk

`func (o *QueryCollectionsForm) GetCollectionNamesOk() (*[]string, bool)`

GetCollectionNamesOk returns a tuple with the CollectionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionNames

`func (o *QueryCollectionsForm) SetCollectionNames(v []string)`

SetCollectionNames sets CollectionNames field to given value.


### GetQuery

`func (o *QueryCollectionsForm) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryCollectionsForm) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryCollectionsForm) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetK

`func (o *QueryCollectionsForm) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *QueryCollectionsForm) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *QueryCollectionsForm) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *QueryCollectionsForm) HasK() bool`

HasK returns a boolean if a field has been set.

### SetKNil

`func (o *QueryCollectionsForm) SetKNil(b bool)`

 SetKNil sets the value for K to be an explicit nil

### UnsetK
`func (o *QueryCollectionsForm) UnsetK()`

UnsetK ensures that no value is present for K, not even an explicit nil
### GetR

`func (o *QueryCollectionsForm) GetR() float32`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *QueryCollectionsForm) GetROk() (*float32, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *QueryCollectionsForm) SetR(v float32)`

SetR sets R field to given value.

### HasR

`func (o *QueryCollectionsForm) HasR() bool`

HasR returns a boolean if a field has been set.

### SetRNil

`func (o *QueryCollectionsForm) SetRNil(b bool)`

 SetRNil sets the value for R to be an explicit nil

### UnsetR
`func (o *QueryCollectionsForm) UnsetR()`

UnsetR ensures that no value is present for R, not even an explicit nil
### GetHybrid

`func (o *QueryCollectionsForm) GetHybrid() bool`

GetHybrid returns the Hybrid field if non-nil, zero value otherwise.

### GetHybridOk

`func (o *QueryCollectionsForm) GetHybridOk() (*bool, bool)`

GetHybridOk returns a tuple with the Hybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrid

`func (o *QueryCollectionsForm) SetHybrid(v bool)`

SetHybrid sets Hybrid field to given value.

### HasHybrid

`func (o *QueryCollectionsForm) HasHybrid() bool`

HasHybrid returns a boolean if a field has been set.

### SetHybridNil

`func (o *QueryCollectionsForm) SetHybridNil(b bool)`

 SetHybridNil sets the value for Hybrid to be an explicit nil

### UnsetHybrid
`func (o *QueryCollectionsForm) UnsetHybrid()`

UnsetHybrid ensures that no value is present for Hybrid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


