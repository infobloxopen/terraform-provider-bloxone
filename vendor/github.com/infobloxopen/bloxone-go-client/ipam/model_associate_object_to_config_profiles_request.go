/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// checks if the AssociateObjectToConfigProfilesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociateObjectToConfigProfilesRequest{}

// AssociateObjectToConfigProfilesRequest AssociateObjectToConfigProfiles associates an object to config profiles.
type AssociateObjectToConfigProfilesRequest struct {
	// The resource identifier.
	ConfigProfileIds []string           `json:"config_profile_ids"`
	Fields           *ProtobufFieldMask `json:"fields,omitempty"`
	// The resource identifier.
	ObjectId             string `json:"object_id"`
	AdditionalProperties map[string]interface{}
}

type _AssociateObjectToConfigProfilesRequest AssociateObjectToConfigProfilesRequest

// NewAssociateObjectToConfigProfilesRequest instantiates a new AssociateObjectToConfigProfilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociateObjectToConfigProfilesRequest(configProfileIds []string, objectId string) *AssociateObjectToConfigProfilesRequest {
	this := AssociateObjectToConfigProfilesRequest{}
	this.ConfigProfileIds = configProfileIds
	this.ObjectId = objectId
	return &this
}

// NewAssociateObjectToConfigProfilesRequestWithDefaults instantiates a new AssociateObjectToConfigProfilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociateObjectToConfigProfilesRequestWithDefaults() *AssociateObjectToConfigProfilesRequest {
	this := AssociateObjectToConfigProfilesRequest{}
	return &this
}

// GetConfigProfileIds returns the ConfigProfileIds field value
func (o *AssociateObjectToConfigProfilesRequest) GetConfigProfileIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfigProfileIds
}

// GetConfigProfileIdsOk returns a tuple with the ConfigProfileIds field value
// and a boolean to check if the value has been set.
func (o *AssociateObjectToConfigProfilesRequest) GetConfigProfileIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigProfileIds, true
}

// SetConfigProfileIds sets field value
func (o *AssociateObjectToConfigProfilesRequest) SetConfigProfileIds(v []string) {
	o.ConfigProfileIds = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AssociateObjectToConfigProfilesRequest) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociateObjectToConfigProfilesRequest) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AssociateObjectToConfigProfilesRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *AssociateObjectToConfigProfilesRequest) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetObjectId returns the ObjectId field value
func (o *AssociateObjectToConfigProfilesRequest) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *AssociateObjectToConfigProfilesRequest) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *AssociateObjectToConfigProfilesRequest) SetObjectId(v string) {
	o.ObjectId = v
}

func (o AssociateObjectToConfigProfilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociateObjectToConfigProfilesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config_profile_ids"] = o.ConfigProfileIds
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["object_id"] = o.ObjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssociateObjectToConfigProfilesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config_profile_ids",
		"object_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAssociateObjectToConfigProfilesRequest := _AssociateObjectToConfigProfilesRequest{}

	err = json.Unmarshal(data, &varAssociateObjectToConfigProfilesRequest)

	if err != nil {
		return err
	}

	*o = AssociateObjectToConfigProfilesRequest(varAssociateObjectToConfigProfilesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config_profile_ids")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "object_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssociateObjectToConfigProfilesRequest struct {
	value *AssociateObjectToConfigProfilesRequest
	isSet bool
}

func (v NullableAssociateObjectToConfigProfilesRequest) Get() *AssociateObjectToConfigProfilesRequest {
	return v.value
}

func (v *NullableAssociateObjectToConfigProfilesRequest) Set(val *AssociateObjectToConfigProfilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociateObjectToConfigProfilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociateObjectToConfigProfilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociateObjectToConfigProfilesRequest(val *AssociateObjectToConfigProfilesRequest) *NullableAssociateObjectToConfigProfilesRequest {
	return &NullableAssociateObjectToConfigProfilesRequest{value: val, isSet: true}
}

func (v NullableAssociateObjectToConfigProfilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociateObjectToConfigProfilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}