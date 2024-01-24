data "bloxone_infra_hosts" "example" {
  filters = {
    display_name = "example-host"
  }
}

resource "bloxone_infra_service" "example" {
  name         = "example_service"
  pool_id      = data.bloxone_infra_hosts.example.results.0.pool_id
  service_type = "dhcp"

  # Other Optional fields
  description     = "DHCP service"
  desired_version = "3.5.0"
  desired_state   = "start"
  tags = {
    site = "Site A"
  }
}
