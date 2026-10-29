# Get CDC Splunk Destinations filtered by an attribute
data "bloxone_cdc_destination_splunks" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_splunk"
  }
}

# Get CDC Splunk Destinations filtered by tag
# Note: tag_filters is currently ignored by the CDC backend for Splunk destinations, so all destinations are returned.
data "bloxone_cdc_destination_splunks" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Splunk Destinations
data "bloxone_cdc_destination_splunks" "example_all" {}
