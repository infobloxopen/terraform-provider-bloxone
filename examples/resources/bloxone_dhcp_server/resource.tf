resource "bloxone_dhcp_server" "example" {
  name = "example_dhcp_server"
}

resource "bloxone_dhcp_server" "example_with_options" {
  name = "example_dhcp_server_with_options"

  #Other Optional Fields
  comment = "dhcp server"
  tags = {
    site = "Site A"
  }
}
