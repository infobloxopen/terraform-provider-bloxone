# A Federated Block is a required parent for next-available overlapping block lookups;
# look up an existing one rather than creating it here.
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Preview the next available overlapping block(s) under the federated block above, without allocating them.
data "bloxone_federation_next_available_overlapping_blocks" "example" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24

  # Number of overlapping blocks to preview. Defaults to 1.
  overlapping_block_count = 2
}

# With the default overlapping_block_count, only one next available overlapping block is previewed.
data "bloxone_federation_next_available_overlapping_blocks" "example_default_count" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24
}
