resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

resource "bloxone_federation_overlapping_block" "example" {
  name            = "example_federation_overlapping_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  cidr            = 24
  address         = "10.10.0.0"

  comment = "This is an example overlapping block"
  tags = {
    site = "Site A"
  }

}

# Manage a Federated Block (Required as Parent for next_available_id)
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.20.0.0"
  cidr            = 16
}

# Allocate an Overlapping Block from the next available address under the Federated Block above,
# without previewing it via a separate data source first.
resource "bloxone_federation_overlapping_block" "example_next_available" {
  name              = "example_federation_overlapping_block_next_available"
  federated_realm   = bloxone_federation_federated_realm.example.id
  next_available_id = bloxone_federation_federated_block.example.id
  cidr              = 24
}
