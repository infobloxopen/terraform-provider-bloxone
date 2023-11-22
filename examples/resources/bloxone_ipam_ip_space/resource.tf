resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_ipam_ip_space" "example_tags" {
  name    = "example_ip_space_tags"
  comment = "Example IP space with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
}