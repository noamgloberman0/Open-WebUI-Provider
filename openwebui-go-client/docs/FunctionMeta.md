# FunctionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Manifest** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFunctionMeta

`func NewFunctionMeta() *FunctionMeta`

NewFunctionMeta instantiates a new FunctionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMetaWithDefaults

`func NewFunctionMetaWithDefaults() *FunctionMeta`

NewFunctionMetaWithDefaults instantiates a new FunctionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FunctionMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FunctionMeta) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FunctionMeta) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetManifest

`func (o *FunctionMeta) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FunctionMeta) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FunctionMeta) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FunctionMeta) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *FunctionMeta) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *FunctionMeta) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


