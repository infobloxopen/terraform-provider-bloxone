/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"time"
)

// checks if the IpamsvcHardwareFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcHardwareFilter{}

// IpamsvcHardwareFilter A __HardwareFilter__ object (_dhcp/hardware_filter_) applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.
type IpamsvcHardwareFilter struct {
	// The list of addresses to match for the hardware filter.
	Addresses []string `json:"addresses,omitempty"`
	// The description for the hardware filter. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The list of DHCP options for the hardware filter. May be either a specific option or a group of options.
	DhcpOptions []IpamsvcOptionItem `json:"dhcp_options,omitempty"`
	// The configuration for header option filename field.
	HeaderOptionFilename *string `json:"header_option_filename,omitempty"`
	// The configuration for header option server address field.
	HeaderOptionServerAddress *string `json:"header_option_server_address,omitempty"`
	// The configuration for header option server name field.
	HeaderOptionServerName *string `json:"header_option_server_name,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The lease lifetime duration in seconds.
	LeaseTime *int64 `json:"lease_time,omitempty"`
	// The name of the hardware filter. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.
	Role *string `json:"role,omitempty"`
	// The tags for the hardware filter in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The resource identifier.
	VendorSpecificOptionOptionSpace *string `json:"vendor_specific_option_option_space,omitempty"`
}

// NewIpamsvcHardwareFilter instantiates a new IpamsvcHardwareFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcHardwareFilter(name string) *IpamsvcHardwareFilter {
	this := IpamsvcHardwareFilter{}
	this.Name = name
	return &this
}

// NewIpamsvcHardwareFilterWithDefaults instantiates a new IpamsvcHardwareFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcHardwareFilterWithDefaults() *IpamsvcHardwareFilter {
	this := IpamsvcHardwareFilter{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *IpamsvcHardwareFilter) SetAddresses(v []string) {
	o.Addresses = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcHardwareFilter) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcHardwareFilter) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetDhcpOptions() []IpamsvcOptionItem {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []IpamsvcOptionItem
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetDhcpOptionsOk() ([]IpamsvcOptionItem, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []IpamsvcOptionItem and assigns it to the DhcpOptions field.
func (o *IpamsvcHardwareFilter) SetDhcpOptions(v []IpamsvcOptionItem) {
	o.DhcpOptions = v
}

// GetHeaderOptionFilename returns the HeaderOptionFilename field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetHeaderOptionFilename() string {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		var ret string
		return ret
	}
	return *o.HeaderOptionFilename
}

// GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetHeaderOptionFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		return nil, false
	}
	return o.HeaderOptionFilename, true
}

// HasHeaderOptionFilename returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasHeaderOptionFilename() bool {
	if o != nil && !IsNil(o.HeaderOptionFilename) {
		return true
	}

	return false
}

// SetHeaderOptionFilename gets a reference to the given string and assigns it to the HeaderOptionFilename field.
func (o *IpamsvcHardwareFilter) SetHeaderOptionFilename(v string) {
	o.HeaderOptionFilename = &v
}

// GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetHeaderOptionServerAddress() string {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		var ret string
		return ret
	}
	return *o.HeaderOptionServerAddress
}

// GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetHeaderOptionServerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		return nil, false
	}
	return o.HeaderOptionServerAddress, true
}

// HasHeaderOptionServerAddress returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasHeaderOptionServerAddress() bool {
	if o != nil && !IsNil(o.HeaderOptionServerAddress) {
		return true
	}

	return false
}

// SetHeaderOptionServerAddress gets a reference to the given string and assigns it to the HeaderOptionServerAddress field.
func (o *IpamsvcHardwareFilter) SetHeaderOptionServerAddress(v string) {
	o.HeaderOptionServerAddress = &v
}

// GetHeaderOptionServerName returns the HeaderOptionServerName field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetHeaderOptionServerName() string {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		var ret string
		return ret
	}
	return *o.HeaderOptionServerName
}

// GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetHeaderOptionServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		return nil, false
	}
	return o.HeaderOptionServerName, true
}

// HasHeaderOptionServerName returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasHeaderOptionServerName() bool {
	if o != nil && !IsNil(o.HeaderOptionServerName) {
		return true
	}

	return false
}

// SetHeaderOptionServerName gets a reference to the given string and assigns it to the HeaderOptionServerName field.
func (o *IpamsvcHardwareFilter) SetHeaderOptionServerName(v string) {
	o.HeaderOptionServerName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcHardwareFilter) SetId(v string) {
	o.Id = &v
}

// GetLeaseTime returns the LeaseTime field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetLeaseTime() int64 {
	if o == nil || IsNil(o.LeaseTime) {
		var ret int64
		return ret
	}
	return *o.LeaseTime
}

// GetLeaseTimeOk returns a tuple with the LeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetLeaseTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LeaseTime) {
		return nil, false
	}
	return o.LeaseTime, true
}

// HasLeaseTime returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasLeaseTime() bool {
	if o != nil && !IsNil(o.LeaseTime) {
		return true
	}

	return false
}

// SetLeaseTime gets a reference to the given int64 and assigns it to the LeaseTime field.
func (o *IpamsvcHardwareFilter) SetLeaseTime(v int64) {
	o.LeaseTime = &v
}

// GetName returns the Name field value
func (o *IpamsvcHardwareFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcHardwareFilter) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *IpamsvcHardwareFilter) SetRole(v string) {
	o.Role = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcHardwareFilter) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcHardwareFilter) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field value if set, zero value otherwise.
func (o *IpamsvcHardwareFilter) GetVendorSpecificOptionOptionSpace() string {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		var ret string
		return ret
	}
	return *o.VendorSpecificOptionOptionSpace
}

// GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHardwareFilter) GetVendorSpecificOptionOptionSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		return nil, false
	}
	return o.VendorSpecificOptionOptionSpace, true
}

// HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.
func (o *IpamsvcHardwareFilter) HasVendorSpecificOptionOptionSpace() bool {
	if o != nil && !IsNil(o.VendorSpecificOptionOptionSpace) {
		return true
	}

	return false
}

// SetVendorSpecificOptionOptionSpace gets a reference to the given string and assigns it to the VendorSpecificOptionOptionSpace field.
func (o *IpamsvcHardwareFilter) SetVendorSpecificOptionOptionSpace(v string) {
	o.VendorSpecificOptionOptionSpace = &v
}

func (o IpamsvcHardwareFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcHardwareFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.HeaderOptionFilename) {
		toSerialize["header_option_filename"] = o.HeaderOptionFilename
	}
	if !IsNil(o.HeaderOptionServerAddress) {
		toSerialize["header_option_server_address"] = o.HeaderOptionServerAddress
	}
	if !IsNil(o.HeaderOptionServerName) {
		toSerialize["header_option_server_name"] = o.HeaderOptionServerName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LeaseTime) {
		toSerialize["lease_time"] = o.LeaseTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VendorSpecificOptionOptionSpace) {
		toSerialize["vendor_specific_option_option_space"] = o.VendorSpecificOptionOptionSpace
	}
	return toSerialize, nil
}

type NullableIpamsvcHardwareFilter struct {
	value *IpamsvcHardwareFilter
	isSet bool
}

func (v NullableIpamsvcHardwareFilter) Get() *IpamsvcHardwareFilter {
	return v.value
}

func (v *NullableIpamsvcHardwareFilter) Set(val *IpamsvcHardwareFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcHardwareFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcHardwareFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcHardwareFilter(val *IpamsvcHardwareFilter) *NullableIpamsvcHardwareFilter {
	return &NullableIpamsvcHardwareFilter{value: val, isSet: true}
}

func (v NullableIpamsvcHardwareFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcHardwareFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
