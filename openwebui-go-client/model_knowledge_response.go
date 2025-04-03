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

// checks if the KnowledgeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnowledgeResponse{}

// KnowledgeResponse struct for KnowledgeResponse
type KnowledgeResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Data map[string]interface{} `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	AccessControl map[string]interface{} `json:"access_control,omitempty"`
	CreatedAt int32 `json:"created_at"`
	UpdatedAt int32 `json:"updated_at"`
	Files []KnowledgeResponseFilesInner `json:"files,omitempty"`
}

type _KnowledgeResponse KnowledgeResponse

// NewKnowledgeResponse instantiates a new KnowledgeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnowledgeResponse(id string, userId string, name string, description string, createdAt int32, updatedAt int32) *KnowledgeResponse {
	this := KnowledgeResponse{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Description = description
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewKnowledgeResponseWithDefaults instantiates a new KnowledgeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnowledgeResponseWithDefaults() *KnowledgeResponse {
	this := KnowledgeResponse{}
	return &this
}

// GetId returns the Id field value
func (o *KnowledgeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KnowledgeResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *KnowledgeResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *KnowledgeResponse) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *KnowledgeResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KnowledgeResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *KnowledgeResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *KnowledgeResponse) SetDescription(v string) {
	o.Description = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KnowledgeResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *KnowledgeResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeResponse) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *KnowledgeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *KnowledgeResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeResponse) GetAccessControl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeResponse) GetAccessControlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return map[string]interface{}{}, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *KnowledgeResponse) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given map[string]interface{} and assigns it to the AccessControl field.
func (o *KnowledgeResponse) SetAccessControl(v map[string]interface{}) {
	o.AccessControl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *KnowledgeResponse) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *KnowledgeResponse) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *KnowledgeResponse) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *KnowledgeResponse) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *KnowledgeResponse) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeResponse) GetFiles() []KnowledgeResponseFilesInner {
	if o == nil {
		var ret []KnowledgeResponseFilesInner
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeResponse) GetFilesOk() ([]KnowledgeResponseFilesInner, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *KnowledgeResponse) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []KnowledgeResponseFilesInner and assigns it to the Files field.
func (o *KnowledgeResponse) SetFiles(v []KnowledgeResponseFilesInner) {
	o.Files = v
}

func (o KnowledgeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnowledgeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.AccessControl != nil {
		toSerialize["access_control"] = o.AccessControl
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

func (o *KnowledgeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"name",
		"description",
		"created_at",
		"updated_at",
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

	varKnowledgeResponse := _KnowledgeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKnowledgeResponse)

	if err != nil {
		return err
	}

	*o = KnowledgeResponse(varKnowledgeResponse)

	return err
}

type NullableKnowledgeResponse struct {
	value *KnowledgeResponse
	isSet bool
}

func (v NullableKnowledgeResponse) Get() *KnowledgeResponse {
	return v.value
}

func (v *NullableKnowledgeResponse) Set(val *KnowledgeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgeResponse(val *KnowledgeResponse) *NullableKnowledgeResponse {
	return &NullableKnowledgeResponse{value: val, isSet: true}
}

func (v NullableKnowledgeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


