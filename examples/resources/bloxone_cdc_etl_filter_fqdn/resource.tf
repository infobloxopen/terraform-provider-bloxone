resource "bloxone_cdc_etl_filter_fqdn" "example" {
  name = "example_cdc_etl_filter_fqdn"
  data = ["example.com"]
}

resource "bloxone_cdc_etl_filter_fqdn" "example_with_optional" {
  name          = "example_cdc_etl_filter_fqdn_optional"
  data          = ["example.com", "example.org"]
  description   = "Example CDC ETL FQDN filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
