# Get Address block/s filtered by an attribute
data "bloxone_ipam_address_blocks" "example_by_attribute" {
  filters = {
    "name" = "example_subnet"
  }
}

# Get Address blocks filtered by tag
data "bloxone_ipam_address_blocks" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all Address blocks
data "bloxone_ipam_address_blocks" "example_all" {}