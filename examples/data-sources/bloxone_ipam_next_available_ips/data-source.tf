data "bloxone_ipam_address_blocks" "example_by_attribute" {
  filters = {
    name = "example_address_block"
  }
}

data "bloxone_ipam_subnets" "example_by_attribute" {
  filters = {
    name = "example_subnet"
  }
}

data "bloxone_ipam_ranges" "example_by_attribute" {
  filters = {
    name = "example_range"
  }
}

// 'ip_count' allows you to get the number of next available ips in the resource specified by 'id'
// If not defined, count would default to 1
data "bloxone_ipam_next_available_ips" "example_next_ip_ab" {
  id       = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
  ip_count = 5
}

data "bloxone_ipam_next_available_ips" "example_next_ip_ab_default_count" {
  id = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
}

data "bloxone_ipam_next_available_ips" "example_next_ip_sub" {
  id       = data.bloxone_ipam_subnets.example_by_attribute.results.0.id
  ip_count = 5
}

data "bloxone_ipam_next_available_ips" "example_next_ip_range" {
  id       = data.bloxone_ipam_ranges.example_by_attribute.results.0.id
  ip_count = 5
}

// next_available_ips by tag
data "bloxone_ipam_next_available_ips" "example_address_block_next_ip_ab_by_tag" {
  tags = {
    name = "example_address_block"
  }
  ip_count      = 5
  resource_type = "address_block"
}
data "bloxone_ipam_next_available_ips" "example_subnet_next_ip_sub_by_tag" {
  tags = {
    name = "example_subnet"
  }
  ip_count      = 5
  resource_type = "subnet"
}
data "bloxone_ipam_next_available_ips" "example_range_next_ip_range_by_tag" {
  tags = {
    name = "example_range"
  }
  ip_count      = 5
  resource_type = "range"
}

