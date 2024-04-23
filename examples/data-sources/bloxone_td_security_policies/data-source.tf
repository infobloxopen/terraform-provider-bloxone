# Get Security Policies filtered by an attribute
data "bloxone_td_security_policies" "example_by_attribute" {
  filters = {
    "name" = "example_security_policy"
  }
}

# Get Security Policies filtered by tag
data "bloxone_td_security_policies" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all Security Policies
data "bloxone_td_security_policies" "example_all" {}
