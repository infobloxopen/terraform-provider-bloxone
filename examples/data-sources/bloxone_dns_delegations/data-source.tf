# Get delegation zone filtered by an attribute

data "bloxone_dns_delegations" "example_by_attribute" {
  filters = {
    fqdn = "tf-acc-Site A.com."
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
