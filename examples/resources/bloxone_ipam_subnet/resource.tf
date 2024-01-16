resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_ipam_address_block" "example" {
  address = "10.0.0.0"
  cidr    = 16
  space   = bloxone_ipam_ip_space.example.id
}

# Static address
resource "bloxone_ipam_subnet" "example" {
  address = "10.0.0.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_subnet"
  comment = "Subnet for test site"
  tags = {
    site = "Test Site"
  }
}

# Next available subnet
resource "bloxone_ipam_subnet" "example_nas" {
  next_available_id = bloxone_ipam_address_block.example.id
  cidr              = 24
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_subnet"
  comment = "Subnet for test site"
  tags = {
    site = "Test Site"
  }
}
