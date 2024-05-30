<!-- BEGIN_TF_DOCS -->
# Terraform Module to create BloxOne Anycast in AWS

This module will retrieve an AWS EC2 instance that uses a BloxOne AMI.
The instance will be configured to setup anycast config profile and add anycast protocols like BGP, OSPF.
Then a DHCP HA group will be created with the provided hosts with anycast as HA configuration.

This module will fetch the already created BloxOne Host and create an anycast config profile with the desired routing protocols.

## Example Usage

```hcl
module "bloxone_anycast" {
 source = "/Users/agadiyarhj/go/src/github.com/infobloxopen/terraform-provider-bloxone/modules/bloxone_anycast"

 hosts = {
   host1 = "active",
   host2 = "passive"
 }
 name      = "ac"
 service   = "DHCP"
 anycast_ip_address = "192.2.2.1"
 routing_protocols = ["BGP", "OSPF"]

 bgp_config = {
   asn           = "6500"
   holddown_secs = 180
   neighbors     = [
     {
       asn        = "6501"
       ip_address = "172.28.4.198"
     }
   ]
 }

 ospf_config = {
   area                = "10.10.0.1"
   area_type           = "STANDARD"
   authentication_type = "Clear"
   interface           = "eth0"
   authentication_key  = "YXV0aGV"
   hello_interval      = 10
   dead_interval       = 40
   retransmit_interval = 5
   transmit_delay      = 1
 }
*}
```

## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | 0.0.1 |

## Resources

| Name | Type |
|------|------|
| [bloxone_anycast_config.ac](https://registry.terraform.io/providers/hashicorp/bloxone/latest/docs/resources/anycast_config) | resource |
| [bloxone_anycast_host.this](https://registry.terraform.io/providers/hashicorp/bloxone/latest/docs/resources/anycast_host) | resource |
| [bloxone_dhcp_ha_group.example_anycast](https://registry.terraform.io/providers/hashicorp/bloxone/latest/docs/resources/dhcp_ha_group) | resource |
| [bloxone_dhcp_hosts.this](https://registry.terraform.io/providers/hashicorp/bloxone/latest/docs/data-sources/dhcp_hosts) | data source |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/hashicorp/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_anycast_config_name"></a> [anycast\_config\_name](#input\_anycast\_config\_name) | Name of the Anycast configuration. | `string` | `"anycast-config-1"` | no |
| <a name="input_anycast_ip_address"></a> [anycast\_ip\_address](#input\_anycast\_ip\_address) | Anycast IP address. | `string` | `"10.10.10.5"` | no |
| <a name="input_bgp_config"></a> [bgp\_config](#input\_bgp\_config) | BGP configuration | <pre>object({<br>    asn            = string<br>    holddown_secs  = number<br>    neighbors      = list(object({<br>      asn        = string<br>      ip_address = string<br>    }))<br>  })</pre> | <pre>{<br>  "asn": "6500",<br>  "holddown_secs": 180,<br>  "neighbors": [<br>    {<br>      "asn": "6501",<br>      "ip_address": "172.28.4.198"<br>    }<br>  ]<br>}</pre> | no |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | Map of hostnames or IP addresses for the Anycast configuration. | `map(string)` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name of the Anycast service. | `string` | `"anycast-service"` | no |
| <a name="input_ospf_config"></a> [ospf\_config](#input\_ospf\_config) | OSPF configuration. | <pre>object({<br>    area                = string<br>    area_type           = string<br>    authentication_type = string<br>    interface           = string<br>    authentication_key  = string<br>    hello_interval      = number<br>    dead_interval       = number<br>    retransmit_interval = number<br>    transmit_delay      = number<br>  })</pre> | <pre>{<br>  "area": "10.10.0.1",<br>  "area_type": "STANDARD",<br>  "authentication_key": "YXV0aGV",<br>  "authentication_type": "Clear",<br>  "dead_interval": 40,<br>  "hello_interval": 10,<br>  "interface": "eth0",<br>  "retransmit_interval": 5,<br>  "transmit_delay": 1<br>}</pre> | no |
| <a name="input_routing_protocols"></a> [routing\_protocols](#input\_routing\_protocols) | List of routing protocols to be configured (e.g., BGP, OSPF). | `list(string)` | <pre>[<br>  "BGP",<br>  "OSPF"<br>]</pre> | no |
| <a name="input_service"></a> [service](#input\_service) | The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`). | `string` | `"DHCP"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_anycast_config_id"></a> [anycast\_config\_id](#output\_anycast\_config\_id) | The ID of the anycast config |
| <a name="output_anycast_config_name"></a> [anycast\_config\_name](#output\_anycast\_config\_name) | The name of the anycast config |
| <a name="output_anycast_host_configs"></a> [anycast\_host\_configs](#output\_anycast\_host\_configs) | The anycast host configurations |
| <a name="output_anycast_ip_address"></a> [anycast\_ip\_address](#output\_anycast\_ip\_address) | The anycast IP address |
| <a name="output_service"></a> [service](#output\_service) | The service of the anycast config |
<!-- END_TF_DOCS -->
