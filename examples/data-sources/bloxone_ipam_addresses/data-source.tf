# Get subnets filtered by an attribute
data "bloxone_ipam_addresses" "example_by_attribute" {
  filters = {
    "address" = "10.0.0.0"
  }
}

# Get addresses filtered by tag
data "bloxone_ipam_addresses" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all addresses
data "bloxone_ipam_addresses" "example_all" {}
