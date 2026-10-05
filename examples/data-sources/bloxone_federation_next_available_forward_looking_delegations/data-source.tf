# Create a Federated Realm
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# Create a Federated Block
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.10.0.0"
  cidr            = 16

  tags = {
    site = "Site A"
  }
}

# Create next available Forward Looking Delegations within a specific Federated Block
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block" {
  federated_block_id = bloxone_federation_federated_block.example.id
  cidr               = 26
  fld_count          = 2
}

# Create next available Forward Looking Delegations globally for IPv4
# depends_on ensures the block exists before the global allocation is attempted
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip4" {
  cidr     = 26
  protocol = "ip4"

  tags = {
    site = "Site A"
  }

  depends_on = [bloxone_federation_federated_block.example]
}
