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

// checks if the StatusResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseSchema{}

// StatusResponseSchema struct for StatusResponseSchema
type StatusResponseSchema struct {
	Data StatusSchema `json:"data"`
}

type _StatusResponseSchema StatusResponseSchema

// NewStatusResponseSchema instantiates a new StatusResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseSchema(data StatusSchema) *StatusResponseSchema {
	this := StatusResponseSchema{}
	this.Data = data
	return &this
}

// NewStatusResponseSchemaWithDefaults instantiates a new StatusResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseSchemaWithDefaults() *StatusResponseSchema {
	this := StatusResponseSchema{}
	return &this
}

// GetData returns the Data field value
func (o *StatusResponseSchema) GetData() StatusSchema {
	if o == nil {
		var ret StatusSchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StatusResponseSchema) GetDataOk() (*StatusSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StatusResponseSchema) SetData(v StatusSchema) {
	o.Data = v
}

func (o StatusResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *StatusResponseSchema) UnmarshalJSON(data []byte) (err error) {
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

	varStatusResponseSchema := _StatusResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatusResponseSchema)

	if err != nil {
		return err
	}

	*o = StatusResponseSchema(varStatusResponseSchema)

	return err
}

type NullableStatusResponseSchema struct {
	value *StatusResponseSchema
	isSet bool
}

func (v NullableStatusResponseSchema) Get() *StatusResponseSchema {
	return v.value
}

func (v *NullableStatusResponseSchema) Set(val *StatusResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseSchema(val *StatusResponseSchema) *NullableStatusResponseSchema {
	return &NullableStatusResponseSchema{value: val, isSet: true}
}

func (v NullableStatusResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


