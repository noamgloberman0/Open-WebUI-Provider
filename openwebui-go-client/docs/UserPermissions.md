# UserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspace** | [**WorkspacePermissions**](WorkspacePermissions.md) |  | 
**Chat** | [**ChatPermissions**](ChatPermissions.md) |  | 
**Features** | [**FeaturesPermissions**](FeaturesPermissions.md) |  | 

## Methods

### NewUserPermissions

`func NewUserPermissions(workspace WorkspacePermissions, chat ChatPermissions, features FeaturesPermissions, ) *UserPermissions`

NewUserPermissions instantiates a new UserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionsWithDefaults

`func NewUserPermissionsWithDefaults() *UserPermissions`

NewUserPermissionsWithDefaults instantiates a new UserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspace

`func (o *UserPermissions) GetWorkspace() WorkspacePermissions`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *UserPermissions) GetWorkspaceOk() (*WorkspacePermissions, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *UserPermissions) SetWorkspace(v WorkspacePermissions)`

SetWorkspace sets Workspace field to given value.


### GetChat

`func (o *UserPermissions) GetChat() ChatPermissions`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *UserPermissions) GetChatOk() (*ChatPermissions, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *UserPermissions) SetChat(v ChatPermissions)`

SetChat sets Chat field to given value.


### GetFeatures

`func (o *UserPermissions) GetFeatures() FeaturesPermissions`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UserPermissions) GetFeaturesOk() (*FeaturesPermissions, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UserPermissions) SetFeatures(v FeaturesPermissions)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


