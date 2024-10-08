resource "bloxone_cloud_discovery_provider" "example_azure" {
  name               = "example_provider_azure"
  provider_type      = "Microsoft Azure"
  account_preference = "auto_discover_multiple"
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
  destination_types_enabled = [
    "IPAM/DHCP",
    "ACCOUNTS"
  ]
  destinations = [
    {
      config           = {}
      destination_type = "ACCOUNTS"
    },
    {
      config = {
        view_id = bloxone_dns_view.example.id
      }
      destination_type = "IPAM/DHCP"
    }
  ]

  # Other Optional fields

  tags = {
    site = "Site A"
  }
}
