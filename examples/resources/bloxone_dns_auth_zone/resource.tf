terraform {
  required_providers {
    bloxone = {
      source  = "registry.terraform.io/infobloxopen/bloxone"
      version = "0.0.1"
      # Other parameters...
    }
  }
}

provider "bloxone" {
  csp_url = "https://stage.csp.infoblox.com"
  api_key = "49e506b53774b20427c6db7bf4a68bb846c3dcdb2cc2de786f19a75f29010a8e"
}

resource "bloxone_keys_tsig" "example_tsig" {
  name = "test-tsig."
}

resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "domain.com."
  primary_type = "cloud"

  # Other optional fields
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Site A"
  }
  transfer_acl = [
    {
      access  = "deny"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "allow"
      element = "ip"
      address = "10.0.0.0/24"
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
      access  = "deny"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "allow"
      element = "ip"
      address = "10.0.0.0/24"
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
  query_acl = [
    {
      access  = "deny"
      element = "ip"
      address = "192.168.1.1"
    },
    {
      access  = "allow"
      element = "ip"
      address = "10.0.0.0/24"
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

}