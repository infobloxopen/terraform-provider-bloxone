# Get hosts filtered by an attribute
data "bloxone_infra_hosts" "example_by_attribute" {
  filters = {
    "name" = "example_host"
  }
}

# Get hosts filtered by tag
data "bloxone_infra_hosts" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all hosts
data "bloxone_infra_hosts" "example_all" {}
