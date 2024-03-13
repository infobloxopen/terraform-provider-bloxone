resource "bloxone_ipam_ip_space" "current" {
  name = "example-space"
}

resource "bloxone_infra_host" "example" {
  display_name = "example_host"

  # Other Optional fields
  description = "An example host"
  ip_space    = bloxone_ipam_ip_space.current.id
  tags = {
    site = "Site A"
  }
}
