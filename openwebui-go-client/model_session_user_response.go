/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SessionUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionUserResponse{}

// SessionUserResponse struct for SessionUserResponse
type SessionUserResponse struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Role string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
	Token string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresAt NullableInt32 `json:"expires_at,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
}

type _SessionUserResponse SessionUserResponse

// NewSessionUserResponse instantiates a new SessionUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionUserResponse(id string, email string, name string, role string, profileImageUrl string, token string, tokenType string) *SessionUserResponse {
	this := SessionUserResponse{}
	this.Id = id
	this.Email = email
	this.Name = name
	this.Role = role
	this.ProfileImageUrl = profileImageUrl
	this.Token = token
	this.TokenType = tokenType
	return &this
}

// NewSessionUserResponseWithDefaults instantiates a new SessionUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionUserResponseWithDefaults() *SessionUserResponse {
	this := SessionUserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SessionUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SessionUserResponse) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *SessionUserResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SessionUserResponse) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *SessionUserResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SessionUserResponse) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *SessionUserResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SessionUserResponse) SetRole(v string) {
	o.Role = v
}

// GetProfileImageUrl returns the ProfileImageUrl field value
func (o *SessionUserResponse) GetProfileImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileImageUrl
}

// GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetProfileImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileImageUrl, true
}

// SetProfileImageUrl sets field value
func (o *SessionUserResponse) SetProfileImageUrl(v string) {
	o.ProfileImageUrl = v
}

// GetToken returns the Token field value
func (o *SessionUserResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SessionUserResponse) SetToken(v string) {
	o.Token = v
}

// GetTokenType returns the TokenType field value
func (o *SessionUserResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *SessionUserResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *SessionUserResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionUserResponse) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionUserResponse) GetExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SessionUserResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt32 and assigns it to the ExpiresAt field.
func (o *SessionUserResponse) SetExpiresAt(v int32) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *SessionUserResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *SessionUserResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionUserResponse) GetPermissions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionUserResponse) GetPermissionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Permissions) {
		return map[string]interface{}{}, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *SessionUserResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]interface{} and assigns it to the Permissions field.
func (o *SessionUserResponse) SetPermissions(v map[string]interface{}) {
	o.Permissions = v
}

func (o SessionUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["profile_image_url"] = o.ProfileImageUrl
	toSerialize["token"] = o.Token
	toSerialize["token_type"] = o.TokenType
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

func (o *SessionUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email",
		"name",
		"role",
		"profile_image_url",
		"token",
		"token_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSessionUserResponse := _SessionUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSessionUserResponse)

	if err != nil {
		return err
	}

	*o = SessionUserResponse(varSessionUserResponse)

	return err
}

type NullableSessionUserResponse struct {
	value *SessionUserResponse
	isSet bool
}

func (v NullableSessionUserResponse) Get() *SessionUserResponse {
	return v.value
}

func (v *NullableSessionUserResponse) Set(val *SessionUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionUserResponse(val *SessionUserResponse) *NullableSessionUserResponse {
	return &NullableSessionUserResponse{value: val, isSet: true}
}

func (v NullableSessionUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


