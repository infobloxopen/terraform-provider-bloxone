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
