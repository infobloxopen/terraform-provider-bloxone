# Get HA Groups filtered by an attribute
data "bloxone_ipam_ha_groups" "example_by_attribute" {
  filters = {
    "name" = "example_ha"
  }
}

# Get HA Groups filtered by tag
data "bloxone_ipam_ha_groups" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all HA Groups
data "bloxone_ipam_ha_groups" "example_all" {}
