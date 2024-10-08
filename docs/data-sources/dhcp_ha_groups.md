---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dhcp_ha_groups Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves information about existing HA Groups.
  The HA Group object represents on-prem hosts that can serve the same leases for HA.
---

# bloxone_dhcp_ha_groups (Data Source)

Retrieves information about existing HA Groups.

The HA Group object represents on-prem hosts that can serve the same leases for HA.

## Example Usage

```terraform
# Get HA Groups filtered by an attribute
data "bloxone_dhcp_ha_groups" "example_by_attribute" {
  filters = {
    "name" = "example_ha"
  }
}

# Get HA Groups filtered by tag with collect_stats enabled
data "bloxone_dhcp_ha_groups" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
  collect_stats = true
}

# Get all HA Groups
data "bloxone_dhcp_ha_groups" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `collect_stats` (Boolean) collect_stats gets the HA group stats(state, status, heartbeat) if set to true. Defaults to false
- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Required:

- `hosts` (Attributes List) The list of hosts. (see [below for nested schema](#nestedatt--results--hosts))
- `mode` (String) The mode of the HA group. Valid values are:
  * _active-active_: Both on-prem hosts remain active.
  * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over.
  * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.
- `name` (String) The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.

Optional:

- `anycast_config_id` (String) The resource identifier.
- `collect_stats` (Boolean) collect_stats gets the HA group stats(state, status, heartbeat) if set to true. Defaults to false
- `comment` (String) The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.
- `tags` (Map of String) The tags for the HA group.

Read-Only:

- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `ip_space` (String) The resource identifier.
- `status` (String) Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
- `status_v6` (String) Status of the DHCPv6 HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
- `tags_all` (Map of String) The tags for the HA group including default tags.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedatt--results--hosts"></a>
### Nested Schema for `results.hosts`

Required:

- `host` (String) The resource identifier.
- `role` (String) The role of this host in the HA relationship: _active_ or _passive_.

Optional:

- `address` (String) The address on which this host listens.

Read-Only:

- `heartbeats` (Attributes List) Last successful heartbeat received from its peer/s. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request. (see [below for nested schema](#nestedatt--results--hosts--heartbeats))
- `port` (Number) The HA port.
- `port_v6` (Number) The HA port used for IPv6 communication.
- `state` (String) The state of DHCP on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
- `state_v6` (String) The state of DHCPv6 on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.

<a id="nestedatt--results--hosts--heartbeats"></a>
### Nested Schema for `results.hosts.heartbeats`

Read-Only:

- `peer` (String) The name of the peer.
- `successful_heartbeat` (String) The timestamp as a string of the last successful heartbeat received from the peer above.
- `successful_heartbeat_v6` (String) The timestamp as a string of the last successful DHCPv6 heartbeat received from the peer above.
