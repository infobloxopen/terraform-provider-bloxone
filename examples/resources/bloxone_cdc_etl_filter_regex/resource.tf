resource "bloxone_cdc_etl_filter_regex" "example" {
  name = "example_cdc_etl_filter_regex"
  data = ["^example.*$"]
}

resource "bloxone_cdc_etl_filter_regex" "example_with_optional" {
  name          = "example_cdc_etl_filter_regex_optional"
  data          = ["^example.*$", "^test.*$"]
  description   = "Example CDC ETL Regex filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
