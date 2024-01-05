resource "bloxone_dhcp_option_space" "option_space" {
  name     = "option_space"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_code" "option_code" {
  code         = 250
  name         = "example_option_code"
  option_space = bloxone_dhcp_option_space.option_space.id
  type         = "int32"
}

resource "bloxone_dhcp_option_code" "option_code_with_options" {
  code         = 251
  name         = "example_option_code_with_options"
  option_space = bloxone_dhcp_option_space.option_space.id
  type         = "int32"

  # Other optional fields
  array   = true
  comment = "Option code example"
}
