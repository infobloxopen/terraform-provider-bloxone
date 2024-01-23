resource "bloxone_dns_view" "example" {
  name = "example-view"
}

resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "domain.com."
  primary_type = "cloud"
  view         = bloxone_dns_view.example.id
}

resource "bloxone_dns_delegation" "example" {
  fqdn = "del.domain.com."
  delegation_servers = [{
    address = "12.0.0.0"
    fqdn    = "ns1.com."
  }]

  # Other optional fields
  view    = bloxone_dns_view.example.id
  comment = "Delegation zone created through Terraform"
  tags = {
    site = "Site A"
  }
  disabled = true

  depends_on = [bloxone_dns_view.example, bloxone_dns_auth_zone.example]
}
