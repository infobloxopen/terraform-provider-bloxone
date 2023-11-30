resource "bloxone_dns_forward_nsg" "example" {
  name = "example_dns_forward_nsg"

  # Other Optional fields
  comment = "An example forward nsg"
  external_forwarders = [
    {
      address = "12.10.2.1",
      fqdn    = "ext.primary.com."
    },
  ]
  tags = {
    site = "Test Site"
  }
}
