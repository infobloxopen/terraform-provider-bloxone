# Create a Federated Realm
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# Create a Forward Looking Delegation
resource "bloxone_federation_forward_looking_delegation" "example" {
  address          = "10.10.0.0"
  cidr             = 24
  federated_realms = [bloxone_federation_federated_realm.example.id]
  name             = "example_forward_looking_delegation"
  comment          = "Example forward looking delegation"

  tags = {
    site = "Site A"
  }
}
