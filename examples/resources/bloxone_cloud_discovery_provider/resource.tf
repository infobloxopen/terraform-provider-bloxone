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

resource "bloxone_cloud_discovery_provider" "example_azure" {
  name               = "example_provider_azure"
  provider_type      = "Microsoft Azure"
  account_preference = "single"
  credential_preference = {
    access_identifier_type = "tenant_id"
    credential_type        = "dynamic"
  }
  source_configs = [
    {
      credential_config = {
        access_identifier = "xyz98765-4321-abcd-efgh-ijklmnopqrst"
      }
      restricted_to_accounts = ["12345678-abcd-efgh-ijkl-901234567890"]
    }
  ]

  # Other Optional fields

  tags = {
    site = "Site A"
  }
}

resource "bloxone_cloud_discovery_provider" "example_gcp" {
  name               = "example_provider_gcp"
  provider_type      = "Google Cloud Platform"
  account_preference = "single"
  credential_preference = {
    access_identifier_type = "project_id"
    credential_type        = "dynamic"
  }
  source_configs = [
    {
      credential_config = {
        access_identifier = "my-bloxone-example-2024"
      }
    }
  ]

  # Other Optional fields

  tags = {
    site = "Site A"
  }
}