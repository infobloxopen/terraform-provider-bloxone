# Get CDC ETL Record Type Filters filtered by an attribute
data "bloxone_cdc_etl_filter_record_types" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_record_type"
  }
}

# Get CDC ETL Record Type Filters filtered by tag
data "bloxone_cdc_etl_filter_record_types" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL Record Type Filters
data "bloxone_cdc_etl_filter_record_types" "example_all" {}
