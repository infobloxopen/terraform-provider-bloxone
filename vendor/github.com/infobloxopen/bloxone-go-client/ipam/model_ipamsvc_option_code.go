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

// checks if the IpamsvcOptionCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcOptionCode{}

// IpamsvcOptionCode An __OptionCode__ (_dhcp/option_code_) defines a DHCP option code.
type IpamsvcOptionCode struct {
	// Indicates whether the option value is an array of the type or not.
	Array *bool `json:"array,omitempty"`
	// The option code.
	Code int64 `json:"code"`
	// The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// The resource identifier.
	OptionSpace string `json:"option_space"`
	// The source for the option code.  Valid values are:  * _dhcp_server_  * _reserved_  * _blox_one_ddi_  * _customer_  Defaults to _customer_.
	Source *string `json:"source,omitempty"`
	// The option type for the option code.  Valid values are: * _address4_ * _address6_ * _boolean_ * _empty_ * _fqdn_ * _int8_ * _int16_ * _int32_ * _text_ * _uint8_ * _uint16_ * _uint32_
	Type string `json:"type"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIpamsvcOptionCode instantiates a new IpamsvcOptionCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcOptionCode(code int64, name string, optionSpace string, type_ string) *IpamsvcOptionCode {
	this := IpamsvcOptionCode{}
	this.Code = code
	this.Name = name
	this.OptionSpace = optionSpace
	this.Type = type_
	return &this
}

// NewIpamsvcOptionCodeWithDefaults instantiates a new IpamsvcOptionCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcOptionCodeWithDefaults() *IpamsvcOptionCode {
	this := IpamsvcOptionCode{}
	return &this
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetArray() bool {
	if o == nil || IsNil(o.Array) {
		var ret bool
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetArrayOk() (*bool, bool) {
	if o == nil || IsNil(o.Array) {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasArray() bool {
	if o != nil && !IsNil(o.Array) {
		return true
	}

	return false
}

// SetArray gets a reference to the given bool and assigns it to the Array field.
func (o *IpamsvcOptionCode) SetArray(v bool) {
	o.Array = &v
}

// GetCode returns the Code field value
func (o *IpamsvcOptionCode) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *IpamsvcOptionCode) SetCode(v int64) {
	o.Code = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcOptionCode) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcOptionCode) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcOptionCode) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IpamsvcOptionCode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcOptionCode) SetName(v string) {
	o.Name = v
}

// GetOptionSpace returns the OptionSpace field value
func (o *IpamsvcOptionCode) GetOptionSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptionSpace
}

// GetOptionSpaceOk returns a tuple with the OptionSpace field value
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetOptionSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptionSpace, true
}

// SetOptionSpace sets field value
func (o *IpamsvcOptionCode) SetOptionSpace(v string) {
	o.OptionSpace = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *IpamsvcOptionCode) SetSource(v string) {
	o.Source = &v
}

// GetType returns the Type field value
func (o *IpamsvcOptionCode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IpamsvcOptionCode) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcOptionCode) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionCode) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcOptionCode) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcOptionCode) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IpamsvcOptionCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcOptionCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Array) {
		toSerialize["array"] = o.Array
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["option_space"] = o.OptionSpace
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIpamsvcOptionCode struct {
	value *IpamsvcOptionCode
	isSet bool
}

func (v NullableIpamsvcOptionCode) Get() *IpamsvcOptionCode {
	return v.value
}

func (v *NullableIpamsvcOptionCode) Set(val *IpamsvcOptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcOptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcOptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcOptionCode(val *IpamsvcOptionCode) *NullableIpamsvcOptionCode {
	return &NullableIpamsvcOptionCode{value: val, isSet: true}
}

func (v NullableIpamsvcOptionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcOptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
