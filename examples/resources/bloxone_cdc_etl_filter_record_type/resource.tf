resource "bloxone_cdc_etl_filter_record_type" "example" {
  name = "example_cdc_etl_filter_record_type"
  data = ["A"]
}

resource "bloxone_cdc_etl_filter_record_type" "example_with_optional" {
  name          = "example_cdc_etl_filter_record_type_optional"
  data          = ["A", "AAAA"]
  description   = "Example CDC ETL Record Type filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
