resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

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
