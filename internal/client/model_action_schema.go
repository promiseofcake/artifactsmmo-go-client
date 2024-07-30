/*
Artifacts API

 Artifacts is an API-based MMO game where you can manage 5 characters to explore, fight, gather resources, craft items and much more.  Website: https://artifactsmmo.com/  Documentation: https://docs.artifactsmmo.com/  OpenAPI Spec: https://api.artifactsmmo.com/openapi.json 

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ActionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionSchema{}

// ActionSchema struct for ActionSchema
type ActionSchema struct {
	// Character name.
	Character string `json:"character"`
	// Account character.
	Account string `json:"account"`
	// Type of action.
	Type string `json:"type"`
	// Description of action.
	Description string `json:"description"`
	Content interface{} `json:"content"`
	// Cooldown in seconds.
	Cooldown int32 `json:"cooldown"`
	// Datetime of cooldown expiration.
	CooldownExpiration time.Time `json:"cooldown_expiration"`
	// Datetime of creation.
	CreatedAt time.Time `json:"created_at"`
}

type _ActionSchema ActionSchema

// NewActionSchema instantiates a new ActionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionSchema(character string, account string, type_ string, description string, content interface{}, cooldown int32, cooldownExpiration time.Time, createdAt time.Time) *ActionSchema {
	this := ActionSchema{}
	this.Character = character
	this.Account = account
	this.Type = type_
	this.Description = description
	this.Content = content
	this.Cooldown = cooldown
	this.CooldownExpiration = cooldownExpiration
	this.CreatedAt = createdAt
	return &this
}

// NewActionSchemaWithDefaults instantiates a new ActionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionSchemaWithDefaults() *ActionSchema {
	this := ActionSchema{}
	return &this
}

// GetCharacter returns the Character field value
func (o *ActionSchema) GetCharacter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Character
}

// GetCharacterOk returns a tuple with the Character field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetCharacterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Character, true
}

// SetCharacter sets field value
func (o *ActionSchema) SetCharacter(v string) {
	o.Character = v
}

// GetAccount returns the Account field value
func (o *ActionSchema) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ActionSchema) SetAccount(v string) {
	o.Account = v
}

// GetType returns the Type field value
func (o *ActionSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActionSchema) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *ActionSchema) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ActionSchema) SetDescription(v string) {
	o.Description = v
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ActionSchema) GetContent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionSchema) GetContentOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ActionSchema) SetContent(v interface{}) {
	o.Content = v
}

// GetCooldown returns the Cooldown field value
func (o *ActionSchema) GetCooldown() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetCooldownOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *ActionSchema) SetCooldown(v int32) {
	o.Cooldown = v
}

// GetCooldownExpiration returns the CooldownExpiration field value
func (o *ActionSchema) GetCooldownExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CooldownExpiration
}

// GetCooldownExpirationOk returns a tuple with the CooldownExpiration field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetCooldownExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CooldownExpiration, true
}

// SetCooldownExpiration sets field value
func (o *ActionSchema) SetCooldownExpiration(v time.Time) {
	o.CooldownExpiration = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ActionSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ActionSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ActionSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o ActionSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["character"] = o.Character
	toSerialize["account"] = o.Account
	toSerialize["type"] = o.Type
	toSerialize["description"] = o.Description
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["cooldown_expiration"] = o.CooldownExpiration
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *ActionSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"character",
		"account",
		"type",
		"description",
		"content",
		"cooldown",
		"cooldown_expiration",
		"created_at",
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

	varActionSchema := _ActionSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionSchema)

	if err != nil {
		return err
	}

	*o = ActionSchema(varActionSchema)

	return err
}

type NullableActionSchema struct {
	value *ActionSchema
	isSet bool
}

func (v NullableActionSchema) Get() *ActionSchema {
	return v.value
}

func (v *NullableActionSchema) Set(val *ActionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableActionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableActionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionSchema(val *ActionSchema) *NullableActionSchema {
	return &NullableActionSchema{value: val, isSet: true}
}

func (v NullableActionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

