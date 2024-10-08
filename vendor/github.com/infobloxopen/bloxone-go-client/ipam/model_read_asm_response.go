/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the ReadASMResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadASMResponse{}

// ReadASMResponse The response format to retrieve the __ASM__ object.
type ReadASMResponse struct {
	// The ASM object.
	Result               *ASM `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadASMResponse ReadASMResponse

// NewReadASMResponse instantiates a new ReadASMResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadASMResponse() *ReadASMResponse {
	this := ReadASMResponse{}
	return &this
}

// NewReadASMResponseWithDefaults instantiates a new ReadASMResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadASMResponseWithDefaults() *ReadASMResponse {
	this := ReadASMResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadASMResponse) GetResult() ASM {
	if o == nil || IsNil(o.Result) {
		var ret ASM
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadASMResponse) GetResultOk() (*ASM, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadASMResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ASM and assigns it to the Result field.
func (o *ReadASMResponse) SetResult(v ASM) {
	o.Result = &v
}

func (o ReadASMResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadASMResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadASMResponse) UnmarshalJSON(data []byte) (err error) {
	varReadASMResponse := _ReadASMResponse{}

	err = json.Unmarshal(data, &varReadASMResponse)

	if err != nil {
		return err
	}

	*o = ReadASMResponse(varReadASMResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadASMResponse struct {
	value *ReadASMResponse
	isSet bool
}

func (v NullableReadASMResponse) Get() *ReadASMResponse {
	return v.value
}

func (v *NullableReadASMResponse) Set(val *ReadASMResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadASMResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadASMResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadASMResponse(val *ReadASMResponse) *NullableReadASMResponse {
	return &NullableReadASMResponse{value: val, isSet: true}
}

func (v NullableReadASMResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadASMResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
