# Automatic1111ConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AUTOMATIC1111BASEURL** | **string** |  | 
**AUTOMATIC1111APIAUTH** | **string** |  | 
**AUTOMATIC1111CFGSCALE** | [**NullableAutomatic1111CfgScale**](Automatic1111CfgScale.md) |  | 
**AUTOMATIC1111SAMPLER** | **NullableString** |  | 
**AUTOMATIC1111SCHEDULER** | **NullableString** |  | 

## Methods

### NewAutomatic1111ConfigForm

`func NewAutomatic1111ConfigForm(aUTOMATIC1111BASEURL string, aUTOMATIC1111APIAUTH string, aUTOMATIC1111CFGSCALE NullableAutomatic1111CfgScale, aUTOMATIC1111SAMPLER NullableString, aUTOMATIC1111SCHEDULER NullableString, ) *Automatic1111ConfigForm`

NewAutomatic1111ConfigForm instantiates a new Automatic1111ConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomatic1111ConfigFormWithDefaults

`func NewAutomatic1111ConfigFormWithDefaults() *Automatic1111ConfigForm`

NewAutomatic1111ConfigFormWithDefaults instantiates a new Automatic1111ConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAUTOMATIC1111BASEURL

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111BASEURL() string`

GetAUTOMATIC1111BASEURL returns the AUTOMATIC1111BASEURL field if non-nil, zero value otherwise.

### GetAUTOMATIC1111BASEURLOk

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111BASEURLOk() (*string, bool)`

GetAUTOMATIC1111BASEURLOk returns a tuple with the AUTOMATIC1111BASEURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATIC1111BASEURL

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111BASEURL(v string)`

SetAUTOMATIC1111BASEURL sets AUTOMATIC1111BASEURL field to given value.


### GetAUTOMATIC1111APIAUTH

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111APIAUTH() string`

GetAUTOMATIC1111APIAUTH returns the AUTOMATIC1111APIAUTH field if non-nil, zero value otherwise.

### GetAUTOMATIC1111APIAUTHOk

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111APIAUTHOk() (*string, bool)`

GetAUTOMATIC1111APIAUTHOk returns a tuple with the AUTOMATIC1111APIAUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATIC1111APIAUTH

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111APIAUTH(v string)`

SetAUTOMATIC1111APIAUTH sets AUTOMATIC1111APIAUTH field to given value.


### GetAUTOMATIC1111CFGSCALE

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111CFGSCALE() Automatic1111CfgScale`

GetAUTOMATIC1111CFGSCALE returns the AUTOMATIC1111CFGSCALE field if non-nil, zero value otherwise.

### GetAUTOMATIC1111CFGSCALEOk

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111CFGSCALEOk() (*Automatic1111CfgScale, bool)`

GetAUTOMATIC1111CFGSCALEOk returns a tuple with the AUTOMATIC1111CFGSCALE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATIC1111CFGSCALE

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111CFGSCALE(v Automatic1111CfgScale)`

SetAUTOMATIC1111CFGSCALE sets AUTOMATIC1111CFGSCALE field to given value.


### SetAUTOMATIC1111CFGSCALENil

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111CFGSCALENil(b bool)`

 SetAUTOMATIC1111CFGSCALENil sets the value for AUTOMATIC1111CFGSCALE to be an explicit nil

### UnsetAUTOMATIC1111CFGSCALE
`func (o *Automatic1111ConfigForm) UnsetAUTOMATIC1111CFGSCALE()`

UnsetAUTOMATIC1111CFGSCALE ensures that no value is present for AUTOMATIC1111CFGSCALE, not even an explicit nil
### GetAUTOMATIC1111SAMPLER

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111SAMPLER() string`

GetAUTOMATIC1111SAMPLER returns the AUTOMATIC1111SAMPLER field if non-nil, zero value otherwise.

### GetAUTOMATIC1111SAMPLEROk

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111SAMPLEROk() (*string, bool)`

GetAUTOMATIC1111SAMPLEROk returns a tuple with the AUTOMATIC1111SAMPLER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATIC1111SAMPLER

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111SAMPLER(v string)`

SetAUTOMATIC1111SAMPLER sets AUTOMATIC1111SAMPLER field to given value.


### SetAUTOMATIC1111SAMPLERNil

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111SAMPLERNil(b bool)`

 SetAUTOMATIC1111SAMPLERNil sets the value for AUTOMATIC1111SAMPLER to be an explicit nil

### UnsetAUTOMATIC1111SAMPLER
`func (o *Automatic1111ConfigForm) UnsetAUTOMATIC1111SAMPLER()`

UnsetAUTOMATIC1111SAMPLER ensures that no value is present for AUTOMATIC1111SAMPLER, not even an explicit nil
### GetAUTOMATIC1111SCHEDULER

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111SCHEDULER() string`

GetAUTOMATIC1111SCHEDULER returns the AUTOMATIC1111SCHEDULER field if non-nil, zero value otherwise.

### GetAUTOMATIC1111SCHEDULEROk

`func (o *Automatic1111ConfigForm) GetAUTOMATIC1111SCHEDULEROk() (*string, bool)`

GetAUTOMATIC1111SCHEDULEROk returns a tuple with the AUTOMATIC1111SCHEDULER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATIC1111SCHEDULER

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111SCHEDULER(v string)`

SetAUTOMATIC1111SCHEDULER sets AUTOMATIC1111SCHEDULER field to given value.


### SetAUTOMATIC1111SCHEDULERNil

`func (o *Automatic1111ConfigForm) SetAUTOMATIC1111SCHEDULERNil(b bool)`

 SetAUTOMATIC1111SCHEDULERNil sets the value for AUTOMATIC1111SCHEDULER to be an explicit nil

### UnsetAUTOMATIC1111SCHEDULER
`func (o *Automatic1111ConfigForm) UnsetAUTOMATIC1111SCHEDULER()`

UnsetAUTOMATIC1111SCHEDULER ensures that no value is present for AUTOMATIC1111SCHEDULER, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


