# Get Federated Block filtered by an attribute
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Get Federated Block filtered by tag
data "bloxone_federation_federated_blocks" "example_by_tag" {
  tag_filters = {
    key1 = "value1"
  }
}

# Get all Federated Block
data "bloxone_federation_federated_blocks" "example_all" {}
