resource "bloxone_dhcp_server" "example" {
  name = "example_dhcp_server"
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

resource "bloxone_dhcp_server" "example_with_options" {
  name = "example_dhcp_server_with_options"

  #Other Optional Fields
  comment = "dhcp server"
  tags = {
    site = "Site A"
  }

  //dhcp options
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "10.0.0.1"
      type         = "option"
    }
  ]
}

data "bloxone_dns_hosts" "example_dns_host" {
  filters = {
    name = "DNSHost01"
  }
}

resource "bloxone_dns_view" "example_dns_view" {
  name = "example_dns_view"

  # Other Optional fields
  comment = "An example view"
  tags = {
    site = "Site A"
  }

}

resource "bloxone_dns_auth_zone" "example_auth_zone" {
  fqdn         = "domain.com."
  primary_type = "cloud"

  # Other optional fields
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Site A"
  }
  internal_secondaries = [
    {
      host = data.bloxone_dns_hosts.example_dns_host.results.0.id
    },
  ]
  view = bloxone_dns_view.example_dns_view.id
}

resource "bloxone_dns_auth_zone" "example_reverse_zone" {
  comment            = "This is a comment about the resource."
  disabled           = false
  fqdn               = "168.192.in-addr.arpa."
  primary_type       = "cloud"
  gss_tsig_enabled   = false
  initial_soa_serial = 1
  internal_secondaries = [
    {
      host = data.bloxone_dns_hosts.example_dns_host.results.0.id
    },
  ]
  view = bloxone_dns_view.example_dns_view.id
}

resource "bloxone_dhcp_server" "example_with_ddns" {
  name = "example_dhcp_server_with_ddns"

  #Other Optional Fields
  comment = "dhcp server setup example for ddns"
  tags = {
    site = "Site A"
  }

  ddns_enabled = "true"
  ddns_domain  = "domain.com."

  # ddns_zones configuration
  ddns_zones = [
    {
      gss_tsig_enabled = false
      tsig_enabled     = false
      tsig_key         = null
      zone             = bloxone_dns_auth_zone.example_reverse_zone.id
    },
    {
      gss_tsig_enabled = false
      tsig_enabled     = false
      tsig_key         = null
      zone             = bloxone_dns_auth_zone.example_auth_zone.id
    }
  ]
}
