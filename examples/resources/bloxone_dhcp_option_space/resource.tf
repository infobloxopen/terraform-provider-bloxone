resource "bloxone_dhcp_option_space" "example" {
  name     = "example_dhcp_option_space"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_space" "example_with_options" {
  name     = "example_dhcp_option_space_with_options"
  protocol = "ip6"
  #Other Optional Fields
  comment = "dhcp option space"
  tags = {
    site = "Test Site"
  }
}