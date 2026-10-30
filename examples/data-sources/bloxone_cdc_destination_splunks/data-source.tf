# Get CDC Splunk Destinations filtered by an attribute
data "bloxone_cdc_destination_splunks" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_splunk"
  }
}

# Get all CDC Splunk Destinations
data "bloxone_cdc_destination_splunks" "example_all" {}
