# Get IPAM Hosts filtered by an attribute
data "bloxone_ipam_hosts" "example_by_attribute" {
  filters = {
    "name" = "example_ipam_host"
  }
}

# Get IPAM Hosts filtered by tag
data "bloxone_ipam_hosts" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all IPAM Hosts
data "bloxone_ipam_hosts" "example_all" {}
