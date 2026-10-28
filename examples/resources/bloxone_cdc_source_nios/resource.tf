resource "bloxone_cdc_source_nios" "example" {
  name    = "example_cdc_source_nios_basic"
  address = "10.1.0.0"
}

# NIOS source with all fields set.
# ca_cert.pem and certkey.pem in this folder are self-signed samples; replace
# them with your own CA certificate and client certificate + private key.
resource "bloxone_cdc_source_nios" "example_with_all_fields" {
  name          = "example_cdc_source_nios_all_fields"
  address       = "10.1.0.1"
  description   = "Example CDC NIOS source with all fields"
  enabled       = true
  nios_username = "admin"
  nios_password = "example_password"

  # Verify the Grid Master's TLS certificate instead of skipping verification
  insecure_mode         = false
  rpz_log_insecure_mode = false
  ca_certificate        = file("${path.module}/ca_cert.pem")
  certificate           = file("${path.module}/certkey.pem")

  # SCP credentials used to transfer files from NIOS
  scp_username = "scp_user"
  scp_password = "example_scp_password"

  tags = {
    site = "Site A"
  }
}
