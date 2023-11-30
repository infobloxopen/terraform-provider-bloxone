# Get DHCP fixed address filtered by an attribute
data "bloxone_dhcp_fixed_addresses" "example_by_attribute" {
  filters = {
    name = "example_fixed_address"
  }
}

# Get DHCP fixed address by tag
data "bloxone_dhcp_fixed_addresses" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all fixed address
data "bloxone_dhcp_fixed_addresses" "example_all" {}
