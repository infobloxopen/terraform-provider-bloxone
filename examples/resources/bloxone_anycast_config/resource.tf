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

  # Other Optional Fields
  description = "anycast configuration example"
  tags = {
    tag1 = "value1"
  }
}
