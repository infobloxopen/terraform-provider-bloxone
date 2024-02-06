resource "bloxone_keys_tsig" "test" {
  name = "test-tsig."
}

resource "bloxone_dns_acl" "test_acl" {
  name = "test-acl"
}

resource "bloxone_dns_acl" "example_acl" {
  name = "example_dns_acl"

  # Other Optional fields
  comment = "An example acl"
  tags = {
    site = "Site A"
  }
  list = [
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
      acl     = bloxone_dns_acl.test_acl.id
    },
    {
      element = "tsig_key"
      access  = "deny"
      tsig_key = {
        key = bloxone_keys_tsig.test.id
      }
    }
  ]
}
