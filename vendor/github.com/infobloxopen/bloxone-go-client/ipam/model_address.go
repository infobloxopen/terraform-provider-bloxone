/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address An __Address__ object (_ipam/address_) represents any single IP address within a given IP space.
type Address struct {
	// The address in form \"a.b.c.d\".
	Address string `json:"address"`
	// The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// The compartment associated with the object. If no compartment is associated with the object, the value defaults to empty.
	CompartmentId *string `json:"compartment_id,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The DHCP information associated with this object.
	DhcpInfo *DHCPInfo `json:"dhcp_info,omitempty"`
	// Read only. Represent the value of the same field in the associated _dhcp/fixed_address_ object.
	DisableDhcp *bool `json:"disable_dhcp,omitempty"`
	// The discovery attributes for this address in JSON format.
	DiscoveryAttrs map[string]interface{} `json:"discovery_attrs,omitempty"`
	// The discovery metadata for this address in JSON format.
	DiscoveryMetadata map[string]interface{} `json:"discovery_metadata,omitempty"`
	// The external keys (source key) for this address in JSON format.
	ExternalKeys map[string]interface{} `json:"external_keys,omitempty"`
	// The resource identifier.
	Host *string `json:"host,omitempty"`
	// The hardware address associated with this IP address.
	Hwaddr *string `json:"hwaddr,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the network interface card (NIC) associated with the address, if any.
	Interface *string `json:"interface,omitempty"`
	// The list of all names associated with this address.
	Names []Name `json:"names,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// The type of protocol (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The resource identifier.
	Range *string `json:"range,omitempty"`
	// The resource identifier.
	Space *string `json:"space,omitempty"`
	// The state of the address (_used_ or _free_).
	State *string `json:"state,omitempty"`
	// The tags for this address in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  Address was created by the IPAM component.  _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.  _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.  _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.  _DHCP_                 |  Address was created by the DHCP component.  _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.  _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.  _DHCP_, _DISABLED_     |  Address is disabled.  _DNS_                  |  Address is used by one or more DNS records.  _DISCOVERED_           |  Address is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.
	Usage                []string `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Address Address

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(address string) *Address {
	this := Address{}
	this.Address = address
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddress returns the Address field value
func (o *Address) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Address) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Address) SetAddress(v string) {
	o.Address = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Address) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Address) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Address) SetComment(v string) {
	o.Comment = &v
}

// GetCompartmentId returns the CompartmentId field value if set, zero value otherwise.
func (o *Address) GetCompartmentId() string {
	if o == nil || IsNil(o.CompartmentId) {
		var ret string
		return ret
	}
	return *o.CompartmentId
}

// GetCompartmentIdOk returns a tuple with the CompartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCompartmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompartmentId) {
		return nil, false
	}
	return o.CompartmentId, true
}

// HasCompartmentId returns a boolean if a field has been set.
func (o *Address) HasCompartmentId() bool {
	if o != nil && !IsNil(o.CompartmentId) {
		return true
	}

	return false
}

// SetCompartmentId gets a reference to the given string and assigns it to the CompartmentId field.
func (o *Address) SetCompartmentId(v string) {
	o.CompartmentId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Address) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Address) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Address) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpInfo returns the DhcpInfo field value if set, zero value otherwise.
func (o *Address) GetDhcpInfo() DHCPInfo {
	if o == nil || IsNil(o.DhcpInfo) {
		var ret DHCPInfo
		return ret
	}
	return *o.DhcpInfo
}

// GetDhcpInfoOk returns a tuple with the DhcpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDhcpInfoOk() (*DHCPInfo, bool) {
	if o == nil || IsNil(o.DhcpInfo) {
		return nil, false
	}
	return o.DhcpInfo, true
}

// HasDhcpInfo returns a boolean if a field has been set.
func (o *Address) HasDhcpInfo() bool {
	if o != nil && !IsNil(o.DhcpInfo) {
		return true
	}

	return false
}

// SetDhcpInfo gets a reference to the given DHCPInfo and assigns it to the DhcpInfo field.
func (o *Address) SetDhcpInfo(v DHCPInfo) {
	o.DhcpInfo = &v
}

// GetDisableDhcp returns the DisableDhcp field value if set, zero value otherwise.
func (o *Address) GetDisableDhcp() bool {
	if o == nil || IsNil(o.DisableDhcp) {
		var ret bool
		return ret
	}
	return *o.DisableDhcp
}

// GetDisableDhcpOk returns a tuple with the DisableDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDisableDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDhcp) {
		return nil, false
	}
	return o.DisableDhcp, true
}

// HasDisableDhcp returns a boolean if a field has been set.
func (o *Address) HasDisableDhcp() bool {
	if o != nil && !IsNil(o.DisableDhcp) {
		return true
	}

	return false
}

// SetDisableDhcp gets a reference to the given bool and assigns it to the DisableDhcp field.
func (o *Address) SetDisableDhcp(v bool) {
	o.DisableDhcp = &v
}

// GetDiscoveryAttrs returns the DiscoveryAttrs field value if set, zero value otherwise.
func (o *Address) GetDiscoveryAttrs() map[string]interface{} {
	if o == nil || IsNil(o.DiscoveryAttrs) {
		var ret map[string]interface{}
		return ret
	}
	return o.DiscoveryAttrs
}

// GetDiscoveryAttrsOk returns a tuple with the DiscoveryAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDiscoveryAttrsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DiscoveryAttrs) {
		return map[string]interface{}{}, false
	}
	return o.DiscoveryAttrs, true
}

