# Get DHCP Option code filtered by an attribute
data "bloxone_dhcp_option_codes" "example_by_name" {
  filters = {
    name = "example-code"
  }
}

# Get DHCP Option code/s by tag
data "bloxone_dhcp_option_codes" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP Option codes
data "bloxone_dhcp_option_codes" "example_all" {}
