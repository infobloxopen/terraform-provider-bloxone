# Get Cloud Discovery Providers filtered by an attribute
data "bloxone_cloud_discovery_providers" "example_by_attribute" {
  filters = {
    "name" = "example_provider_azure"
  }
}

# Get Cloud Discovery Providers filtered by tag
data "bloxone_cloud_discovery_providers" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Cloud Discovery Providers
data "bloxone_cloud_discovery_providers" "example_all" {}
