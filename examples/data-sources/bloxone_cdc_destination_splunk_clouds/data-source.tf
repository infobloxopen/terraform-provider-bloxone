# Get CDC Splunk Cloud Destinations filtered by an attribute
data "bloxone_cdc_destination_splunk_clouds" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_splunk_cloud"
  }
}

# Get CDC Splunk Cloud Destinations filtered by tag
data "bloxone_cdc_destination_splunk_clouds" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Splunk Cloud Destinations
data "bloxone_cdc_destination_splunk_clouds" "example_all" {}
