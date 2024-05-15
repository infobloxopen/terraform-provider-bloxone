# Get Category Filters filtered by an attribute
data "bloxone_td_category_filters" "example_by_attribute" {
  filters = {
    "name" = "example_category_filter"
  }
}

# Get Category Filters filtered by tag
data "bloxone_td_category_filters" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all Category Filters
data "bloxone_td_category_filters" "example_all" {}
