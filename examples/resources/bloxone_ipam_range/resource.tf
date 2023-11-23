resource "bloxone_ip_space" "example_space" {
  name = "example_space"
  comment = "Example IP space created by the terraform provider"
  tags = {
    location   = "site1"
  }
}

resource "bloxone_subnet" "example_subnet" {
  name = "example_subnet"
  space = bloxone_ip_space.example_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet created by the terraform provider"
  tags = {
    location   = "site1"
  }
}

resource "bloxone_range" "example_range" {
  start = "192.168.1.15"
  end = "192.168.1.30"
  space = bloxone_ip_space.example_space.id

  # Other optional fields
  name = "example_tf_range"
  comment = "Example Range created by the terraform provider"
  tags = {
    location   = "site1"
  }
  depends_on = [bloxone_subnet.example_subnet]
}
