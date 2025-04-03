# CodeInterpreterConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ENABLE_CODE_EXECUTION** | **bool** |  | 
**CODE_EXECUTION_ENGINE** | **string** |  | 
**CODE_EXECUTION_JUPYTER_URL** | **NullableString** |  | 
**CODE_EXECUTION_JUPYTER_AUTH** | **NullableString** |  | 
**CODE_EXECUTION_JUPYTER_AUTH_TOKEN** | **NullableString** |  | 
**CODE_EXECUTION_JUPYTER_AUTH_PASSWORD** | **NullableString** |  | 
**CODE_EXECUTION_JUPYTER_TIMEOUT** | **NullableInt32** |  | 
**ENABLE_CODE_INTERPRETER** | **bool** |  | 
**CODE_INTERPRETER_ENGINE** | **string** |  | 
**CODE_INTERPRETER_PROMPT_TEMPLATE** | **NullableString** |  | 
**CODE_INTERPRETER_JUPYTER_URL** | **NullableString** |  | 
**CODE_INTERPRETER_JUPYTER_AUTH** | **NullableString** |  | 
**CODE_INTERPRETER_JUPYTER_AUTH_TOKEN** | **NullableString** |  | 
**CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD** | **NullableString** |  | 
**CODE_INTERPRETER_JUPYTER_TIMEOUT** | **NullableInt32** |  | 

## Methods

### NewCodeInterpreterConfigForm

`func NewCodeInterpreterConfigForm(eNABLECODEEXECUTION bool, cODEEXECUTIONENGINE string, cODEEXECUTIONJUPYTERURL NullableString, cODEEXECUTIONJUPYTERAUTH NullableString, cODEEXECUTIONJUPYTERAUTHTOKEN NullableString, cODEEXECUTIONJUPYTERAUTHPASSWORD NullableString, cODEEXECUTIONJUPYTERTIMEOUT NullableInt32, eNABLECODEINTERPRETER bool, cODEINTERPRETERENGINE string, cODEINTERPRETERPROMPTTEMPLATE NullableString, cODEINTERPRETERJUPYTERURL NullableString, cODEINTERPRETERJUPYTERAUTH NullableString, cODEINTERPRETERJUPYTERAUTHTOKEN NullableString, cODEINTERPRETERJUPYTERAUTHPASSWORD NullableString, cODEINTERPRETERJUPYTERTIMEOUT NullableInt32, ) *CodeInterpreterConfigForm`

NewCodeInterpreterConfigForm instantiates a new CodeInterpreterConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeInterpreterConfigFormWithDefaults

`func NewCodeInterpreterConfigFormWithDefaults() *CodeInterpreterConfigForm`

NewCodeInterpreterConfigFormWithDefaults instantiates a new CodeInterpreterConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetENABLE_CODE_EXECUTION

`func (o *CodeInterpreterConfigForm) GetENABLE_CODE_EXECUTION() bool`

GetENABLE_CODE_EXECUTION returns the ENABLE_CODE_EXECUTION field if non-nil, zero value otherwise.

### GetENABLE_CODE_EXECUTIONOk

`func (o *CodeInterpreterConfigForm) GetENABLE_CODE_EXECUTIONOk() (*bool, bool)`

GetENABLE_CODE_EXECUTIONOk returns a tuple with the ENABLE_CODE_EXECUTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_CODE_EXECUTION

`func (o *CodeInterpreterConfigForm) SetENABLE_CODE_EXECUTION(v bool)`

SetENABLE_CODE_EXECUTION sets ENABLE_CODE_EXECUTION field to given value.


### GetCODE_EXECUTION_ENGINE

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_ENGINE() string`

GetCODE_EXECUTION_ENGINE returns the CODE_EXECUTION_ENGINE field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_ENGINEOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_ENGINEOk() (*string, bool)`

GetCODE_EXECUTION_ENGINEOk returns a tuple with the CODE_EXECUTION_ENGINE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_ENGINE

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_ENGINE(v string)`

SetCODE_EXECUTION_ENGINE sets CODE_EXECUTION_ENGINE field to given value.


### GetCODE_EXECUTION_JUPYTER_URL

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_URL() string`

GetCODE_EXECUTION_JUPYTER_URL returns the CODE_EXECUTION_JUPYTER_URL field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_JUPYTER_URLOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_URLOk() (*string, bool)`

GetCODE_EXECUTION_JUPYTER_URLOk returns a tuple with the CODE_EXECUTION_JUPYTER_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_JUPYTER_URL

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_URL(v string)`

