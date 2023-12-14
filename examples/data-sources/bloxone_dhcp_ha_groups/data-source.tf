# Get HA Groups filtered by an attribute
data "bloxone_dhcp_ha_groups" "example_by_attribute" {
  filters = {
    "name" = "example_ha"
  }
}

# Get HA Groups filtered by tag with collect_stats enabled
data "bloxone_dhcp_ha_groups" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
  collect_stats = true
}

# Get all HA Groups
data "bloxone_dhcp_ha_groups" "example_all" {}
