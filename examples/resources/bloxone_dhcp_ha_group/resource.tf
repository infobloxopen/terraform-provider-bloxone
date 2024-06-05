# Get first host for the HA group.
data "bloxone_dhcp_hosts" "example_host_1" {
  filters = {
    name = "Host-1"
  }
}

# Get second host for the HA group.
data "bloxone_dhcp_hosts" "example_host_2" {
  filters = {
    name = "Host-2"
  }
}

resource "bloxone_dhcp_ha_group" "example_tags" {
  name = "example_ha_group_tags"
  mode = "active-active"
  hosts = [
    {
      host = data.bloxone_dhcp_hosts.example_host_1.results.0.id,
      role = "active"
    },
    {
      host = data.bloxone_dhcp_hosts.example_host_2.results.0.id,
      role = "active"
    }
  ]

  # Optional fields
  comment = "Example HA Group with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
}
