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

// checks if the IpamsvcInheritedDHCPConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcInheritedDHCPConfig{}

// IpamsvcInheritedDHCPConfig The inheritance configuration for a field of type _DHCPConfig_.
type IpamsvcInheritedDHCPConfig struct {
	AbandonedReclaimTime   *InheritanceInheritedUInt32        `json:"abandoned_reclaim_time,omitempty"`
	AbandonedReclaimTimeV6 *InheritanceInheritedUInt32        `json:"abandoned_reclaim_time_v6,omitempty"`
	AllowUnknown           *InheritanceInheritedBool          `json:"allow_unknown,omitempty"`
	AllowUnknownV6         *InheritanceInheritedBool          `json:"allow_unknown_v6,omitempty"`
	Filters                *InheritedDHCPConfigFilterList     `json:"filters,omitempty"`
	FiltersV6              *InheritedDHCPConfigFilterList     `json:"filters_v6,omitempty"`
	IgnoreClientUid        *InheritanceInheritedBool          `json:"ignore_client_uid,omitempty"`
	IgnoreList             *InheritedDHCPConfigIgnoreItemList `json:"ignore_list,omitempty"`
	LeaseTime              *InheritanceInheritedUInt32        `json:"lease_time,omitempty"`
	LeaseTimeV6            *InheritanceInheritedUInt32        `json:"lease_time_v6,omitempty"`
}

// NewIpamsvcInheritedDHCPConfig instantiates a new IpamsvcInheritedDHCPConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcInheritedDHCPConfig() *IpamsvcInheritedDHCPConfig {
	this := IpamsvcInheritedDHCPConfig{}
	return &this
}

// NewIpamsvcInheritedDHCPConfigWithDefaults instantiates a new IpamsvcInheritedDHCPConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcInheritedDHCPConfigWithDefaults() *IpamsvcInheritedDHCPConfig {
	this := IpamsvcInheritedDHCPConfig{}
	return &this
}

// GetAbandonedReclaimTime returns the AbandonedReclaimTime field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetAbandonedReclaimTime() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.AbandonedReclaimTime) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.AbandonedReclaimTime
}

// GetAbandonedReclaimTimeOk returns a tuple with the AbandonedReclaimTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetAbandonedReclaimTimeOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.AbandonedReclaimTime) {
		return nil, false
	}
	return o.AbandonedReclaimTime, true
}

// HasAbandonedReclaimTime returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasAbandonedReclaimTime() bool {
	if o != nil && !IsNil(o.AbandonedReclaimTime) {
		return true
	}

	return false
}

// SetAbandonedReclaimTime gets a reference to the given InheritanceInheritedUInt32 and assigns it to the AbandonedReclaimTime field.
func (o *IpamsvcInheritedDHCPConfig) SetAbandonedReclaimTime(v InheritanceInheritedUInt32) {
	o.AbandonedReclaimTime = &v
}

// GetAbandonedReclaimTimeV6 returns the AbandonedReclaimTimeV6 field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetAbandonedReclaimTimeV6() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.AbandonedReclaimTimeV6) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.AbandonedReclaimTimeV6
}

// GetAbandonedReclaimTimeV6Ok returns a tuple with the AbandonedReclaimTimeV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetAbandonedReclaimTimeV6Ok() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.AbandonedReclaimTimeV6) {
		return nil, false
	}
	return o.AbandonedReclaimTimeV6, true
}

// HasAbandonedReclaimTimeV6 returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasAbandonedReclaimTimeV6() bool {
	if o != nil && !IsNil(o.AbandonedReclaimTimeV6) {
		return true
	}

	return false
}

// SetAbandonedReclaimTimeV6 gets a reference to the given InheritanceInheritedUInt32 and assigns it to the AbandonedReclaimTimeV6 field.
func (o *IpamsvcInheritedDHCPConfig) SetAbandonedReclaimTimeV6(v InheritanceInheritedUInt32) {
	o.AbandonedReclaimTimeV6 = &v
}

// GetAllowUnknown returns the AllowUnknown field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetAllowUnknown() InheritanceInheritedBool {
	if o == nil || IsNil(o.AllowUnknown) {
		var ret InheritanceInheritedBool
		return ret
	}
	return *o.AllowUnknown
}

// GetAllowUnknownOk returns a tuple with the AllowUnknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetAllowUnknownOk() (*InheritanceInheritedBool, bool) {
	if o == nil || IsNil(o.AllowUnknown) {
		return nil, false
	}
	return o.AllowUnknown, true
}

