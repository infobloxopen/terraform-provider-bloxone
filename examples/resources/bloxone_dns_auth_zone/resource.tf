resource "bloxone_keys_tsig" "example_tsig" {
  name = "example_tsig.domain.com."
}

resource "bloxone_dns_acl" "example_acl" {
  name = "example_acl"
  list = [
    {
      access  = "deny"
      element = "ip"
      address = "192.168.1.0/24"
    },
  ]
}

resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "domain.com."
  primary_type = "cloud"

  # Other optional fields
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Site A"
  }
  query_acl = [
    {
      access  = "deny"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      element = "acl"
      acl     = bloxone_dns_acl.example_acl.id
    },
    {
      access  = "allow"
      element = "tsig_key"
      tsig_key = {
        key = bloxone_keys_tsig.example_tsig.id
      }
    },
    {
      access  = "deny"
      element = "any"
    },
  ]
  update_acl = [
    {
      access  = "allow"
      element = "any"
    },
  ]
  transfer_acl = [
    {
      access  = "allow"
      element = "any"
    },
  ]
}
