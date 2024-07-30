/*
Artifacts API

 Artifacts is an API-based MMO game where you can manage 5 characters to explore, fight, gather resources, craft items and much more.  Website: https://artifactsmmo.com/  Documentation: https://docs.artifactsmmo.com/  OpenAPI Spec: https://api.artifactsmmo.com/openapi.json 

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TaskResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResponseSchema{}

// TaskResponseSchema struct for TaskResponseSchema
type TaskResponseSchema struct {
	Data TaskDataSchema `json:"data"`
}

type _TaskResponseSchema TaskResponseSchema

// NewTaskResponseSchema instantiates a new TaskResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResponseSchema(data TaskDataSchema) *TaskResponseSchema {
	this := TaskResponseSchema{}
	this.Data = data
	return &this
}

// NewTaskResponseSchemaWithDefaults instantiates a new TaskResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResponseSchemaWithDefaults() *TaskResponseSchema {
	this := TaskResponseSchema{}
	return &this
}

// GetData returns the Data field value
func (o *TaskResponseSchema) GetData() TaskDataSchema {
	if o == nil {
		var ret TaskDataSchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaskResponseSchema) GetDataOk() (*TaskDataSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaskResponseSchema) SetData(v TaskDataSchema) {
	o.Data = v
}

func (o TaskResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TaskResponseSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varTaskResponseSchema := _TaskResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskResponseSchema)

	if err != nil {
		return err
	}

	*o = TaskResponseSchema(varTaskResponseSchema)

	return err
}

type NullableTaskResponseSchema struct {
	value *TaskResponseSchema
	isSet bool
}

func (v NullableTaskResponseSchema) Get() *TaskResponseSchema {
	return v.value
}

func (v *NullableTaskResponseSchema) Set(val *TaskResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResponseSchema(val *TaskResponseSchema) *NullableTaskResponseSchema {
	return &NullableTaskResponseSchema{value: val, isSet: true}
}

func (v NullableTaskResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

