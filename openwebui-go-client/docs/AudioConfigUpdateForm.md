# AudioConfigUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tts** | [**TTSConfigForm**](TTSConfigForm.md) |  | 
**Stt** | [**STTConfigForm**](STTConfigForm.md) |  | 

## Methods

### NewAudioConfigUpdateForm

`func NewAudioConfigUpdateForm(tts TTSConfigForm, stt STTConfigForm, ) *AudioConfigUpdateForm`

NewAudioConfigUpdateForm instantiates a new AudioConfigUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioConfigUpdateFormWithDefaults

`func NewAudioConfigUpdateFormWithDefaults() *AudioConfigUpdateForm`

NewAudioConfigUpdateFormWithDefaults instantiates a new AudioConfigUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTts

`func (o *AudioConfigUpdateForm) GetTts() TTSConfigForm`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *AudioConfigUpdateForm) GetTtsOk() (*TTSConfigForm, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *AudioConfigUpdateForm) SetTts(v TTSConfigForm)`

SetTts sets Tts field to given value.


### GetStt

`func (o *AudioConfigUpdateForm) GetStt() STTConfigForm`

GetStt returns the Stt field if non-nil, zero value otherwise.

### GetSttOk

`func (o *AudioConfigUpdateForm) GetSttOk() (*STTConfigForm, bool)`

GetSttOk returns a tuple with the Stt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStt

`func (o *AudioConfigUpdateForm) SetStt(v STTConfigForm)`

SetStt sets Stt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


