# Federated Realm
resource "bloxone_federation_federated_realm" "example" {
  name = "example_federation_federated_realm"
}

# IPv4 Federated Block with a unique tag for global FLD filtering
resource "bloxone_federation_federated_block" "example" {
  name            = "example_federation_federated_block"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "10.10.0.0"
  cidr            = 16

  # Tag used to filter this block in the global data sources below.
  # Without a tags filter, the global endpoint may fail if any block in the
  # environment has empty-valued tags.
  tags = {
    example_fld_ip4 = "example"
  }
}

# IPv6 Federated Block with a unique tag for global FLD filtering
resource "bloxone_federation_federated_block" "example_ipv6" {
  name            = "example_federation_federated_block_ipv6"
  federated_realm = bloxone_federation_federated_realm.example.id
  address         = "2001:db8::"
  cidr            = 32

  tags = {
    example_fld_ip6 = "example"
  }
}

# Create next available Forward Looking Delegations globally for IPv4.
# Note: this data source creates FLDs as a side effect; those FLDs are NOT
# tracked in Terraform state and must be deleted before the parent block can be
# destroyed. Use `terraform destroy -refresh=false` after removing child FLDs.
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip4" {
  cidr     = 26
  protocol = "ip4"
  tags = {
    example_fld_ip4 = "example"
  }
  depends_on = [bloxone_federation_federated_block.example]
}

# Create next available Forward Looking Delegations globally for IPv6
data "bloxone_federation_next_available_forward_looking_delegations" "example_global_ip6" {
  cidr     = 48
  protocol = "ip6"
  tags = {
    example_fld_ip6 = "example"
  }
  depends_on = [bloxone_federation_federated_block.example_ipv6]
}

# Create next available Forward Looking Delegations within a specific Federated Block
data "bloxone_federation_next_available_forward_looking_delegations" "example_by_block" {
  federated_block_id = bloxone_federation_federated_block.example.id
  cidr               = 26
  fld_count          = 2
  comment            = "Example next available FLD under block"
}
