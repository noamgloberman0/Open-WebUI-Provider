# PushModelForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Insecure** | Pointer to **NullableBool** |  | [optional] 
**Stream** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPushModelForm

`func NewPushModelForm(name string, ) *PushModelForm`

NewPushModelForm instantiates a new PushModelForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushModelFormWithDefaults

`func NewPushModelFormWithDefaults() *PushModelForm`

NewPushModelFormWithDefaults instantiates a new PushModelForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PushModelForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushModelForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushModelForm) SetName(v string)`

SetName sets Name field to given value.


### GetInsecure

`func (o *PushModelForm) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *PushModelForm) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *PushModelForm) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *PushModelForm) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### SetInsecureNil

`func (o *PushModelForm) SetInsecureNil(b bool)`

 SetInsecureNil sets the value for Insecure to be an explicit nil

### UnsetInsecure
`func (o *PushModelForm) UnsetInsecure()`

UnsetInsecure ensures that no value is present for Insecure, not even an explicit nil
### GetStream

`func (o *PushModelForm) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *PushModelForm) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *PushModelForm) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *PushModelForm) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *PushModelForm) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *PushModelForm) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


