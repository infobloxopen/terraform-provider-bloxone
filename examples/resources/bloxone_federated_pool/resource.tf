# Create a Federated Realm to associate the pool with
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# Create a Federated Pool
resource "bloxone_federation_federated_pool" "example" {
  name            = "example_federation_federated_pool"
  federated_realm = bloxone_federation_federated_realm.example.id
  protocol        = "ip4"
  region          = "us-east-1"

  # Other optional fields
  description = "Example Federated Pool"

  network_compliance = {
    default_netmask_length = 26
    minimum_netmask_length = 25
    maximum_netmask_length = 28
  }

  tags = {
    site = "Site A"
  }
}
