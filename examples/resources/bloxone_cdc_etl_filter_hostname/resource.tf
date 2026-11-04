resource "bloxone_cdc_etl_filter_hostname" "example" {
  name = "example_cdc_etl_filter_hostname"
  data = ["host1"]
}

resource "bloxone_cdc_etl_filter_hostname" "example_with_optional" {
  name          = "example_cdc_etl_filter_hostname_optional"
  data          = ["host1", "host2"]
  description   = "Example CDC ETL Hostname filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