// HasAllowUnknown returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasAllowUnknown() bool {
	if o != nil && !IsNil(o.AllowUnknown) {
		return true
	}

	return false
}

// SetAllowUnknown gets a reference to the given InheritanceInheritedBool and assigns it to the AllowUnknown field.
func (o *IpamsvcInheritedDHCPConfig) SetAllowUnknown(v InheritanceInheritedBool) {
	o.AllowUnknown = &v
}

// GetAllowUnknownV6 returns the AllowUnknownV6 field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetAllowUnknownV6() InheritanceInheritedBool {
	if o == nil || IsNil(o.AllowUnknownV6) {
		var ret InheritanceInheritedBool
		return ret
	}
	return *o.AllowUnknownV6
}

// GetAllowUnknownV6Ok returns a tuple with the AllowUnknownV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetAllowUnknownV6Ok() (*InheritanceInheritedBool, bool) {
	if o == nil || IsNil(o.AllowUnknownV6) {
		return nil, false
	}
	return o.AllowUnknownV6, true
}

// HasAllowUnknownV6 returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasAllowUnknownV6() bool {
	if o != nil && !IsNil(o.AllowUnknownV6) {
		return true
	}

	return false
}

// SetAllowUnknownV6 gets a reference to the given InheritanceInheritedBool and assigns it to the AllowUnknownV6 field.
func (o *IpamsvcInheritedDHCPConfig) SetAllowUnknownV6(v InheritanceInheritedBool) {
	o.AllowUnknownV6 = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetFilters() InheritedDHCPConfigFilterList {
	if o == nil || IsNil(o.Filters) {
		var ret InheritedDHCPConfigFilterList
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetFiltersOk() (*InheritedDHCPConfigFilterList, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given InheritedDHCPConfigFilterList and assigns it to the Filters field.
func (o *IpamsvcInheritedDHCPConfig) SetFilters(v InheritedDHCPConfigFilterList) {
	o.Filters = &v
}

// GetFiltersV6 returns the FiltersV6 field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetFiltersV6() InheritedDHCPConfigFilterList {
	if o == nil || IsNil(o.FiltersV6) {
		var ret InheritedDHCPConfigFilterList
		return ret
	}
	return *o.FiltersV6
}

// GetFiltersV6Ok returns a tuple with the FiltersV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetFiltersV6Ok() (*InheritedDHCPConfigFilterList, bool) {
	if o == nil || IsNil(o.FiltersV6) {
		return nil, false
	}
	return o.FiltersV6, true
}

// HasFiltersV6 returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasFiltersV6() bool {
	if o != nil && !IsNil(o.FiltersV6) {
		return true
	}

	return false
}

// SetFiltersV6 gets a reference to the given InheritedDHCPConfigFilterList and assigns it to the FiltersV6 field.
func (o *IpamsvcInheritedDHCPConfig) SetFiltersV6(v InheritedDHCPConfigFilterList) {
	o.FiltersV6 = &v
}

// GetIgnoreClientUid returns the IgnoreClientUid field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetIgnoreClientUid() InheritanceInheritedBool {
	if o == nil || IsNil(o.IgnoreClientUid) {
		var ret InheritanceInheritedBool
		return ret
	}
	return *o.IgnoreClientUid
}

// GetIgnoreClientUidOk returns a tuple with the IgnoreClientUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetIgnoreClientUidOk() (*InheritanceInheritedBool, bool) {
	if o == nil || IsNil(o.IgnoreClientUid) {
		return nil, false
	}
	return o.IgnoreClientUid, true
}

// HasIgnoreClientUid returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasIgnoreClientUid() bool {
	if o != nil && !IsNil(o.IgnoreClientUid) {
		return true
	}

	return false
}

// SetIgnoreClientUid gets a reference to the given InheritanceInheritedBool and assigns it to the IgnoreClientUid field.
func (o *IpamsvcInheritedDHCPConfig) SetIgnoreClientUid(v InheritanceInheritedBool) {
	o.IgnoreClientUid = &v
}

// GetIgnoreList returns the IgnoreList field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetIgnoreList() InheritedDHCPConfigIgnoreItemList {
	if o == nil || IsNil(o.IgnoreList) {
		var ret InheritedDHCPConfigIgnoreItemList
		return ret
	}
	return *o.IgnoreList
}

