resource "bloxone_td_internal_domain_list" "example" {
  name             = "example_list"
  internal_domains = ["example.somedomain.com", "187.13.5.64/32"]

  # Other optional fields
  description = "Example of an Internal Domain Lists"
  tags = {
    site = "Site A"
  }
}