# Get internal domain lists filtered by an attribute
data "bloxone_td_internal_domain_lists" "example_by_attribute" {
  filters = {
    name = "example_list"
  }
}

# Get internal domain lists filtered by tag
data "bloxone_td_internal_domain_lists" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all auth zones
data "bloxone_td_internal_domain_lists" "example_all" {}
