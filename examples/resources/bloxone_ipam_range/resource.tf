resource "bloxone_ipam_ip_space" "example" {
  name    = "example"
  comment = "Example IP space created by the terraform provider"
  tags = {
    location = "site1"
  }
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
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

  //dhcp options
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "1.1.1.1"
      type         = "option"
    }
  ]
}

resource "bloxone_ipam_range" "example" {
  start = "192.168.1.15"
  end   = "192.168.1.30"
  space = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example"
  comment = "Example Range created by the terraform provider"
  tags = {
    location = "site1"
  }
  exclusion_ranges = [
    {
      start = "192.168.1.17"
      end   = "192.168.1.20"
    }
  ]
  depends_on = [bloxone_ipam_subnet.example]
}