// HasDiscoveryAttrs returns a boolean if a field has been set.
func (o *Address) HasDiscoveryAttrs() bool {
	if o != nil && !IsNil(o.DiscoveryAttrs) {
		return true
	}

	return false
}

// SetDiscoveryAttrs gets a reference to the given map[string]interface{} and assigns it to the DiscoveryAttrs field.
func (o *Address) SetDiscoveryAttrs(v map[string]interface{}) {
	o.DiscoveryAttrs = v
}

// GetDiscoveryMetadata returns the DiscoveryMetadata field value if set, zero value otherwise.
func (o *Address) GetDiscoveryMetadata() map[string]interface{} {
	if o == nil || IsNil(o.DiscoveryMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.DiscoveryMetadata
}

// GetDiscoveryMetadataOk returns a tuple with the DiscoveryMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDiscoveryMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DiscoveryMetadata) {
		return map[string]interface{}{}, false
	}
	return o.DiscoveryMetadata, true
}

// HasDiscoveryMetadata returns a boolean if a field has been set.
func (o *Address) HasDiscoveryMetadata() bool {
	if o != nil && !IsNil(o.DiscoveryMetadata) {
		return true
	}

	return false
}

// SetDiscoveryMetadata gets a reference to the given map[string]interface{} and assigns it to the DiscoveryMetadata field.
func (o *Address) SetDiscoveryMetadata(v map[string]interface{}) {
	o.DiscoveryMetadata = v
}

// GetExternalKeys returns the ExternalKeys field value if set, zero value otherwise.
func (o *Address) GetExternalKeys() map[string]interface{} {
	if o == nil || IsNil(o.ExternalKeys) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalKeys
}

// GetExternalKeysOk returns a tuple with the ExternalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetExternalKeysOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalKeys) {
		return map[string]interface{}{}, false
	}
	return o.ExternalKeys, true
}

// HasExternalKeys returns a boolean if a field has been set.
func (o *Address) HasExternalKeys() bool {
	if o != nil && !IsNil(o.ExternalKeys) {
		return true
	}

	return false
}

// SetExternalKeys gets a reference to the given map[string]interface{} and assigns it to the ExternalKeys field.
func (o *Address) SetExternalKeys(v map[string]interface{}) {
	o.ExternalKeys = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Address) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Address) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Address) SetHost(v string) {
	o.Host = &v
}

// GetHwaddr returns the Hwaddr field value if set, zero value otherwise.
func (o *Address) GetHwaddr() string {
	if o == nil || IsNil(o.Hwaddr) {
		var ret string
		return ret
	}
	return *o.Hwaddr
}

// GetHwaddrOk returns a tuple with the Hwaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetHwaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Hwaddr) {
		return nil, false
	}
	return o.Hwaddr, true
}

// HasHwaddr returns a boolean if a field has been set.
func (o *Address) HasHwaddr() bool {
	if o != nil && !IsNil(o.Hwaddr) {
		return true
	}

	return false
}

// SetHwaddr gets a reference to the given string and assigns it to the Hwaddr field.
func (o *Address) SetHwaddr(v string) {
	o.Hwaddr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Address) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Address) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Address) SetId(v string) {
	o.Id = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *Address) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *Address) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *Address) SetInterface(v string) {
	o.Interface = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *Address) GetNames() []Name {
	if o == nil || IsNil(o.Names) {
		var ret []Name
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNamesOk() ([]Name, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *Address) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []Name and assigns it to the Names field.
func (o *Address) SetNames(v []Name) {
	o.Names = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Address) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Address) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *Address) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Address) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Address) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Address) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Address) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Address) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *Address) SetRange(v string) {
	o.Range = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *Address) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *Address) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *Address) SetSpace(v string) {
	o.Space = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Address) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Address) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Address) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Address) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Address) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Address) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Address) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Address) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *Address) GetUsage() []string {
	if o == nil || IsNil(o.Usage) {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetUsageOk() ([]string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *Address) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *Address) SetUsage(v []string) {
	o.Usage = v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CompartmentId) {
		toSerialize["compartment_id"] = o.CompartmentId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpInfo) {
		toSerialize["dhcp_info"] = o.DhcpInfo
	}
	if !IsNil(o.DisableDhcp) {
		toSerialize["disable_dhcp"] = o.DisableDhcp
	}
	if !IsNil(o.DiscoveryAttrs) {
		toSerialize["discovery_attrs"] = o.DiscoveryAttrs
	}
	if !IsNil(o.DiscoveryMetadata) {
		toSerialize["discovery_metadata"] = o.DiscoveryMetadata
	}
	if !IsNil(o.ExternalKeys) {
		toSerialize["external_keys"] = o.ExternalKeys
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Hwaddr) {
		toSerialize["hwaddr"] = o.Hwaddr
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Address) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddress := _Address{}

	err = json.Unmarshal(data, &varAddress)

	if err != nil {
		return err
	}

	*o = Address(varAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "compartment_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "dhcp_info")
		delete(additionalProperties, "disable_dhcp")
		delete(additionalProperties, "discovery_attrs")
		delete(additionalProperties, "discovery_metadata")
		delete(additionalProperties, "external_keys")
		delete(additionalProperties, "host")
		delete(additionalProperties, "hwaddr")
		delete(additionalProperties, "id")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "names")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "range")
		delete(additionalProperties, "space")
		delete(additionalProperties, "state")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
