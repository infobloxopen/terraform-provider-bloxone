# Get Network Lists filtered by an attribute
data "bloxone_td_network_lists" "example_by_attribute" {
  filters = {
    "name" = "example_network_list"
  }
}

# Get all Network Lists
data "bloxone_td_network_lists" "example_all" {}
