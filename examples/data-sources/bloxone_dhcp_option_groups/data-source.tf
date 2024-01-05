# Get DHCP Option group filtered by an attribute
data "bloxone_dhcp_option_groups" "example_by_name" {
  filters = {
    name = "example-group"
  }
}

# Get DHCP Option group/s by tag
data "bloxone_dhcp_option_groups" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP Option groups
data "bloxone_dhcp_option_groups" "example_all" {}
