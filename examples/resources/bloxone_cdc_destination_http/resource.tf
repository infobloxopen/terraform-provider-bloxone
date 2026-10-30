# Replace credential_id with the ID of your own credential configured in the Portal
resource "bloxone_cdc_destination_http" "example" {
  name          = "example_cdc_destination_http"
  address       = "https://example1.com/ingest"
  credential_id = "00000000-0000-0000-0000-000000000000"
}

resource "bloxone_cdc_destination_http" "example_with_optional" {
  name               = "example_cdc_destination_http_optional"
  address            = "https://example2.com/ingest"
  credential_id      = "00000000-0000-0000-0000-000000000000"
  output_data_format = "CIM"
  transport_protocol = "HTTPS"
  description        = "Example CDC HTTP destination"
  enabled            = true

  tags = {
    site = "Site A"
  }
}
