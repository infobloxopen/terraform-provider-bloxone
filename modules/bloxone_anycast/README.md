<!-- BEGIN_TF_DOCS -->
# Terraform Module to Create BloxOne Anycast Configurations

This Terraform module is designed to configure BloxOne Anycast services with DHCP HA pairs and DNS configurations based on the specified services. It retrieves existing BloxOne hosts, sets up anycast configuration profiles, and adds protocols like BGP and OSPF.

## Features

   Anycast Configuration Profile: The module creates an anycast configuration profile and applies specified routing protocols.

   DHCP HA Group: When the service is set to DHCP and at least two hosts are provided, a DHCP HA group is created with anycast configuration.

   DNS Configuration: When the service is set to DNS, DNS resources including DNS views and authoritative zones are created.

## Module Workflow

   Retrieve BloxOne Hosts: The module fetches existing BloxOne hosts based on the provided host names or IP addresses.

   Create Anycast Configuration Profile: An anycast configuration profile is created with the desired anycast IP address and routing protocols.

   Configure Anycast Hosts: Each host is configured with the specified BGP and OSPF settings.

   DHCP HA Group Creation: If the service includes DHCP, a DHCP HA group is created using the provided hosts with anycast as the HA configuration.

   DNS Resources Creation: If the service includes DNS, DNS views and authoritative zones are created with anycast-enabled hosts.

## Example Usage

```hcl
  module "bloxone_anycast" {

  anycast_config_name = "anycast_config"
  hosts = {
    host1 = {
      role             = "active",
      routing_protocols = ["BGP", "OSPF"]
    },
    host2 = {
      role             = "passive",
      routing_protocols = ["OSPF"]
    },
    host3 = {
      role             = "active",
      routing_protocols = ["BGP"]
    },
    host4 = {
      role             = "active",
      routing_protocols = ["OSPF"]
    }
  }
  service             = "DHCP"
  anycast_ip_address  = "192.2.2.1"

  bgp_configs = {
    host1 = {
      asn           = "65001",
      holddown_secs = 180,
      neighbors = [
        {
          asn        = "65002",
          ip_address = "192.0.2.1"
        }
      ]
    },
    host3 = {
      asn           = "65003",
      holddown_secs = 180,
      neighbors = [
        {
          asn        = "65004",
          ip_address = "192.0.2.2"
        }
      ]
    }
  }

  ospf_configs = {
    host1 = {
      area                = "0.0.0.0",
      area_type           = "STANDARD",
      authentication_type = "Clear",
      authentication_key  = "YXV0aGVk",
      interface           = "ens5",
      hello_interval      = 10,
      dead_interval       = 40,
      retransmit_interval = 5,
      transmit_delay      = 1
    },
    host2 = {
      area                = "0.0.0.1",
      area_type           = "STANDARD",
      authentication_type = "Clear",
      authentication_key  = "YXV0aGVk",
      interface           = "ens5",
      hello_interval      = 10,
      dead_interval       = 40,
      retransmit_interval = 5,
      transmit_delay      = 1
    },
    host4 = {
      area                = "0.0.0.3",
      area_type           = "STANDARD",
      authentication_type = "Clear",
      authentication_key  = "YXV0aGVk",
      interface           = "ens5",
      hello_interval      = 11,
      dead_interval       = 40,
      retransmit_interval = 5,
      transmit_delay      = 1
    }
  }

  ha_name    = "example_ha_group"
  view_name  = "example_view"
  fqdn       = "example.com"
  primary_type = "cloud"
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_bloxone"></a> [bloxone](#requirement\_bloxone) | >= 1.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | 0.0.1 |

## Resources

| Name | Type |
|------|------|
| [bloxone_anycast_config.ac](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/anycast_config) | resource |
| [bloxone_anycast_host.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/anycast_host) | resource |
| [bloxone_dhcp_ha_group.example_anycast](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dhcp_ha_group) | resource |
| [bloxone_dns_auth_zone.authzone](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dns_auth_zone) | resource |
| [bloxone_dns_view.example](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dns_view) | resource |
| [bloxone_dhcp_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dhcp_hosts) | data source |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_anycast_config_name"></a> [anycast\_config\_name](#input\_anycast\_config\_name) | Name of the Anycast configuration. | `string` | n/a | yes |
| <a name="input_anycast_ip_address"></a> [anycast\_ip\_address](#input\_anycast\_ip\_address) | Anycast IP address. | `string` | n/a | yes |
| <a name="input_bgp_configs"></a> [bgp\_configs](#input\_bgp\_configs) | Map of BGP configurations per host. | <pre>map(object({<br>    asn            = string<br>    holddown_secs  = number<br>    neighbors      = list(object({<br>      asn        = string<br>      ip_address = string<br>    }))<br>  }))</pre> | `{}` | no |
| <a name="input_fqdn"></a> [fqdn](#input\_fqdn) | FQDN of the Anycast service. | `string` | n/a | yes |
| <a name="input_ha_name"></a> [ha\_name](#input\_ha\_name) | Name of the Anycast service. | `string` | n/a | yes |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | Map of hostnames or IP addresses for the Anycast configuration. | <pre>map(object({<br>    role             = string,<br>    routing_protocols = list(string)  # This will now handle multiple protocols<br>  }))</pre> | `{}` | no |
| <a name="input_ospf_configs"></a> [ospf\_configs](#input\_ospf\_configs) | Map of OSPF configurations per host. | <pre>map(object({<br>    area_type           = string<br>    area                = string<br>    authentication_type = string<br>    interface           = string<br>    authentication_key  = string<br>    hello_interval      = number<br>    dead_interval       = number<br>    retransmit_interval = number<br>    transmit_delay      = number<br>  }))</pre> | `{}` | no |
| <a name="input_primary_type"></a> [primary\_type](#input\_primary\_type) | Primary type of the Anycast service. | `string` | n/a | yes |
| <a name="input_service"></a> [service](#input\_service) | The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`). | `string` | `"DHCP"` | no |
| <a name="input_view_name"></a> [view\_name](#input\_view\_name) | Name of the Anycast service. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_anycast_config_id"></a> [anycast\_config\_id](#output\_anycast\_config\_id) | The ID of the anycast config |
| <a name="output_anycast_config_name"></a> [anycast\_config\_name](#output\_anycast\_config\_name) | The name of the anycast config |
| <a name="output_anycast_host_configs"></a> [anycast\_host\_configs](#output\_anycast\_host\_configs) | n/a |
| <a name="output_anycast_ip_address"></a> [anycast\_ip\_address](#output\_anycast\_ip\_address) | The anycast IP address |
| <a name="output_service"></a> [service](#output\_service) | The service of the anycast config |
<!-- END_TF_DOCS -->
