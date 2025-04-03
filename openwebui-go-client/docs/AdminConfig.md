# AdminConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SHOW_ADMIN_DETAILS** | **bool** |  | 
**WEBUI_URL** | **string** |  | 
**ENABLE_SIGNUP** | **bool** |  | 
**ENABLE_API_KEY** | **bool** |  | 
**ENABLE_API_KEY_ENDPOINT_RESTRICTIONS** | **bool** |  | 
**API_KEY_ALLOWED_ENDPOINTS** | **string** |  | 
**ENABLE_CHANNELS** | **bool** |  | 
**DEFAULT_USER_ROLE** | **string** |  | 
**JWT_EXPIRES_IN** | **string** |  | 
**ENABLE_COMMUNITY_SHARING** | **bool** |  | 
**ENABLE_MESSAGE_RATING** | **bool** |  | 

## Methods

### NewAdminConfig

`func NewAdminConfig(sHOWADMINDETAILS bool, wEBUIURL string, eNABLESIGNUP bool, eNABLEAPIKEY bool, eNABLEAPIKEYENDPOINTRESTRICTIONS bool, aPIKEYALLOWEDENDPOINTS string, eNABLECHANNELS bool, dEFAULTUSERROLE string, jWTEXPIRESIN string, eNABLECOMMUNITYSHARING bool, eNABLEMESSAGERATING bool, ) *AdminConfig`

NewAdminConfig instantiates a new AdminConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminConfigWithDefaults

`func NewAdminConfigWithDefaults() *AdminConfig`

NewAdminConfigWithDefaults instantiates a new AdminConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSHOW_ADMIN_DETAILS

`func (o *AdminConfig) GetSHOW_ADMIN_DETAILS() bool`

GetSHOW_ADMIN_DETAILS returns the SHOW_ADMIN_DETAILS field if non-nil, zero value otherwise.

### GetSHOW_ADMIN_DETAILSOk

`func (o *AdminConfig) GetSHOW_ADMIN_DETAILSOk() (*bool, bool)`

GetSHOW_ADMIN_DETAILSOk returns a tuple with the SHOW_ADMIN_DETAILS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHOW_ADMIN_DETAILS

`func (o *AdminConfig) SetSHOW_ADMIN_DETAILS(v bool)`

SetSHOW_ADMIN_DETAILS sets SHOW_ADMIN_DETAILS field to given value.


### GetWEBUI_URL

`func (o *AdminConfig) GetWEBUI_URL() string`

GetWEBUI_URL returns the WEBUI_URL field if non-nil, zero value otherwise.

### GetWEBUI_URLOk

`func (o *AdminConfig) GetWEBUI_URLOk() (*string, bool)`

GetWEBUI_URLOk returns a tuple with the WEBUI_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWEBUI_URL

`func (o *AdminConfig) SetWEBUI_URL(v string)`

SetWEBUI_URL sets WEBUI_URL field to given value.


### GetENABLE_SIGNUP

`func (o *AdminConfig) GetENABLE_SIGNUP() bool`

GetENABLE_SIGNUP returns the ENABLE_SIGNUP field if non-nil, zero value otherwise.

### GetENABLE_SIGNUPOk

`func (o *AdminConfig) GetENABLE_SIGNUPOk() (*bool, bool)`

GetENABLE_SIGNUPOk returns a tuple with the ENABLE_SIGNUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_SIGNUP

`func (o *AdminConfig) SetENABLE_SIGNUP(v bool)`

SetENABLE_SIGNUP sets ENABLE_SIGNUP field to given value.


### GetENABLE_API_KEY

`func (o *AdminConfig) GetENABLE_API_KEY() bool`

GetENABLE_API_KEY returns the ENABLE_API_KEY field if non-nil, zero value otherwise.

### GetENABLE_API_KEYOk

`func (o *AdminConfig) GetENABLE_API_KEYOk() (*bool, bool)`

GetENABLE_API_KEYOk returns a tuple with the ENABLE_API_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_API_KEY

`func (o *AdminConfig) SetENABLE_API_KEY(v bool)`

SetENABLE_API_KEY sets ENABLE_API_KEY field to given value.


### GetENABLE_API_KEY_ENDPOINT_RESTRICTIONS

`func (o *AdminConfig) GetENABLE_API_KEY_ENDPOINT_RESTRICTIONS() bool`

GetENABLE_API_KEY_ENDPOINT_RESTRICTIONS returns the ENABLE_API_KEY_ENDPOINT_RESTRICTIONS field if non-nil, zero value otherwise.

### GetENABLE_API_KEY_ENDPOINT_RESTRICTIONSOk

