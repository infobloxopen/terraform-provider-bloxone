
resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_dns_view" "example" {
  name = "example_dns_view"

  # Other Optional fields
  comment   = "An example view"
  ip_spaces = [bloxone_ipam_ip_space.example.id]
  tags = {
    site = "Site A"
  }
  match_clients_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.10.10"
    }
  ]
  custom_root_ns_enabled = true
  custom_root_ns = [
    {
      address = "192.168.11.11"
      fqdn    = "example.com."
    }
  ]
  ecs_enabled    = true
  ecs_forwarding = false
  ecs_prefix_v4  = 24
  ecs_prefix_v6  = 56
  ecs_zones = [
    {
      access = "allow"
      fqdn   = "example.com."
    }
  ]
  recursion_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.1.1"
    }
  ]
  dnssec_enabled         = true
  dnssec_validate_expiry = true
  dnssec_trust_anchors = [
    {
      algorithm  = 8
      public_key = "AwEAAejpWrcCPGWEoiebhWKSdT6LcMGBsoXadKu1XNthMZUvx3P92HNE4J3q3EtAX8pnTsNShrsDvvgn4hmCsrURMLx/g+76JtLU5pdbtrGFjelHAuMrzLgFzpuA5Ct9THth5Hto6c0rl4yzz3qT3+I/rnUYrL/zd9zKWyMp1A9KlHqwCA3JbFZfl4IKBD2/g+GScEcpnDfUUVDU+7qRZkZ4BhBQ4a6Em73zggz/crcDtwc1cHcRP0DGbekZhF29+yjTPW4zKqGUHW8ZtP49ZMXOTY42epeiddFNy0Ze2jbTg99CnKvAxIKzYInUaPJ04rgMyeuVWpRKsVetJemhCaj9lEs="
      zone       = "example.com."
      sep        = true
    }
  ]
  zone_authority = {
    default_ttl       = 28800
    expire            = 2419200
    mname             = "ns.b1ddi.example.com"
    negative_ttl      = 900
    refresh           = 10800
    retry             = 3600
    rname             = "hostmaster@example.com"
    use_default_mname = false
  }

}
