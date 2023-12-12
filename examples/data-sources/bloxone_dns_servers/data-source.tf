# Get DNS Servers filtered by an attribute
data "bloxone_dns_servers" "example_by_attribute" {
  filters = {
    "name" = "example_server"
  }
}

# Get DNS Servers filtered by tag
data "bloxone_dns_servers" "example_by_tag" {
  tag_filters = {
    site = "Test Site"
  }
}

# Get all DNS Servers
data "bloxone_dns_servers" "example_all" {
}
