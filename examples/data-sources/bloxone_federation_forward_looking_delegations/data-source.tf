# Get Forward Looking Delegation filtered by an attribute
data "bloxone_federation_forward_looking_delegations" "example_by_attribute" {
  filters = {
    name = "example_forward_looking_delegation"
  }
}

# Get Forward Looking Delegation filtered by tag
data "bloxone_federation_forward_looking_delegations" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Forward Looking Delegations
data "bloxone_federation_forward_looking_delegations" "example_all" {}
