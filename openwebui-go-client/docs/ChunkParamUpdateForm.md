# ChunkParamUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TextSplitter** | Pointer to **NullableString** |  | [optional] 
**ChunkSize** | **int32** |  | 
**ChunkOverlap** | **int32** |  | 

## Methods

### NewChunkParamUpdateForm

`func NewChunkParamUpdateForm(chunkSize int32, chunkOverlap int32, ) *ChunkParamUpdateForm`

NewChunkParamUpdateForm instantiates a new ChunkParamUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunkParamUpdateFormWithDefaults

`func NewChunkParamUpdateFormWithDefaults() *ChunkParamUpdateForm`

NewChunkParamUpdateFormWithDefaults instantiates a new ChunkParamUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTextSplitter

`func (o *ChunkParamUpdateForm) GetTextSplitter() string`

GetTextSplitter returns the TextSplitter field if non-nil, zero value otherwise.

### GetTextSplitterOk

`func (o *ChunkParamUpdateForm) GetTextSplitterOk() (*string, bool)`

GetTextSplitterOk returns a tuple with the TextSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSplitter

`func (o *ChunkParamUpdateForm) SetTextSplitter(v string)`

SetTextSplitter sets TextSplitter field to given value.

### HasTextSplitter

`func (o *ChunkParamUpdateForm) HasTextSplitter() bool`

HasTextSplitter returns a boolean if a field has been set.

### SetTextSplitterNil

`func (o *ChunkParamUpdateForm) SetTextSplitterNil(b bool)`

 SetTextSplitterNil sets the value for TextSplitter to be an explicit nil

### UnsetTextSplitter
`func (o *ChunkParamUpdateForm) UnsetTextSplitter()`

UnsetTextSplitter ensures that no value is present for TextSplitter, not even an explicit nil
### GetChunkSize

`func (o *ChunkParamUpdateForm) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *ChunkParamUpdateForm) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *ChunkParamUpdateForm) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.


### GetChunkOverlap

`func (o *ChunkParamUpdateForm) GetChunkOverlap() int32`

GetChunkOverlap returns the ChunkOverlap field if non-nil, zero value otherwise.

### GetChunkOverlapOk

`func (o *ChunkParamUpdateForm) GetChunkOverlapOk() (*int32, bool)`

GetChunkOverlapOk returns a tuple with the ChunkOverlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkOverlap

`func (o *ChunkParamUpdateForm) SetChunkOverlap(v int32)`

SetChunkOverlap sets ChunkOverlap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


