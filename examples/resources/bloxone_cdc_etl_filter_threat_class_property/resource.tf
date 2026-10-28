# Each entry in 'data' must be a valid threat class (e.g. "Malicious", "DDoS") or
# threat property name from Infoblox Threat Intelligence (TIDE).
# The full lists are available from the TIDE API:
#   GET /tide/api/data/threat_classes
#   GET /tide/api/data/properties
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
