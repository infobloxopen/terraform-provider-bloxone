resource "bloxone_ipam_ip_space" "example" {
  name = "example"
}

resource "bloxone_ipam_subnet" "example" {
  address = "10.0.0.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.example.id
}

// Passing a static address for the IPAM Host
resource "bloxone_ipam_host" "example" {
  name = "example_ipam_host"
  addresses = [
    {
      address = "10.0.0.1"
      space   = bloxone_ipam_ip_space.example.id
    }
  ]
  #Other Optional Fields
  comment = "IPAM Host"
  tags = {
    site = "Site A"
  }
}

// Dynamically getting the IPAM Host address using Next Available IP
resource "bloxone_ipam_host" "example_naip" {
  name = "example_ipam_host_naip"
  addresses = [
    {
      next_available_id = bloxone_ipam_subnet.example.id
    }
  ]
  #Other Optional Fields
  comment = "IPAM Host"
  tags = {
    site = "Site A"
  }
}
