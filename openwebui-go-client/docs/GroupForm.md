# GroupForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGroupForm

`func NewGroupForm(name string, description string, ) *GroupForm`

NewGroupForm instantiates a new GroupForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupFormWithDefaults

`func NewGroupFormWithDefaults() *GroupForm`

NewGroupFormWithDefaults instantiates a new GroupForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupForm) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupForm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissions

`func (o *GroupForm) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupForm) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupForm) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupForm) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *GroupForm) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GroupForm) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


