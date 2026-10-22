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

# Create the next available Forward Looking Delegation globally for IPv6.
# Without a Federated Block or Pool, the tags select the Federated Blocks to allocate from,
# so at least one IPv6 Federated Block must carry the same tags.
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip6" {
  cidr     = 48
  protocol = "ip6"
  tags = {
    site = "Site A"
  }
}
