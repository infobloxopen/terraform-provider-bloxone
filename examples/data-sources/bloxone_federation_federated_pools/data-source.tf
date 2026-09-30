# Get Federated Pool filtered by an attribute
data "bloxone_federation_federated_pools" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_pool"
  }
}

# Get Federated Pool filtered by tag
data "bloxone_federation_federated_pools" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Federated Pools
data "bloxone_federation_federated_pools" "example_all" {}
