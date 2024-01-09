resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"
}

resource "bloxone_dns_dname_record" "example" {
  rdata = {
    target = "example."
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "dname"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
