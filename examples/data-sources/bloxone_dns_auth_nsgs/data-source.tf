# Get Auth NSGs filtered by an attribute
data "bloxone_dns_auth_nsgs" "example_by_attribute" {
  filters = {
    "name" = "example_auth_nsg"
  }
}

# Get Auth NSGs filtered by tag
data "bloxone_dns_auth_nsgs" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Auth NSGs
data "bloxone_dns_auth_nsgs" "example_all" {}
