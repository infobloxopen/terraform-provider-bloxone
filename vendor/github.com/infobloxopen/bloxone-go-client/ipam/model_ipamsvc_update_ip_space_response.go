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

// checks if the IpamsvcUpdateIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcUpdateIPSpaceResponse{}

// IpamsvcUpdateIPSpaceResponse The response format to update the __IPSpace__ object.
type IpamsvcUpdateIPSpaceResponse struct {
	Result *IpamsvcIPSpace `json:"result,omitempty"`
}

// NewIpamsvcUpdateIPSpaceResponse instantiates a new IpamsvcUpdateIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcUpdateIPSpaceResponse() *IpamsvcUpdateIPSpaceResponse {
	this := IpamsvcUpdateIPSpaceResponse{}
	return &this
}

// NewIpamsvcUpdateIPSpaceResponseWithDefaults instantiates a new IpamsvcUpdateIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcUpdateIPSpaceResponseWithDefaults() *IpamsvcUpdateIPSpaceResponse {
	this := IpamsvcUpdateIPSpaceResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IpamsvcUpdateIPSpaceResponse) GetResult() IpamsvcIPSpace {
	if o == nil || IsNil(o.Result) {
		var ret IpamsvcIPSpace
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcUpdateIPSpaceResponse) GetResultOk() (*IpamsvcIPSpace, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IpamsvcUpdateIPSpaceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IpamsvcIPSpace and assigns it to the Result field.
func (o *IpamsvcUpdateIPSpaceResponse) SetResult(v IpamsvcIPSpace) {
	o.Result = &v
}

func (o IpamsvcUpdateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcUpdateIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableIpamsvcUpdateIPSpaceResponse struct {
	value *IpamsvcUpdateIPSpaceResponse
	isSet bool
}

func (v NullableIpamsvcUpdateIPSpaceResponse) Get() *IpamsvcUpdateIPSpaceResponse {
	return v.value
}

func (v *NullableIpamsvcUpdateIPSpaceResponse) Set(val *IpamsvcUpdateIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcUpdateIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcUpdateIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcUpdateIPSpaceResponse(val *IpamsvcUpdateIPSpaceResponse) *NullableIpamsvcUpdateIPSpaceResponse {
	return &NullableIpamsvcUpdateIPSpaceResponse{value: val, isSet: true}
}

func (v NullableIpamsvcUpdateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcUpdateIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
