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

// checks if the DataPageMapSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPageMapSchema{}

// DataPageMapSchema struct for DataPageMapSchema
type DataPageMapSchema struct {
	Data []MapSchema `json:"data"`
	Total NullableInt32 `json:"total"`
	Page NullableInt32 `json:"page"`
	Size NullableInt32 `json:"size"`
	Pages NullableInt32 `json:"pages,omitempty"`
}

type _DataPageMapSchema DataPageMapSchema

// NewDataPageMapSchema instantiates a new DataPageMapSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPageMapSchema(data []MapSchema, total NullableInt32, page NullableInt32, size NullableInt32) *DataPageMapSchema {
	this := DataPageMapSchema{}
	this.Data = data
	this.Total = total
	this.Page = page
	this.Size = size
	return &this
}

// NewDataPageMapSchemaWithDefaults instantiates a new DataPageMapSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPageMapSchemaWithDefaults() *DataPageMapSchema {
	this := DataPageMapSchema{}
	return &this
}

// GetData returns the Data field value
func (o *DataPageMapSchema) GetData() []MapSchema {
	if o == nil {
		var ret []MapSchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataPageMapSchema) GetDataOk() ([]MapSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DataPageMapSchema) SetData(v []MapSchema) {
	o.Data = v
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DataPageMapSchema) GetTotal() int32 {
	if o == nil || o.Total.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPageMapSchema) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// SetTotal sets field value
func (o *DataPageMapSchema) SetTotal(v int32) {
	o.Total.Set(&v)
}

// GetPage returns the Page field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DataPageMapSchema) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPageMapSchema) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// SetPage sets field value
func (o *DataPageMapSchema) SetPage(v int32) {
	o.Page.Set(&v)
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DataPageMapSchema) GetSize() int32 {
	if o == nil || o.Size.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPageMapSchema) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// SetSize sets field value
func (o *DataPageMapSchema) SetSize(v int32) {
	o.Size.Set(&v)
}

// GetPages returns the Pages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPageMapSchema) GetPages() int32 {
	if o == nil || IsNil(o.Pages.Get()) {
		var ret int32
		return ret
	}
	return *o.Pages.Get()
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPageMapSchema) GetPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages.Get(), o.Pages.IsSet()
}

// HasPages returns a boolean if a field has been set.
func (o *DataPageMapSchema) HasPages() bool {
	if o != nil && o.Pages.IsSet() {
		return true
	}

	return false
}

// SetPages gets a reference to the given NullableInt32 and assigns it to the Pages field.
func (o *DataPageMapSchema) SetPages(v int32) {
	o.Pages.Set(&v)
}
// SetPagesNil sets the value for Pages to be an explicit nil
func (o *DataPageMapSchema) SetPagesNil() {
	o.Pages.Set(nil)
}

// UnsetPages ensures that no value is present for Pages, not even an explicit nil
func (o *DataPageMapSchema) UnsetPages() {
	o.Pages.Unset()
}

func (o DataPageMapSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPageMapSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["total"] = o.Total.Get()
	toSerialize["page"] = o.Page.Get()
	toSerialize["size"] = o.Size.Get()
	if o.Pages.IsSet() {
		toSerialize["pages"] = o.Pages.Get()
	}
	return toSerialize, nil
}

func (o *DataPageMapSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"total",
		"page",
		"size",
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

	varDataPageMapSchema := _DataPageMapSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataPageMapSchema)

	if err != nil {
		return err
	}

	*o = DataPageMapSchema(varDataPageMapSchema)

	return err
}

type NullableDataPageMapSchema struct {
	value *DataPageMapSchema
	isSet bool
}

func (v NullableDataPageMapSchema) Get() *DataPageMapSchema {
	return v.value
}

func (v *NullableDataPageMapSchema) Set(val *DataPageMapSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPageMapSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPageMapSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPageMapSchema(val *DataPageMapSchema) *NullableDataPageMapSchema {
	return &NullableDataPageMapSchema{value: val, isSet: true}
}

func (v NullableDataPageMapSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPageMapSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


