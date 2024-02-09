# Get "HTTPS" records filtered by an attribute
data "bloxone_dns_https_records" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

# Get "HTTPS" records filtered by tag
data "bloxone_dns_https_records" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all "HTTPS" records
data "bloxone_dns_https_records" "example_all" {}
