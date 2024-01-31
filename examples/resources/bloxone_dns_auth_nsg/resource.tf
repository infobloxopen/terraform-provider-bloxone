resource "bloxone_dns_auth_nsg" "example" {
  name = "example_dns_auth_nsg"

  # Other Optional fields
  comment = "An example auth nsg"
  external_primaries = [
    {
      address = "12.10.2.1",
      fqdn    = "ext.primary.com."
      type    = "primary"
    },
  ]
  external_secondaries = [
    {
      address = "12.10.2.2",
      fqdn    = "ext.primary.com."
    }
  ]
  tags = {
    site = "Site A"
  }
}

