# ChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Title** | **string** |  | 
**Chat** | **map[string]interface{}** |  | 
**UpdatedAt** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**ShareId** | Pointer to **NullableString** |  | [optional] 
**Archived** | **bool** |  | 
**Pinned** | Pointer to **NullableBool** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**FolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChatResponse

`func NewChatResponse(id string, userId string, title string, chat map[string]interface{}, updatedAt int32, createdAt int32, archived bool, ) *ChatResponse`

NewChatResponse instantiates a new ChatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatResponseWithDefaults

`func NewChatResponseWithDefaults() *ChatResponse`

NewChatResponseWithDefaults instantiates a new ChatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ChatResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChatResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChatResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTitle

`func (o *ChatResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChatResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChatResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetChat

`func (o *ChatResponse) GetChat() map[string]interface{}`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatResponse) GetChatOk() (*map[string]interface{}, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatResponse) SetChat(v map[string]interface{})`

SetChat sets Chat field to given value.


### GetUpdatedAt

`func (o *ChatResponse) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChatResponse) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChatResponse) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ChatResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetShareId

`func (o *ChatResponse) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *ChatResponse) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *ChatResponse) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *ChatResponse) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareIdNil

`func (o *ChatResponse) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *ChatResponse) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil
### GetArchived

`func (o *ChatResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ChatResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ChatResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetPinned

`func (o *ChatResponse) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *ChatResponse) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *ChatResponse) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *ChatResponse) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### SetPinnedNil

`func (o *ChatResponse) SetPinnedNil(b bool)`

 SetPinnedNil sets the value for Pinned to be an explicit nil

### UnsetPinned
`func (o *ChatResponse) UnsetPinned()`

UnsetPinned ensures that no value is present for Pinned, not even an explicit nil
### GetMeta

`func (o *ChatResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChatResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChatResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChatResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetFolderId

`func (o *ChatResponse) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ChatResponse) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ChatResponse) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ChatResponse) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ChatResponse) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ChatResponse) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


