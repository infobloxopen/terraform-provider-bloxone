# Get Overlapping Block filtered by an attribute
data "bloxone_federation_overlapping_blocks" "example_by_attribute" {
  filters = {
    name = "example_federation_overlapping_block"
  }
}

# Get Overlapping Block filtered by tag
data "bloxone_federation_overlapping_blocks" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Overlapping Block
data "bloxone_federation_overlapping_blocks" "example_all" {}