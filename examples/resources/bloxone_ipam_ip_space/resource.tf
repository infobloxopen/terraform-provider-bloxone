resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
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

resource "bloxone_ipam_ip_space" "example_tags" {
  name    = "example_ip_space_tags"
  comment = "Example IP space with tags created by the terraform provider"
  tags = {
    location = "site1"
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
  default_realms = [data.bloxone_federation_federated_realms.federated_realm.results.0.id]
}
