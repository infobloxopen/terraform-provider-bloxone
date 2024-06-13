<!-- BEGIN_TF_DOCS -->
# Terraform Module to Create BloxOne Anycast Configurations

This Terraform module is designed to configure BloxOne Anycast services with DHCP HA pairs and DNS, DFP configurations based on the specified services. It retrieves existing BloxOne hosts, sets up anycast configuration profiles, and adds protocols like BGP and OSPF.

## Features

   Anycast Configuration Profile: The module creates an anycast configuration profile and applies specified routing protocols.

   DHCP HA Group: When the service is set to DHCP and at least two hosts are provided, a DHCP HA group is created with anycast configuration.

   DNS Configuration: When the service is set to DNS, Anycast config profile is created with DNS

   DFP Configuration: When the service is set to DFP, Anycast config profile is created with DFP

## Module Workflow

   Retrieve BloxOne Hosts: The module fetches existing BloxOne hosts based on the provided host names or IP addresses.

   Create Anycast Configuration Profile: An anycast configuration profile is created with the desired anycast IP address and routing protocols.

   Configure Anycast Hosts: Each host is configured with the specified BGP and OSPF settings.

   DHCP HA Group Creation: If the service includes DHCP, a DHCP HA group is created using the provided hosts with anycast as the HA configuration.

   DNS Resources Creation: If the service includes DNS, DNS anycast config profile is created with anycast-enabled hosts.

   DFP Resources Creation: If the service includes DFP, DFP anycast config profile is created with anycast-enabled hosts.

## Example Usage

```hcl
 # Create a BloxOne Anycast Configuration for DHCP
module "bloxone_anycast" {

 anycast_config_name = "ac"

 hosts = {
   host1 = {
     role              = "active",
     routing_protocols = ["BGP", "OSPF"]
     bgp_config = {
       asn           = "65001"
       holddown_secs = 180
       neighbors = [
         { asn = "65002", ip_address = "172.28.4.198" }
       ]
     }
     ospf_config = {
       area                = "0.0.0.0"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   },
   host2 = {
     role              = "passive",
     routing_protocols = ["OSPF"]
     ospf_config = {
       area                = "0.0.0.1"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   }
 }

 service            = "dhcp"
 anycast_ip_address = "192.2.2.1"
 ha_name            = "example_ha_group"
 }

# Create a BloxOne Anycast Configuration for DNS
module "bloxone_anycast" {

 anycast_config_name = "ac"

 hosts = {
   host1 = {
     role              = "active",
     routing_protocols = ["BGP", "OSPF"]
     bgp_config = {
       asn           = "65001"
       holddown_secs = 180
       neighbors = [
         { asn = "65002", ip_address = "172.28.4.198" }
       ]
     }
     ospf_config = {
       area                = "0.0.0.0"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   },
   host2 = {
     role              = "passive",
     routing_protocols = ["OSPF"]
     ospf_config = {
       area                = "0.0.0.1"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   }
 }

 service            = "dns"
 anycast_ip_address = "192.2.2.1"
 ha_name            = null
}

# Create a BloxOne Anycast Configuration for DFP
module "bloxone_anycast" {

 anycast_config_name = "ac"

 hosts = {
   host1 = {
     role              = "active",
     routing_protocols = ["BGP", "OSPF"]
     bgp_config = {
       asn           = "65001"
       holddown_secs = 180
       neighbors = [
         { asn = "65002", ip_address = "172.28.4.198" }
       ]
     }
     ospf_config = {
       area                = "0.0.0.0"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   },
   host2 = {
     role              = "passive",
     routing_protocols = ["OSPF"]
     ospf_config = {
       area                = "0.0.0.1"
       area_type           = "STANDARD"
       authentication_type = "Clear"
       authentication_key  = "YXV0aGVk"
       interface           = "ens5"
       hello_interval      = 10
       dead_interval       = 40
       retransmit_interval = 5
       transmit_delay      = 1
     }
   }
   host3 = {
      role              = "passive",
      routing_protocols = ["OSPF"]
      ospf_config = {
         area                = "0.0.0.1"
         area_type           = "STANDARD"
         authentication_type = "Clear"
         authentication_key  = "YXV0aGVk"
         interface           = "ens5"
         hello_interval      = 10
         dead_interval       = 40
         retransmit_interval = 5
         transmit_delay      = 1
     }
   }

 service            = "dfp"
 anycast_ip_address = "192.2.2.1"
 ha_name            = null
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_bloxone"></a> [bloxone](#requirement\_bloxone) | >= 1.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | >= 1.1.0 |

## Resources

| Name | Type |
|------|------|
| [bloxone_anycast_config.ac](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/anycast_config) | resource |
| [bloxone_anycast_host.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/anycast_host) | resource |
| [bloxone_dhcp_ha_group.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dhcp_ha_group) | resource |
| [bloxone_infra_service.anycast](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service) | resource |
| [bloxone_dhcp_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dhcp_hosts) | data source |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_anycast_config_name"></a> [anycast\_config\_name](#input\_anycast\_config\_name) | Name of the Anycast configuration. | `string` | n/a | yes |
| <a name="input_anycast_ip_address"></a> [anycast\_ip\_address](#input\_anycast\_ip\_address) | Anycast IP address. | `string` | n/a | yes |
| <a name="input_ha_name"></a> [ha\_name](#input\_ha\_name) | Name of the HA group. | `string` | `null` | no |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | Map of hostnames with their roles, routing protocols, BGP, and OSPF configurations. | <pre>map(object({<br>    role               = string<br>    routing_protocols  = list(string)<br>    bgp_config = optional(object({<br>      asn            = string<br>      holddown_secs  = number<br>      neighbors      = list(object({<br>        asn        = string<br>        ip_address = string<br>      }))<br>    }))<br>    ospf_config = optional(object({<br>      area                = string<br>      area_type           = string<br>      authentication_type = string<br>      interface           = string<br>      authentication_key  = string<br>      hello_interval      = number<br>      dead_interval       = number<br>      retransmit_interval = number<br>      transmit_delay      = number<br>    }))<br>  }))</pre> | n/a | yes |
| <a name="input_service"></a> [service](#input\_service) | The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`). | `string` | `"dhcp"` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br>    create = string<br>    update = string<br>    read   = string<br>  })</pre> | `null` | no |
| <a name="input_wait_for_state"></a> [wait\_for\_state](#input\_wait\_for\_state) | If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API. | `bool` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_anycast_config"></a> [anycast\_config](#output\_anycast\_config) | The anycast config |
| <a name="output_anycast_hosts"></a> [anycast\_hosts](#output\_anycast\_hosts) | Map of anycast hosts |
| <a name="output_dhcp_ha_group"></a> [dhcp\_ha\_group](#output\_dhcp\_ha\_group) | The DHCP HA group |
| <a name="output_infra_services"></a> [infra\_services](#output\_infra\_services) | Map of infrastructure services |
<!-- END_TF_DOCS -->
