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

// checks if the ConfigACLItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigACLItem{}

// ConfigACLItem Element in an ACL.   Error if both _acl_ and _address_ are given.
type ConfigACLItem struct {
	// Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
	Access string `json:"access"`
	// The resource identifier.
	Acl *string `json:"acl,omitempty"`
	// Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
	Address *string `json:"address,omitempty"`
	// Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.
	Element string         `json:"element"`
	TsigKey *ConfigTSIGKey `json:"tsig_key,omitempty"`
}

// NewConfigACLItem instantiates a new ConfigACLItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigACLItem(access string, element string) *ConfigACLItem {
	this := ConfigACLItem{}
	this.Access = access
	this.Element = element
	return &this
}

// NewConfigACLItemWithDefaults instantiates a new ConfigACLItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigACLItemWithDefaults() *ConfigACLItem {
	this := ConfigACLItem{}
	return &this
}

// GetAccess returns the Access field value
func (o *ConfigACLItem) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ConfigACLItem) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ConfigACLItem) SetAccess(v string) {
	o.Access = v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *ConfigACLItem) GetAcl() string {
	if o == nil || IsNil(o.Acl) {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigACLItem) GetAclOk() (*string, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *ConfigACLItem) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *ConfigACLItem) SetAcl(v string) {
	o.Acl = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ConfigACLItem) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigACLItem) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ConfigACLItem) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ConfigACLItem) SetAddress(v string) {
	o.Address = &v
}

// GetElement returns the Element field value
func (o *ConfigACLItem) GetElement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *ConfigACLItem) GetElementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element, true
}

// SetElement sets field value
func (o *ConfigACLItem) SetElement(v string) {
	o.Element = v
}

// GetTsigKey returns the TsigKey field value if set, zero value otherwise.
func (o *ConfigACLItem) GetTsigKey() ConfigTSIGKey {
	if o == nil || IsNil(o.TsigKey) {
		var ret ConfigTSIGKey
		return ret
	}
	return *o.TsigKey
}

// GetTsigKeyOk returns a tuple with the TsigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigACLItem) GetTsigKeyOk() (*ConfigTSIGKey, bool) {
	if o == nil || IsNil(o.TsigKey) {
		return nil, false
	}
	return o.TsigKey, true
}

// HasTsigKey returns a boolean if a field has been set.
func (o *ConfigACLItem) HasTsigKey() bool {
	if o != nil && !IsNil(o.TsigKey) {
		return true
	}

	return false
}

// SetTsigKey gets a reference to the given ConfigTSIGKey and assigns it to the TsigKey field.
func (o *ConfigACLItem) SetTsigKey(v ConfigTSIGKey) {
	o.TsigKey = &v
}

func (o ConfigACLItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigACLItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["element"] = o.Element
	if !IsNil(o.TsigKey) {
		toSerialize["tsig_key"] = o.TsigKey
	}
	return toSerialize, nil
}

type NullableConfigACLItem struct {
	value *ConfigACLItem
	isSet bool
}

func (v NullableConfigACLItem) Get() *ConfigACLItem {
	return v.value
}

func (v *NullableConfigACLItem) Set(val *ConfigACLItem) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigACLItem) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigACLItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigACLItem(val *ConfigACLItem) *NullableConfigACLItem {
	return &NullableConfigACLItem{value: val, isSet: true}
}

func (v NullableConfigACLItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigACLItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
