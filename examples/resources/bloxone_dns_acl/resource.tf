resource "bloxone_dns_acl" "example_acl" {
  name = "example_dns_acl"

  # Other Optional fields
  comment = "An example acl"
  tags = {
    site = "Site A"
  }
  list = [
    {
      access  = "allow"
      element = "ip"
      address = "192.168.1.1"
    },
  ]
}
