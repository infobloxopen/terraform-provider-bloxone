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

// checks if the IpamsvcDDNSHostnameBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcDDNSHostnameBlock{}

// IpamsvcDDNSHostnameBlock The dynamic DNS Hostname configuration.
type IpamsvcDDNSHostnameBlock struct {
	// Indicates if DDNS should generate a hostname when not supplied by the client.
	DdnsGenerateName *bool `json:"ddns_generate_name,omitempty"`
	// The prefix used in the generation of an FQDN.
	DdnsGeneratedPrefix *string `json:"ddns_generated_prefix,omitempty"`
}

// NewIpamsvcDDNSHostnameBlock instantiates a new IpamsvcDDNSHostnameBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcDDNSHostnameBlock() *IpamsvcDDNSHostnameBlock {
	this := IpamsvcDDNSHostnameBlock{}
	return &this
}

// NewIpamsvcDDNSHostnameBlockWithDefaults instantiates a new IpamsvcDDNSHostnameBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcDDNSHostnameBlockWithDefaults() *IpamsvcDDNSHostnameBlock {
	this := IpamsvcDDNSHostnameBlock{}
	return &this
}

// GetDdnsGenerateName returns the DdnsGenerateName field value if set, zero value otherwise.
func (o *IpamsvcDDNSHostnameBlock) GetDdnsGenerateName() bool {
	if o == nil || IsNil(o.DdnsGenerateName) {
		var ret bool
		return ret
	}
	return *o.DdnsGenerateName
}

// GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDDNSHostnameBlock) GetDdnsGenerateNameOk() (*bool, bool) {
	if o == nil || IsNil(o.DdnsGenerateName) {
		return nil, false
	}
	return o.DdnsGenerateName, true
}

// HasDdnsGenerateName returns a boolean if a field has been set.
func (o *IpamsvcDDNSHostnameBlock) HasDdnsGenerateName() bool {
	if o != nil && !IsNil(o.DdnsGenerateName) {
		return true
	}

	return false
}

// SetDdnsGenerateName gets a reference to the given bool and assigns it to the DdnsGenerateName field.
func (o *IpamsvcDDNSHostnameBlock) SetDdnsGenerateName(v bool) {
	o.DdnsGenerateName = &v
}

// GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field value if set, zero value otherwise.
func (o *IpamsvcDDNSHostnameBlock) GetDdnsGeneratedPrefix() string {
	if o == nil || IsNil(o.DdnsGeneratedPrefix) {
		var ret string
		return ret
	}
	return *o.DdnsGeneratedPrefix
}

// GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDDNSHostnameBlock) GetDdnsGeneratedPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.DdnsGeneratedPrefix) {
		return nil, false
	}
	return o.DdnsGeneratedPrefix, true
}

// HasDdnsGeneratedPrefix returns a boolean if a field has been set.
func (o *IpamsvcDDNSHostnameBlock) HasDdnsGeneratedPrefix() bool {
	if o != nil && !IsNil(o.DdnsGeneratedPrefix) {
		return true
	}

	return false
}

// SetDdnsGeneratedPrefix gets a reference to the given string and assigns it to the DdnsGeneratedPrefix field.
func (o *IpamsvcDDNSHostnameBlock) SetDdnsGeneratedPrefix(v string) {
	o.DdnsGeneratedPrefix = &v
}

func (o IpamsvcDDNSHostnameBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcDDNSHostnameBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdnsGenerateName) {
		toSerialize["ddns_generate_name"] = o.DdnsGenerateName
	}
	if !IsNil(o.DdnsGeneratedPrefix) {
		toSerialize["ddns_generated_prefix"] = o.DdnsGeneratedPrefix
	}
	return toSerialize, nil
}

type NullableIpamsvcDDNSHostnameBlock struct {
	value *IpamsvcDDNSHostnameBlock
	isSet bool
}

func (v NullableIpamsvcDDNSHostnameBlock) Get() *IpamsvcDDNSHostnameBlock {
	return v.value
}

func (v *NullableIpamsvcDDNSHostnameBlock) Set(val *IpamsvcDDNSHostnameBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcDDNSHostnameBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcDDNSHostnameBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcDDNSHostnameBlock(val *IpamsvcDDNSHostnameBlock) *NullableIpamsvcDDNSHostnameBlock {
	return &NullableIpamsvcDDNSHostnameBlock{value: val, isSet: true}
}

func (v NullableIpamsvcDDNSHostnameBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcDDNSHostnameBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
