# BatchProcessFilesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** |  | 
**Status** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBatchProcessFilesResult

`func NewBatchProcessFilesResult(fileId string, status string, ) *BatchProcessFilesResult`

NewBatchProcessFilesResult instantiates a new BatchProcessFilesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProcessFilesResultWithDefaults

`func NewBatchProcessFilesResultWithDefaults() *BatchProcessFilesResult`

NewBatchProcessFilesResultWithDefaults instantiates a new BatchProcessFilesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *BatchProcessFilesResult) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *BatchProcessFilesResult) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *BatchProcessFilesResult) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetStatus

`func (o *BatchProcessFilesResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchProcessFilesResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchProcessFilesResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *BatchProcessFilesResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchProcessFilesResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchProcessFilesResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchProcessFilesResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BatchProcessFilesResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BatchProcessFilesResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


