/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// Data struct for Data
type Data struct {
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Data) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny);
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Data)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Data) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableData struct {
	value *Data
	isSet bool
}

func (v NullableData) Get() *Data {
	return v.value
}

func (v *NullableData) Set(val *Data) {
	v.value = val
	v.isSet = true
}

func (v NullableData) IsSet() bool {
	return v.isSet
}

func (v *NullableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableData(val *Data) *NullableData {
	return &NullableData{value: val, isSet: true}
}

func (v NullableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


