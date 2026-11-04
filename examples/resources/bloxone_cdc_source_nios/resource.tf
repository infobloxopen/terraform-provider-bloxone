resource "bloxone_cdc_source_nios" "example" {
  name    = "example_cdc_source_nios_basic"
  address = "10.1.0.0"
}

# NIOS source with all fields set.
# ca_cert.pem and certkey.pem in this folder are self-signed samples; replace
# them with your own Grid Master CA certificate and RPZ logs certificate.
resource "bloxone_cdc_source_nios" "example_with_all_fields" {
  name        = "example_cdc_source_nios_all_fields"
  description = "Example CDC NIOS source with all fields"
  enabled     = true

  # Credentials for Grid Master configuration
  address        = "10.1.0.1"
  nios_username  = "admin"
  nios_password  = "example_password"
  insecure_mode  = false
  ca_certificate = file("${path.module}/ca_cert.pem")

  # SCP credentials for DNS query logs transfer
  scp_username = "scp_user"
  scp_password = "example_scp_password"

  # Certificate for RPZ logs
  certificate           = file("${path.module}/certkey.pem")
  rpz_log_insecure_mode = false

  tags = {
    site = "Site A"
  }
}
