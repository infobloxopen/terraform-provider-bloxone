<!-- BEGIN_TF_DOCS -->
# Terraform Module to create BloxOne Host in Azure

This module will provision an Azure virtual machine that uses a BloxOne image.
The virtual machine will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
If a join token is not provided, a new one will be created.

The BloxOne Host created in the CSP is created automatically, and cannot be managed through terraform.
A `bloxone_infra_hosts` data source is provided to retrieve the host information from the CSP.
The data source will use the `tags` variable to filter the hosts.
A `tf_module_host_id` tag will be added to the tags variable so that the data source can uniquely find the host.

This module will also create a BloxOne Infra Service for each service type provided in the `services` variable.
The service will be named `<service_type>_<host_display_name>`.

## Example Usage

```hcl
module "bloxone_infra_host_azure" {
  source = "github.com/infobloxopen/terraform-provider-bloxone/modules/bloxone_infra_host_azure"

  vm_name                   = "bloxone-vm"
  location                  = "eastus"
  resource_group_name       = "my-resource-group"
  subnet_id                 = "subnet-id"
  vnet_id                   = "vnet-id"
  vm_network_security_group = "nsg-id"

  azure_instance_tags       = {
    environment = "dev"
  }

  tags                      = {
  location = "office1"
  }

  services = {
    dhcp = "start"
    dns = "start"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | n/a |
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | n/a |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Resources

| Name | Type |
|------|------|
| [azurerm_network_interface.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface) | resource |
| [azurerm_virtual_machine.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_machine) | resource |
| [bloxone_infra_join_token.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_join_token) | resource |
| [bloxone_infra_service.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service) | resource |
| [random_uuid.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/uuid) | resource |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_azure_instance_tags"></a> [azure\_instance\_tags](#input\_azure\_instance\_tags) | The tags to use for the Azure virtual machine. | `map(string)` | `{}` | no |
| <a name="input_join_token"></a> [join\_token](#input\_join\_token) | The join token to use for the BloxOne Host. If not provided, a join token will be created. | <pre>object({<br>    join_token = string<br>  })</pre> | `null` | no |
| <a name="input_location"></a> [location](#input\_location) | The location where the resources will be created | `string` | `"eastus"` | no |
| <a name="input_managed_disk_type"></a> [managed\_disk\_type](#input\_managed\_disk\_type) | Type of managed disk for the VMs that will be part of this compute group. Allowable values are 'Standard\_LRS' or 'Premium\_LRS'. | `string` | `"Standard_LRS"` | no |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | The name of the resource group in which the resources will be created | `string` | n/a | yes |
| <a name="input_services"></a> [services](#input\_services) | The services to provision on the BloxOne Host. The services must be a map of valid service type with values of "start" or "stop". Valid service types are "dhcp" and "dns". | `map(string)` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | The tags to use for the BloxOne Host. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br>    create = string<br>    update = string<br>    read   = string<br>  })</pre> | `null` | no |
| <a name="input_vm_name"></a> [vm\_name](#input\_vm\_name) | The name of the virtual machine | `string` | n/a | yes |
| <a name="input_vm_network_security_group_name"></a> [vm\_network\_security\_group\_name](#input\_vm\_network\_security\_group\_name) | The name of the network security group that will be created and associated to the BloxOne Host | `string` | n/a | yes |
| <a name="input_vm_os_offer"></a> [vm\_os\_offer](#input\_vm\_os\_offer) | The name of the offer of the image that you want to deploy | `string` | `"infoblox-bloxone-34"` | no |
| <a name="input_vm_os_version"></a> [vm\_os\_version](#input\_vm\_os\_version) | The version of the image that you want to deploy. | `string` | `"latest"` | no |
| <a name="input_vm_size"></a> [vm\_size](#input\_vm\_size) | Size of the Virtual Machine based on Azure sizing | `string` | `"Standard_F4s_v2"` | no |
| <a name="input_vnet_subnet_id"></a> [vnet\_subnet\_id](#input\_vnet\_subnet\_id) | The subnet id of the virtual network on which the BloxOne Host will be connected | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_azurerm_network_interface"></a> [azurerm\_network\_interface](#output\_azurerm\_network\_interface) | The Azure network interface object for the instance |
| <a name="output_azurerm_virtual_machine"></a> [azurerm\_virtual\_machine](#output\_azurerm\_virtual\_machine) | The Azure virtual machine object for the instance |
| <a name="output_host"></a> [host](#output\_host) | The `bloxone_infra_host` object for the instance |
| <a name="output_services"></a> [services](#output\_services) | The `bloxone_infra_service` objects for the instance. May be empty if no services were specified. |
<!-- END_TF_DOCS -->