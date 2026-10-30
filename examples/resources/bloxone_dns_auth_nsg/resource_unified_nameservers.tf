# DNS Server Groups that use the Unified Nameservers model (version = "v2")
# assign every nameserver through the single "nameservers" list.
resource "bloxone_dns_auth_nsg" "example_unified" {
  name = "example_unified_dns_auth_nsg"

  nameservers = [
    {
      role    = "primary"
      address = "12.10.2.1"
      fqdn    = "ns1.primary.com."
    },
    {
      role         = "secondary"
      address      = "12.10.2.2"
      fqdn         = "ns2.secondary.com."
      stealth      = false
      tsig_enabled = false
    },
  ]

  # Other optional fields
  comment = "An example auth nsg using Unified Nameservers"
  tags = {
    site = "Site A"
  }
}
