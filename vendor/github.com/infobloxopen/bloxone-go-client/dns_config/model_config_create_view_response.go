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

// checks if the ConfigCreateViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCreateViewResponse{}

// ConfigCreateViewResponse The View object create response format.
type ConfigCreateViewResponse struct {
	Result *ConfigView `json:"result,omitempty"`
}

// NewConfigCreateViewResponse instantiates a new ConfigCreateViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCreateViewResponse() *ConfigCreateViewResponse {
	this := ConfigCreateViewResponse{}
	return &this
}

// NewConfigCreateViewResponseWithDefaults instantiates a new ConfigCreateViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCreateViewResponseWithDefaults() *ConfigCreateViewResponse {
	this := ConfigCreateViewResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigCreateViewResponse) GetResult() ConfigView {
	if o == nil || IsNil(o.Result) {
		var ret ConfigView
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCreateViewResponse) GetResultOk() (*ConfigView, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigCreateViewResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigView and assigns it to the Result field.
func (o *ConfigCreateViewResponse) SetResult(v ConfigView) {
	o.Result = &v
}

func (o ConfigCreateViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCreateViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigCreateViewResponse struct {
	value *ConfigCreateViewResponse
	isSet bool
}

func (v NullableConfigCreateViewResponse) Get() *ConfigCreateViewResponse {
	return v.value
}

func (v *NullableConfigCreateViewResponse) Set(val *ConfigCreateViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCreateViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCreateViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCreateViewResponse(val *ConfigCreateViewResponse) *NullableConfigCreateViewResponse {
	return &NullableConfigCreateViewResponse{value: val, isSet: true}
}

func (v NullableConfigCreateViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCreateViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
