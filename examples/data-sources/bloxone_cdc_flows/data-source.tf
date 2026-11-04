# Get CDC Flows filtered by an attribute
data "bloxone_cdc_flows" "example_by_attribute" {
  filters = {
    name = "example_cdc_flow"
  }
}

# Get CDC Flows filtered by tag
data "bloxone_cdc_flows" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Flows
data "bloxone_cdc_flows" "example_all" {}
