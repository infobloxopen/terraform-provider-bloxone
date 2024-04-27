/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"encoding/json"
)

// checks if the ServiceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceConfig{}

// ServiceConfig struct for ServiceConfig
type ServiceConfig struct {
	Config               *ServiceConfigObject `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceConfig ServiceConfig

// NewServiceConfig instantiates a new ServiceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConfig() *ServiceConfig {
	this := ServiceConfig{}
	return &this
}

// NewServiceConfigWithDefaults instantiates a new ServiceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConfigWithDefaults() *ServiceConfig {
	this := ServiceConfig{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ServiceConfig) GetConfig() ServiceConfigObject {
	if o == nil || IsNil(o.Config) {
		var ret ServiceConfigObject
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConfig) GetConfigOk() (*ServiceConfigObject, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ServiceConfig) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ServiceConfigObject and assigns it to the Config field.
func (o *ServiceConfig) SetConfig(v ServiceConfigObject) {
	o.Config = &v
}

func (o ServiceConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceConfig) UnmarshalJSON(data []byte) (err error) {
	varServiceConfig := _ServiceConfig{}

	err = json.Unmarshal(data, &varServiceConfig)

	if err != nil {
		return err
	}

	*o = ServiceConfig(varServiceConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceConfig struct {
	value *ServiceConfig
	isSet bool
}

func (v NullableServiceConfig) Get() *ServiceConfig {
	return v.value
}

func (v *NullableServiceConfig) Set(val *ServiceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConfig(val *ServiceConfig) *NullableServiceConfig {
	return &NullableServiceConfig{value: val, isSet: true}
}

func (v NullableServiceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}