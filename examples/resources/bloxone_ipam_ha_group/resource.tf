data "bloxone_ipam_dhcp_hosts" "example_host_1" {
  filter = {
    name = "TF_TEST_HOST_01"
  }
}

data "bloxone_ipam_dhcp_hosts" "example_host_2" {
  filter = {
    name = "TF_TEST_HOST_02"
  }
}

resource "bloxone_ipam_ha_group" "example" {
  name = "example_ha_group"
}

resource "bloxone_ipam_ha_group" "example_tags" {
  name = "example_ha_group_tags"
  mode = "active-active"
  hosts = [
    {
      host = data.bloxone_ipam_dhcp_hosts.example_host_1.results.0.id,
      role = "active"
    },
    {
      host = data.bloxone_ipam_dhcp_hosts.example_host_2.results.0.id,
      role = "active"
    }
  ]

  # Optional fields
  comment = "Example HA Group with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
}
