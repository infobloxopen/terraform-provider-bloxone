# A Federated Block is a required parent for next-available reserved block lookups;
# look up an existing one rather than creating it here.
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Preview the next available reserved block(s) under the federated block above, without allocating them.
data "bloxone_federation_next_available_reserved_blocks" "example" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24

  # Number of reserved blocks to preview. Defaults to 1.
  reserved_block_count = 2
}

data "bloxone_federation_next_available_reserved_blocks" "example_default_count" {
  id   = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr = 24
}
