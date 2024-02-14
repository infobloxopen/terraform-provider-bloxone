
data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

resource "bloxone_dns_server" "example_server" {
  name = "example_dns_server"

  # Other Optional fields
  comment = "An example server"
  tags = {
    site = "Site A"
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
