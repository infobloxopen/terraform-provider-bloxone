# Get CDC Application Destinations filtered by an attribute
data "bloxone_cdc_destination_applications" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_application"
  }
}

# Get CDC Application Destinations filtered by tag
data "bloxone_cdc_destination_applications" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Application Destinations
data "bloxone_cdc_destination_applications" "example_all" {}
