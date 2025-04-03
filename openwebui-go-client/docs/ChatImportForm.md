# ChatImportForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | **map[string]interface{}** |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Pinned** | Pointer to **NullableBool** |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChatImportForm

`func NewChatImportForm(chat map[string]interface{}, ) *ChatImportForm`

NewChatImportForm instantiates a new ChatImportForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatImportFormWithDefaults

`func NewChatImportFormWithDefaults() *ChatImportForm`

NewChatImportFormWithDefaults instantiates a new ChatImportForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatImportForm) GetChat() map[string]interface{}`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatImportForm) GetChatOk() (*map[string]interface{}, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatImportForm) SetChat(v map[string]interface{})`

SetChat sets Chat field to given value.


### GetMeta

`func (o *ChatImportForm) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChatImportForm) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChatImportForm) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChatImportForm) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ChatImportForm) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ChatImportForm) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetPinned

`func (o *ChatImportForm) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *ChatImportForm) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *ChatImportForm) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *ChatImportForm) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### SetPinnedNil

`func (o *ChatImportForm) SetPinnedNil(b bool)`

 SetPinnedNil sets the value for Pinned to be an explicit nil

### UnsetPinned
`func (o *ChatImportForm) UnsetPinned()`

UnsetPinned ensures that no value is present for Pinned, not even an explicit nil
### GetFolderId

`func (o *ChatImportForm) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ChatImportForm) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ChatImportForm) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ChatImportForm) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ChatImportForm) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ChatImportForm) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


