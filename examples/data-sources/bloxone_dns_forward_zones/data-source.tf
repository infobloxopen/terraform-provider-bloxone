# Get forward zones filtered by an attribute
data "bloxone_dns_forward_zones" "example_by_attribute" {
  filters = {
    fqdn = "tf-acc-test.com."
  }
}

# Get forward zones filtered by tag
data "bloxone_dns_forward_zones" "example_by_tag" {
  tag_filters = {
    region = "eu"
  }
}

# Get all forward zones
data "bloxone_dns_forward_zones" "example_all" {}
