# TaskConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TASK_MODEL** | **NullableString** |  | 
**TASK_MODEL_EXTERNAL** | **NullableString** |  | 
**ENABLE_TITLE_GENERATION** | **bool** |  | 
**TITLE_GENERATION_PROMPT_TEMPLATE** | **string** |  | 
**IMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE** | **string** |  | 
**ENABLE_AUTOCOMPLETE_GENERATION** | **bool** |  | 
**AUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH** | **int32** |  | 
**TAGS_GENERATION_PROMPT_TEMPLATE** | **string** |  | 
**ENABLE_TAGS_GENERATION** | **bool** |  | 
**ENABLE_SEARCH_QUERY_GENERATION** | **bool** |  | 
**ENABLE_RETRIEVAL_QUERY_GENERATION** | **bool** |  | 
**QUERY_GENERATION_PROMPT_TEMPLATE** | **string** |  | 
**TOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE** | **string** |  | 

## Methods

### NewTaskConfigForm

`func NewTaskConfigForm(tASKMODEL NullableString, tASKMODELEXTERNAL NullableString, eNABLETITLEGENERATION bool, tITLEGENERATIONPROMPTTEMPLATE string, iMAGEPROMPTGENERATIONPROMPTTEMPLATE string, eNABLEAUTOCOMPLETEGENERATION bool, aUTOCOMPLETEGENERATIONINPUTMAXLENGTH int32, tAGSGENERATIONPROMPTTEMPLATE string, eNABLETAGSGENERATION bool, eNABLESEARCHQUERYGENERATION bool, eNABLERETRIEVALQUERYGENERATION bool, qUERYGENERATIONPROMPTTEMPLATE string, tOOLSFUNCTIONCALLINGPROMPTTEMPLATE string, ) *TaskConfigForm`

NewTaskConfigForm instantiates a new TaskConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskConfigFormWithDefaults

`func NewTaskConfigFormWithDefaults() *TaskConfigForm`

NewTaskConfigFormWithDefaults instantiates a new TaskConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTASK_MODEL

`func (o *TaskConfigForm) GetTASK_MODEL() string`

GetTASK_MODEL returns the TASK_MODEL field if non-nil, zero value otherwise.

### GetTASK_MODELOk

`func (o *TaskConfigForm) GetTASK_MODELOk() (*string, bool)`

GetTASK_MODELOk returns a tuple with the TASK_MODEL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTASK_MODEL

`func (o *TaskConfigForm) SetTASK_MODEL(v string)`

SetTASK_MODEL sets TASK_MODEL field to given value.


### SetTASK_MODELNil

`func (o *TaskConfigForm) SetTASK_MODELNil(b bool)`

 SetTASK_MODELNil sets the value for TASK_MODEL to be an explicit nil

### UnsetTASK_MODEL
`func (o *TaskConfigForm) UnsetTASK_MODEL()`

UnsetTASK_MODEL ensures that no value is present for TASK_MODEL, not even an explicit nil
### GetTASK_MODEL_EXTERNAL

`func (o *TaskConfigForm) GetTASK_MODEL_EXTERNAL() string`

GetTASK_MODEL_EXTERNAL returns the TASK_MODEL_EXTERNAL field if non-nil, zero value otherwise.

### GetTASK_MODEL_EXTERNALOk

`func (o *TaskConfigForm) GetTASK_MODEL_EXTERNALOk() (*string, bool)`

GetTASK_MODEL_EXTERNALOk returns a tuple with the TASK_MODEL_EXTERNAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTASK_MODEL_EXTERNAL

`func (o *TaskConfigForm) SetTASK_MODEL_EXTERNAL(v string)`

SetTASK_MODEL_EXTERNAL sets TASK_MODEL_EXTERNAL field to given value.


### SetTASK_MODEL_EXTERNALNil

`func (o *TaskConfigForm) SetTASK_MODEL_EXTERNALNil(b bool)`

 SetTASK_MODEL_EXTERNALNil sets the value for TASK_MODEL_EXTERNAL to be an explicit nil

### UnsetTASK_MODEL_EXTERNAL
`func (o *TaskConfigForm) UnsetTASK_MODEL_EXTERNAL()`

