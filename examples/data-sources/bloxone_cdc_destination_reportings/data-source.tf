# Get CDC NIOS Reporting Destinations filtered by an attribute
data "bloxone_cdc_destination_reportings" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_reporting"
  }
}

# Get CDC NIOS Reporting Destinations filtered by tag
# Note: tag_filters is currently ignored by the CDC backend for reporting destinations, so all destinations are returned.
data "bloxone_cdc_destination_reportings" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC NIOS Reporting Destinations
data "bloxone_cdc_destination_reportings" "example_all" {}
