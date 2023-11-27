# Get DHCP Host filtered by an attribute
data "bloxone_dhcp_hosts" "dhcp_host_by_name" {
  filters = {
    name = "example-host"
  }
}

# Get DHCP Host by tag
data "bloxone_dhcp_hosts" "dhcp_host_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP hosts
data "bloxone_dhcp_hosts" "all_hosts" {}
