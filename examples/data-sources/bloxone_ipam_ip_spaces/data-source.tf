# Get IP Spaces filtered by an attribute
data "bloxone_ipam_ip_spaces" "example_by_attribute" {
  filters = {
    "name" = "example_ip_space"
  }
}

# Get IP Spaces filtered by tag
data "bloxone_ipam_ip_spaces" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all IP Spaces
data "bloxone_ipam_ip_spaces" "example_all" {}
