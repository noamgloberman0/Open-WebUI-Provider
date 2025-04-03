# ModelMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileImageUrl** | Pointer to [**ProfileImageUrl**](ProfileImageUrl.md) |  | [optional] [default to /static/favicon.png]
**Description** | Pointer to [**Description**](Description.md) |  | [optional] 
**Capabilities** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 

## Methods

### NewModelMeta

`func NewModelMeta() *ModelMeta`

NewModelMeta instantiates a new ModelMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetaWithDefaults

`func NewModelMetaWithDefaults() *ModelMeta`

NewModelMetaWithDefaults instantiates a new ModelMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileImageUrl

`func (o *ModelMeta) GetProfileImageUrl() ProfileImageUrl`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *ModelMeta) GetProfileImageUrlOk() (*ProfileImageUrl, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *ModelMeta) SetProfileImageUrl(v ProfileImageUrl)`

SetProfileImageUrl sets ProfileImageUrl field to given value.

### HasProfileImageUrl

`func (o *ModelMeta) HasProfileImageUrl() bool`

HasProfileImageUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ModelMeta) GetDescription() Description`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelMeta) GetDescriptionOk() (*Description, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelMeta) SetDescription(v Description)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCapabilities

`func (o *ModelMeta) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelMeta) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelMeta) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelMeta) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


