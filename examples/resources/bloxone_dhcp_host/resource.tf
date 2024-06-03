data "bloxone_infra_hosts" "host" {
  filters = {
    display_name = "my-host"
  }
}

resource "bloxone_dhcp_server" "dhcp_server" {
  name = "my-dhcp-server"
}

resource "bloxone_dhcp_host" "example_dhcp_host" {
  id     = data.bloxone_infra_hosts.host.results.0.legacy_id
  server = bloxone_dhcp_server.dhcp_server.id

  # Other Optional fields
  tags = {
    site = "Site A"
  }
}