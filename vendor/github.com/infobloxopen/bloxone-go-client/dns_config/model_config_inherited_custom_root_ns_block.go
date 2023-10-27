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

// checks if the ConfigInheritedCustomRootNSBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigInheritedCustomRootNSBlock{}

// ConfigInheritedCustomRootNSBlock Inheritance block for fields: _custom_root_ns_enabled_, _custom_root_ns_.
type ConfigInheritedCustomRootNSBlock struct {
	// Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// Human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source *string                  `json:"source,omitempty"`
	Value  *ConfigCustomRootNSBlock `json:"value,omitempty"`
}

// NewConfigInheritedCustomRootNSBlock instantiates a new ConfigInheritedCustomRootNSBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInheritedCustomRootNSBlock() *ConfigInheritedCustomRootNSBlock {
	this := ConfigInheritedCustomRootNSBlock{}
	return &this
}

// NewConfigInheritedCustomRootNSBlockWithDefaults instantiates a new ConfigInheritedCustomRootNSBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInheritedCustomRootNSBlockWithDefaults() *ConfigInheritedCustomRootNSBlock {
	this := ConfigInheritedCustomRootNSBlock{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ConfigInheritedCustomRootNSBlock) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedCustomRootNSBlock) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ConfigInheritedCustomRootNSBlock) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ConfigInheritedCustomRootNSBlock) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConfigInheritedCustomRootNSBlock) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedCustomRootNSBlock) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConfigInheritedCustomRootNSBlock) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConfigInheritedCustomRootNSBlock) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConfigInheritedCustomRootNSBlock) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedCustomRootNSBlock) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConfigInheritedCustomRootNSBlock) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConfigInheritedCustomRootNSBlock) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigInheritedCustomRootNSBlock) GetValue() ConfigCustomRootNSBlock {
	if o == nil || IsNil(o.Value) {
		var ret ConfigCustomRootNSBlock
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedCustomRootNSBlock) GetValueOk() (*ConfigCustomRootNSBlock, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigInheritedCustomRootNSBlock) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ConfigCustomRootNSBlock and assigns it to the Value field.
func (o *ConfigInheritedCustomRootNSBlock) SetValue(v ConfigCustomRootNSBlock) {
	o.Value = &v
}

func (o ConfigInheritedCustomRootNSBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigInheritedCustomRootNSBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableConfigInheritedCustomRootNSBlock struct {
	value *ConfigInheritedCustomRootNSBlock
	isSet bool
}

func (v NullableConfigInheritedCustomRootNSBlock) Get() *ConfigInheritedCustomRootNSBlock {
	return v.value
}

func (v *NullableConfigInheritedCustomRootNSBlock) Set(val *ConfigInheritedCustomRootNSBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInheritedCustomRootNSBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInheritedCustomRootNSBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInheritedCustomRootNSBlock(val *ConfigInheritedCustomRootNSBlock) *NullableConfigInheritedCustomRootNSBlock {
	return &NullableConfigInheritedCustomRootNSBlock{value: val, isSet: true}
}

func (v NullableConfigInheritedCustomRootNSBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInheritedCustomRootNSBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}