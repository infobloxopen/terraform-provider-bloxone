# Get CDC HTTP Destinations filtered by an attribute
data "bloxone_cdc_destination_https" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_http"
  }
}

# Get CDC HTTP Destinations filtered by tag
data "bloxone_cdc_destination_https" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC HTTP Destinations
data "bloxone_cdc_destination_https" "example_all" {}
