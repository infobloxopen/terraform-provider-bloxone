resource "bloxone_cdc_destination_syslog" "example" {
  name    = "example_cdc_destination_syslog"
  address = "syslog.dest.example.com"
}

resource "bloxone_cdc_destination_syslog" "example_with_optional" {
  name               = "example_cdc_destination_syslog_optional"
  address            = "syslog.dest1.example.com"
  port               = 514
  transport_protocol = "UDP"
  output_data_format = "CEF"
  description        = "Example CDC syslog destination"
  enabled            = true

  tags = {
    site = "Site A"
  }
}

# Syslog over TLS (TCP with a CA certificate)
resource "bloxone_cdc_destination_syslog" "example_with_tls" {
  name               = "example_cdc_destination_syslog_tls"
  address            = "syslog.dest2.example.com"
  port               = 6514
  insecure_mode      = false
  transport_protocol = "TCP"
  ca_certificate     = file("${path.module}/ca_cert.pem")
}
