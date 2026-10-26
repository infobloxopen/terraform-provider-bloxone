# Get Federated Block filtered by an attribute
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Get Federated Block filtered by tag
data "bloxone_federation_federated_blocks" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get the Federated Blocks in a Federated Pool
data "bloxone_federation_federated_pools" "example_pool" {
  filters = {
    name = "example_federation_federated_pool"
  }
}

data "bloxone_federation_federated_blocks" "example_by_pool" {
  filters = {
    federated_pool_id = data.bloxone_federation_federated_pools.example_pool.results[0].id
  }
}

# Get all Federated Block
data "bloxone_federation_federated_blocks" "example_all" {}
