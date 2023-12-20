# Get records filtered by an attribute
data "bloxone_dns_records" "example_by_attribute" {
  type = "TYPE256"
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

# Get records filtered by tag
data "bloxone_dns_records" "example_by_tag" {
  type = "TYPE256"
  tag_filters = {
    "region" = "eu"
  }
}

# Get all records
data "bloxone_dns_records" "example_all" {}
