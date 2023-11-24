
resource "bloxone_ipam_ip_space" "example_ip_space" {
  name = "example_ip_space"
}

resource "bloxone_dns_view" "example_dns_view" {
  name = "example_dns_view"

  # Other Optional fields
  comment  = "An example view"
  ip_space = bloxone_ipam_ip_space.example_ip_space.id
  tags = {
    site = "Test Site"
  }
  match_clients_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.10.10"
    }
  ]
  default_ttl = 800

}