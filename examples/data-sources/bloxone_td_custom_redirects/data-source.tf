# Get Custom Redirects filtered by an attribute
data "bloxone_td_custom_redirects" "example_by_attribute" {
  filters = {
    "name" = "example_custom_redirect"
  }
}

# Get Custom Redirects filtered by tag
data "bloxone_td_custom_redirects" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all Custom Redirects
data "bloxone_td_custom_redirects" "example_all" {}
