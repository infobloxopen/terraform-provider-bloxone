# Create a Federated Realm ( Required As Parent )
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

# Create an IPv6 Forward Looking Delegation
# Note: specify the address without a prefix length; use `cidr` for the prefix length.
resource "bloxone_federation_forward_looking_delegation" "example_ipv6" {
  address          = "2001:db8:a1::"
  cidr             = 48
  federated_realms = [bloxone_federation_federated_realm.example.id]
  name             = "example_forward_looking_delegation_ipv6"
}
