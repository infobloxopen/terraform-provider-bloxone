resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

resource "bloxone_ipam_ip_space" "example_tags" {
  name    = "example_ip_space_tags"
  comment = "Example IP space with tags created by the terraform provider"
  tags = {
    location = "site1"
  }

  //dhcp options
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "10.0.0.1"
      type         = "option"
    }
  ]
}