UnsetTASK_MODEL_EXTERNAL ensures that no value is present for TASK_MODEL_EXTERNAL, not even an explicit nil
### GetENABLE_TITLE_GENERATION

`func (o *TaskConfigForm) GetENABLE_TITLE_GENERATION() bool`

GetENABLE_TITLE_GENERATION returns the ENABLE_TITLE_GENERATION field if non-nil, zero value otherwise.

### GetENABLE_TITLE_GENERATIONOk

`func (o *TaskConfigForm) GetENABLE_TITLE_GENERATIONOk() (*bool, bool)`

GetENABLE_TITLE_GENERATIONOk returns a tuple with the ENABLE_TITLE_GENERATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_TITLE_GENERATION

`func (o *TaskConfigForm) SetENABLE_TITLE_GENERATION(v bool)`

SetENABLE_TITLE_GENERATION sets ENABLE_TITLE_GENERATION field to given value.


### GetTITLE_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) GetTITLE_GENERATION_PROMPT_TEMPLATE() string`

GetTITLE_GENERATION_PROMPT_TEMPLATE returns the TITLE_GENERATION_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetTITLE_GENERATION_PROMPT_TEMPLATEOk

`func (o *TaskConfigForm) GetTITLE_GENERATION_PROMPT_TEMPLATEOk() (*string, bool)`

GetTITLE_GENERATION_PROMPT_TEMPLATEOk returns a tuple with the TITLE_GENERATION_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTITLE_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) SetTITLE_GENERATION_PROMPT_TEMPLATE(v string)`

SetTITLE_GENERATION_PROMPT_TEMPLATE sets TITLE_GENERATION_PROMPT_TEMPLATE field to given value.


### GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE() string`

GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE returns the IMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATEOk

`func (o *TaskConfigForm) GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATEOk() (*string, bool)`

GetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATEOk returns a tuple with the IMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) SetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE(v string)`

SetIMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE sets IMAGE_PROMPT_GENERATION_PROMPT_TEMPLATE field to given value.


### GetENABLE_AUTOCOMPLETE_GENERATION

`func (o *TaskConfigForm) GetENABLE_AUTOCOMPLETE_GENERATION() bool`

GetENABLE_AUTOCOMPLETE_GENERATION returns the ENABLE_AUTOCOMPLETE_GENERATION field if non-nil, zero value otherwise.

### GetENABLE_AUTOCOMPLETE_GENERATIONOk

`func (o *TaskConfigForm) GetENABLE_AUTOCOMPLETE_GENERATIONOk() (*bool, bool)`

GetENABLE_AUTOCOMPLETE_GENERATIONOk returns a tuple with the ENABLE_AUTOCOMPLETE_GENERATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_AUTOCOMPLETE_GENERATION

`func (o *TaskConfigForm) SetENABLE_AUTOCOMPLETE_GENERATION(v bool)`

SetENABLE_AUTOCOMPLETE_GENERATION sets ENABLE_AUTOCOMPLETE_GENERATION field to given value.


### GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH

`func (o *TaskConfigForm) GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH() int32`

GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH returns the AUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH field if non-nil, zero value otherwise.

### GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTHOk

`func (o *TaskConfigForm) GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTHOk() (*int32, bool)`

GetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTHOk returns a tuple with the AUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH

`func (o *TaskConfigForm) SetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH(v int32)`

SetAUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH sets AUTOCOMPLETE_GENERATION_INPUT_MAX_LENGTH field to given value.


### GetTAGS_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) GetTAGS_GENERATION_PROMPT_TEMPLATE() string`

GetTAGS_GENERATION_PROMPT_TEMPLATE returns the TAGS_GENERATION_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetTAGS_GENERATION_PROMPT_TEMPLATEOk

`func (o *TaskConfigForm) GetTAGS_GENERATION_PROMPT_TEMPLATEOk() (*string, bool)`

