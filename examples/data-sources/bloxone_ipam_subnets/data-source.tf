# Get subnets filtered by an attribute
data "bloxone_ipam_subnets" "example_by_attribute" {
  filters = {
    "address" = "10.0.0.0"
    "cidr"    = "24"
  }
}

# Get subnets filtered by tag
data "bloxone_ipam_subnets" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all subnets
data "bloxone_ipam_subnets" "example_all" {}
