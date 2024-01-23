resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"
}

resource "bloxone_dns_aaaa_record" "example" {
  rdata = {
    address = "2001:db8::1"
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "aaaa"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
