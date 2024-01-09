resource "bloxone_dhcp_option_group" "example" {
  name     = "example_dhcp_option_group"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_space" "option_space" {
  name     = "option_space"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_code" "option_code" {
  code         = 234
  name         = "option_code"
  option_space = bloxone_dhcp_option_space.option_space.id
  type         = "boolean"
}

resource "bloxone_dhcp_option_group" "example_with_options" {
  name     = "example_dhcp_option_group_with_options"
  protocol = "ip6"

  # Other Optional Fields
  dhcp_options = [
    {
      type         = "option"
      option_code  = bloxone_dhcp_option_code.option_code.id
      option_value = "true"
    }
  ]
  comment = "dhcp option group"
  tags = {
    site = "Test Site"
  }
}
