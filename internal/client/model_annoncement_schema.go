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

// checks if the AnnoncementSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnoncementSchema{}

// AnnoncementSchema struct for AnnoncementSchema
type AnnoncementSchema struct {
	// Annoncement text.
	Message string `json:"message"`
	CreatedAt NullableTime `json:"created_at,omitempty"`
}

type _AnnoncementSchema AnnoncementSchema

// NewAnnoncementSchema instantiates a new AnnoncementSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnoncementSchema(message string) *AnnoncementSchema {
	this := AnnoncementSchema{}
	this.Message = message
	return &this
}

// NewAnnoncementSchemaWithDefaults instantiates a new AnnoncementSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnoncementSchemaWithDefaults() *AnnoncementSchema {
	this := AnnoncementSchema{}
	return &this
}

// GetMessage returns the Message field value
func (o *AnnoncementSchema) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AnnoncementSchema) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AnnoncementSchema) SetMessage(v string) {
	o.Message = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnnoncementSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnnoncementSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnnoncementSchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *AnnoncementSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *AnnoncementSchema) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *AnnoncementSchema) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

func (o AnnoncementSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnoncementSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	return toSerialize, nil
}

func (o *AnnoncementSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varAnnoncementSchema := _AnnoncementSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnnoncementSchema)

	if err != nil {
		return err
	}

	*o = AnnoncementSchema(varAnnoncementSchema)

	return err
}

type NullableAnnoncementSchema struct {
	value *AnnoncementSchema
	isSet bool
}

func (v NullableAnnoncementSchema) Get() *AnnoncementSchema {
	return v.value
}

func (v *NullableAnnoncementSchema) Set(val *AnnoncementSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnoncementSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnoncementSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnoncementSchema(val *AnnoncementSchema) *NullableAnnoncementSchema {
	return &NullableAnnoncementSchema{value: val, isSet: true}
}

func (v NullableAnnoncementSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnoncementSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

