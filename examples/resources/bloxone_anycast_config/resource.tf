data "bloxone_infra_hosts" "anycast_host" {
  filters = {
    display_name = "my_host"
  }
}

#Create anycast configuration with necessary fields
resource "bloxone_anycast_config" "example" {

  name               = "anycast_example"
  service            = "DNS"
  anycast_ip_address = "192.2.2.1"

  tags = {
    tag1 = "value1"
  }

  onprem_hosts = [
    {
      id   = data.bloxone_infra_hosts.anycast_host.results.0.legacy_id
      name = data.bloxone_infra_hosts.anycast_host.results.0.display_name
    }
  ]
}
