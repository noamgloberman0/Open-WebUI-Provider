# RatingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to [**Rating**](Rating.md) |  | [optional] 
**ModelId** | Pointer to [**ModelId**](ModelId.md) |  | [optional] 
**SiblingModelIds** | Pointer to [**SiblingModelIds**](SiblingModelIds.md) |  | [optional] 
**Reason** | Pointer to [**Reason**](Reason.md) |  | [optional] 
**Comment** | Pointer to [**Comment**](Comment.md) |  | [optional] 

## Methods

### NewRatingData

`func NewRatingData() *RatingData`

NewRatingData instantiates a new RatingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingDataWithDefaults

`func NewRatingDataWithDefaults() *RatingData`

NewRatingDataWithDefaults instantiates a new RatingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *RatingData) GetRating() Rating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *RatingData) GetRatingOk() (*Rating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *RatingData) SetRating(v Rating)`

SetRating sets Rating field to given value.

### HasRating

`func (o *RatingData) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetModelId

`func (o *RatingData) GetModelId() ModelId`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *RatingData) GetModelIdOk() (*ModelId, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *RatingData) SetModelId(v ModelId)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *RatingData) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetSiblingModelIds

`func (o *RatingData) GetSiblingModelIds() SiblingModelIds`

GetSiblingModelIds returns the SiblingModelIds field if non-nil, zero value otherwise.

### GetSiblingModelIdsOk

`func (o *RatingData) GetSiblingModelIdsOk() (*SiblingModelIds, bool)`

GetSiblingModelIdsOk returns a tuple with the SiblingModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiblingModelIds

`func (o *RatingData) SetSiblingModelIds(v SiblingModelIds)`

SetSiblingModelIds sets SiblingModelIds field to given value.

### HasSiblingModelIds

`func (o *RatingData) HasSiblingModelIds() bool`

HasSiblingModelIds returns a boolean if a field has been set.

### GetReason

`func (o *RatingData) GetReason() Reason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RatingData) GetReasonOk() (*Reason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RatingData) SetReason(v Reason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RatingData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *RatingData) GetComment() Comment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RatingData) GetCommentOk() (*Comment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RatingData) SetComment(v Comment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RatingData) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


