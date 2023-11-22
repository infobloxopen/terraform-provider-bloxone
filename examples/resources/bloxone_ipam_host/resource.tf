resource "bloxone_ipam_host" "example" {
  name = "example_ipam_host"

  #Other Optional Fields
  comment = "IPAM Host"
  tags = {
    site = "Test Site"
  }
}
