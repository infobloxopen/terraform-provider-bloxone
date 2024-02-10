# Get DNS ACLs filtered by an attribute
data "bloxone_dns_acls" "example_by_attribute" {
  filters = {
    "name" = "example_acl"
  }
}

# Get DNS ACLs filtered by tag
data "bloxone_dns_acls" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all DNS ACLs
data "bloxone_dns_acls" "example_all" {
}
