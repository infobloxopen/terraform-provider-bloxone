# Repro for the follow-up to #263: adding, at subnet level, an option that the
# parent address block already provides with the SAME value makes the API drop
# the subnet's copy, and the apply fails with
# "Provider produced inconsistent result after apply".
#
# Credentials come from INFOBLOX_PORTAL_URL / INFOBLOX_PORTAL_KEY env vars.
#
# Steps (use the released provider; remove the bloxone dev_overrides entry from
# ~/.terraformrc if you want to test v1.6.0 instead of a local build):
#   1. terraform apply                          -> works (routers + servers on subnet)
#   2. terraform apply -var add_duplicate=true  -> FAILS on the bug
#      (adds domain-name = "test.example", the same value the block provides)
#
# Control cases that work:
#   terraform apply -var add_duplicate=true -var duplicate_value=other.example

terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = "1.6.0"
    }
  }
}

# Credentials come from INFOBLOX_PORTAL_URL / INFOBLOX_PORTAL_KEY env vars.
provider "bloxone" {
  csp_url = "https://stage.csp.infoblox.com"
  api_key = "4a815e1e1c86a208efab3a5bfcc6f1f73259c009c43d30b13b337786ca9b3328"

  # Other Optional Fields
  default_tags = {
    managed_by = "terraform"
    site       = "Site A"
  }
}

variable "add_duplicate" {
  type    = bool
  default = false
}

variable "duplicate_value" {
  type    = string
  default = "test.example" # same as the address block
}

resource "bloxone_ipam_ip_space" "repro" {
  name = "repro-dup-ip-space"
}

# Custom option space/codes so the repro does not depend on tenant defaults.
resource "bloxone_dhcp_option_space" "repro" {
  name     = "repro-dup-option-space"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_code" "routers" {
  code         = 250
  name         = "repro-dup-routers"
  option_space = bloxone_dhcp_option_space.repro.id
  type         = "text"
}

resource "bloxone_dhcp_option_code" "servers" {
  code         = 251
  name         = "repro-dup-servers"
  option_space = bloxone_dhcp_option_space.repro.id
  type         = "text"
}

resource "bloxone_dhcp_option_code" "domain" {
  code         = 252
  name         = "repro-dup-domain"
  option_space = bloxone_dhcp_option_space.repro.id
  type         = "text"
}

# Parent block provides servers (2 values) and domain.
resource "bloxone_ipam_address_block" "repro" {
  address = "203.0.113.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.repro.id

  dhcp_options = [
    {
      option_code  = bloxone_dhcp_option_code.servers.id
      option_value = "192.0.2.53,192.0.2.54"
      type         = "option"
    },
    {
      option_code  = bloxone_dhcp_option_code.domain.id
      option_value = "test.example"
      type         = "option"
    },
  ]
}

# Subnet: routers + servers (value differs from the block) from the start.
# With add_duplicate=true, domain is added with the block's value.
resource "bloxone_ipam_subnet" "repro" {
  address = "203.0.113.128"
  cidr    = 26
  space   = bloxone_ipam_ip_space.repro.id

  depends_on = [bloxone_ipam_address_block.repro]

  dhcp_options = concat(
    [
      {
        option_code  = bloxone_dhcp_option_code.routers.id
        option_value = "203.0.113.129"
        type         = "option"
      },
      {
        option_code  = bloxone_dhcp_option_code.servers.id
        option_value = "192.0.2.53"
        type         = "option"
      },
    ],
    var.add_duplicate ? [
      {
        option_code  = bloxone_dhcp_option_code.domain.id
        option_value = var.duplicate_value
        type         = "option"
      },
    ] : []
  )
}
