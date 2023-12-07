resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"
}

resource "bloxone_dns_srv_record" "example" {
  rdata = {
    port     = 80
    priority = 10
    target   = "example.com"
    weight   = 10
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "srv"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