SetCODE_EXECUTION_JUPYTER_URL sets CODE_EXECUTION_JUPYTER_URL field to given value.


### SetCODE_EXECUTION_JUPYTER_URLNil

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_URLNil(b bool)`

 SetCODE_EXECUTION_JUPYTER_URLNil sets the value for CODE_EXECUTION_JUPYTER_URL to be an explicit nil

### UnsetCODE_EXECUTION_JUPYTER_URL
`func (o *CodeInterpreterConfigForm) UnsetCODE_EXECUTION_JUPYTER_URL()`

UnsetCODE_EXECUTION_JUPYTER_URL ensures that no value is present for CODE_EXECUTION_JUPYTER_URL, not even an explicit nil
### GetCODE_EXECUTION_JUPYTER_AUTH

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH() string`

GetCODE_EXECUTION_JUPYTER_AUTH returns the CODE_EXECUTION_JUPYTER_AUTH field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_JUPYTER_AUTHOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTHOk() (*string, bool)`

GetCODE_EXECUTION_JUPYTER_AUTHOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_JUPYTER_AUTH

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH(v string)`

SetCODE_EXECUTION_JUPYTER_AUTH sets CODE_EXECUTION_JUPYTER_AUTH field to given value.


### SetCODE_EXECUTION_JUPYTER_AUTHNil

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTHNil(b bool)`

 SetCODE_EXECUTION_JUPYTER_AUTHNil sets the value for CODE_EXECUTION_JUPYTER_AUTH to be an explicit nil

### UnsetCODE_EXECUTION_JUPYTER_AUTH
`func (o *CodeInterpreterConfigForm) UnsetCODE_EXECUTION_JUPYTER_AUTH()`

UnsetCODE_EXECUTION_JUPYTER_AUTH ensures that no value is present for CODE_EXECUTION_JUPYTER_AUTH, not even an explicit nil
### GetCODE_EXECUTION_JUPYTER_AUTH_TOKEN

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_TOKEN() string`

GetCODE_EXECUTION_JUPYTER_AUTH_TOKEN returns the CODE_EXECUTION_JUPYTER_AUTH_TOKEN field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_JUPYTER_AUTH_TOKENOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_TOKENOk() (*string, bool)`

GetCODE_EXECUTION_JUPYTER_AUTH_TOKENOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_JUPYTER_AUTH_TOKEN

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_TOKEN(v string)`

SetCODE_EXECUTION_JUPYTER_AUTH_TOKEN sets CODE_EXECUTION_JUPYTER_AUTH_TOKEN field to given value.


### SetCODE_EXECUTION_JUPYTER_AUTH_TOKENNil

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_TOKENNil(b bool)`

 SetCODE_EXECUTION_JUPYTER_AUTH_TOKENNil sets the value for CODE_EXECUTION_JUPYTER_AUTH_TOKEN to be an explicit nil

### UnsetCODE_EXECUTION_JUPYTER_AUTH_TOKEN
`func (o *CodeInterpreterConfigForm) UnsetCODE_EXECUTION_JUPYTER_AUTH_TOKEN()`

UnsetCODE_EXECUTION_JUPYTER_AUTH_TOKEN ensures that no value is present for CODE_EXECUTION_JUPYTER_AUTH_TOKEN, not even an explicit nil
### GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD() string`

GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD returns the CODE_EXECUTION_JUPYTER_AUTH_PASSWORD field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDOk() (*string, bool)`

GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD(v string)`

SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD sets CODE_EXECUTION_JUPYTER_AUTH_PASSWORD field to given value.


### SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDNil

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDNil(b bool)`

 SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDNil sets the value for CODE_EXECUTION_JUPYTER_AUTH_PASSWORD to be an explicit nil

### UnsetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD
`func (o *CodeInterpreterConfigForm) UnsetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD()`

UnsetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD ensures that no value is present for CODE_EXECUTION_JUPYTER_AUTH_PASSWORD, not even an explicit nil
### GetCODE_EXECUTION_JUPYTER_TIMEOUT

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_TIMEOUT() int32`

GetCODE_EXECUTION_JUPYTER_TIMEOUT returns the CODE_EXECUTION_JUPYTER_TIMEOUT field if non-nil, zero value otherwise.

### GetCODE_EXECUTION_JUPYTER_TIMEOUTOk

