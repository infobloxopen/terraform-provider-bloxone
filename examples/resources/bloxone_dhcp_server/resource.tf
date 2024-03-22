
data "bloxone_dns_hosts" "my_host" {
  filters = {
    name = "my_host"
  }
}

resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "domain.com."
  primary_type = "cloud"
  internal_secondaries = [
    {
      host = data.bloxone_dns_hosts.my_host.results.0.id
    },
  ]
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

resource "bloxone_dhcp_server" "example" {
  name = "example"

  ddns_enabled = "true"
  ddns_domain  = "domain.com."

  # ddns_zones configuration
  ddns_zones = [
    {
      gss_tsig_enabled = false
      tsig_enabled     = false
      tsig_key         = null
      zone             = bloxone_dns_auth_zone.example.id
    }
  ]

  //dhcp options configuration
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "10.0.0.1"
      type         = "option"
    }
  ]
}