GetTAGS_GENERATION_PROMPT_TEMPLATEOk returns a tuple with the TAGS_GENERATION_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTAGS_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) SetTAGS_GENERATION_PROMPT_TEMPLATE(v string)`

SetTAGS_GENERATION_PROMPT_TEMPLATE sets TAGS_GENERATION_PROMPT_TEMPLATE field to given value.


### GetENABLE_TAGS_GENERATION

`func (o *TaskConfigForm) GetENABLE_TAGS_GENERATION() bool`

GetENABLE_TAGS_GENERATION returns the ENABLE_TAGS_GENERATION field if non-nil, zero value otherwise.

### GetENABLE_TAGS_GENERATIONOk

`func (o *TaskConfigForm) GetENABLE_TAGS_GENERATIONOk() (*bool, bool)`

GetENABLE_TAGS_GENERATIONOk returns a tuple with the ENABLE_TAGS_GENERATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_TAGS_GENERATION

`func (o *TaskConfigForm) SetENABLE_TAGS_GENERATION(v bool)`

SetENABLE_TAGS_GENERATION sets ENABLE_TAGS_GENERATION field to given value.


### GetENABLE_SEARCH_QUERY_GENERATION

`func (o *TaskConfigForm) GetENABLE_SEARCH_QUERY_GENERATION() bool`

GetENABLE_SEARCH_QUERY_GENERATION returns the ENABLE_SEARCH_QUERY_GENERATION field if non-nil, zero value otherwise.

### GetENABLE_SEARCH_QUERY_GENERATIONOk

`func (o *TaskConfigForm) GetENABLE_SEARCH_QUERY_GENERATIONOk() (*bool, bool)`

GetENABLE_SEARCH_QUERY_GENERATIONOk returns a tuple with the ENABLE_SEARCH_QUERY_GENERATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_SEARCH_QUERY_GENERATION

`func (o *TaskConfigForm) SetENABLE_SEARCH_QUERY_GENERATION(v bool)`

SetENABLE_SEARCH_QUERY_GENERATION sets ENABLE_SEARCH_QUERY_GENERATION field to given value.


### GetENABLE_RETRIEVAL_QUERY_GENERATION

`func (o *TaskConfigForm) GetENABLE_RETRIEVAL_QUERY_GENERATION() bool`

GetENABLE_RETRIEVAL_QUERY_GENERATION returns the ENABLE_RETRIEVAL_QUERY_GENERATION field if non-nil, zero value otherwise.

### GetENABLE_RETRIEVAL_QUERY_GENERATIONOk

`func (o *TaskConfigForm) GetENABLE_RETRIEVAL_QUERY_GENERATIONOk() (*bool, bool)`

GetENABLE_RETRIEVAL_QUERY_GENERATIONOk returns a tuple with the ENABLE_RETRIEVAL_QUERY_GENERATION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_RETRIEVAL_QUERY_GENERATION

`func (o *TaskConfigForm) SetENABLE_RETRIEVAL_QUERY_GENERATION(v bool)`

SetENABLE_RETRIEVAL_QUERY_GENERATION sets ENABLE_RETRIEVAL_QUERY_GENERATION field to given value.


### GetQUERY_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) GetQUERY_GENERATION_PROMPT_TEMPLATE() string`

GetQUERY_GENERATION_PROMPT_TEMPLATE returns the QUERY_GENERATION_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetQUERY_GENERATION_PROMPT_TEMPLATEOk

`func (o *TaskConfigForm) GetQUERY_GENERATION_PROMPT_TEMPLATEOk() (*string, bool)`

GetQUERY_GENERATION_PROMPT_TEMPLATEOk returns a tuple with the QUERY_GENERATION_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQUERY_GENERATION_PROMPT_TEMPLATE

`func (o *TaskConfigForm) SetQUERY_GENERATION_PROMPT_TEMPLATE(v string)`

SetQUERY_GENERATION_PROMPT_TEMPLATE sets QUERY_GENERATION_PROMPT_TEMPLATE field to given value.


### GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE

`func (o *TaskConfigForm) GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE() string`

GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE returns the TOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATEOk

`func (o *TaskConfigForm) GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATEOk() (*string, bool)`

GetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATEOk returns a tuple with the TOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE

`func (o *TaskConfigForm) SetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE(v string)`

SetTOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE sets TOOLS_FUNCTION_CALLING_PROMPT_TEMPLATE field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


