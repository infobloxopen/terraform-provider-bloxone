output "anycast_config_id" {
  description = "The ID of the anycast config"
  value       = bloxone_anycast_config.ac.id
}

output "anycast_config_name" {
  description = "The name of the anycast config"
  value       = bloxone_anycast_config.ac.name
}

output "anycast_ip_address" {
  description = "The anycast IP address"
  value       = bloxone_anycast_config.ac.anycast_ip_address
}

output "service" {
  description = "The service of the anycast config"
  value       = bloxone_anycast_config.ac.service
}

output "anycast_host_configs" {
  value = {
    for host_key, host in data.bloxone_infra_hosts.this : host_key => {
      id             = host.results[0].legacy_id,
    }
  }
}

output "anycast_host_ids" {
  value = { for k, v in bloxone_anycast_host.this : k => v.id }
  description = "The IDs of the created anycast hosts."
}

output "dhcp_ha_group_id" {
  value       = length(bloxone_dhcp_ha_group.this) > 0 ? bloxone_dhcp_ha_group.this[0].id : null
  description = "The ID of the created DHCP HA group."
}

output "infra_service_ids" {
  value = { for k, v in bloxone_infra_service.anycast : k => v.id }
  description = "The IDs of the created infrastructure services."
}
