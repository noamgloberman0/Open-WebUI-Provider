# BatchProcessFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BatchProcessFilesResult**](BatchProcessFilesResult.md) |  | 
**Errors** | [**[]BatchProcessFilesResult**](BatchProcessFilesResult.md) |  | 

## Methods

### NewBatchProcessFilesResponse

`func NewBatchProcessFilesResponse(results []BatchProcessFilesResult, errors []BatchProcessFilesResult, ) *BatchProcessFilesResponse`

NewBatchProcessFilesResponse instantiates a new BatchProcessFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProcessFilesResponseWithDefaults

`func NewBatchProcessFilesResponseWithDefaults() *BatchProcessFilesResponse`

NewBatchProcessFilesResponseWithDefaults instantiates a new BatchProcessFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchProcessFilesResponse) GetResults() []BatchProcessFilesResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchProcessFilesResponse) GetResultsOk() (*[]BatchProcessFilesResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchProcessFilesResponse) SetResults(v []BatchProcessFilesResult)`

SetResults sets Results field to given value.


### GetErrors

`func (o *BatchProcessFilesResponse) GetErrors() []BatchProcessFilesResult`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchProcessFilesResponse) GetErrorsOk() (*[]BatchProcessFilesResult, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchProcessFilesResponse) SetErrors(v []BatchProcessFilesResult)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


