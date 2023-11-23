# Get all Ranges
data "b1ddi_ranges" "example_all_ranges" {}

## Get specific Range by start and end values
data "b1ddi_ranges" "example_range_by_start_end" {
  filters = {
    "start" = "192.168.1.15",
    "end" = "192.168.1.30"
  }
}

## Get specific Range by name
data "b1ddi_ranges" "example_range_by_name" {
  filters = {
    "name" = "example_range"
  }
}

# Get Range by tag
data "b1ddi_ranges" "example_range_by_tag"{
  tag_filters = {
    location = "site1"
  }
}