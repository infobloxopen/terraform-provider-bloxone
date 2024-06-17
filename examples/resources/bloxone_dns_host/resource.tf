data "bloxone_infra_hosts" "host" {
  filters = {
    display_name = "my-host"
  }
}

resource "bloxone_dns_server" "dhcp_server" {
  name = "my-dns-server"
}

resource "bloxone_dns_host" "example_dns_host" {
  id     = data.bloxone_infra_hosts.host.results.0.legacy_id
  server = bloxone_dns_server.dhcp_server.id

  # Other Optional fields
  tags = {
    site = "Site A"
  }
}