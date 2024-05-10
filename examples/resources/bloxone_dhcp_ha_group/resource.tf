data "bloxone_dhcp_hosts" "example_host_1" {
  filters = {
    name = "Your Host name"
  }
}

data "bloxone_dhcp_hosts" "example_host_2" {
  filters = {
    name = "Your host name"
  }
}

# Get the anycast configuration for the service
data "bloxone_anycast_configs" "example_service" {
  service = "DNS"
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
