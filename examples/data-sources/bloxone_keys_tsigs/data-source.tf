# Get TSIG Keys filtered by an attribute
data "bloxone_keys_tsigs" "example_by_attribute" {
  filters = {
    "name" = "example_tsig_key"
  }
}

# Get TSIG Keys filtered by tag
data "bloxone_keys_tsigs" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all TSIG Keys
data "bloxone_keys_tsigs" "example_all" {}