// GetIgnoreListOk returns a tuple with the IgnoreList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetIgnoreListOk() (*InheritedDHCPConfigIgnoreItemList, bool) {
	if o == nil || IsNil(o.IgnoreList) {
		return nil, false
	}
	return o.IgnoreList, true
}

// HasIgnoreList returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasIgnoreList() bool {
	if o != nil && !IsNil(o.IgnoreList) {
		return true
	}

	return false
}

// SetIgnoreList gets a reference to the given InheritedDHCPConfigIgnoreItemList and assigns it to the IgnoreList field.
func (o *IpamsvcInheritedDHCPConfig) SetIgnoreList(v InheritedDHCPConfigIgnoreItemList) {
	o.IgnoreList = &v
}

// GetLeaseTime returns the LeaseTime field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetLeaseTime() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.LeaseTime) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.LeaseTime
}

// GetLeaseTimeOk returns a tuple with the LeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetLeaseTimeOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.LeaseTime) {
		return nil, false
	}
	return o.LeaseTime, true
}

// HasLeaseTime returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasLeaseTime() bool {
	if o != nil && !IsNil(o.LeaseTime) {
		return true
	}

	return false
}

// SetLeaseTime gets a reference to the given InheritanceInheritedUInt32 and assigns it to the LeaseTime field.
func (o *IpamsvcInheritedDHCPConfig) SetLeaseTime(v InheritanceInheritedUInt32) {
	o.LeaseTime = &v
}

// GetLeaseTimeV6 returns the LeaseTimeV6 field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPConfig) GetLeaseTimeV6() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.LeaseTimeV6) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.LeaseTimeV6
}

// GetLeaseTimeV6Ok returns a tuple with the LeaseTimeV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPConfig) GetLeaseTimeV6Ok() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.LeaseTimeV6) {
		return nil, false
	}
	return o.LeaseTimeV6, true
}

// HasLeaseTimeV6 returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPConfig) HasLeaseTimeV6() bool {
	if o != nil && !IsNil(o.LeaseTimeV6) {
		return true
	}

	return false
}

// SetLeaseTimeV6 gets a reference to the given InheritanceInheritedUInt32 and assigns it to the LeaseTimeV6 field.
func (o *IpamsvcInheritedDHCPConfig) SetLeaseTimeV6(v InheritanceInheritedUInt32) {
	o.LeaseTimeV6 = &v
}

func (o IpamsvcInheritedDHCPConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcInheritedDHCPConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbandonedReclaimTime) {
		toSerialize["abandoned_reclaim_time"] = o.AbandonedReclaimTime
	}
	if !IsNil(o.AbandonedReclaimTimeV6) {
		toSerialize["abandoned_reclaim_time_v6"] = o.AbandonedReclaimTimeV6
	}
	if !IsNil(o.AllowUnknown) {
		toSerialize["allow_unknown"] = o.AllowUnknown
	}
	if !IsNil(o.AllowUnknownV6) {
		toSerialize["allow_unknown_v6"] = o.AllowUnknownV6
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.FiltersV6) {
		toSerialize["filters_v6"] = o.FiltersV6
	}
	if !IsNil(o.IgnoreClientUid) {
		toSerialize["ignore_client_uid"] = o.IgnoreClientUid
	}
	if !IsNil(o.IgnoreList) {
		toSerialize["ignore_list"] = o.IgnoreList
	}
	if !IsNil(o.LeaseTime) {
		toSerialize["lease_time"] = o.LeaseTime
	}
	if !IsNil(o.LeaseTimeV6) {
		toSerialize["lease_time_v6"] = o.LeaseTimeV6
	}
	return toSerialize, nil
}

type NullableIpamsvcInheritedDHCPConfig struct {
	value *IpamsvcInheritedDHCPConfig
	isSet bool
}

func (v NullableIpamsvcInheritedDHCPConfig) Get() *IpamsvcInheritedDHCPConfig {
	return v.value
}

func (v *NullableIpamsvcInheritedDHCPConfig) Set(val *IpamsvcInheritedDHCPConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcInheritedDHCPConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcInheritedDHCPConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcInheritedDHCPConfig(val *IpamsvcInheritedDHCPConfig) *NullableIpamsvcInheritedDHCPConfig {
	return &NullableIpamsvcInheritedDHCPConfig{value: val, isSet: true}
}

func (v NullableIpamsvcInheritedDHCPConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcInheritedDHCPConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
