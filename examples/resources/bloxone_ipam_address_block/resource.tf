resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
  tags = {
    location = "site1"
  }
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

resource "bloxone_ipam_address_block" "example" {
  address = "192.168.1.0"
  cidr    = 24
  name    = "example_address_block"
  space   = bloxone_ipam_ip_space.example.id

  //dhcp options
  dhcp_options = [
    {
      option_code  = data.bloxone_dhcp_option_codes.option_code.results.0.id
      option_value = "10.0.0.1"
      type         = "option"
    }
  ]
}

resource "bloxone_ipam_address_block" "example_tags" {
  address = "10.0.0.0"
  cidr    = 8
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_address_block_tags"
  comment = "Example address block with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
  asm_config = {
    asm_threshold       = 90
    enable              = "true"
    enable_notification = "true"
    forecast_period     = 10
    growth_factor       = 10
    growth_type         = "percent"
    history             = 30
    min_total           = 2
    min_unused          = 10
    reenable_date       = "2024-01-24T10:10:00+00:00"
  }
  dhcp_config = {
    allow_unknown = true
    ignore_list = [
      {
        type  = "hardware"
        value = "aa:bb:cc:dd:ee:ff"
      },
      {
        type  = "client_text"
        value = "001d.a18b.36d0"
      },
      {
        type  = "client_hex"
        value = "333964392D4769302F31"
      }
    ]
  }

  //federated realms
  federated_realms = [data.bloxone_federation_federated_realms.federated_realm.results.0.id]
}

# Next available address block
resource "bloxone_ipam_address_block" "example_na_ab" {
  next_available_id = bloxone_ipam_address_block.example.id
  cidr              = 26
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  name    = "example_address_block_tags"
  comment = "Example address block with tags created by the terraform provider"
  tags = {
    location = "site1"
  }
  asm_config = {
    asm_threshold       = 90
    enable              = "true"
    enable_notification = "true"
    forecast_period     = 10
    growth_factor       = 10
    growth_type         = "percent"
    history             = 30
    min_total           = 2
    min_unused          = 10
    reenable_date       = "2024-01-24T10:10:00+00:00"
  }
  dhcp_config = {
    allow_unknown = true
    ignore_list = [
      {
        type  = "hardware"
        value = "aa:bb:cc:dd:ee:ff"
      },
      {
        type  = "client_text"
        value = "001d.a18b.36d0"
      },
      {
        type  = "client_hex"
        value = "333964392D4769302F31"
      }
    ]
  }
}
