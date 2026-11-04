# Manage a Federated Realm (Required as Parent)
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# Manage a Federated Reserved Block
resource "bloxone_federation_reserved_block" "example" {
  name            = "example_reserved_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.20.0.0"
  cidr            = 24
  comment         = "Reserved for third-party allocations"

  tags = {
    site = "Site A"
  }
}

# Manage a Federated Block (Required as Parent for next_available_id)
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.10.0.0"
  cidr            = 16
}

# Allocate a Reserved Block from the next available address under the Federated Block above,
# without previewing it via a separate data source first.
resource "bloxone_federation_reserved_block" "example_next_available" {
  name              = "example_reserved_block_next_available"
  federated_realm   = bloxone_federation_federated_realm.example.id
  next_available_id = bloxone_federation_federated_block.example.id
  cidr              = 24
}
