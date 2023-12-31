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

// checks if the IpamsvcInheritedASMConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcInheritedASMConfig{}

// IpamsvcInheritedASMConfig The inheritance configuration for the __ASMConfig__ object.
type IpamsvcInheritedASMConfig struct {
	AsmEnableBlock *IpamsvcInheritedAsmEnableBlock `json:"asm_enable_block,omitempty"`
	AsmGrowthBlock *IpamsvcInheritedAsmGrowthBlock `json:"asm_growth_block,omitempty"`
	AsmThreshold   *InheritanceInheritedUInt32     `json:"asm_threshold,omitempty"`
	ForecastPeriod *InheritanceInheritedUInt32     `json:"forecast_period,omitempty"`
	History        *InheritanceInheritedUInt32     `json:"history,omitempty"`
	MinTotal       *InheritanceInheritedUInt32     `json:"min_total,omitempty"`
	MinUnused      *InheritanceInheritedUInt32     `json:"min_unused,omitempty"`
}

// NewIpamsvcInheritedASMConfig instantiates a new IpamsvcInheritedASMConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcInheritedASMConfig() *IpamsvcInheritedASMConfig {
	this := IpamsvcInheritedASMConfig{}
	return &this
}

// NewIpamsvcInheritedASMConfigWithDefaults instantiates a new IpamsvcInheritedASMConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcInheritedASMConfigWithDefaults() *IpamsvcInheritedASMConfig {
	this := IpamsvcInheritedASMConfig{}
	return &this
}

// GetAsmEnableBlock returns the AsmEnableBlock field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetAsmEnableBlock() IpamsvcInheritedAsmEnableBlock {
	if o == nil || IsNil(o.AsmEnableBlock) {
		var ret IpamsvcInheritedAsmEnableBlock
		return ret
	}
	return *o.AsmEnableBlock
}

// GetAsmEnableBlockOk returns a tuple with the AsmEnableBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetAsmEnableBlockOk() (*IpamsvcInheritedAsmEnableBlock, bool) {
	if o == nil || IsNil(o.AsmEnableBlock) {
		return nil, false
	}
	return o.AsmEnableBlock, true
}

// HasAsmEnableBlock returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasAsmEnableBlock() bool {
	if o != nil && !IsNil(o.AsmEnableBlock) {
		return true
	}

	return false
}

// SetAsmEnableBlock gets a reference to the given IpamsvcInheritedAsmEnableBlock and assigns it to the AsmEnableBlock field.
func (o *IpamsvcInheritedASMConfig) SetAsmEnableBlock(v IpamsvcInheritedAsmEnableBlock) {
	o.AsmEnableBlock = &v
}

// GetAsmGrowthBlock returns the AsmGrowthBlock field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetAsmGrowthBlock() IpamsvcInheritedAsmGrowthBlock {
	if o == nil || IsNil(o.AsmGrowthBlock) {
		var ret IpamsvcInheritedAsmGrowthBlock
		return ret
	}
	return *o.AsmGrowthBlock
}

// GetAsmGrowthBlockOk returns a tuple with the AsmGrowthBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetAsmGrowthBlockOk() (*IpamsvcInheritedAsmGrowthBlock, bool) {
	if o == nil || IsNil(o.AsmGrowthBlock) {
		return nil, false
	}
	return o.AsmGrowthBlock, true
}

// HasAsmGrowthBlock returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasAsmGrowthBlock() bool {
	if o != nil && !IsNil(o.AsmGrowthBlock) {
		return true
	}

	return false
}

// SetAsmGrowthBlock gets a reference to the given IpamsvcInheritedAsmGrowthBlock and assigns it to the AsmGrowthBlock field.
func (o *IpamsvcInheritedASMConfig) SetAsmGrowthBlock(v IpamsvcInheritedAsmGrowthBlock) {
	o.AsmGrowthBlock = &v
}

// GetAsmThreshold returns the AsmThreshold field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetAsmThreshold() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.AsmThreshold) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.AsmThreshold
}

// GetAsmThresholdOk returns a tuple with the AsmThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetAsmThresholdOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.AsmThreshold) {
		return nil, false
	}
	return o.AsmThreshold, true
}

// HasAsmThreshold returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasAsmThreshold() bool {
	if o != nil && !IsNil(o.AsmThreshold) {
		return true
	}

	return false
}

