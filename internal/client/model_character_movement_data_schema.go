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

// checks if the CharacterMovementDataSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterMovementDataSchema{}

// CharacterMovementDataSchema struct for CharacterMovementDataSchema
type CharacterMovementDataSchema struct {
	// Cooldown details
	Cooldown CooldownSchema `json:"cooldown"`
	// Destination details.
	Destination DestinationResponseSchema `json:"destination"`
	// Character details.
	Character CharacterSchema `json:"character"`
}

type _CharacterMovementDataSchema CharacterMovementDataSchema

// NewCharacterMovementDataSchema instantiates a new CharacterMovementDataSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterMovementDataSchema(cooldown CooldownSchema, destination DestinationResponseSchema, character CharacterSchema) *CharacterMovementDataSchema {
	this := CharacterMovementDataSchema{}
	this.Cooldown = cooldown
	this.Destination = destination
	this.Character = character
	return &this
}

// NewCharacterMovementDataSchemaWithDefaults instantiates a new CharacterMovementDataSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterMovementDataSchemaWithDefaults() *CharacterMovementDataSchema {
	this := CharacterMovementDataSchema{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *CharacterMovementDataSchema) GetCooldown() CooldownSchema {
	if o == nil {
		var ret CooldownSchema
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *CharacterMovementDataSchema) GetCooldownOk() (*CooldownSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *CharacterMovementDataSchema) SetCooldown(v CooldownSchema) {
	o.Cooldown = v
}

// GetDestination returns the Destination field value
func (o *CharacterMovementDataSchema) GetDestination() DestinationResponseSchema {
	if o == nil {
		var ret DestinationResponseSchema
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *CharacterMovementDataSchema) GetDestinationOk() (*DestinationResponseSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *CharacterMovementDataSchema) SetDestination(v DestinationResponseSchema) {
	o.Destination = v
}

// GetCharacter returns the Character field value
func (o *CharacterMovementDataSchema) GetCharacter() CharacterSchema {
	if o == nil {
		var ret CharacterSchema
		return ret
	}

	return o.Character
}

// GetCharacterOk returns a tuple with the Character field value
// and a boolean to check if the value has been set.
func (o *CharacterMovementDataSchema) GetCharacterOk() (*CharacterSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Character, true
}

// SetCharacter sets field value
func (o *CharacterMovementDataSchema) SetCharacter(v CharacterSchema) {
	o.Character = v
}

func (o CharacterMovementDataSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterMovementDataSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["destination"] = o.Destination
	toSerialize["character"] = o.Character
	return toSerialize, nil
}

func (o *CharacterMovementDataSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"destination",
		"character",
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

	varCharacterMovementDataSchema := _CharacterMovementDataSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCharacterMovementDataSchema)

	if err != nil {
		return err
	}

	*o = CharacterMovementDataSchema(varCharacterMovementDataSchema)

	return err
}

type NullableCharacterMovementDataSchema struct {
	value *CharacterMovementDataSchema
	isSet bool
}

func (v NullableCharacterMovementDataSchema) Get() *CharacterMovementDataSchema {
	return v.value
}

func (v *NullableCharacterMovementDataSchema) Set(val *CharacterMovementDataSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterMovementDataSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterMovementDataSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterMovementDataSchema(val *CharacterMovementDataSchema) *NullableCharacterMovementDataSchema {
	return &NullableCharacterMovementDataSchema{value: val, isSet: true}
}

func (v NullableCharacterMovementDataSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterMovementDataSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


