<!-- BEGIN_TF_DOCS -->
# Terraform Module to Create BloxOne Anycast Configurations

This Terraform module configures BloxOne Anycast services for DHCP HA pairs, DNS, and DFP based on the specified service type. It fetches BloxOne hosts by provided names, creates an anycast configuration profile with the desired IP address and routing protocols, and if the service type is `dhcp`, it also creates a DHCP HA group.

Note: The module only creates the `anycast` service object and assumes pre-existing hosts in BloxOne and pre-configured `dhcp`, `dns`, or `dfp` services.
## Example Usage

### Anycast Configuration for DHCP
```hcl
module "bloxone_anycast" {
 anycast_config_name = "ac"

 hosts = {
   host1 = {
     ha_role              = "active",
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
     ha_role              = "passive",
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
 ha_group_name      = "example_ha_group"
 }
```
### Anycast Configuration for DNS
```hcl

 module "bloxone_anycast" {
  anycast_config_name = "ac"

  hosts = {
    host1 = {
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
 }
```
### Anycast Configuration for DFP
```hcl
module "bloxone_anycast" {
 anycast_config_name = "ac"

 hosts = {
   host1 = {
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
 }
 service            = "dfp"
 anycast_ip_address = "192.2.2.1"
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
| <a name="input_ha_group_name"></a> [ha\_group\_name](#input\_ha\_group\_name) | Name of the HA group. | `string` | `null` | no |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | Map of hostnames with their roles, routing protocols, BGP, and OSPF configurations. | <pre>map(object({<br>    ha_role           = optional(string)<br>    routing_protocols = list(string)<br>    bgp_config = optional(object({<br>      asn           = optional(string)<br>      holddown_secs = optional(number)<br>      neighbors = optional(list(object({<br>        asn        = string<br>        ip_address = string<br>      })))<br>    }))<br>    ospf_config = optional(object({<br>      area                = optional(string)<br>      area_type           = optional(string)<br>      authentication_type = optional(string)<br>      authentication_key  = optional(string)<br>      interface           = optional(string)<br>      hello_interval      = optional(number)<br>      dead_interval       = optional(number)<br>      retransmit_interval = optional(number)<br>      transmit_delay      = optional(number)<br>    }))<br>  }))</pre> | `{}` | no |
| <a name="input_service"></a> [service](#input\_service) | The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`). | `string` | `"dhcp"` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br>    create = string<br>    update = string<br>    read   = string<br>  })</pre> | `null` | no |
| <a name="input_wait_for_state"></a> [wait\_for\_state](#input\_wait\_for\_state) | If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API. | `bool` | `true` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_anycast_config"></a> [anycast\_config](#output\_anycast\_config) | The anycast config |
| <a name="output_anycast_hosts"></a> [anycast\_hosts](#output\_anycast\_hosts) | Map of anycast hosts |
| <a name="output_dhcp_ha_group"></a> [dhcp\_ha\_group](#output\_dhcp\_ha\_group) | The DHCP HA group |
| <a name="output_infra_services"></a> [infra\_services](#output\_infra\_services) | Map of infrastructure services |
<!-- END_TF_DOCS -->
