output "services" {
  value       = bloxone_infra_service.this
  description = "The `bloxone_infra_service` objects for the instance. May be empty if no services were specified."
}

output "host" {
  value       = data.bloxone_infra_hosts.this.results.0
  description = "The `bloxone_infra_host` object for the instance"
}

output "azurerm_linux_virtual_machine" {
  value       = azurerm_linux_virtual_machine.this
  description = "The Azure virtual machine object for the instance"
}

