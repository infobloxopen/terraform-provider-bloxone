# Get Named Lists filtered by an attribute
data "bloxone_td_named_lists" "example_by_attribute" {
  filters = {
    "name" = "example_named_list"
  }
}

# Get Named Lists filtered by tag
data "bloxone_td_named_lists" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all Named Lists
data "bloxone_td_named_lists" "example_all" {}
