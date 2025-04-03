# BannerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | **string** |  | 
**Dismissible** | **bool** |  | 
**Timestamp** | **int32** |  | 

## Methods

### NewBannerModel

`func NewBannerModel(id string, type_ string, content string, dismissible bool, timestamp int32, ) *BannerModel`

NewBannerModel instantiates a new BannerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannerModelWithDefaults

`func NewBannerModelWithDefaults() *BannerModel`

NewBannerModelWithDefaults instantiates a new BannerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BannerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BannerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BannerModel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BannerModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BannerModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BannerModel) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *BannerModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BannerModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BannerModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BannerModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BannerModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BannerModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *BannerModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BannerModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BannerModel) SetContent(v string)`

SetContent sets Content field to given value.


### GetDismissible

`func (o *BannerModel) GetDismissible() bool`

GetDismissible returns the Dismissible field if non-nil, zero value otherwise.

### GetDismissibleOk

`func (o *BannerModel) GetDismissibleOk() (*bool, bool)`

GetDismissibleOk returns a tuple with the Dismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissible

`func (o *BannerModel) SetDismissible(v bool)`

SetDismissible sets Dismissible field to given value.


### GetTimestamp

`func (o *BannerModel) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BannerModel) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BannerModel) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


