# Splunk Cloud requires the Splunk forwarder certificate and a CA certificate.
# Replace forwarder_cert.pem and ca_cert.pem with your own certificates.
resource "bloxone_cdc_destination_splunk_cloud" "example" {
  name           = "example_cdc_destination_splunk_cloud"
  address        = "splunkcloud.example.com:9197"
  indexer_name   = "main"
  certificate    = file("${path.module}/forwarder_cert.pem")
  ca_certificate = file("${path.module}/ca_cert.pem")
}

resource "bloxone_cdc_destination_splunk_cloud" "example_with_optional" {
  name                       = "example_cdc_destination_splunk_cloud_optional"
  address                    = "10.10.10.1:9497,10.10.10.2:9397"
  indexer_name               = "main"
  certificate                = file("${path.module}/forwarder_cert.pem")
  ca_certificate             = file("${path.module}/ca_cert.pem")
  certificate_key_passphrase = "example_passphrase"
  config_option              = "MANUAL"
  output_data_format         = "CIM"
  description                = "Example CDC Splunk Cloud destination"
  enabled                    = true

  tags = {
    site = "Site A"
  }
}
