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
  description    = "The anycast host configurations"
  value = {
    anycast_config = bloxone_anycast_config.ac.name,
    bgp_asn        = var.bgp_config.asn,
    ospf_area      = var.ospf_config.area
  }
}
