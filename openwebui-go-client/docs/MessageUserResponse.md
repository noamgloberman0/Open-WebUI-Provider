# MessageUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**ChannelId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Content** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **int32** |  | 
**UpdatedAt** | **int32** |  | 
**LatestReplyAt** | **NullableInt32** |  | 
**ReplyCount** | **int32** |  | 
**Reactions** | [**[]Reactions**](Reactions.md) |  | 
**User** | [**UserNameResponse**](UserNameResponse.md) |  | 

## Methods

### NewMessageUserResponse

`func NewMessageUserResponse(id string, userId string, content string, createdAt int32, updatedAt int32, latestReplyAt NullableInt32, replyCount int32, reactions []Reactions, user UserNameResponse, ) *MessageUserResponse`

NewMessageUserResponse instantiates a new MessageUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUserResponseWithDefaults

`func NewMessageUserResponseWithDefaults() *MessageUserResponse`

NewMessageUserResponseWithDefaults instantiates a new MessageUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *MessageUserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessageUserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessageUserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannelId

`func (o *MessageUserResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MessageUserResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MessageUserResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MessageUserResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MessageUserResponse) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MessageUserResponse) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetParentId

`func (o *MessageUserResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MessageUserResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MessageUserResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MessageUserResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *MessageUserResponse) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *MessageUserResponse) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetContent

`func (o *MessageUserResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageUserResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageUserResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetData

`func (o *MessageUserResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessageUserResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MessageUserResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *MessageUserResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *MessageUserResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *MessageUserResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *MessageUserResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MessageUserResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MessageUserResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MessageUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MessageUserResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MessageUserResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetCreatedAt

`func (o *MessageUserResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageUserResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageUserResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MessageUserResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageUserResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageUserResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLatestReplyAt

`func (o *MessageUserResponse) GetLatestReplyAt() int32`

GetLatestReplyAt returns the LatestReplyAt field if non-nil, zero value otherwise.

### GetLatestReplyAtOk

`func (o *MessageUserResponse) GetLatestReplyAtOk() (*int32, bool)`

GetLatestReplyAtOk returns a tuple with the LatestReplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReplyAt

`func (o *MessageUserResponse) SetLatestReplyAt(v int32)`

SetLatestReplyAt sets LatestReplyAt field to given value.


### SetLatestReplyAtNil

`func (o *MessageUserResponse) SetLatestReplyAtNil(b bool)`

 SetLatestReplyAtNil sets the value for LatestReplyAt to be an explicit nil

### UnsetLatestReplyAt
`func (o *MessageUserResponse) UnsetLatestReplyAt()`

UnsetLatestReplyAt ensures that no value is present for LatestReplyAt, not even an explicit nil
### GetReplyCount

`func (o *MessageUserResponse) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *MessageUserResponse) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *MessageUserResponse) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.


### GetReactions

`func (o *MessageUserResponse) GetReactions() []Reactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MessageUserResponse) GetReactionsOk() (*[]Reactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MessageUserResponse) SetReactions(v []Reactions)`

SetReactions sets Reactions field to given value.


### GetUser

`func (o *MessageUserResponse) GetUser() UserNameResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MessageUserResponse) GetUserOk() (*UserNameResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MessageUserResponse) SetUser(v UserNameResponse)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


