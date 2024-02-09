data "bloxone_ipam_address_blocks" "example_by_attribute" {
  filters = {
    name = "example_address_block"
  }
}

// 'address_block_count' allows you to get the number of next available address blocks in the address block specified by 'id'
// If not defined, count would default to 1
data "bloxone_ipam_next_available_address_blocks" "example_next_available_ab" {
  id                  = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
  address_block_count = 5
}

data "bloxone_ipam_next_available_address_blocks" "example_next_available_ab_default_count" {
  id = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
}
