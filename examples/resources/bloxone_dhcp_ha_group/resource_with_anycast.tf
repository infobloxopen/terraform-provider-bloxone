# Get first host for the HA group.
# This host must have 'Anycast' enabled.
data "bloxone_dhcp_hosts" "example_host_1" {
  filters = {
    name = "Host-1"
  }
}

# Get second host for the HA group.
# This host must have 'Anycast' enabled.
data "bloxone_dhcp_hosts" "example_host_2" {
  filters = {
    name = "Host-2"
  }
}

# Get the Anycast configuration for the service
data "bloxone_anycast_configs" "example_service" {
  service = "DHCP"
}

resource "bloxone_dhcp_ha_group" "example_anycast" {
  name              = "example_ha_group_anycast"
  mode              = "anycast"
  anycast_config_id = format("accm/ac_configs/%s", data.bloxone_anycast_configs.example_service.results.0.id)

  hosts = [
    {
      host = data.bloxone_dhcp_hosts.example_host_1.results.0.id,
      role = "active"
    },
    {
      host = data.bloxone_dhcp_hosts.example_host_2.results.0.id,
      role = "passive"
    }
  ]
}
