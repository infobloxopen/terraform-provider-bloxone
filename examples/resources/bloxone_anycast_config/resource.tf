data "bloxone_infra_services" "anycast_services" {
  filters = {
    service_type = "anycast"
  }
}

data "bloxone_infra_hosts" "anycast_hosts" {
  filters = {
    pool_id = data.bloxone_infra_services.anycast_services.results.0.pool_id
  }
}

#Create anycast configuration with necessary fields
resource "bloxone_anycast_ac_config" "example" {

  name               = "anycast_config_test"
  service            = "DNS"
  anycast_ip_address = "192.2.2.1"

  tags = {
    tag1 = "value1"
  }

  onprem_hosts = [
    {
      id   = data.bloxone_infra_hosts.anycast_hosts.results.0.legacy_id
      name = data.bloxone_infra_hosts.anycast_hosts.results.0.display_name
    }
  ]
}
