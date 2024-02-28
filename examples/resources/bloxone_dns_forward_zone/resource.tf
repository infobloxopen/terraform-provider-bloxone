resource "bloxone_dns_view" "example_view" {
  name = "example_dns_view"
}

resource "bloxone_dns_forward_zone" "example" {
  fqdn = "domain.com."

  # Other optional fields
  comment = "Example of a Forward Zone"
  tags = {
    site = "Site A"
  }
  view = bloxone_dns_view.example_view.id
}
