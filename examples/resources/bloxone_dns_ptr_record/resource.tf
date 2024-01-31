resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "192.in-addr.arpa."
  primary_type = "cloud"
}

resource "bloxone_dns_ptr_record" "example" {
  rdata = {
    dname = "example.com"
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "1.0.168"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
