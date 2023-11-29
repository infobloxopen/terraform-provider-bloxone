# Get delegation zone filtered by an attribute
data "bloxone_dns_views" "example_by_attribute" {
  filters = {
    "name" = "example_view"
  }
}

data "bloxone_dns_delegations" "example_by_attribute" {
  filters = {
    fqdn = "tf-acc-test.com."
    view = data.bloxone_dns_views.example_by_attribute.results.0.id
  }
}

# Get delegation zones filtered by tag
data "bloxone_dns_delegations" "example_by_tag" {
  tag_filters = {
    region = "eu"
  }
}

# Get all delegation zones
data "bloxone_dns_delegations" "example_all" {}
