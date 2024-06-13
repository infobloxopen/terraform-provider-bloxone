/**
 /**
 * # Terraform Module to Create BloxOne Anycast Configurations
 *
 * This Terraform module is designed to configure BloxOne Anycast services with DHCP HA pairs and DNS, DFP configurations based on the specified services. It retrieves existing BloxOne hosts, sets up anycast configuration profiles, and adds protocols like BGP and OSPF.
 *
 * ## Features
 *
 *    Anycast Configuration Profile: The module creates an anycast configuration profile and applies specified routing protocols.
 *
 *    DHCP HA Group: When the service is set to DHCP and at least two hosts are provided, a DHCP HA group is created with anycast configuration.
 *
 *    DNS Configuration: When the service is set to DNS, Anycast config profile is created with DNS
 *
 *    DFP Configuration: When the service is set to DFP, Anycast config profile is created with DFP
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
 *    DNS Resources Creation: If the service includes DNS, DNS anycast config profile is created with anycast-enabled hosts.
 *
 *    DFP Resources Creation: If the service includes DFP, DFP anycast config profile is created with anycast-enabled hosts.
 *
 * ## Example Usage
 *
 * ```hcl
 *  # Create a BloxOne Anycast Configuration for DHCP
 * module "bloxone_anycast" {
 *
 *  anycast_config_name = "ac"
 *
 *  hosts = {
 *    host1 = {
 *      role              = "active",
 *      routing_protocols = ["BGP", "OSPF"]
 *      bgp_config = {
 *        asn           = "65001"
 *        holddown_secs = 180
 *        neighbors = [
 *          { asn = "65002", ip_address = "172.28.4.198" }
 *        ]
 *      }
 *      ospf_config = {
 *        area                = "0.0.0.0"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    },
 *    host2 = {
 *      role              = "passive",
 *      routing_protocols = ["OSPF"]
 *      ospf_config = {
 *        area                = "0.0.0.1"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    }
 *  }
 *
 *  service            = "dhcp"
 *  anycast_ip_address = "192.2.2.1"
 *  ha_name            = "example_ha_group"
 *  }
 *
 * # Create a BloxOne Anycast Configuration for DNS
 * module "bloxone_anycast" {
 *
 *  anycast_config_name = "ac"
 *
 *  hosts = {
 *    host1 = {
 *      role              = "active",
 *      routing_protocols = ["BGP", "OSPF"]
 *      bgp_config = {
 *        asn           = "65001"
 *        holddown_secs = 180
 *        neighbors = [
 *          { asn = "65002", ip_address = "172.28.4.198" }
 *        ]
 *      }
 *      ospf_config = {
 *        area                = "0.0.0.0"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    },
 *    host2 = {
 *      role              = "passive",
 *      routing_protocols = ["OSPF"]
 *      ospf_config = {
 *        area                = "0.0.0.1"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    }
 *  }
 *
 *  service            = "dns"
 *  anycast_ip_address = "192.2.2.1"
 *  ha_name            = null
 * }
 *
 * # Create a BloxOne Anycast Configuration for DNS
 * module "bloxone_anycast" {
 *
 *  anycast_config_name = "ac"
 *
 *  hosts = {
 *    host1 = {
 *      role              = "active",
 *      routing_protocols = ["BGP", "OSPF"]
 *      bgp_config = {
 *        asn           = "65001"
 *        holddown_secs = 180
 *        neighbors = [
 *          { asn = "65002", ip_address = "172.28.4.198" }
 *        ]
 *      }
 *      ospf_config = {
 *        area                = "0.0.0.0"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    },
 *    host2 = {
 *      role              = "passive",
 *      routing_protocols = ["OSPF"]
 *      ospf_config = {
 *        area                = "0.0.0.1"
 *        area_type           = "STANDARD"
 *        authentication_type = "Clear"
 *        authentication_key  = "YXV0aGVk"
 *        interface           = "ens5"
 *        hello_interval      = 10
 *        dead_interval       = 40
 *        retransmit_interval = 5
 *        transmit_delay      = 1
 *      }
 *    }
 *    host3 = {
 *       role              = "passive",
 *       routing_protocols = ["OSPF"]
 *       ospf_config = {
 *          area                = "0.0.0.1"
 *          area_type           = "STANDARD"
 *          authentication_type = "Clear"
 *          authentication_key  = "YXV0aGVk"
 *          interface           = "ens5"
 *          hello_interval      = 10
 *          dead_interval       = 40
 *          retransmit_interval = 5
 *          transmit_delay      = 1
 *      }
 *    }
 *
 *  service            = "dfp"
 *  anycast_ip_address = "192.2.2.1"
 *  ha_name            = null
 * }
 ```
 */

locals {
  service_type_to_anycast_service_type = {
    "dhcp" = "DHCP"
    "dns"  = "DNS"
    "dfp"  = "DFP"
  }
}

data "bloxone_infra_hosts" "this" {
  for_each = var.hosts
  filters = {
    "display_name" = each.key
  }

  lifecycle {
    postcondition {
      condition = self.results != null
      error_message = "Host not found for ${each.key}"
    }

    postcondition {
      condition = contains(self.results[0].configs[*].service_type, var.service)
      error_message = "${var.service} for ${each.key} is not configured"
    }
  }
}

# Create an anycast config profile with on-prem hosts
resource "bloxone_anycast_config" "ac" {
  anycast_ip_address = var.anycast_ip_address
  name               = var.anycast_config_name
  service            = local.service_type_to_anycast_service_type[var.service]
}

resource "bloxone_infra_service" "anycast" {
  for_each       = var.hosts
  name           = format("%s_anycast", each.key)
  pool_id        = data.bloxone_infra_hosts.this[each.key].results[0].pool_id
  service_type   = "anycast"
  desired_state  = "start"
  wait_for_state = false
}

# Adding an anycast host with BGP and OSPF routing protocol
resource "bloxone_anycast_host" "this" {
  for_each  = data.bloxone_infra_hosts.this
  id        = one(data.bloxone_infra_hosts.this[each.key].results).legacy_id

  # Adding the anycast config profile and enabling BGP routing protocol
  anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.ac.name
      routing_protocols   = var.hosts[each.key].routing_protocols
    }
  ]

  # Adding the BGP configuration if specified
  config_bgp = contains(var.hosts[each.key].routing_protocols, "BGP") && var.hosts[each.key].bgp_config != null ? var.hosts[each.key].bgp_config : null

  # Adding the OSPF configuration if specified
  config_ospf = contains(var.hosts[each.key].routing_protocols, "OSPF") && var.hosts[each.key].ospf_config != null ? var.hosts[each.key].ospf_config : null
}

data "bloxone_dhcp_hosts" "this" {
  for_each = var.hosts
  filters = {
    name = each.key
  }

  lifecycle {
    postcondition {
      condition = self.results != null
      error_message = "Host not found for ${each.key}"
    }
  }
}

# Define the HA group resource
resource "bloxone_dhcp_ha_group" "this" {
  count = var.service == "dhcp" ? 1 : 0
  name  = var.ha_name
  mode  = "anycast"
  anycast_config_id = format("accm/ac_configs/%s", bloxone_anycast_config.ac.id)

  hosts = [
    for k, v in var.hosts : {
      host = data.bloxone_dhcp_hosts.this[k].results[0].id
      role = v.role
    }
  ]

  lifecycle {
    precondition {
      condition = length(var.hosts) >= 2
      error_message = "At least two hosts are required for DHCP HA"
    }
  }
}
