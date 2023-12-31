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

// checks if the IpamsvcLeaseRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcLeaseRange{}

// IpamsvcLeaseRange struct for IpamsvcLeaseRange
type IpamsvcLeaseRange struct {
	// The resource identifier.
	Id *string `json:"id,omitempty"`
}

// NewIpamsvcLeaseRange instantiates a new IpamsvcLeaseRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcLeaseRange() *IpamsvcLeaseRange {
	this := IpamsvcLeaseRange{}
	return &this
}

// NewIpamsvcLeaseRangeWithDefaults instantiates a new IpamsvcLeaseRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcLeaseRangeWithDefaults() *IpamsvcLeaseRange {
	this := IpamsvcLeaseRange{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcLeaseRange) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcLeaseRange) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcLeaseRange) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcLeaseRange) SetId(v string) {
	o.Id = &v
}

func (o IpamsvcLeaseRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcLeaseRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableIpamsvcLeaseRange struct {
	value *IpamsvcLeaseRange
	isSet bool
}

func (v NullableIpamsvcLeaseRange) Get() *IpamsvcLeaseRange {
	return v.value
}

func (v *NullableIpamsvcLeaseRange) Set(val *IpamsvcLeaseRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcLeaseRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcLeaseRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcLeaseRange(val *IpamsvcLeaseRange) *NullableIpamsvcLeaseRange {
	return &NullableIpamsvcLeaseRange{value: val, isSet: true}
}

func (v NullableIpamsvcLeaseRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcLeaseRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
