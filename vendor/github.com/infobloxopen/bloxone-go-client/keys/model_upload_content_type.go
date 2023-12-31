/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
	"fmt"
)

// UploadContentType  - UNKNOWN: Unknown type.  - KEYTAB: Keytab file containing Kerberos keys.
type UploadContentType string

// List of uploadContentType
const (
	UPLOADCONTENTTYPE_UNKNOWN UploadContentType = "UNKNOWN"
	UPLOADCONTENTTYPE_KEYTAB  UploadContentType = "KEYTAB"
)

// All allowed values of UploadContentType enum
var AllowedUploadContentTypeEnumValues = []UploadContentType{
	"UNKNOWN",
	"KEYTAB",
}

func (v *UploadContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UploadContentType(value)
	for _, existing := range AllowedUploadContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UploadContentType", value)
}

// NewUploadContentTypeFromValue returns a pointer to a valid UploadContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUploadContentTypeFromValue(v string) (*UploadContentType, error) {
	ev := UploadContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UploadContentType: valid values are %v", v, AllowedUploadContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UploadContentType) IsValid() bool {
	for _, existing := range AllowedUploadContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to uploadContentType value
func (v UploadContentType) Ptr() *UploadContentType {
	return &v
}

type NullableUploadContentType struct {
	value *UploadContentType
	isSet bool
}

func (v NullableUploadContentType) Get() *UploadContentType {
	return v.value
}

func (v *NullableUploadContentType) Set(val *UploadContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadContentType(val *UploadContentType) *NullableUploadContentType {
	return &NullableUploadContentType{value: val, isSet: true}
}

func (v NullableUploadContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
