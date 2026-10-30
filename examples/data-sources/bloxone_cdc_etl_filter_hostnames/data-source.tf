# Get CDC ETL Hostname Filters filtered by an attribute
data "bloxone_cdc_etl_filter_hostnames" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_hostname"
  }
}

# Get CDC ETL Hostname Filters filtered by tag
data "bloxone_cdc_etl_filter_hostnames" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL Hostname Filters
data "bloxone_cdc_etl_filter_hostnames" "example_all" {}
