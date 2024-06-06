/**
 /**
 * # Terraform Module to Create BloxOne Anycast Configurations
 *
 * This Terraform module is designed to configure BloxOne Anycast services with DHCP HA pairs and DNS configurations based on the specified services. It retrieves existing BloxOne hosts, sets up anycast configuration profiles, and adds protocols like BGP and OSPF.
 *
 * ## Features
 *
 *    Anycast Configuration Profile: The module creates an anycast configuration profile and applies specified routing protocols.
 *
 *    DHCP HA Group: When the service is set to DHCP and at least two hosts are provided, a DHCP HA group is created with anycast configuration.
 *
 *    DNS Configuration: When the service is set to DNS, DNS resources including DNS views and authoritative zones are created.
 *
 * ## Module Workflow
 *
 *    Retrieve BloxOne Hosts: The module fetches existing BloxOne hosts based on the provided host names or IP addresses.
 *
 *    Create Anycast Configuration Profile: An anycast configuration profile is created with the desired anycast IP address and routing protocols.
 *
 *    Configure Anycast Hosts: Each host is configured with the specified BGP and OSPF settings.
 *
 *    DHCP HA Group Creation: If the service includes DHCP, a DHCP HA group is created using the provided hosts with anycast as the HA configuration.
 *
 *    DNS Resources Creation: If the service includes DNS, DNS views and authoritative zones are created with anycast-enabled hosts.
 *
 * ## Example Usage
 *
 * ```hcl
 *   module "bloxone_anycast" {
 *
 *   anycast_config_name = "anycast_config"
 *   hosts = {
 *     host1 = {
 *       role             = "active",
 *       routing_protocols = ["BGP", "OSPF"]
 *     },
 *     host2 = {
 *       role             = "passive",
 *       routing_protocols = ["OSPF"]
 *     },
 *     host3 = {
 *       role             = "active",
 *       routing_protocols = ["BGP"]
 *     },
 *     host4 = {
 *       role             = "active",
 *       routing_protocols = ["OSPF"]
 *     }
 *   }
 *   service             = "DHCP"
 *   anycast_ip_address  = "192.2.2.1"
 *
 *   bgp_configs = {
 *     host1 = {
 *       asn           = "65001",
 *       holddown_secs = 180,
 *       neighbors = [
 *         {
 *           asn        = "65002",
 *           ip_address = "192.0.2.1"
 *         }
 *       ]
 *     },
 *     host3 = {
 *       asn           = "65003",
 *       holddown_secs = 180,
 *       neighbors = [
 *         {
 *           asn        = "65004",
 *           ip_address = "192.0.2.2"
 *         }
 *       ]
 *     }
 *   }
 *
 *   ospf_configs = {
 *     host1 = {
 *       area                = "0.0.0.0",
 *       area_type           = "STANDARD",
 *       authentication_type = "Clear",
 *       authentication_key  = "YXV0aGVk",
 *       interface           = "ens5",
 *       hello_interval      = 10,
 *       dead_interval       = 40,
 *       retransmit_interval = 5,
 *       transmit_delay      = 1
 *     },
 *     host2 = {
 *       area                = "0.0.0.1",
 *       area_type           = "STANDARD",
 *       authentication_type = "Clear",
 *       authentication_key  = "YXV0aGVk",
 *       interface           = "ens5",
 *       hello_interval      = 10,
 *       dead_interval       = 40,
 *       retransmit_interval = 5,
 *       transmit_delay      = 1
 *     },
 *     host4 = {
 *       area                = "0.0.0.3",
 *       area_type           = "STANDARD",
 *       authentication_type = "Clear",
 *       authentication_key  = "YXV0aGVk",
 *       interface           = "ens5",
 *       hello_interval      = 11,
 *       dead_interval       = 40,
 *       retransmit_interval = 5,
 *       transmit_delay      = 1
 *     }
 *   }
 *
 *   ha_name    = "example_ha_group"
 *   view_name  = "example_view"
 *   fqdn       = "example.com"
 *   primary_type = "cloud"
 * }
 ```
 */

# Ensure at least 2 hosts for DHCP
locals {
  effective_hosts = var.service == "DHCP" ? { for k, v in var.hosts : k => v if length(var.hosts) >= 2 } : var.hosts
}

data "bloxone_infra_hosts" "this" {
  for_each = local.effective_hosts
  filters = {
    "display_name" = each.key
  }
}

# Create an anycast config profile with on-prem hosts
resource "bloxone_anycast_config" "ac" {
  anycast_ip_address = var.anycast_ip_address
  name               = var.anycast_config_name
  service            = var.service
}

# Adding an anycast host with BGP routing protocol
resource "bloxone_anycast_host" "this" {
  for_each  = data.bloxone_infra_hosts.this
  id        = data.bloxone_infra_hosts.this[each.key].results[0].legacy_id

  # Adding the anycast config profile and enabling BGP routing protocol
  anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.ac.name
      routing_protocols   = var.hosts[each.key].routing_protocols
    }
  ]

  # Adding the BGP configuration
  config_bgp = contains(var.hosts[each.key].routing_protocols, "BGP") ? var.bgp_configs[each.key] : null

  # Adding the OSPF configuration
  config_ospf = contains(var.hosts[each.key].routing_protocols, "OSPF") ? var.ospf_configs[each.key] : null
}

data "bloxone_dhcp_hosts" "this" {
  for_each = local.effective_hosts
  filters = {
    name = each.key
  }
}

# Define the HA group resource
resource "bloxone_dhcp_ha_group" "example_anycast" {
  count = var.service == "DHCP" && length(var.hosts) >= 2 ? 1 : 0
  name  = var.ha_name
  mode  = "anycast"
  anycast_config_id = format("accm/ac_configs/%s", bloxone_anycast_config.ac.id)

  hosts = [
    for host_key in slice(keys(var.hosts), 0, 2) : {
      host = data.bloxone_dhcp_hosts.this[host_key].results[0].id
      role = var.hosts[host_key].role
    }
  ]
}

# Create a DNS view resource if the service is DNS
resource "bloxone_dns_view" "example" {
  count    = var.service == "DNS" ? 1 : 0
  name     = var.view_name
}

# Authzone with anycast host
resource "bloxone_dns_auth_zone" "authzone" {
  count = var.service == "DNS" ? 1 : 0
  fqdn               = "${var.fqdn}."
  primary_type       = var.primary_type
  view               = bloxone_dns_view.example[0].id

  internal_secondaries = [
    for host in keys(var.hosts) : {
      host = format("dns/host/%s", data.bloxone_infra_hosts.this[host].results[0].legacy_id)
    }
  ]
}