`func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_TIMEOUTOk() (*int32, bool)`

GetCODE_EXECUTION_JUPYTER_TIMEOUTOk returns a tuple with the CODE_EXECUTION_JUPYTER_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_EXECUTION_JUPYTER_TIMEOUT

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_TIMEOUT(v int32)`

SetCODE_EXECUTION_JUPYTER_TIMEOUT sets CODE_EXECUTION_JUPYTER_TIMEOUT field to given value.


### SetCODE_EXECUTION_JUPYTER_TIMEOUTNil

`func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_TIMEOUTNil(b bool)`

 SetCODE_EXECUTION_JUPYTER_TIMEOUTNil sets the value for CODE_EXECUTION_JUPYTER_TIMEOUT to be an explicit nil

### UnsetCODE_EXECUTION_JUPYTER_TIMEOUT
`func (o *CodeInterpreterConfigForm) UnsetCODE_EXECUTION_JUPYTER_TIMEOUT()`

UnsetCODE_EXECUTION_JUPYTER_TIMEOUT ensures that no value is present for CODE_EXECUTION_JUPYTER_TIMEOUT, not even an explicit nil
### GetENABLE_CODE_INTERPRETER

`func (o *CodeInterpreterConfigForm) GetENABLE_CODE_INTERPRETER() bool`

GetENABLE_CODE_INTERPRETER returns the ENABLE_CODE_INTERPRETER field if non-nil, zero value otherwise.

### GetENABLE_CODE_INTERPRETEROk

`func (o *CodeInterpreterConfigForm) GetENABLE_CODE_INTERPRETEROk() (*bool, bool)`

GetENABLE_CODE_INTERPRETEROk returns a tuple with the ENABLE_CODE_INTERPRETER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_CODE_INTERPRETER

`func (o *CodeInterpreterConfigForm) SetENABLE_CODE_INTERPRETER(v bool)`

SetENABLE_CODE_INTERPRETER sets ENABLE_CODE_INTERPRETER field to given value.


### GetCODE_INTERPRETER_ENGINE

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_ENGINE() string`

GetCODE_INTERPRETER_ENGINE returns the CODE_INTERPRETER_ENGINE field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_ENGINEOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_ENGINEOk() (*string, bool)`

GetCODE_INTERPRETER_ENGINEOk returns a tuple with the CODE_INTERPRETER_ENGINE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_ENGINE

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_ENGINE(v string)`

SetCODE_INTERPRETER_ENGINE sets CODE_INTERPRETER_ENGINE field to given value.


### GetCODE_INTERPRETER_PROMPT_TEMPLATE

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_PROMPT_TEMPLATE() string`

GetCODE_INTERPRETER_PROMPT_TEMPLATE returns the CODE_INTERPRETER_PROMPT_TEMPLATE field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_PROMPT_TEMPLATEOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_PROMPT_TEMPLATEOk() (*string, bool)`

GetCODE_INTERPRETER_PROMPT_TEMPLATEOk returns a tuple with the CODE_INTERPRETER_PROMPT_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_PROMPT_TEMPLATE

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_PROMPT_TEMPLATE(v string)`

SetCODE_INTERPRETER_PROMPT_TEMPLATE sets CODE_INTERPRETER_PROMPT_TEMPLATE field to given value.


### SetCODE_INTERPRETER_PROMPT_TEMPLATENil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_PROMPT_TEMPLATENil(b bool)`

 SetCODE_INTERPRETER_PROMPT_TEMPLATENil sets the value for CODE_INTERPRETER_PROMPT_TEMPLATE to be an explicit nil

### UnsetCODE_INTERPRETER_PROMPT_TEMPLATE
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_PROMPT_TEMPLATE()`

UnsetCODE_INTERPRETER_PROMPT_TEMPLATE ensures that no value is present for CODE_INTERPRETER_PROMPT_TEMPLATE, not even an explicit nil
### GetCODE_INTERPRETER_JUPYTER_URL

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_URL() string`

GetCODE_INTERPRETER_JUPYTER_URL returns the CODE_INTERPRETER_JUPYTER_URL field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_JUPYTER_URLOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_URLOk() (*string, bool)`

GetCODE_INTERPRETER_JUPYTER_URLOk returns a tuple with the CODE_INTERPRETER_JUPYTER_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_JUPYTER_URL

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_URL(v string)`

SetCODE_INTERPRETER_JUPYTER_URL sets CODE_INTERPRETER_JUPYTER_URL field to given value.


