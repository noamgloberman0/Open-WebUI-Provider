# GroupUpdateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupUpdateForm

`func NewGroupUpdateForm(name string, description string, ) *GroupUpdateForm`

NewGroupUpdateForm instantiates a new GroupUpdateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateFormWithDefaults

`func NewGroupUpdateFormWithDefaults() *GroupUpdateForm`

NewGroupUpdateFormWithDefaults instantiates a new GroupUpdateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupUpdateForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdateForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdateForm) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupUpdateForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupUpdateForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupUpdateForm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissions

`func (o *GroupUpdateForm) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupUpdateForm) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupUpdateForm) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupUpdateForm) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *GroupUpdateForm) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GroupUpdateForm) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetUserIds

`func (o *GroupUpdateForm) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GroupUpdateForm) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GroupUpdateForm) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GroupUpdateForm) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *GroupUpdateForm) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *GroupUpdateForm) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


