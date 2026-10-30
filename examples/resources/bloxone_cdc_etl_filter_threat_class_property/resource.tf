resource "bloxone_cdc_etl_filter_threat_class_property" "example" {
  name = "example_cdc_etl_filter_threat_class_property"
  data = ["Malicious"]
}

resource "bloxone_cdc_etl_filter_threat_class_property" "example_with_optional" {
  name          = "example_cdc_etl_filter_threat_class_property_optional"
  data          = ["DDoS", "Whitelist", "Malicious"]
  description   = "Example CDC ETL Threat Class Property filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
