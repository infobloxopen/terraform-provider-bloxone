
resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_dns_view" "example" {
  name = "example_dns_view"

  # Other Optional fields
  comment   = "An example view"
  ip_spaces = [bloxone_ipam_ip_space.example.id]
  tags = {
    site = "Site A"
  }
  match_clients_acl = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.10.10"
    }
  ]
  lame_ttl = 80

}
