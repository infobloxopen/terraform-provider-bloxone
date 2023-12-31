/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_provision

import (
	"encoding/json"
)

// checks if the HostactivationCreateJoinTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostactivationCreateJoinTokenResponse{}

// HostactivationCreateJoinTokenResponse struct for HostactivationCreateJoinTokenResponse
type HostactivationCreateJoinTokenResponse struct {
	JoinToken *string                  `json:"join_token,omitempty"`
	Result    *HostactivationJoinToken `json:"result,omitempty"`
}

// NewHostactivationCreateJoinTokenResponse instantiates a new HostactivationCreateJoinTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostactivationCreateJoinTokenResponse() *HostactivationCreateJoinTokenResponse {
	this := HostactivationCreateJoinTokenResponse{}
	return &this
}

// NewHostactivationCreateJoinTokenResponseWithDefaults instantiates a new HostactivationCreateJoinTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostactivationCreateJoinTokenResponseWithDefaults() *HostactivationCreateJoinTokenResponse {
	this := HostactivationCreateJoinTokenResponse{}
	return &this
}

// GetJoinToken returns the JoinToken field value if set, zero value otherwise.
func (o *HostactivationCreateJoinTokenResponse) GetJoinToken() string {
	if o == nil || IsNil(o.JoinToken) {
		var ret string
		return ret
	}
	return *o.JoinToken
}

// GetJoinTokenOk returns a tuple with the JoinToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationCreateJoinTokenResponse) GetJoinTokenOk() (*string, bool) {
	if o == nil || IsNil(o.JoinToken) {
		return nil, false
	}
	return o.JoinToken, true
}

// HasJoinToken returns a boolean if a field has been set.
func (o *HostactivationCreateJoinTokenResponse) HasJoinToken() bool {
	if o != nil && !IsNil(o.JoinToken) {
		return true
	}

	return false
}

// SetJoinToken gets a reference to the given string and assigns it to the JoinToken field.
func (o *HostactivationCreateJoinTokenResponse) SetJoinToken(v string) {
	o.JoinToken = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *HostactivationCreateJoinTokenResponse) GetResult() HostactivationJoinToken {
	if o == nil || IsNil(o.Result) {
		var ret HostactivationJoinToken
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationCreateJoinTokenResponse) GetResultOk() (*HostactivationJoinToken, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *HostactivationCreateJoinTokenResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given HostactivationJoinToken and assigns it to the Result field.
func (o *HostactivationCreateJoinTokenResponse) SetResult(v HostactivationJoinToken) {
	o.Result = &v
}

func (o HostactivationCreateJoinTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostactivationCreateJoinTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JoinToken) {
		toSerialize["join_token"] = o.JoinToken
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableHostactivationCreateJoinTokenResponse struct {
	value *HostactivationCreateJoinTokenResponse
	isSet bool
}

func (v NullableHostactivationCreateJoinTokenResponse) Get() *HostactivationCreateJoinTokenResponse {
	return v.value
}

func (v *NullableHostactivationCreateJoinTokenResponse) Set(val *HostactivationCreateJoinTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHostactivationCreateJoinTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHostactivationCreateJoinTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostactivationCreateJoinTokenResponse(val *HostactivationCreateJoinTokenResponse) *NullableHostactivationCreateJoinTokenResponse {
	return &NullableHostactivationCreateJoinTokenResponse{value: val, isSet: true}
}

func (v NullableHostactivationCreateJoinTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostactivationCreateJoinTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
