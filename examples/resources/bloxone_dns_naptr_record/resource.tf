resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"
}

resource "bloxone_dns_naptr_record" "example" {
  rdata = {
    flags       = "U"
    order       = 100
    preference  = 10
    regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
    replacement = "."
    services    = "SIP+D2U"
  }
  zone = bloxone_dns_auth_zone.example.id

  # Other optional fields
  name_in_zone = "naptr"
  comment      = "Example comment"
  disabled     = false
  ttl          = 3600
  tags = {
    location = "site1"
  }
}
