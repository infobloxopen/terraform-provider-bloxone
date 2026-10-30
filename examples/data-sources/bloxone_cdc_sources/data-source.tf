# Get the built-in Infoblox Cloud Source
data "bloxone_cdc_sources" "cloud" {
  filters = {
    type = "SOURCE_BLOXONE"
  }
}

# Get the built-in Schedule Source
data "bloxone_cdc_sources" "schedule" {
  filters = {
    type = "SOURCE_SCHEDULE"
  }
}

# Get CDC Sources filtered by name
data "bloxone_cdc_sources" "example_by_name" {
  filters = {
    name = "example_cdc_source_nios"
  }
}

# Get all CDC Sources
data "bloxone_cdc_sources" "example_all" {}
