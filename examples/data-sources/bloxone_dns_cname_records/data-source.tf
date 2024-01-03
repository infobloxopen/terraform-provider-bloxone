# Get "CNAME" records filtered by an attribute
data "bloxone_dns_cname_records" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

# Get "CNAME" records filtered by tag
data "bloxone_dns_cname_records" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all "CNAME" records
data "bloxone_dns_cname_records" "example_all" {}
