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