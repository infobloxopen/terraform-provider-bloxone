
terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = "0.1.0"
    }
  }
}

provider "bloxone" {
  csp_url = "https://csp.infoblox.com"
  api_key = "f77723789095d2a9ca1d8665d78b4fbdaad4783eaaca94ea4474d9177e06babc"
}

data "bloxone_dhcp_hosts" "example_host_1" {
  filters = {
    name = "Goutham-HA1"
  }
}

data "bloxone_dhcp_hosts" "example_host_2" {
  filters = {
    name = "Goutham-HA2"
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
