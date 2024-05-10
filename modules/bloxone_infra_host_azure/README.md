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
  source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_azure"

  vm_name                   = "bloxone-vm"
  location                  = "eastus"
  resource_group_name       = "my-resource-group"
  subnet_id                 = "subnet-id"
  vnet_id                   = "vnet-id"
  vm_network_interface_ids  = ["nic-id"]

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
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 3.0.0 |
| <a name="requirement_bloxone"></a> [bloxone](#requirement\_bloxone) | >= 1.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | >= 3.0.0 |
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | >= 1.0.0 |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Resources

| Name | Type |
|------|------|
| [azurerm_linux_virtual_machine.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_virtual_machine) | resource |
| [bloxone_infra_join_token.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_join_token) | resource |
| [bloxone_infra_service.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service) | resource |
| [random_uuid.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/uuid) | resource |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_admin_username"></a> [admin\_username](#input\_admin\_username) | The username to use for the BloxOne Host. | `string` | `"infobloxadmin"` | no |
| <a name="input_azure_instance_tags"></a> [azure\_instance\_tags](#input\_azure\_instance\_tags) | The tags to use for the Azure virtual machine. | `map(string)` | `{}` | no |
| <a name="input_join_token"></a> [join\_token](#input\_join\_token) | The join token to use for the BloxOne Host. If not provided, a join token will be created. | `string` | `null` | no |
| <a name="input_location"></a> [location](#input\_location) | The location where the resources will be created | `string` | n/a | yes |
| <a name="input_os_disk_caching"></a> [os\_disk\_caching](#input\_os\_disk\_caching) | The caching type to use for the OS disk. | `string` | `"ReadWrite"` | no |
| <a name="input_os_disk_storage_account_type"></a> [os\_disk\_storage\_account\_type](#input\_os\_disk\_storage\_account\_type) | The storage account type to use for the OS disk. | `string` | `"Standard_LRS"` | no |
| <a name="input_plan_name"></a> [plan\_name](#input\_plan\_name) | The name of the plan to use for the BloxOne Host. | `string` | `"infoblox-bloxone"` | no |
| <a name="input_plan_product"></a> [plan\_product](#input\_plan\_product) | The product to use for the BloxOne Host. | `string` | `"infoblox-bloxone-34"` | no |
| <a name="input_plan_publisher"></a> [plan\_publisher](#input\_plan\_publisher) | The publisher to use for the BloxOne Host. | `string` | `"infoblox"` | no |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | The name of the resource group in which the resources will be created | `string` | n/a | yes |
| <a name="input_services"></a> [services](#input\_services) | The services to provision on the BloxOne Host. The services must be a map of valid service type with values of "start" or "stop". Valid service types are "dhcp", "dns", "anycast", "dfp". | `map(string)` | n/a | yes |
| <a name="input_source_image_reference_offer"></a> [source\_image\_reference\_offer](#input\_source\_image\_reference\_offer) | The offer of the image that you want to deploy | `string` | `"infoblox-bloxone-34"` | no |
| <a name="input_source_image_reference_publisher"></a> [source\_image\_reference\_publisher](#input\_source\_image\_reference\_publisher) | The publisher of the image that you want to deploy | `string` | `"infoblox"` | no |
| <a name="input_source_image_reference_sku"></a> [source\_image\_reference\_sku](#input\_source\_image\_reference\_sku) | The sku of the image that you want to deploy | `string` | `"infoblox-bloxone"` | no |
| <a name="input_source_image_reference_version"></a> [source\_image\_reference\_version](#input\_source\_image\_reference\_version) | The version of the image that you want to deploy. | `string` | `"3.8.1"` | no |
| <a name="input_ssh_public_key_path"></a> [ssh\_public\_key\_path](#input\_ssh\_public\_key\_path) | The path to the SSH public key to use for the BloxOne Host. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | The tags to use for the BloxOne Host. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br>    create = string<br>    update = string<br>    read   = string<br>  })</pre> | `null` | no |
| <a name="input_vm_name"></a> [vm\_name](#input\_vm\_name) | The name of the virtual machine | `string` | n/a | yes |
| <a name="input_vm_network_interface_ids"></a> [vm\_network\_interface\_ids](#input\_vm\_network\_interface\_ids) | The network interface ids that will be associated to the BloxOne Host | `list(string)` | n/a | yes |
| <a name="input_vm_size"></a> [vm\_size](#input\_vm\_size) | Size of the Virtual Machine based on Azure sizing | `string` | `"Standard_F8s"` | no |
| <a name="input_wait_for_state"></a> [wait\_for\_state](#input\_wait\_for\_state) | If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API. | `bool` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_azurerm_linux_virtual_machine"></a> [azurerm\_linux\_virtual\_machine](#output\_azurerm\_linux\_virtual\_machine) | The Azure virtual machine object for the instance |
| <a name="output_host"></a> [host](#output\_host) | The `bloxone_infra_host` object for the instance |
| <a name="output_services"></a> [services](#output\_services) | The `bloxone_infra_service` objects for the instance. May be empty if no services were specified. |
<!-- END_TF_DOCS -->