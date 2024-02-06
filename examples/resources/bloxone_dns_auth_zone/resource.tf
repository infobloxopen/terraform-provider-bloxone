resource "bloxone_keys_tsig" "test" {
  name = "test-tsig."
}

resource "bloxone_dns_acl" "test" {
  name = "test-acl"
}

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
    },
    {
      element = "acl"
      acl = bloxone_dns_acl.test.id
    },
    {
      element = "tsig_key"
      access = "deny"
      tsig_key = {
        key = bloxone_keys_tsig.test.id
      }
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
    },
    {
      element = "acl"
      acl = bloxone_dns_acl.test.id
    },
    {
      element = "tsig_key"
      access = "deny"
      tsig_key = {
        key = bloxone_keys_tsig.test.id
      }
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
    },
    {
      element = "acl"
      acl = bloxone_dns_acl.test.id
    },
    {
      element = "tsig_key"
      access = "deny"
      tsig_key = {
        key = bloxone_keys_tsig.test.id
      }
    }
  ]

}