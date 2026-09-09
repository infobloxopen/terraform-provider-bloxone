resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  cidr            = 24
  address         = "10.10.0.0"

  network_compliance = {
    default_netmask_length = 26
    minimum_netmask_length = 25
    maximum_netmask_length = 28
  }

  tags = {
    site = "Site A"
  }
}