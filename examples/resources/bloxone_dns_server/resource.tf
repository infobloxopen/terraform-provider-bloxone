
resource "bloxone_dns_server" "example_server" {
  name = "example_dns_server"

  # Other Optional fields
  comment = "An example server"
  tags = {
    site = "Test Site"
  }
}
