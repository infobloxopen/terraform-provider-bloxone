


/**
 * # Terraform Module to create BloxOne Anycast in AWS
 *
 * This module will retrieve an AWS EC2 instance that uses a BloxOne AMI.
 * The instance will be configured to setup anycast config profile and add anycast protocols like BGP, OSPF.
 *
 *
 * This module will fetch the already created BloxOne Host and create an anycast config profile with the desired routing protocols.
 *
 * ## Example Usage
 *
 * ```hcl
 * module "bloxone_anycast" {
 *   host_name = "MY_HOST_NAME"
 *
 *  hosts = [
 *    "HOST_1"
 *    "HOST_2"
 *     ]
 *
 *   name = "ac"
 *   service = "DNS"
 *   ip_address = "192.2.2.1"
 *   routing_protocols   = ["BGP", "OSPF"]
 *   # Adding the BGP configuration
 *     config_bgp = {
 *      ...
      }
 *     config_ospf = {
 *     ...
 *    }
 * }
 * ```
 * 
 */

//Fetch the aws instance
# data "bloxone_infra_hosts" "anycast_host_1" {
#   filters = {
#     "name" = "example_host_1"
#   }
# }

data "bloxone_infra_hosts" "this" {
  for_each       = var.hosts
  filters = {
    "display_name" = each.value
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
  id        = data.bloxone_infra_hosts.this[each.key].results.0.legacy_id

  # Adding the anycast config profile and enabling BGP routing protocol
  anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.ac.name
      routing_protocols   = ["BGP", "OSPF"]
    }
  ]

  # Adding the BGP configuration
  config_bgp = {
    asn          = var.asn
    holddown_secs = var.holddown_secs
    neighbors    = var.bgp_neighbors
  }

  # Adding the OSPF configuration
  config_ospf = var.ospf_config
}

data "bloxone_dhcp_hosts" "this" {
  for_each       = var.hosts
  filters = {
    name = each.value
  }
}

# Define the HA group resource
resource "bloxone_dhcp_ha_group" "example_anycast" {
  name              = "example_ha_group_anycast"
  mode              = "anycast"
  anycast_config_id = format("accm/ac_configs/%s", bloxone_anycast_config.ac.id)

  hosts = [
    {
      host = data.bloxone_dhcp_hosts.this["host1"].results[0].id
      role = "active"
    },
    {
      host = data.bloxone_dhcp_hosts.this["host2"].results[0].id
      role = "passive"
    }
  ]
}

output "first_host_id" {
  value = data.bloxone_dhcp_hosts.this["host1"].results[0].id
}