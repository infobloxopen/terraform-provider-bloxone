# Get join tokens filtered by an attribute
data "bloxone_infra_join_tokens" "example_by_attribute" {
  filters = {
    "name" = "example_join_token"
  }
}

# Get join tokens filtered by tag
data "bloxone_infra_join_tokens" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all join tokens
data "bloxone_infra_join_tokens" "example_all" {}