`func (o *AdminConfig) GetENABLE_API_KEY_ENDPOINT_RESTRICTIONSOk() (*bool, bool)`

GetENABLE_API_KEY_ENDPOINT_RESTRICTIONSOk returns a tuple with the ENABLE_API_KEY_ENDPOINT_RESTRICTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_API_KEY_ENDPOINT_RESTRICTIONS

`func (o *AdminConfig) SetENABLE_API_KEY_ENDPOINT_RESTRICTIONS(v bool)`

SetENABLE_API_KEY_ENDPOINT_RESTRICTIONS sets ENABLE_API_KEY_ENDPOINT_RESTRICTIONS field to given value.


### GetAPI_KEY_ALLOWED_ENDPOINTS

`func (o *AdminConfig) GetAPI_KEY_ALLOWED_ENDPOINTS() string`

GetAPI_KEY_ALLOWED_ENDPOINTS returns the API_KEY_ALLOWED_ENDPOINTS field if non-nil, zero value otherwise.

### GetAPI_KEY_ALLOWED_ENDPOINTSOk

`func (o *AdminConfig) GetAPI_KEY_ALLOWED_ENDPOINTSOk() (*string, bool)`

GetAPI_KEY_ALLOWED_ENDPOINTSOk returns a tuple with the API_KEY_ALLOWED_ENDPOINTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_KEY_ALLOWED_ENDPOINTS

`func (o *AdminConfig) SetAPI_KEY_ALLOWED_ENDPOINTS(v string)`

SetAPI_KEY_ALLOWED_ENDPOINTS sets API_KEY_ALLOWED_ENDPOINTS field to given value.


### GetENABLE_CHANNELS

`func (o *AdminConfig) GetENABLE_CHANNELS() bool`

GetENABLE_CHANNELS returns the ENABLE_CHANNELS field if non-nil, zero value otherwise.

### GetENABLE_CHANNELSOk

`func (o *AdminConfig) GetENABLE_CHANNELSOk() (*bool, bool)`

GetENABLE_CHANNELSOk returns a tuple with the ENABLE_CHANNELS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_CHANNELS

`func (o *AdminConfig) SetENABLE_CHANNELS(v bool)`

SetENABLE_CHANNELS sets ENABLE_CHANNELS field to given value.


### GetDEFAULT_USER_ROLE

`func (o *AdminConfig) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *AdminConfig) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *AdminConfig) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.


### GetJWT_EXPIRES_IN

`func (o *AdminConfig) GetJWT_EXPIRES_IN() string`

GetJWT_EXPIRES_IN returns the JWT_EXPIRES_IN field if non-nil, zero value otherwise.

### GetJWT_EXPIRES_INOk

`func (o *AdminConfig) GetJWT_EXPIRES_INOk() (*string, bool)`

GetJWT_EXPIRES_INOk returns a tuple with the JWT_EXPIRES_IN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJWT_EXPIRES_IN

`func (o *AdminConfig) SetJWT_EXPIRES_IN(v string)`

SetJWT_EXPIRES_IN sets JWT_EXPIRES_IN field to given value.


### GetENABLE_COMMUNITY_SHARING

`func (o *AdminConfig) GetENABLE_COMMUNITY_SHARING() bool`

GetENABLE_COMMUNITY_SHARING returns the ENABLE_COMMUNITY_SHARING field if non-nil, zero value otherwise.

### GetENABLE_COMMUNITY_SHARINGOk

`func (o *AdminConfig) GetENABLE_COMMUNITY_SHARINGOk() (*bool, bool)`

GetENABLE_COMMUNITY_SHARINGOk returns a tuple with the ENABLE_COMMUNITY_SHARING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_COMMUNITY_SHARING

`func (o *AdminConfig) SetENABLE_COMMUNITY_SHARING(v bool)`

SetENABLE_COMMUNITY_SHARING sets ENABLE_COMMUNITY_SHARING field to given value.


### GetENABLE_MESSAGE_RATING

`func (o *AdminConfig) GetENABLE_MESSAGE_RATING() bool`

GetENABLE_MESSAGE_RATING returns the ENABLE_MESSAGE_RATING field if non-nil, zero value otherwise.

### GetENABLE_MESSAGE_RATINGOk

`func (o *AdminConfig) GetENABLE_MESSAGE_RATINGOk() (*bool, bool)`

GetENABLE_MESSAGE_RATINGOk returns a tuple with the ENABLE_MESSAGE_RATING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_MESSAGE_RATING

`func (o *AdminConfig) SetENABLE_MESSAGE_RATING(v bool)`

SetENABLE_MESSAGE_RATING sets ENABLE_MESSAGE_RATING field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


