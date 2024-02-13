/**
 * # Terraform Module to create BloxOne Host in Azure
 *
 * This module will provision an Azure virtual machine that uses a BloxOne image.
 * The virtual machine will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
 * If a join token is not provided, a new one will be created.
 *
 * The BloxOne Host created in the CSP is created automatically, and cannot be managed through terraform.
 * A `bloxone_infra_hosts` data source is provided to retrieve the host information from the CSP.
 * The data source will use the `tags` variable to filter the hosts.
 * A `tf_module_host_id` tag will be added to the tags variable so that the data source can uniquely find the host.
 *
 * This module will also create a BloxOne Infra Service for each service type provided in the `services` variable.
 * The service will be named `<service_type>_<host_display_name>`.
 *
 * ## Example Usage
 *
 * ```hcl
 * module "bloxone_infra_host_azure" {
 *   source = "github.com/infobloxopen/terraform-provider-bloxone/modules/bloxone_infra_host_azure"
 *
 *   vm_name                   = "bloxone-vm"
 *   location                  = "eastus"
 *   resource_group_name       = "my-resource-group"
 *   subnet_id                 = "subnet-id"
 *   vnet_id                   = "vnet-id"
 *   vm_network_security_group = "nsg-id"
 *
 *   azure_instance_tags       = {
 *     environment = "dev"
 *   }
 *
 *   tags                      = {
 *     location = "office1"
 *   }
 *
 *   services = {
 *     dhcp = "start"
 *     dns = "start"
 *   }
 * }
 * ```
 * 
 */

resource "random_uuid" "this" {}

locals {
  join_token = var.join_token == null ? bloxone_infra_join_token.this[0].join_token : var.join_token.join_token
  tags       = merge(
    var.tags,
    {
      "tf_module_host_id" = "bloxone-host-${random_uuid.this.result}"
    }
  )
}

resource "bloxone_infra_join_token" "this" {
  count = var.join_token == null ? 1 : 0
  name  = "jt-${random_uuid.this.result}"
}

resource "azurerm_network_interface" "this" {
  name                = "${var.vm_name}-nic"
  resource_group_name = var.resource_group_name
  location            = var.location
  tags                = var.azure_instance_tags

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.vnet_subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "this" {

  depends_on = [azurerm_network_interface.this]

  name                  = var.vm_name
  location              = var.location
  resource_group_name   = var.resource_group_name
  network_interface_ids = [azurerm_network_interface.this.id]
  vm_size               = var.vm_size

  storage_image_reference {
    publisher = "infoblox"
    offer     = var.vm_os_offer
    sku       = "infoblox-bloxone"
    version   = var.vm_os_version
  }

  delete_os_disk_on_termination    = true
  delete_data_disks_on_termination = true
  tags                             = var.azure_instance_tags

  os_profile_linux_config {
    disable_password_authentication = false
  }

  os_profile {
    computer_name  = var.vm_name
    admin_username = "dummy"
    admin_password = "Dummy@123"
    custom_data    = templatefile("${path.module}/userdata.tftpl", {
      join_token = local.join_token
      tags : local.tags
    })
  }

  storage_os_disk {
    name              = "${var.vm_name}-osdisk"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  plan {
    name      = "infoblox-bloxone"
    product   = "infoblox-bloxone-34"
    publisher = "infoblox"
  }
}

data "bloxone_infra_hosts" "this" {
  filters = {
    host_type = "3" // For BloxOne VM
  }
  tag_filters        = local.tags
  retry_if_not_found = true
  timeouts           = var.timeouts

  lifecycle {
    postcondition {
      condition     = self.results == null ? false : length(self.results) == 1
      error_message = "BloxOne Host not found in CSP."
    }
  }

  depends_on = [
    azurerm_virtual_machine.this
  ]
}

resource "bloxone_infra_service" "this" {
  for_each      = var.services
  name          = format("%s_%s", each.key, data.bloxone_infra_hosts.this.results[0].display_name)
  pool_id       = data.bloxone_infra_hosts.this.results[0].pool_id
  service_type  = each.key
  desired_state = each.value
  tags          = local.tags
  timeouts      = var.timeouts
}
