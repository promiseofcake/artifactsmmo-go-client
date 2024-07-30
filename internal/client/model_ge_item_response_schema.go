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

// checks if the GEItemResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GEItemResponseSchema{}

// GEItemResponseSchema struct for GEItemResponseSchema
type GEItemResponseSchema struct {
	Data GEItemSchema `json:"data"`
}

type _GEItemResponseSchema GEItemResponseSchema

// NewGEItemResponseSchema instantiates a new GEItemResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGEItemResponseSchema(data GEItemSchema) *GEItemResponseSchema {
	this := GEItemResponseSchema{}
	this.Data = data
	return &this
}

// NewGEItemResponseSchemaWithDefaults instantiates a new GEItemResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGEItemResponseSchemaWithDefaults() *GEItemResponseSchema {
	this := GEItemResponseSchema{}
	return &this
}

// GetData returns the Data field value
func (o *GEItemResponseSchema) GetData() GEItemSchema {
	if o == nil {
		var ret GEItemSchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GEItemResponseSchema) GetDataOk() (*GEItemSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GEItemResponseSchema) SetData(v GEItemSchema) {
	o.Data = v
}

func (o GEItemResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GEItemResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GEItemResponseSchema) UnmarshalJSON(data []byte) (err error) {
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

	varGEItemResponseSchema := _GEItemResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGEItemResponseSchema)

	if err != nil {
		return err
	}

	*o = GEItemResponseSchema(varGEItemResponseSchema)

	return err
}

type NullableGEItemResponseSchema struct {
	value *GEItemResponseSchema
	isSet bool
}

func (v NullableGEItemResponseSchema) Get() *GEItemResponseSchema {
	return v.value
}

func (v *NullableGEItemResponseSchema) Set(val *GEItemResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGEItemResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGEItemResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGEItemResponseSchema(val *GEItemResponseSchema) *NullableGEItemResponseSchema {
	return &NullableGEItemResponseSchema{value: val, isSet: true}
}

func (v NullableGEItemResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGEItemResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

