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

// checks if the GenerateImageForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateImageForm{}

// GenerateImageForm struct for GenerateImageForm
type GenerateImageForm struct {
	Model NullableString `json:"model,omitempty"`
	Prompt string `json:"prompt"`
	Size NullableString `json:"size,omitempty"`
	N *int32 `json:"n,omitempty"`
	NegativePrompt NullableString `json:"negative_prompt,omitempty"`
}

type _GenerateImageForm GenerateImageForm

// NewGenerateImageForm instantiates a new GenerateImageForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateImageForm(prompt string) *GenerateImageForm {
	this := GenerateImageForm{}
	this.Prompt = prompt
	var n int32 = 1
	this.N = &n
	return &this
}

// NewGenerateImageFormWithDefaults instantiates a new GenerateImageForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateImageFormWithDefaults() *GenerateImageForm {
	this := GenerateImageForm{}
	var n int32 = 1
	this.N = &n
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateImageForm) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateImageForm) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *GenerateImageForm) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *GenerateImageForm) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *GenerateImageForm) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *GenerateImageForm) UnsetModel() {
	o.Model.Unset()
}

// GetPrompt returns the Prompt field value
func (o *GenerateImageForm) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *GenerateImageForm) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *GenerateImageForm) SetPrompt(v string) {
	o.Prompt = v
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateImageForm) GetSize() string {
	if o == nil || IsNil(o.Size.Get()) {
		var ret string
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateImageForm) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *GenerateImageForm) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableString and assigns it to the Size field.
func (o *GenerateImageForm) SetSize(v string) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *GenerateImageForm) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GenerateImageForm) UnsetSize() {
	o.Size.Unset()
}

// GetN returns the N field value if set, zero value otherwise.
func (o *GenerateImageForm) GetN() int32 {
	if o == nil || IsNil(o.N) {
		var ret int32
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateImageForm) GetNOk() (*int32, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *GenerateImageForm) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given int32 and assigns it to the N field.
func (o *GenerateImageForm) SetN(v int32) {
	o.N = &v
}

// GetNegativePrompt returns the NegativePrompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateImageForm) GetNegativePrompt() string {
	if o == nil || IsNil(o.NegativePrompt.Get()) {
		var ret string
		return ret
	}
	return *o.NegativePrompt.Get()
}

// GetNegativePromptOk returns a tuple with the NegativePrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateImageForm) GetNegativePromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativePrompt.Get(), o.NegativePrompt.IsSet()
}

// HasNegativePrompt returns a boolean if a field has been set.
func (o *GenerateImageForm) HasNegativePrompt() bool {
	if o != nil && o.NegativePrompt.IsSet() {
		return true
	}

	return false
}

// SetNegativePrompt gets a reference to the given NullableString and assigns it to the NegativePrompt field.
func (o *GenerateImageForm) SetNegativePrompt(v string) {
	o.NegativePrompt.Set(&v)
}
// SetNegativePromptNil sets the value for NegativePrompt to be an explicit nil
func (o *GenerateImageForm) SetNegativePromptNil() {
	o.NegativePrompt.Set(nil)
}

// UnsetNegativePrompt ensures that no value is present for NegativePrompt, not even an explicit nil
func (o *GenerateImageForm) UnsetNegativePrompt() {
	o.NegativePrompt.Unset()
}

func (o GenerateImageForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateImageForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	toSerialize["prompt"] = o.Prompt
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if o.NegativePrompt.IsSet() {
		toSerialize["negative_prompt"] = o.NegativePrompt.Get()
	}
	return toSerialize, nil
}

func (o *GenerateImageForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt",
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

	varGenerateImageForm := _GenerateImageForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateImageForm)

	if err != nil {
		return err
	}

	*o = GenerateImageForm(varGenerateImageForm)

	return err
}

type NullableGenerateImageForm struct {
	value *GenerateImageForm
	isSet bool
}

func (v NullableGenerateImageForm) Get() *GenerateImageForm {
	return v.value
}

func (v *NullableGenerateImageForm) Set(val *GenerateImageForm) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateImageForm) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateImageForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateImageForm(val *GenerateImageForm) *NullableGenerateImageForm {
	return &NullableGenerateImageForm{value: val, isSet: true}
}

func (v NullableGenerateImageForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateImageForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


