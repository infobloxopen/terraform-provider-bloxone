# Get CDC Syslog Destinations filtered by an attribute
data "bloxone_cdc_destination_syslogs" "example_by_attribute" {
  filters = {
    name = "example_cdc_destination_syslog"
  }
}

# Get CDC Syslog Destinations filtered by tag
data "bloxone_cdc_destination_syslogs" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC Syslog Destinations
data "bloxone_cdc_destination_syslogs" "example_all" {}