// SetAsmThreshold gets a reference to the given InheritanceInheritedUInt32 and assigns it to the AsmThreshold field.
func (o *IpamsvcInheritedASMConfig) SetAsmThreshold(v InheritanceInheritedUInt32) {
	o.AsmThreshold = &v
}

// GetForecastPeriod returns the ForecastPeriod field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetForecastPeriod() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.ForecastPeriod) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.ForecastPeriod
}

// GetForecastPeriodOk returns a tuple with the ForecastPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetForecastPeriodOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.ForecastPeriod) {
		return nil, false
	}
	return o.ForecastPeriod, true
}

// HasForecastPeriod returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasForecastPeriod() bool {
	if o != nil && !IsNil(o.ForecastPeriod) {
		return true
	}

	return false
}

// SetForecastPeriod gets a reference to the given InheritanceInheritedUInt32 and assigns it to the ForecastPeriod field.
func (o *IpamsvcInheritedASMConfig) SetForecastPeriod(v InheritanceInheritedUInt32) {
	o.ForecastPeriod = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetHistory() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.History) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetHistoryOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given InheritanceInheritedUInt32 and assigns it to the History field.
func (o *IpamsvcInheritedASMConfig) SetHistory(v InheritanceInheritedUInt32) {
	o.History = &v
}

// GetMinTotal returns the MinTotal field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetMinTotal() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.MinTotal) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.MinTotal
}

// GetMinTotalOk returns a tuple with the MinTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetMinTotalOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.MinTotal) {
		return nil, false
	}
	return o.MinTotal, true
}

// HasMinTotal returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasMinTotal() bool {
	if o != nil && !IsNil(o.MinTotal) {
		return true
	}

	return false
}

// SetMinTotal gets a reference to the given InheritanceInheritedUInt32 and assigns it to the MinTotal field.
func (o *IpamsvcInheritedASMConfig) SetMinTotal(v InheritanceInheritedUInt32) {
	o.MinTotal = &v
}

// GetMinUnused returns the MinUnused field value if set, zero value otherwise.
func (o *IpamsvcInheritedASMConfig) GetMinUnused() InheritanceInheritedUInt32 {
	if o == nil || IsNil(o.MinUnused) {
		var ret InheritanceInheritedUInt32
		return ret
	}
	return *o.MinUnused
}

// GetMinUnusedOk returns a tuple with the MinUnused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedASMConfig) GetMinUnusedOk() (*InheritanceInheritedUInt32, bool) {
	if o == nil || IsNil(o.MinUnused) {
		return nil, false
	}
	return o.MinUnused, true
}

// HasMinUnused returns a boolean if a field has been set.
func (o *IpamsvcInheritedASMConfig) HasMinUnused() bool {
	if o != nil && !IsNil(o.MinUnused) {
		return true
	}

	return false
}

// SetMinUnused gets a reference to the given InheritanceInheritedUInt32 and assigns it to the MinUnused field.
func (o *IpamsvcInheritedASMConfig) SetMinUnused(v InheritanceInheritedUInt32) {
	o.MinUnused = &v
}

func (o IpamsvcInheritedASMConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcInheritedASMConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AsmEnableBlock) {
		toSerialize["asm_enable_block"] = o.AsmEnableBlock
	}
	if !IsNil(o.AsmGrowthBlock) {
		toSerialize["asm_growth_block"] = o.AsmGrowthBlock
	}
	if !IsNil(o.AsmThreshold) {
		toSerialize["asm_threshold"] = o.AsmThreshold
	}
	if !IsNil(o.ForecastPeriod) {
		toSerialize["forecast_period"] = o.ForecastPeriod
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.MinTotal) {
		toSerialize["min_total"] = o.MinTotal
	}
	if !IsNil(o.MinUnused) {
		toSerialize["min_unused"] = o.MinUnused
	}
	return toSerialize, nil
}

type NullableIpamsvcInheritedASMConfig struct {
	value *IpamsvcInheritedASMConfig
	isSet bool
}

func (v NullableIpamsvcInheritedASMConfig) Get() *IpamsvcInheritedASMConfig {
	return v.value
}

func (v *NullableIpamsvcInheritedASMConfig) Set(val *IpamsvcInheritedASMConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcInheritedASMConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcInheritedASMConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcInheritedASMConfig(val *IpamsvcInheritedASMConfig) *NullableIpamsvcInheritedASMConfig {
	return &NullableIpamsvcInheritedASMConfig{value: val, isSet: true}
}

func (v NullableIpamsvcInheritedASMConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcInheritedASMConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
