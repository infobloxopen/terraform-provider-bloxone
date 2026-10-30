# Get CDC NIOS Reporting Destinations filtered by an attribute
data "bloxone_cdc_destination_reportings" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_reporting"
  }
}

# Get all CDC NIOS Reporting Destinations
data "bloxone_cdc_destination_reportings" "example_all" {}
