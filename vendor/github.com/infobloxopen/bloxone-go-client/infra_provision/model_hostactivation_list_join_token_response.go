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

// checks if the HostactivationListJoinTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostactivationListJoinTokenResponse{}

// HostactivationListJoinTokenResponse struct for HostactivationListJoinTokenResponse
type HostactivationListJoinTokenResponse struct {
	Page    *ApiPageInfo              `json:"page,omitempty"`
	Results []HostactivationJoinToken `json:"results,omitempty"`
}

// NewHostactivationListJoinTokenResponse instantiates a new HostactivationListJoinTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostactivationListJoinTokenResponse() *HostactivationListJoinTokenResponse {
	this := HostactivationListJoinTokenResponse{}
	return &this
}

// NewHostactivationListJoinTokenResponseWithDefaults instantiates a new HostactivationListJoinTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostactivationListJoinTokenResponseWithDefaults() *HostactivationListJoinTokenResponse {
	this := HostactivationListJoinTokenResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *HostactivationListJoinTokenResponse) GetPage() ApiPageInfo {
	if o == nil || IsNil(o.Page) {
		var ret ApiPageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationListJoinTokenResponse) GetPageOk() (*ApiPageInfo, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *HostactivationListJoinTokenResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ApiPageInfo and assigns it to the Page field.
func (o *HostactivationListJoinTokenResponse) SetPage(v ApiPageInfo) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HostactivationListJoinTokenResponse) GetResults() []HostactivationJoinToken {
	if o == nil || IsNil(o.Results) {
		var ret []HostactivationJoinToken
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationListJoinTokenResponse) GetResultsOk() ([]HostactivationJoinToken, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HostactivationListJoinTokenResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HostactivationJoinToken and assigns it to the Results field.
func (o *HostactivationListJoinTokenResponse) SetResults(v []HostactivationJoinToken) {
	o.Results = v
}

func (o HostactivationListJoinTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostactivationListJoinTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableHostactivationListJoinTokenResponse struct {
	value *HostactivationListJoinTokenResponse
	isSet bool
}

func (v NullableHostactivationListJoinTokenResponse) Get() *HostactivationListJoinTokenResponse {
	return v.value
}

func (v *NullableHostactivationListJoinTokenResponse) Set(val *HostactivationListJoinTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHostactivationListJoinTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHostactivationListJoinTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostactivationListJoinTokenResponse(val *HostactivationListJoinTokenResponse) *NullableHostactivationListJoinTokenResponse {
	return &NullableHostactivationListJoinTokenResponse{value: val, isSet: true}
}

func (v NullableHostactivationListJoinTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostactivationListJoinTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
