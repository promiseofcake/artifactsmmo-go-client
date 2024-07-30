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

// checks if the EventSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSchema{}

// EventSchema struct for EventSchema
type EventSchema struct {
	// Name of the event.
	Name string `json:"name"`
	// Map of the event.
	Map MapSchema `json:"map"`
	// Previous map skin.
	PreviousSkin string `json:"previous_skin"`
	// Duration in minutes.
	Duration int32 `json:"duration"`
	// Expiration datetime.
	Expiration time.Time `json:"expiration"`
	// Start datetime.
	CreatedAt time.Time `json:"created_at"`
}

type _EventSchema EventSchema

// NewEventSchema instantiates a new EventSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSchema(name string, map_ MapSchema, previousSkin string, duration int32, expiration time.Time, createdAt time.Time) *EventSchema {
	this := EventSchema{}
	this.Name = name
	this.Map = map_
	this.PreviousSkin = previousSkin
	this.Duration = duration
	this.Expiration = expiration
	this.CreatedAt = createdAt
	return &this
}

// NewEventSchemaWithDefaults instantiates a new EventSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSchemaWithDefaults() *EventSchema {
	this := EventSchema{}
	return &this
}

// GetName returns the Name field value
func (o *EventSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventSchema) SetName(v string) {
	o.Name = v
}

// GetMap returns the Map field value
func (o *EventSchema) GetMap() MapSchema {
	if o == nil {
		var ret MapSchema
		return ret
	}

	return o.Map
}

// GetMapOk returns a tuple with the Map field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetMapOk() (*MapSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Map, true
}

// SetMap sets field value
func (o *EventSchema) SetMap(v MapSchema) {
	o.Map = v
}

// GetPreviousSkin returns the PreviousSkin field value
func (o *EventSchema) GetPreviousSkin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousSkin
}

// GetPreviousSkinOk returns a tuple with the PreviousSkin field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetPreviousSkinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousSkin, true
}

// SetPreviousSkin sets field value
func (o *EventSchema) SetPreviousSkin(v string) {
	o.PreviousSkin = v
}

// GetDuration returns the Duration field value
func (o *EventSchema) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *EventSchema) SetDuration(v int32) {
	o.Duration = v
}

// GetExpiration returns the Expiration field value
func (o *EventSchema) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *EventSchema) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EventSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EventSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o EventSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["map"] = o.Map
	toSerialize["previous_skin"] = o.PreviousSkin
	toSerialize["duration"] = o.Duration
	toSerialize["expiration"] = o.Expiration
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *EventSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"map",
		"previous_skin",
		"duration",
		"expiration",
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

	varEventSchema := _EventSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventSchema)

	if err != nil {
		return err
	}

	*o = EventSchema(varEventSchema)

	return err
}

type NullableEventSchema struct {
	value *EventSchema
	isSet bool
}

func (v NullableEventSchema) Get() *EventSchema {
	return v.value
}

func (v *NullableEventSchema) Set(val *EventSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSchema(val *EventSchema) *NullableEventSchema {
	return &NullableEventSchema{value: val, isSet: true}
}

func (v NullableEventSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

