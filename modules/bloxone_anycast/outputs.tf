output "anycast_config" {
  description = "The anycast config"
  value       = bloxone_anycast_config.ac
}

output "anycast_hosts" {
  description = "Map of anycast hosts"
  value       = bloxone_anycast_host.this
}

output "infra_services" {
  description = "Map of infrastructure services"
  value       =  bloxone_infra_service.anycast
}

output "dhcp_ha_group" {
  description = "The DHCP HA group"
  value       = bloxone_dhcp_ha_group.this
}