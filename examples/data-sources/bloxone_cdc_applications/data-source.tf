# Get CDC Applications filtered by an attribute
data "bloxone_cdc_applications" "example_by_attribute" {
  filters = {
    name = "example_cdc_application"
  }
}

# Get CDC Applications filtered by tag
data "bloxone_cdc_applications" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Applications
data "bloxone_cdc_applications" "example_all" {}
