# TagFilterForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Skip** | Pointer to **NullableInt32** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTagFilterForm

`func NewTagFilterForm(name string, ) *TagFilterForm`

NewTagFilterForm instantiates a new TagFilterForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFilterFormWithDefaults

`func NewTagFilterFormWithDefaults() *TagFilterForm`

NewTagFilterFormWithDefaults instantiates a new TagFilterForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagFilterForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagFilterForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagFilterForm) SetName(v string)`

SetName sets Name field to given value.


### GetSkip

`func (o *TagFilterForm) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *TagFilterForm) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *TagFilterForm) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *TagFilterForm) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### SetSkipNil

`func (o *TagFilterForm) SetSkipNil(b bool)`

 SetSkipNil sets the value for Skip to be an explicit nil

### UnsetSkip
`func (o *TagFilterForm) UnsetSkip()`

UnsetSkip ensures that no value is present for Skip, not even an explicit nil
### GetLimit

`func (o *TagFilterForm) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TagFilterForm) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TagFilterForm) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TagFilterForm) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *TagFilterForm) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *TagFilterForm) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


