
resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "tf-acc-test.com."
  primary_type = "cloud"

  # Other optional fields
  name    = "example_auth_zone"
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Test Site"
  }
}
