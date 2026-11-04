resource "bloxone_cdc_etl_filter_ip_network" "example" {
  name = "example_cdc_etl_filter_ip_network"
  data = ["10.0.0.0/24"]
}

resource "bloxone_cdc_etl_filter_ip_network" "example_with_optional" {
  name          = "example_cdc_etl_filter_ip_network_optional"
  data          = ["10.0.0.0/24", "192.168.1.0/24"]
  description   = "Example CDC ETL IP Network filter"
  enabled       = true
  filter_source = "SOURCE_DEFAULT"

  tags = {
    site = "Site A"
  }
}
