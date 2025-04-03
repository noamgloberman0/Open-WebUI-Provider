# QueryDocForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** |  | 
**Query** | **string** |  | 
**K** | Pointer to **NullableInt32** |  | [optional] 
**R** | Pointer to **NullableFloat32** |  | [optional] 
**Hybrid** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewQueryDocForm

`func NewQueryDocForm(collectionName string, query string, ) *QueryDocForm`

NewQueryDocForm instantiates a new QueryDocForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDocFormWithDefaults

`func NewQueryDocFormWithDefaults() *QueryDocForm`

NewQueryDocFormWithDefaults instantiates a new QueryDocForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *QueryDocForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *QueryDocForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *QueryDocForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetQuery

`func (o *QueryDocForm) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryDocForm) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryDocForm) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetK

`func (o *QueryDocForm) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *QueryDocForm) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *QueryDocForm) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *QueryDocForm) HasK() bool`

HasK returns a boolean if a field has been set.

### SetKNil

`func (o *QueryDocForm) SetKNil(b bool)`

 SetKNil sets the value for K to be an explicit nil

### UnsetK
`func (o *QueryDocForm) UnsetK()`

UnsetK ensures that no value is present for K, not even an explicit nil
### GetR

`func (o *QueryDocForm) GetR() float32`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *QueryDocForm) GetROk() (*float32, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *QueryDocForm) SetR(v float32)`

SetR sets R field to given value.

### HasR

`func (o *QueryDocForm) HasR() bool`

HasR returns a boolean if a field has been set.

### SetRNil

`func (o *QueryDocForm) SetRNil(b bool)`

 SetRNil sets the value for R to be an explicit nil

### UnsetR
`func (o *QueryDocForm) UnsetR()`

UnsetR ensures that no value is present for R, not even an explicit nil
### GetHybrid

`func (o *QueryDocForm) GetHybrid() bool`

GetHybrid returns the Hybrid field if non-nil, zero value otherwise.

### GetHybridOk

`func (o *QueryDocForm) GetHybridOk() (*bool, bool)`

GetHybridOk returns a tuple with the Hybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrid

`func (o *QueryDocForm) SetHybrid(v bool)`

SetHybrid sets Hybrid field to given value.

### HasHybrid

`func (o *QueryDocForm) HasHybrid() bool`

HasHybrid returns a boolean if a field has been set.

### SetHybridNil

`func (o *QueryDocForm) SetHybridNil(b bool)`

 SetHybridNil sets the value for Hybrid to be an explicit nil

### UnsetHybrid
`func (o *QueryDocForm) UnsetHybrid()`

UnsetHybrid ensures that no value is present for Hybrid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


