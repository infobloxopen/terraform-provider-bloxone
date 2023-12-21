resource "bloxone_ipam_ip_space" "example" {
  name    = "example"
  comment = "Example IP space created by the terraform provider"
  tags = {
    location = "site1"
  }
}

resource "bloxone_ipam_address_block" "example" {
  address = "192.160.0.0"
  cidr    = 16
  space   = bloxone_ipam_ip_space.example.id
  comment = "Example Address Block created by the terraform provider"
  tags = {
    location = "site1"
  }
}

resource "bloxone_ipam_subnet" "example" {
  name    = "example"
  space   = bloxone_ipam_ip_space.example.id
  address = "192.168.1.0"
  cidr    = 24
  comment = "Example Subnet created by the terraform provider"
  tags = {
    location = "site1"
  }
}

resource "bloxone_dhcp_fixed_address" "example_fixed_address" {
  name        = "example_fixed_address"
  address     = "192.168.1.1"
  ip_space    = bloxone_ipam_ip_space.example.id
  match_type  = "mac"
  match_value = "00:00:00:00:00:00"
  comment     = "Example Fixed Address created by the terraform provider"
  tags = {
    location = "site1"
  }
  depends_on = [bloxone_ipam_subnet.example]
}
