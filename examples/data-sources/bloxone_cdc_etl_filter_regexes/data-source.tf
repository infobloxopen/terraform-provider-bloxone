# Get CDC ETL Regex Filters filtered by an attribute
data "bloxone_cdc_etl_filter_regexes" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_regex"
  }
}

# Get CDC ETL Regex Filters filtered by tag
data "bloxone_cdc_etl_filter_regexes" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all CDC ETL Regex Filters
data "bloxone_cdc_etl_filter_regexes" "example_all" {}
