# Get DHCP Option spaces filtered by an attribute
data "bloxone_dhcp_option_spaces" "example_by_name" {
  filters = {
    name = "example-space"
  }
}

# Get DHCP Option space/s by tag
data "bloxone_dhcp_option_spaces" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP Option spaces
data "bloxone_dhcp_option_spaces" "example_all" {}
