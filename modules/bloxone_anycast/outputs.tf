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
      bgp_asn        = contains(var.hosts[host_key].routing_protocols , "BGP") ? var.bgp_configs[host_key].asn : null,
      ospf_area      = contains(var.hosts[host_key].routing_protocols, "OSPF") ? var.ospf_configs[host_key].area : null
    }
  }
}
