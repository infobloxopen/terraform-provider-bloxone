# Get dhcp servers filtered by an attribute
data "bloxone_dhcp servers" "example_by_attribute" {
  filters = {
    "name" = "example_dhcp_server"
  }
}

# Get dhcp servers filtered by tag
data "bloxone_dhcp servers" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all dhcp servers
data "bloxone_dhcp servers" "example_all" {}
