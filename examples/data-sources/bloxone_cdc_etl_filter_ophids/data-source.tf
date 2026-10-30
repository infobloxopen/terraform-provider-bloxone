# Get CDC ETL OPHID Filters filtered by an attribute
data "bloxone_cdc_etl_filter_ophids" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_ophid"
  }
}

# Get CDC ETL OPHID Filters filtered by tag
data "bloxone_cdc_etl_filter_ophids" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL OPHID Filters
data "bloxone_cdc_etl_filter_ophids" "example_all" {}
