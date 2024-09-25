resource "bloxone_dns_view" "example" {
  name = "example_dns_view"
}


resource "bloxone_cloud_discovery_provider" "example_aws" {
  name               = "example_provider_aws"
  provider_type      = "Amazon Web Services"
  account_preference = "single"
  credential_preference = {
    access_identifier_type = "role_arn"
    credential_type        = "dynamic"
  }
  source_configs = [
    {
      credential_config = {
        access_identifier = "arn:aws:iam::123456789012:role/role-name"
      }
    }
  ]

  # Other Optional fields
  destinations = [
    {
      config           = {}
      destination_type = "IPAM/DHCP"
    },
    {
      config = {
        view_id = bloxone_dns_view.example.id
      }
      destination_type = "DNS"
    }
  ]

  tags = {
    site = "Site A"
  }

}
