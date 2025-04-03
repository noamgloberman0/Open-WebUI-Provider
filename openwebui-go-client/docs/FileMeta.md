# FileMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**Size** | Pointer to [**Size**](Size.md) |  | [optional] 

## Methods

### NewFileMeta

`func NewFileMeta() *FileMeta`

NewFileMeta instantiates a new FileMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetaWithDefaults

`func NewFileMetaWithDefaults() *FileMeta`

NewFileMetaWithDefaults instantiates a new FileMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileMeta) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileMeta) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileMeta) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *FileMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *FileMeta) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FileMeta) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FileMeta) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FileMeta) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *FileMeta) GetSize() Size`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileMeta) GetSizeOk() (*Size, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileMeta) SetSize(v Size)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileMeta) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


