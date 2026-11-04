resource "bloxone_cdc_etl_filter_ophid" "example" {
  name = "example_cdc_etl_filter_ophid"
  data = ["abcd1234"]
}

resource "bloxone_cdc_etl_filter_ophid" "example_with_optional" {
  name          = "example_cdc_etl_filter_ophid_optional"
  data          = ["abcd1234", "id-1"]
  description   = "Example CDC ETL OPHID filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
