resource "bloxone_cdc_destination_reporting" "example" {
  name                        = "example_cdc_destination_reporting"
  address                     = "nios.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}

resource "bloxone_cdc_destination_reporting" "example_with_optional" {
  name                        = "example_cdc_destination_reporting_optional"
  address                     = "nios2.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
  description                 = "Example CDC NIOS reporting destination"
  enabled                     = true

  tags = {
    site = "Site A"
  }
}
