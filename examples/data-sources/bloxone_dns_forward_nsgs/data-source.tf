# Get Forward NSGs filtered by an attribute
data "bloxone_dns_forward_nsgs" "example_by_attribute" {
  filters = {
    "name" = "example_forward_nsg"
  }
}

# Get Forward NSGs filtered by tag
data "bloxone_dns_forward_nsgs" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Forward NSGs
data "bloxone_dns_forward_nsgs" "example_all" {}
