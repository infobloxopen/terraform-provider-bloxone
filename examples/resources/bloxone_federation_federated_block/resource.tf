# Create a Federated Realm to associate the pool and the blocks with
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_block_realm"
}

# Create a basic Federated Block without a Federated Pool
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  cidr            = 24
  address         = "10.20.0.0"
}

# Create a Federated Pool to create a block in
resource "bloxone_federation_federated_pool" "example" {
  name            = "example_federation_federated_block_pool"
  federated_realm = bloxone_federation_federated_realm.example.id
  protocol        = "ip4"
}

# Create a Federated Block in the Federated Pool
resource "bloxone_federation_federated_block" "example_with_pool" {
  name              = "example_federation_federated_block_with_pool"
  federated_realm   = bloxone_federation_federated_realm.example.id
  federated_pool_id = bloxone_federation_federated_pool.example.id
  cidr              = 24
  address           = "10.10.0.0"

  # network_compliance can be configured for NIOS_X provider only
  network_compliance = {
    default_netmask_length = 26
    minimum_netmask_length = 25
    maximum_netmask_length = 28
  }

  tags = {
    site = "Site A"
  }
}

# Allocate a Federated Block from the next available address under the parent block above,
# without previewing it via a separate data source first.
resource "bloxone_federation_federated_block" "example_next_available" {
  name              = "example_federation_federated_block_next_available"
  federated_realm   = bloxone_federation_federated_realm.example.id
  next_available_id = bloxone_federation_federated_block.example.id
  cidr              = 25
}

# Allocate a Federated Block directly from a Federated Pool's remaining space. This is only
# supported for AWS-backed pools, and has no preview possible since there is no corresponding
# data source for this case.
resource "bloxone_federation_federated_block" "example_next_available_from_pool" {
  name              = "example_federation_federated_block_next_available_from_pool"
  federated_realm   = bloxone_federation_federated_realm.example.id
  next_available_id = bloxone_federation_federated_pool.example.id
  cidr              = 24
}
