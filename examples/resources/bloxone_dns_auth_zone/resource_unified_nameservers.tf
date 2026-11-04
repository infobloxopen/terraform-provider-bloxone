# Zones that use the Unified Nameservers model (version = "v2") assign their
# nameservers inline and must not set primary_type.
resource "bloxone_dns_auth_zone" "example_unified" {
  fqdn = "unified.domain.com."

  nameservers = [
    {
      role    = "primary"
      address = "192.168.10.10"
      fqdn    = "ns1.domain.com."
    },
    {
      role    = "secondary"
      address = "192.168.10.11"
      fqdn    = "ns2.domain.com."
      stealth = false
    },
  ]

  # Other optional fields
  comment = "Example of an Authoritative Zone using Unified Nameservers"
  tags = {
    site = "Site A"
  }
}

# A zone may instead reference a single DNS Server Group. "nsg" and
# "nameservers" are mutually exclusive.
resource "bloxone_dns_auth_nsg" "example_unified_nsg" {
  name = "example_unified_nsg"
  nameservers = [
    {
      role    = "primary"
      address = "192.168.10.10"
      fqdn    = "ns1.domain.com."
    },
  ]
}

resource "bloxone_dns_auth_zone" "example_unified_with_nsg" {
  fqdn = "unified-nsg.domain.com."
  nsg  = bloxone_dns_auth_nsg.example_unified_nsg.id
}
