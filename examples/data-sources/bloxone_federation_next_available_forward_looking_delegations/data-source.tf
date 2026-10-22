# NOTE: Forward Looking Delegations are created each time this data source is read (every plan and refresh).
# Terraform does not track or delete them: delete them through the API or the portal before destroying the
# parent Federated Block, Federated Pool or Federated Realm.

# A Federated Block is a required parent for the next available Forward Looking Delegations;
# look up an existing one rather than creating it here.
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block"
  }
}

# Existing Forward Looking Delegation inside the Federated Block above
resource "bloxone_federation_forward_looking_delegation" "example" {
  address          = "10.10.0.0"
  cidr             = 24
  federated_realms = [data.bloxone_federation_federated_blocks.example_by_attribute.results.0.federated_realm]
  name             = "example_forward_looking_delegation"
  comment          = "Example forward looking delegation"
}

# Create the next available Forward Looking Delegations within the Federated Block
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block" {
  federated_block_id = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr               = 26
  comment            = "Example next available FLD under block"

  # Number of Forward Looking Delegations to create. Defaults to 1.
  fld_count = 2

  depends_on = [bloxone_federation_forward_looking_delegation.example]
}

# Look up an existing IPv6 Federated Block
data "bloxone_federation_federated_blocks" "example_by_attribute_ip6" {
  filters = {
    name = "example_federation_federated_block_ipv6"
  }
}

# Create the next available Forward Looking Delegations within the IPv6 Federated Block
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block_ip6" {
  federated_block_id = data.bloxone_federation_federated_blocks.example_by_attribute_ip6.results.0.id
  cidr               = 48
  fld_count          = 2

  tags = {
    site = "Site A"
  }
}
