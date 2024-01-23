resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"
}

resource "bloxone_dns_txt_record" "example" {
  rdata = {
    text = "example.com"
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "txt"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
