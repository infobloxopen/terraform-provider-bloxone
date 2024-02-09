resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "example.com."
  primary_type = "cloud"

  # Other optional fields
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Site A"
  }
  transfer_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "deny"
      element = "any"
    }
  ]
  update_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "deny"
      element = "any"
    }
  ]
  query_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "deny"
      element = "any"
    }
  ]

}