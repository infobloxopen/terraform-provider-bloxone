# Get Application Filters filtered by an attribute
data "bloxone_td_application_filters" "example_by_attribute" {
  filters = {
    "name" = "example_application_filter"
  }
}

# Get Application Filters filtered by tag
data "bloxone_td_application_filters" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all Application Filters
data "bloxone_td_application_filters" "example_all" {}
