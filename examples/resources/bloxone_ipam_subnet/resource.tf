resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_ipam_address_block" "example" {
  address = "10.0.0.0"
  cidr    = 16
  space   = bloxone_ipam_ip_space.example.id
}

data "bloxone_dhcp_option_codes" "option_code" {
  filters = {
    name = "domain-name-servers"
  }
}

data "bloxone_federation_federated_realms" "federated_realm" {
  filters = {
    name = "example_federation_federated_realm"
  }
}

# Static address
resource "bloxone_ipam_subnet" "example" {
  address = "10.0.0.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_subnet"
  comment = "Subnet for Site A"
  tags = {
    site = "Site A"
  }
  //dhcp options
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "10.0.0.1"
      type         = "option"
    }
  ]

  //federated realms
  federated_realms = [data.bloxone_federation_federated_realms.federated_realm.results.0.id]
}

# Next available subnet
resource "bloxone_ipam_subnet" "example_na_s" {
  next_available_id = bloxone_ipam_address_block.example.id
  cidr              = 24
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_subnet"
  comment = "Subnet for Site A"
  tags = {
    site = "Site A"
  }
}
