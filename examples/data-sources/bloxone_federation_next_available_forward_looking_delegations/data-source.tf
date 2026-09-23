# Federated Realm
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# IPv4 Federated Block
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.10.0.0"
  cidr            = 16
}

# IPv6 Federated Block
resource "bloxone_federation_federated_block" "example_ipv6" {
  name            = "example_federation_federated_block_ipv6"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "2001:db8::"
  cidr            = 32
}

# Create next available Forward Looking Delegations globally for IPv4
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip4" {
  cidr       = 26
  protocol   = "ip4"
  depends_on = [bloxone_federation_federated_block.example]
}

# Create next available Forward Looking Delegations globally for IPv6
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip6" {
  cidr       = 48
  protocol   = "ip6"
  depends_on = [bloxone_federation_federated_block.example_ipv6]
}

# Create next available Forward Looking Delegations within a Federated Block
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block" {
  federated_block_id = bloxone_federation_federated_block.example.id
  cidr               = 26
  fld_count          = 2
  comment            = "Example next available FLD under block"
}
