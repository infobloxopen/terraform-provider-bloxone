output "anycast_config" {
  description = "The anycast config"
  value       = bloxone_anycast_config.ac
}

output "anycast_hosts" {
  description = "Map of anycast hosts"
  value       = { for k, v in bloxone_anycast_host.this : k => v }
}

output "infra_services" {
  description = "Map of infrastructure services"
  value       = { for k, v in bloxone_infra_service.anycast : k => v }
}

output "dhcp_ha_group" {
  description = "The DHCP HA group"
  value       = bloxone_dhcp_ha_group.this
}