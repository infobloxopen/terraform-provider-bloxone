# Get "AAAA" records filtered by an attribute
data "bloxone_dns_aaaa_records" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

# Get "AAAA" records filtered by tag
data "bloxone_dns_aaaa_records" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all "AAAA" records
data "bloxone_dns_aaaa_records" "example_all" {}
