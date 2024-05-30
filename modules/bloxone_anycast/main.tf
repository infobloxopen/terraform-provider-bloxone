/**
 * # Terraform Module to create BloxOne Anycast in AWS
 *
 * This module will retrieve an AWS EC2 instance that uses a BloxOne AMI.
 * The instance will be configured to setup anycast config profile and add anycast protocols like BGP, OSPF.
 * Then a DHCP HA group will be created with the provided hosts with anycast as HA configuration.
 *
 * This module will fetch the already created BloxOne Host and create an anycast config profile with the desired routing protocols.
 *
 * ## Example Usage
 *
 * ```hcl
 * module "bloxone_anycast" {
 *  source = "/Users/agadiyarhj/go/src/github.com/infobloxopen/terraform-provider-bloxone/modules/bloxone_anycast"
 *
 *  hosts = {
 *    host1 = "active",
 *    host2 = "passive"
 *  }
 *  name      = "ac"
 *  service   = "DHCP"
 *  anycast_ip_address = "192.2.2.1"
 *  routing_protocols = ["BGP", "OSPF"]
 *
 *  bgp_config = {
 *    asn           = "6500"
 *    holddown_secs = 180
 *    neighbors     = [
 *      {
 *        asn        = "6501"
 *        ip_address = "172.28.4.198"
 *      }
 *    ]
 *  }
 *
 *  ospf_config = {
 *    area                = "10.10.0.1"
 *    area_type           = "STANDARD"
 *    authentication_type = "Clear"
 *    interface           = "eth0"
 *    authentication_key  = "YXV0aGV"
 *    hello_interval      = 10
 *    dead_interval       = 40
 *    retransmit_interval = 5
 *    transmit_delay      = 1
 *  }
 *}
 ```
 */

data "bloxone_infra_hosts" "this" {
  for_each       = var.hosts
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
      routing_protocols   = ["BGP", "OSPF"]
    }
  ]

  # Adding the BGP configuration
  config_bgp = var.bgp_config

  # Adding the OSPF configuration
  config_ospf = var.ospf_config
}

data "bloxone_dhcp_hosts" "this" {
  for_each       = var.hosts
  filters = {
    name = each.key
  }
}

# Define the HA group resource
resource "bloxone_dhcp_ha_group" "example_anycast" {
  name              = "example_ha_group_anycast"
  mode              = "anycast"
  anycast_config_id = format("accm/ac_configs/%s", bloxone_anycast_config.ac.id)

  hosts = [
    for host, role in var.hosts : {
      host  = data.bloxone_dhcp_hosts.this[host].results[0].id
      role = role
    }
  ]
}
