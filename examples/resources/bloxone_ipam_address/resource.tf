resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_ipam_subnet" "example" {
  address = "10.0.0.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.example.id
}

resource "bloxone_ipam_address" "example" {
  address = "10.0.0.1"
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  comment = "reservation for test site"
  names = [
    {
      name = "bby-1"
      type = "user"
    }
  ]
  tags = {
    site = "Test Site"
  }

  depends_on = [bloxone_ipam_subnet.example]
}

# Next available address in subnet
resource "bloxone_ipam_address" "example_nas" {
  next_available_id = bloxone_ipam_subnet.example.id
  space             = bloxone_ipam_ip_space.example.id
}
