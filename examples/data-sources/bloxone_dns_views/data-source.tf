
# Get DNS Views filtered by an attribute
data "bloxone_dns_views" "example_by_attribute" {
  filters = {
    "name" = "example_ip_space"
  }
}

# Get DNS Views filtered by tag
data "bloxone_dns_views" "example_by_tag" {
  tag_filters = {
    site = "Test Site"
  }
}

# Get all DNS Views
data "bloxone_dns_views" "example_all" {}
