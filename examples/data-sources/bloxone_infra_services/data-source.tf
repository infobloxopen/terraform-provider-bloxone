# Get services filtered by an attribute
data "bloxone_infra_services" "example_by_attribute" {
  filters = {
    "name" = "example_service"
  }
}

# Get services filtered by tag
data "bloxone_infra_services" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all services
data "bloxone_infra_services" "example_all" {}
