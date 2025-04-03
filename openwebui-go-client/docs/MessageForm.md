# MessageForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMessageForm

`func NewMessageForm(content string, ) *MessageForm`

NewMessageForm instantiates a new MessageForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageFormWithDefaults

`func NewMessageFormWithDefaults() *MessageForm`

NewMessageFormWithDefaults instantiates a new MessageForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MessageForm) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageForm) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageForm) SetContent(v string)`

SetContent sets Content field to given value.


### GetParentId

`func (o *MessageForm) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MessageForm) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MessageForm) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MessageForm) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *MessageForm) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *MessageForm) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetData

`func (o *MessageForm) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessageForm) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MessageForm) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *MessageForm) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *MessageForm) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *MessageForm) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *MessageForm) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MessageForm) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MessageForm) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MessageForm) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MessageForm) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MessageForm) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


