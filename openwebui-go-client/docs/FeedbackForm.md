# FeedbackForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** | Pointer to [**FeedbackFormData**](FeedbackFormData.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Snapshot** | Pointer to [**FeedbackFormSnapshot**](FeedbackFormSnapshot.md) |  | [optional] 

## Methods

### NewFeedbackForm

`func NewFeedbackForm(type_ string, ) *FeedbackForm`

NewFeedbackForm instantiates a new FeedbackForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFormWithDefaults

`func NewFeedbackFormWithDefaults() *FeedbackForm`

NewFeedbackFormWithDefaults instantiates a new FeedbackForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FeedbackForm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedbackForm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedbackForm) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *FeedbackForm) GetData() FeedbackFormData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FeedbackForm) GetDataOk() (*FeedbackFormData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FeedbackForm) SetData(v FeedbackFormData)`

SetData sets Data field to given value.

### HasData

`func (o *FeedbackForm) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *FeedbackForm) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FeedbackForm) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FeedbackForm) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FeedbackForm) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSnapshot

`func (o *FeedbackForm) GetSnapshot() FeedbackFormSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *FeedbackForm) GetSnapshotOk() (*FeedbackFormSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *FeedbackForm) SetSnapshot(v FeedbackFormSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *FeedbackForm) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


