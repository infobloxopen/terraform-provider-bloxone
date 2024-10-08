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

// checks if the CreateOptionSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOptionSpaceResponse{}

// CreateOptionSpaceResponse The response format to create the __OptionSpace__ object.
type CreateOptionSpaceResponse struct {
	// The created OptionSpace object.
	Result               *OptionSpace `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOptionSpaceResponse CreateOptionSpaceResponse

// NewCreateOptionSpaceResponse instantiates a new CreateOptionSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOptionSpaceResponse() *CreateOptionSpaceResponse {
	this := CreateOptionSpaceResponse{}
	return &this
}

// NewCreateOptionSpaceResponseWithDefaults instantiates a new CreateOptionSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOptionSpaceResponseWithDefaults() *CreateOptionSpaceResponse {
	this := CreateOptionSpaceResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateOptionSpaceResponse) GetResult() OptionSpace {
	if o == nil || IsNil(o.Result) {
		var ret OptionSpace
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOptionSpaceResponse) GetResultOk() (*OptionSpace, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateOptionSpaceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OptionSpace and assigns it to the Result field.
func (o *CreateOptionSpaceResponse) SetResult(v OptionSpace) {
	o.Result = &v
}

func (o CreateOptionSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOptionSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOptionSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateOptionSpaceResponse := _CreateOptionSpaceResponse{}

	err = json.Unmarshal(data, &varCreateOptionSpaceResponse)

	if err != nil {
		return err
	}

	*o = CreateOptionSpaceResponse(varCreateOptionSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOptionSpaceResponse struct {
	value *CreateOptionSpaceResponse
	isSet bool
}

func (v NullableCreateOptionSpaceResponse) Get() *CreateOptionSpaceResponse {
	return v.value
}

func (v *NullableCreateOptionSpaceResponse) Set(val *CreateOptionSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOptionSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOptionSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOptionSpaceResponse(val *CreateOptionSpaceResponse) *NullableCreateOptionSpaceResponse {
	return &NullableCreateOptionSpaceResponse{value: val, isSet: true}
}

func (v NullableCreateOptionSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOptionSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
