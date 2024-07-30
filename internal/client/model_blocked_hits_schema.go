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

// checks if the BlockedHitsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockedHitsSchema{}

// BlockedHitsSchema struct for BlockedHitsSchema
type BlockedHitsSchema struct {
	// The amount of fire hits blocked.
	Fire int32 `json:"fire"`
	// The amount of earth hits blocked.
	Earth int32 `json:"earth"`
	// The amount of water hits blocked.
	Water int32 `json:"water"`
	// The amount of air hits blocked.
	Air int32 `json:"air"`
	// The amount of total hits blocked.
	Total int32 `json:"total"`
}

type _BlockedHitsSchema BlockedHitsSchema

// NewBlockedHitsSchema instantiates a new BlockedHitsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockedHitsSchema(fire int32, earth int32, water int32, air int32, total int32) *BlockedHitsSchema {
	this := BlockedHitsSchema{}
	this.Fire = fire
	this.Earth = earth
	this.Water = water
	this.Air = air
	this.Total = total
	return &this
}

// NewBlockedHitsSchemaWithDefaults instantiates a new BlockedHitsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockedHitsSchemaWithDefaults() *BlockedHitsSchema {
	this := BlockedHitsSchema{}
	return &this
}

// GetFire returns the Fire field value
func (o *BlockedHitsSchema) GetFire() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fire
}

// GetFireOk returns a tuple with the Fire field value
// and a boolean to check if the value has been set.
func (o *BlockedHitsSchema) GetFireOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fire, true
}

// SetFire sets field value
func (o *BlockedHitsSchema) SetFire(v int32) {
	o.Fire = v
}

// GetEarth returns the Earth field value
func (o *BlockedHitsSchema) GetEarth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Earth
}

// GetEarthOk returns a tuple with the Earth field value
// and a boolean to check if the value has been set.
func (o *BlockedHitsSchema) GetEarthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Earth, true
}

// SetEarth sets field value
func (o *BlockedHitsSchema) SetEarth(v int32) {
	o.Earth = v
}

// GetWater returns the Water field value
func (o *BlockedHitsSchema) GetWater() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Water
}

// GetWaterOk returns a tuple with the Water field value
// and a boolean to check if the value has been set.
func (o *BlockedHitsSchema) GetWaterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Water, true
}

// SetWater sets field value
func (o *BlockedHitsSchema) SetWater(v int32) {
	o.Water = v
}

// GetAir returns the Air field value
func (o *BlockedHitsSchema) GetAir() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Air
}

// GetAirOk returns a tuple with the Air field value
// and a boolean to check if the value has been set.
func (o *BlockedHitsSchema) GetAirOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Air, true
}

// SetAir sets field value
func (o *BlockedHitsSchema) SetAir(v int32) {
	o.Air = v
}

// GetTotal returns the Total field value
func (o *BlockedHitsSchema) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BlockedHitsSchema) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BlockedHitsSchema) SetTotal(v int32) {
	o.Total = v
}

func (o BlockedHitsSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockedHitsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fire"] = o.Fire
	toSerialize["earth"] = o.Earth
	toSerialize["water"] = o.Water
	toSerialize["air"] = o.Air
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *BlockedHitsSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fire",
		"earth",
		"water",
		"air",
		"total",
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

	varBlockedHitsSchema := _BlockedHitsSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockedHitsSchema)

	if err != nil {
		return err
	}

	*o = BlockedHitsSchema(varBlockedHitsSchema)

	return err
}

type NullableBlockedHitsSchema struct {
	value *BlockedHitsSchema
	isSet bool
}

func (v NullableBlockedHitsSchema) Get() *BlockedHitsSchema {
	return v.value
}

func (v *NullableBlockedHitsSchema) Set(val *BlockedHitsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockedHitsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockedHitsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockedHitsSchema(val *BlockedHitsSchema) *NullableBlockedHitsSchema {
	return &NullableBlockedHitsSchema{value: val, isSet: true}
}

func (v NullableBlockedHitsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockedHitsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


