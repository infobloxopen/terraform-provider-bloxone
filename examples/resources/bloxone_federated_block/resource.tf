resource "bloxone_federated_realm" "example" {
  name = "example_federated_realm"
}

resource "bloxone_federated_block" "example" {
  name            = "example_federated_block"
  federated_realm = bloxone_federated_realm.test_name.id
  cidr            = 24
  address         = "10.10.0.0"
  //tags
  tags = {
    key1 = "value1"
    key2 = "value2"
  }
}