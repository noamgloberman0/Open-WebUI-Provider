# BatchProcessFilesForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]FileModel**](FileModel.md) |  | 
**CollectionName** | **string** |  | 

## Methods

### NewBatchProcessFilesForm

`func NewBatchProcessFilesForm(files []FileModel, collectionName string, ) *BatchProcessFilesForm`

NewBatchProcessFilesForm instantiates a new BatchProcessFilesForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProcessFilesFormWithDefaults

`func NewBatchProcessFilesFormWithDefaults() *BatchProcessFilesForm`

NewBatchProcessFilesFormWithDefaults instantiates a new BatchProcessFilesForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *BatchProcessFilesForm) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *BatchProcessFilesForm) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *BatchProcessFilesForm) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.


### GetCollectionName

`func (o *BatchProcessFilesForm) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *BatchProcessFilesForm) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *BatchProcessFilesForm) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


