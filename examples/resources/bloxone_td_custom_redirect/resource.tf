resource "bloxone_td_custom_redirect" "example" {
  name = "example_custom_redirect"
  data = "156.2.3.10"

  # Other optional fields
  description = "Example of a Custom Redirect"
  tags = {
    site = "Site A"
  }
}