/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigConvertRNameResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigConvertRNameResponse{}

// ConfigConvertRNameResponse struct for ConfigConvertRNameResponse
type ConfigConvertRNameResponse struct {
	// The SOA RNAME field converted from the provided email address.
	Rname *string `json:"rname,omitempty"`
}

// NewConfigConvertRNameResponse instantiates a new ConfigConvertRNameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigConvertRNameResponse() *ConfigConvertRNameResponse {
	this := ConfigConvertRNameResponse{}
	return &this
}

// NewConfigConvertRNameResponseWithDefaults instantiates a new ConfigConvertRNameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigConvertRNameResponseWithDefaults() *ConfigConvertRNameResponse {
	this := ConfigConvertRNameResponse{}
	return &this
}

// GetRname returns the Rname field value if set, zero value otherwise.
func (o *ConfigConvertRNameResponse) GetRname() string {
	if o == nil || IsNil(o.Rname) {
		var ret string
		return ret
	}
	return *o.Rname
}

// GetRnameOk returns a tuple with the Rname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigConvertRNameResponse) GetRnameOk() (*string, bool) {
	if o == nil || IsNil(o.Rname) {
		return nil, false
	}
	return o.Rname, true
}

// HasRname returns a boolean if a field has been set.
func (o *ConfigConvertRNameResponse) HasRname() bool {
	if o != nil && !IsNil(o.Rname) {
		return true
	}

	return false
}

// SetRname gets a reference to the given string and assigns it to the Rname field.
func (o *ConfigConvertRNameResponse) SetRname(v string) {
	o.Rname = &v
}

func (o ConfigConvertRNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigConvertRNameResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rname) {
		toSerialize["rname"] = o.Rname
	}
	return toSerialize, nil
}

type NullableConfigConvertRNameResponse struct {
	value *ConfigConvertRNameResponse
	isSet bool
}

func (v NullableConfigConvertRNameResponse) Get() *ConfigConvertRNameResponse {
	return v.value
}

func (v *NullableConfigConvertRNameResponse) Set(val *ConfigConvertRNameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigConvertRNameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigConvertRNameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigConvertRNameResponse(val *ConfigConvertRNameResponse) *NullableConfigConvertRNameResponse {
	return &NullableConfigConvertRNameResponse{value: val, isSet: true}
}

func (v NullableConfigConvertRNameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigConvertRNameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
