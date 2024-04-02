resource "bloxone_dhcp_server" "example" {
  name = "example_dhcp_server"
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

resource "bloxone_dhcp_server" "example_with_options" {
  name = "example_dhcp_server_with_options"

  #Other Optional Fields
  comment = "dhcp server"
  tags = {
    site = "Site A"
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
