resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.10.0.0"
  cidr            = 16
}

# Preview the next available reserved block(s) under the federated block above, without allocating them.
data "bloxone_next_available_reserved_blocks" "example" {
  id   = bloxone_federation_federated_block.example.id
  cidr = 24

  # Number of reserved blocks to preview. Defaults to 1.
  reserved_block_count = 2
}