### SetCODE_INTERPRETER_JUPYTER_URLNil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_URLNil(b bool)`

 SetCODE_INTERPRETER_JUPYTER_URLNil sets the value for CODE_INTERPRETER_JUPYTER_URL to be an explicit nil

### UnsetCODE_INTERPRETER_JUPYTER_URL
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_JUPYTER_URL()`

UnsetCODE_INTERPRETER_JUPYTER_URL ensures that no value is present for CODE_INTERPRETER_JUPYTER_URL, not even an explicit nil
### GetCODE_INTERPRETER_JUPYTER_AUTH

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH() string`

GetCODE_INTERPRETER_JUPYTER_AUTH returns the CODE_INTERPRETER_JUPYTER_AUTH field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_JUPYTER_AUTHOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTHOk() (*string, bool)`

GetCODE_INTERPRETER_JUPYTER_AUTHOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_JUPYTER_AUTH

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH(v string)`

SetCODE_INTERPRETER_JUPYTER_AUTH sets CODE_INTERPRETER_JUPYTER_AUTH field to given value.


### SetCODE_INTERPRETER_JUPYTER_AUTHNil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTHNil(b bool)`

 SetCODE_INTERPRETER_JUPYTER_AUTHNil sets the value for CODE_INTERPRETER_JUPYTER_AUTH to be an explicit nil

### UnsetCODE_INTERPRETER_JUPYTER_AUTH
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_JUPYTER_AUTH()`

UnsetCODE_INTERPRETER_JUPYTER_AUTH ensures that no value is present for CODE_INTERPRETER_JUPYTER_AUTH, not even an explicit nil
### GetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN() string`

GetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN returns the CODE_INTERPRETER_JUPYTER_AUTH_TOKEN field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_JUPYTER_AUTH_TOKENOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_TOKENOk() (*string, bool)`

GetCODE_INTERPRETER_JUPYTER_AUTH_TOKENOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN(v string)`

SetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN sets CODE_INTERPRETER_JUPYTER_AUTH_TOKEN field to given value.


### SetCODE_INTERPRETER_JUPYTER_AUTH_TOKENNil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_TOKENNil(b bool)`

 SetCODE_INTERPRETER_JUPYTER_AUTH_TOKENNil sets the value for CODE_INTERPRETER_JUPYTER_AUTH_TOKEN to be an explicit nil

### UnsetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN()`

UnsetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN ensures that no value is present for CODE_INTERPRETER_JUPYTER_AUTH_TOKEN, not even an explicit nil
### GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD() string`

GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD returns the CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDOk() (*string, bool)`

GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD(v string)`

SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD sets CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD field to given value.


### SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDNil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDNil(b bool)`

 SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDNil sets the value for CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD to be an explicit nil

### UnsetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD()`

UnsetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD ensures that no value is present for CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD, not even an explicit nil
### GetCODE_INTERPRETER_JUPYTER_TIMEOUT

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_TIMEOUT() int32`

GetCODE_INTERPRETER_JUPYTER_TIMEOUT returns the CODE_INTERPRETER_JUPYTER_TIMEOUT field if non-nil, zero value otherwise.

### GetCODE_INTERPRETER_JUPYTER_TIMEOUTOk

`func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_TIMEOUTOk() (*int32, bool)`

GetCODE_INTERPRETER_JUPYTER_TIMEOUTOk returns a tuple with the CODE_INTERPRETER_JUPYTER_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCODE_INTERPRETER_JUPYTER_TIMEOUT

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_TIMEOUT(v int32)`

SetCODE_INTERPRETER_JUPYTER_TIMEOUT sets CODE_INTERPRETER_JUPYTER_TIMEOUT field to given value.


### SetCODE_INTERPRETER_JUPYTER_TIMEOUTNil

`func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_TIMEOUTNil(b bool)`

 SetCODE_INTERPRETER_JUPYTER_TIMEOUTNil sets the value for CODE_INTERPRETER_JUPYTER_TIMEOUT to be an explicit nil

### UnsetCODE_INTERPRETER_JUPYTER_TIMEOUT
`func (o *CodeInterpreterConfigForm) UnsetCODE_INTERPRETER_JUPYTER_TIMEOUT()`

UnsetCODE_INTERPRETER_JUPYTER_TIMEOUT ensures that no value is present for CODE_INTERPRETER_JUPYTER_TIMEOUT, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


