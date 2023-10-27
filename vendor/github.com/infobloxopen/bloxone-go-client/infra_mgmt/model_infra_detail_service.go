/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_mgmt

import (
	"encoding/json"
	"time"
)

// checks if the InfraDetailService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraDetailService{}

// InfraDetailService struct for InfraDetailService
type InfraDetailService struct {
	// Composite State of this Service (`started`, `stopped`, `stopping`, `starting`, `error`).
	CompositeState *string `json:"composite_state,omitempty"`
	// Composite Status of this Service (`online`, `stopped`, `degraded`, `error`).
	CompositeStatus *string `json:"composite_status,omitempty"`
	// Timestamp of creation of Service.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Current version of this Service.
	CurrentVersion *string `json:"current_version,omitempty"`
	// The description of the Service.
	Description *string `json:"description,omitempty"`
	// The desired state of the Service (`\"start\"` or `\"stop\"`).
	DesiredState *string `json:"desired_state,omitempty"`
	// The desired version of the Service.
	DesiredVersion *string `json:"desired_version,omitempty"`
	// Configuration for the interfaces through which this Service can send outgoing traffic.
	Destinations map[string]map[string]interface{} `json:"destinations,omitempty"`
	// List of Hosts on which this Service is deployed.
	Hosts []InfraDetailServiceHost `json:"hosts,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// List of interfaces on which this Service can operate.
	InterfaceLabels []string             `json:"interface_labels,omitempty"`
	Location        *InfraDetailLocation `json:"location,omitempty"`
	// The name of the Service.
	Name *string        `json:"name,omitempty"`
	Pool *InfraPoolInfo `json:"pool,omitempty"`
	// The type of the Service deployed on the Host (`dns`, `cdc`, etc.).
	ServiceType *string `json:"service_type,omitempty"`
	// Configuration for the interfaces through which this Service can take incoming traffic.
	SourceInterfaces map[string]map[string]interface{} `json:"source_interfaces,omitempty"`
	// Tags associated with this Service.
	Tags map[string]map[string]interface{} `json:"tags,omitempty"`
	// Timestamp of the latest update on Service.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewInfraDetailService instantiates a new InfraDetailService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraDetailService() *InfraDetailService {
	this := InfraDetailService{}
	return &this
}

// NewInfraDetailServiceWithDefaults instantiates a new InfraDetailService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraDetailServiceWithDefaults() *InfraDetailService {
	this := InfraDetailService{}
	return &this
}

// GetCompositeState returns the CompositeState field value if set, zero value otherwise.
func (o *InfraDetailService) GetCompositeState() string {
	if o == nil || IsNil(o.CompositeState) {
		var ret string
		return ret
	}
	return *o.CompositeState
}

// GetCompositeStateOk returns a tuple with the CompositeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetCompositeStateOk() (*string, bool) {
	if o == nil || IsNil(o.CompositeState) {
		return nil, false
	}
	return o.CompositeState, true
}

// HasCompositeState returns a boolean if a field has been set.
func (o *InfraDetailService) HasCompositeState() bool {
	if o != nil && !IsNil(o.CompositeState) {
		return true
	}

	return false
}

// SetCompositeState gets a reference to the given string and assigns it to the CompositeState field.
func (o *InfraDetailService) SetCompositeState(v string) {
	o.CompositeState = &v
}

// GetCompositeStatus returns the CompositeStatus field value if set, zero value otherwise.
func (o *InfraDetailService) GetCompositeStatus() string {
	if o == nil || IsNil(o.CompositeStatus) {
		var ret string
		return ret
	}
	return *o.CompositeStatus
}

// GetCompositeStatusOk returns a tuple with the CompositeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetCompositeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompositeStatus) {
		return nil, false
	}
	return o.CompositeStatus, true
}

// HasCompositeStatus returns a boolean if a field has been set.
func (o *InfraDetailService) HasCompositeStatus() bool {
	if o != nil && !IsNil(o.CompositeStatus) {
		return true
	}

	return false
}

// SetCompositeStatus gets a reference to the given string and assigns it to the CompositeStatus field.
func (o *InfraDetailService) SetCompositeStatus(v string) {
	o.CompositeStatus = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InfraDetailService) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InfraDetailService) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InfraDetailService) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InfraDetailService) GetCurrentVersion() string {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetCurrentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InfraDetailService) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *InfraDetailService) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfraDetailService) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfraDetailService) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfraDetailService) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredState returns the DesiredState field value if set, zero value otherwise.
func (o *InfraDetailService) GetDesiredState() string {
	if o == nil || IsNil(o.DesiredState) {
		var ret string
		return ret
	}
	return *o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetDesiredStateOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredState) {
		return nil, false
	}
	return o.DesiredState, true
}

// HasDesiredState returns a boolean if a field has been set.
func (o *InfraDetailService) HasDesiredState() bool {
	if o != nil && !IsNil(o.DesiredState) {
		return true
	}

	return false
}

// SetDesiredState gets a reference to the given string and assigns it to the DesiredState field.
func (o *InfraDetailService) SetDesiredState(v string) {
	o.DesiredState = &v
}

// GetDesiredVersion returns the DesiredVersion field value if set, zero value otherwise.
func (o *InfraDetailService) GetDesiredVersion() string {
	if o == nil || IsNil(o.DesiredVersion) {
		var ret string
		return ret
	}
	return *o.DesiredVersion
}

// GetDesiredVersionOk returns a tuple with the DesiredVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetDesiredVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredVersion) {
		return nil, false
	}
	return o.DesiredVersion, true
}

// HasDesiredVersion returns a boolean if a field has been set.
func (o *InfraDetailService) HasDesiredVersion() bool {
	if o != nil && !IsNil(o.DesiredVersion) {
		return true
	}

	return false
}

// SetDesiredVersion gets a reference to the given string and assigns it to the DesiredVersion field.
func (o *InfraDetailService) SetDesiredVersion(v string) {
	o.DesiredVersion = &v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InfraDetailService) GetDestinations() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Destinations) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetDestinationsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Destinations) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InfraDetailService) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given map[string]map[string]interface{} and assigns it to the Destinations field.
func (o *InfraDetailService) SetDestinations(v map[string]map[string]interface{}) {
	o.Destinations = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *InfraDetailService) GetHosts() []InfraDetailServiceHost {
	if o == nil || IsNil(o.Hosts) {
		var ret []InfraDetailServiceHost
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetHostsOk() ([]InfraDetailServiceHost, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *InfraDetailService) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []InfraDetailServiceHost and assigns it to the Hosts field.
func (o *InfraDetailService) SetHosts(v []InfraDetailServiceHost) {
	o.Hosts = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InfraDetailService) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InfraDetailService) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InfraDetailService) SetId(v string) {
	o.Id = &v
}

// GetInterfaceLabels returns the InterfaceLabels field value if set, zero value otherwise.
func (o *InfraDetailService) GetInterfaceLabels() []string {
	if o == nil || IsNil(o.InterfaceLabels) {
		var ret []string
		return ret
	}
	return o.InterfaceLabels
}

// GetInterfaceLabelsOk returns a tuple with the InterfaceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetInterfaceLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.InterfaceLabels) {
		return nil, false
	}
	return o.InterfaceLabels, true
}

// HasInterfaceLabels returns a boolean if a field has been set.
func (o *InfraDetailService) HasInterfaceLabels() bool {
	if o != nil && !IsNil(o.InterfaceLabels) {
		return true
	}

	return false
}

// SetInterfaceLabels gets a reference to the given []string and assigns it to the InterfaceLabels field.
func (o *InfraDetailService) SetInterfaceLabels(v []string) {
	o.InterfaceLabels = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InfraDetailService) GetLocation() InfraDetailLocation {
	if o == nil || IsNil(o.Location) {
		var ret InfraDetailLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetLocationOk() (*InfraDetailLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InfraDetailService) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given InfraDetailLocation and assigns it to the Location field.
func (o *InfraDetailService) SetLocation(v InfraDetailLocation) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfraDetailService) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InfraDetailService) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfraDetailService) SetName(v string) {
	o.Name = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *InfraDetailService) GetPool() InfraPoolInfo {
	if o == nil || IsNil(o.Pool) {
		var ret InfraPoolInfo
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetPoolOk() (*InfraPoolInfo, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *InfraDetailService) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given InfraPoolInfo and assigns it to the Pool field.
func (o *InfraDetailService) SetPool(v InfraPoolInfo) {
	o.Pool = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *InfraDetailService) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *InfraDetailService) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *InfraDetailService) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetSourceInterfaces returns the SourceInterfaces field value if set, zero value otherwise.
func (o *InfraDetailService) GetSourceInterfaces() map[string]map[string]interface{} {
	if o == nil || IsNil(o.SourceInterfaces) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.SourceInterfaces
}

// GetSourceInterfacesOk returns a tuple with the SourceInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetSourceInterfacesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceInterfaces) {
		return map[string]map[string]interface{}{}, false
	}
	return o.SourceInterfaces, true
}

// HasSourceInterfaces returns a boolean if a field has been set.
func (o *InfraDetailService) HasSourceInterfaces() bool {
	if o != nil && !IsNil(o.SourceInterfaces) {
		return true
	}

	return false
}

// SetSourceInterfaces gets a reference to the given map[string]map[string]interface{} and assigns it to the SourceInterfaces field.
func (o *InfraDetailService) SetSourceInterfaces(v map[string]map[string]interface{}) {
	o.SourceInterfaces = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InfraDetailService) GetTags() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetTagsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InfraDetailService) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]map[string]interface{} and assigns it to the Tags field.
func (o *InfraDetailService) SetTags(v map[string]map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InfraDetailService) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailService) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InfraDetailService) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InfraDetailService) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InfraDetailService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraDetailService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompositeState) {
		toSerialize["composite_state"] = o.CompositeState
	}
	if !IsNil(o.CompositeStatus) {
		toSerialize["composite_status"] = o.CompositeStatus
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CurrentVersion) {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DesiredState) {
		toSerialize["desired_state"] = o.DesiredState
	}
	if !IsNil(o.DesiredVersion) {
		toSerialize["desired_version"] = o.DesiredVersion
	}
	if !IsNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InterfaceLabels) {
		toSerialize["interface_labels"] = o.InterfaceLabels
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.SourceInterfaces) {
		toSerialize["source_interfaces"] = o.SourceInterfaces
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableInfraDetailService struct {
	value *InfraDetailService
	isSet bool
}

func (v NullableInfraDetailService) Get() *InfraDetailService {
	return v.value
}

func (v *NullableInfraDetailService) Set(val *InfraDetailService) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraDetailService) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraDetailService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraDetailService(val *InfraDetailService) *NullableInfraDetailService {
	return &NullableInfraDetailService{value: val, isSet: true}
}

func (v NullableInfraDetailService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraDetailService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}