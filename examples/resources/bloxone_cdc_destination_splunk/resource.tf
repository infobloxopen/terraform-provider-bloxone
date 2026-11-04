resource "bloxone_cdc_destination_splunk" "example" {
  name         = "example_cdc_destination_splunk"
  address      = "splunk.example.com:9997"
  indexer_name = "main"
}

resource "bloxone_cdc_destination_splunk" "example_with_optional" {
  name               = "example_cdc_destination_splunk_optional"
  address            = "10.10.10.1:9997,10.10.10.2:9997"
  indexer_name       = "main"
  output_data_format = "CIM"
  description        = "Example CDC Splunk destination"
  enabled            = true

  tags = {
    site = "Site A"
  }
}

# Splunk with a forwarder certificate, verifying the indexer certificate against a CA certificate
resource "bloxone_cdc_destination_splunk" "example_with_certificates" {
  name                       = "example_cdc_destination_splunk_certificates"
  address                    = "splunk2.example.com:9997"
  indexer_name               = "main"
  insecure_mode              = false
  ca_certificate             = file("${path.module}/ca_cert.pem")
  certificate                = file("${path.module}/forwarder_cert.pem")
  certificate_key_passphrase = "example_passphrase"
}
