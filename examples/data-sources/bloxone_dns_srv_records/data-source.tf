# Get "SRV" records filtered by an attribute
data "bloxone_dns_srv_records" "example_by_attribute" {
  filters = {
    "absolute_name_spec" = "abc.example.com"
  }
}

# Get "SRV" records filtered by tag
data "bloxone_dns_srv_records" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all "SRV" records
data "bloxone_dns_srv_records" "example_all" {}
