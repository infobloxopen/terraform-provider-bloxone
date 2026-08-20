# Get DNS Views filtered by an attribute
data "bloxone_dns_views" "example_by_attribute" {
  filters = {
    "name" = "example_view"
  }
}

# Get DNS Views filtered by tag
data "bloxone_dns_views" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get DNS Views filtered by a NIOS tag
data "bloxone_dns_views" "example_by_nios_tag" {
  tag_filters = {
    "nios/imported" = "true"
  }
}

# Get all DNS Views
data "bloxone_dns_views" "example_all" {}
