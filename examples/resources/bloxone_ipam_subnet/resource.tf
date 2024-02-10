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
  comment = "Subnet for Site A"
  tags = {
    site = "Site A"
  }
}

resource "bloxone_dhcp_option_code" "option_code" {
  code         = 250
  name         = "example_option_code"
  option_space = bloxone_dhcp_option_space.option_space.id
  type         = "int32"
}

# Next available subnet
resource "bloxone_ipam_subnet" "example_na_s" {
  next_available_id = bloxone_ipam_address_block.example.id
  cidr              = 24
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_subnet"
  comment = "Subnet for Site A"
  tags = {
    site = "Site A"
  }
  #dhcp options
  dhcp_options = [
    {
      option_code  = bloxone_dhcp_option_code.option_code.id
      option_value = "true"
      type         = "option"
    }
  ]
}
