# Get CDC ETL IP Network Filters filtered by an attribute
data "bloxone_cdc_etl_filter_ip_networks" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_ip_network"
  }
}

# Get CDC ETL IP Network Filters filtered by tag
data "bloxone_cdc_etl_filter_ip_networks" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL IP Network Filters
data "bloxone_cdc_etl_filter_ip_networks" "example_all" {}
