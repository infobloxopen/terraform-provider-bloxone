# Get CDC ETL Threat Class Property Filters filtered by an attribute
data "bloxone_cdc_etl_filter_threat_class_properties" "example_by_attribute" {
  filters = {
    name = "example_cdc_etl_filter_threat_class_property"
  }
}

# Get CDC ETL Threat Class Property Filters filtered by tag
data "bloxone_cdc_etl_filter_threat_class_properties" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}
