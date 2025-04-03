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

// checks if the UrlForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlForm{}

// UrlForm struct for UrlForm
type UrlForm struct {
	Url string `json:"url"`
}

type _UrlForm UrlForm

// NewUrlForm instantiates a new UrlForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlForm(url string) *UrlForm {
	this := UrlForm{}
	this.Url = url
	return &this
}

// NewUrlFormWithDefaults instantiates a new UrlForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlFormWithDefaults() *UrlForm {
	this := UrlForm{}
	return &this
}

// GetUrl returns the Url field value
func (o *UrlForm) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UrlForm) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UrlForm) SetUrl(v string) {
	o.Url = v
}

func (o UrlForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *UrlForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varUrlForm := _UrlForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUrlForm)

	if err != nil {
		return err
	}

	*o = UrlForm(varUrlForm)

	return err
}

type NullableUrlForm struct {
	value *UrlForm
	isSet bool
}

func (v NullableUrlForm) Get() *UrlForm {
	return v.value
}

func (v *NullableUrlForm) Set(val *UrlForm) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlForm) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlForm(val *UrlForm) *NullableUrlForm {
	return &NullableUrlForm{value: val, isSet: true}
}

func (v NullableUrlForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


