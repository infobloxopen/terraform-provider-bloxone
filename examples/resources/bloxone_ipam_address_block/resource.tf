resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
  tags = {
    location = "site1"
  }
}

resource "bloxone_ipam_address_block" "example" {
  address = "192.168.1.0"
  cidr    = 24
  name    = "example_address_block"
  space   = bloxone_ipam_ip_space.example.id
}

resource "bloxone_ipam_address_block" "example_tags" {
  address = "10.0.0.0"
  cidr    = 8
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_address_block_tags"
  comment = "Example address block with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
}
