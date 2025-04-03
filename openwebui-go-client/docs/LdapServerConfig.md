# LdapServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Host** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**AttributeForMail** | Pointer to **string** |  | [optional] [default to "mail"]
**AttributeForUsername** | Pointer to **string** |  | [optional] [default to "uid"]
**AppDn** | **string** |  | 
**AppDnPassword** | **string** |  | 
**SearchBase** | **string** |  | 
**SearchFilters** | Pointer to **string** |  | [optional] [default to ""]
**UseTls** | Pointer to **bool** |  | [optional] [default to true]
**CertificatePath** | Pointer to **NullableString** |  | [optional] 
**Ciphers** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLdapServerConfig

`func NewLdapServerConfig(label string, host string, appDn string, appDnPassword string, searchBase string, ) *LdapServerConfig`

NewLdapServerConfig instantiates a new LdapServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapServerConfigWithDefaults

`func NewLdapServerConfigWithDefaults() *LdapServerConfig`

NewLdapServerConfigWithDefaults instantiates a new LdapServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LdapServerConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LdapServerConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LdapServerConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetHost

`func (o *LdapServerConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LdapServerConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LdapServerConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *LdapServerConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapServerConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapServerConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LdapServerConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LdapServerConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LdapServerConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetAttributeForMail

`func (o *LdapServerConfig) GetAttributeForMail() string`

GetAttributeForMail returns the AttributeForMail field if non-nil, zero value otherwise.

### GetAttributeForMailOk

`func (o *LdapServerConfig) GetAttributeForMailOk() (*string, bool)`

GetAttributeForMailOk returns a tuple with the AttributeForMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeForMail

`func (o *LdapServerConfig) SetAttributeForMail(v string)`

SetAttributeForMail sets AttributeForMail field to given value.

### HasAttributeForMail

`func (o *LdapServerConfig) HasAttributeForMail() bool`

HasAttributeForMail returns a boolean if a field has been set.

### GetAttributeForUsername

`func (o *LdapServerConfig) GetAttributeForUsername() string`

GetAttributeForUsername returns the AttributeForUsername field if non-nil, zero value otherwise.

### GetAttributeForUsernameOk

`func (o *LdapServerConfig) GetAttributeForUsernameOk() (*string, bool)`

GetAttributeForUsernameOk returns a tuple with the AttributeForUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeForUsername

`func (o *LdapServerConfig) SetAttributeForUsername(v string)`

SetAttributeForUsername sets AttributeForUsername field to given value.

### HasAttributeForUsername

`func (o *LdapServerConfig) HasAttributeForUsername() bool`

HasAttributeForUsername returns a boolean if a field has been set.

### GetAppDn

`func (o *LdapServerConfig) GetAppDn() string`

GetAppDn returns the AppDn field if non-nil, zero value otherwise.

### GetAppDnOk

`func (o *LdapServerConfig) GetAppDnOk() (*string, bool)`

GetAppDnOk returns a tuple with the AppDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDn

`func (o *LdapServerConfig) SetAppDn(v string)`

SetAppDn sets AppDn field to given value.


### GetAppDnPassword

`func (o *LdapServerConfig) GetAppDnPassword() string`

GetAppDnPassword returns the AppDnPassword field if non-nil, zero value otherwise.

### GetAppDnPasswordOk

`func (o *LdapServerConfig) GetAppDnPasswordOk() (*string, bool)`

GetAppDnPasswordOk returns a tuple with the AppDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDnPassword

`func (o *LdapServerConfig) SetAppDnPassword(v string)`

SetAppDnPassword sets AppDnPassword field to given value.


### GetSearchBase

`func (o *LdapServerConfig) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *LdapServerConfig) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *LdapServerConfig) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetSearchFilters

`func (o *LdapServerConfig) GetSearchFilters() string`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *LdapServerConfig) GetSearchFiltersOk() (*string, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *LdapServerConfig) SetSearchFilters(v string)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *LdapServerConfig) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### GetUseTls

`func (o *LdapServerConfig) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *LdapServerConfig) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *LdapServerConfig) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *LdapServerConfig) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetCertificatePath

`func (o *LdapServerConfig) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *LdapServerConfig) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *LdapServerConfig) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *LdapServerConfig) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### SetCertificatePathNil

`func (o *LdapServerConfig) SetCertificatePathNil(b bool)`

 SetCertificatePathNil sets the value for CertificatePath to be an explicit nil

### UnsetCertificatePath
`func (o *LdapServerConfig) UnsetCertificatePath()`

UnsetCertificatePath ensures that no value is present for CertificatePath, not even an explicit nil
### GetCiphers

`func (o *LdapServerConfig) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *LdapServerConfig) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *LdapServerConfig) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *LdapServerConfig) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### SetCiphersNil

`func (o *LdapServerConfig) SetCiphersNil(b bool)`

 SetCiphersNil sets the value for Ciphers to be an explicit nil

### UnsetCiphers
`func (o *LdapServerConfig) UnsetCiphers()`

UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


