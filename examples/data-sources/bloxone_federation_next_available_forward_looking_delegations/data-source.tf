# NOTE: Forward Looking Delegations are created each time this data source is read (every plan and refresh).
# Terraform does not track or delete them: delete them through the API or the portal before destroying the
# parent Federated Block, Federated Pool or Federated Realm.

# Look up an existing Federated Block
data "bloxone_federation_federated_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block_with_pool"
  }
}

# Create the next available Forward Looking Delegations within the Federated Block.
# The protocol (ip4 or ip6) is taken from the block, so it is not set here.
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block" {
  federated_block_id = data.bloxone_federation_federated_blocks.example_by_attribute.results.0.id
  cidr               = 28
  name               = "example_next_available_fld_by_block"
  comment            = "Example next available FLD under block"

  # Number of Forward Looking Delegations to create. Defaults to 1.
  fld_count = 2

  tags = {
    site = "Site A"
  }
}

# Look up an existing IPv6 Federated Block
data "bloxone_federation_federated_blocks" "example_by_attribute_ipv6" {
  filters = {
    name = "example_federation_federated_block_with_pool_ipv6"
  }
}

# Create the next available IPv6 Forward Looking Delegations within the IPv6 Federated Block.
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block_ipv6" {
  federated_block_id = data.bloxone_federation_federated_blocks.example_by_attribute_ipv6.results.0.id
  cidr               = 56
  fld_count          = 2
}

# Look up an existing Federated Pool
data "bloxone_federation_federated_pools" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_block_pool"
  }
}

# Create the next available Forward Looking Delegations within the Federated Pool.
# The protocol (ip4 or ip6) must match the protocol of the blocks in the pool.
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_pool" {
  federated_pool_id = data.bloxone_federation_federated_pools.example_by_attribute.results.0.id
  cidr              = 28
  protocol          = "ip4"
}

# Look up an existing IPv6 Federated Pool
data "bloxone_federation_federated_pools" "example_by_attribute_ipv6" {
  filters = {
    name = "example_federation_federated_block_pool_ipv6"
  }
}

# Create the next available IPv6 Forward Looking Delegations within the IPv6 Federated Pool.
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_pool_ipv6" {
  federated_pool_id = data.bloxone_federation_federated_pools.example_by_attribute_ipv6.results.0.id
  cidr              = 56
  protocol          = "ip6"
}

# Create the next available Forward Looking Delegations globally. Neither federated_block_id nor
# federated_pool_id is set, so the tags select the candidate Federated Blocks: at least one Federated Block
# that belongs to a Federated Pool must carry the same tags.
# If several Federated Blocks (in the same or different Federated Pools) carry the same tags, all of them are
# candidates: the blocks are filled one after another, and the delegations of a single request with a count
# greater than 1 can be spread across them. Use a distinctive tag, or set federated_block_id or
# federated_pool_id, to control where the Forward Looking Delegations are created.
# The protocol also filters the candidate Federated Blocks. Set it whenever the blocks matching the tags can have
# different protocols (ip4 and ip6): blocks of the other protocol are skipped, and the request fails if no matching
# block has this protocol.
data "bloxone_federation_next_available_forward_looking_delegations" "example_global" {
  cidr     = 28
  protocol = "ip4"

  tags = {
    site = "Site A"
  }
}

# The same tags can be carried by IPv4 and IPv6 Federated Blocks, so the protocol selects which of them is used.
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ipv6" {
  cidr     = 56
  protocol = "ip6"

  tags = {
    site = "Site A"
  }
}
