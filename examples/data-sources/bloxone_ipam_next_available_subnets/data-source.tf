data "bloxone_ipam_address_blocks" "example_by_attribute" {
  filters = {
    name = "example_address_block"
  }
}

// List the subnets available in the above address block
// subnet_count = number of subnets to be created, if not specified defaults to 1
// cidr = size of subnet
data "bloxone_ipam_next_available_subnets" "example_tf_subs" {
  id           = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
  name         = "example_subnet"
  cidr         = 29
  subnet_count = 5
}