# The Next Available Federated Block Lookups are performed under the existing Federated Block
# Here we retrieve an existing federated block
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Preview the next available federated block(s) under the federated block above, without allocating them.
data "bloxone_federation_next_available_federated_blocks" "example" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24

  # Number of federated blocks to preview. Defaults to 1.
  federated_block_count = 2
}

# With the default federated_block_count, only one next available federated block is previewed.
data "bloxone_federation_next_available_federated_blocks" "example_default_count" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24
}

# Instead of specifying a parent id directly, discover candidate parent Federated Blocks by tags.
data "bloxone_federation_next_available_federated_blocks" "example_by_tag_filters" {
  tag_filters = {
    env = "Site A"
  }
  cidr = 24
}
