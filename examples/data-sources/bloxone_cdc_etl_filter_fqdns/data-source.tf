# Get CDC ETL FQDN Filters filtered by an attribute
data "bloxone_cdc_etl_filter_fqdns" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_fqdn"
  }
}

# Get CDC ETL FQDN Filters filtered by tag
data "bloxone_cdc_etl_filter_fqdns" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL FQDN Filters
data "bloxone_cdc_etl_filter_fqdns" "example_all" {}
