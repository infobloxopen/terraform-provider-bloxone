# Get Reserved Block filtered by an attribute
data "bloxone_reserved_blocks" "example_by_attribute" {
  filters = {
    name = "example_reserved_block"
  }
}

# Get Reserved Block filtered by tag
data "bloxone_reserved_blocks" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Reserved Block
data "bloxone_reserved_blocks" "example_all" {}
