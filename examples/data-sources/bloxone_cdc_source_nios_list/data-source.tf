# Get CDC NIOS Sources filtered by name
data "bloxone_cdc_source_nios_list" "example_by_name" {
  filters = {
    name = "example_cdc_source_nios_basic"
  }
}

# Get CDC NIOS Sources filtered by address
data "bloxone_cdc_source_nios_list" "example_by_address" {
  filters = {
    address = "10.1.0.0"
  }
}

# Get CDC NIOS Sources filtered by tag
# Note: tag_filters is currently ignored by the CDC backend for NIOS sources, so all sources are returned.
data "bloxone_cdc_source_nios_list" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC NIOS Sources
data "bloxone_cdc_source_nios_list" "example_all" {}
