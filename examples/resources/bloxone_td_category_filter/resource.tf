data "bloxone_td_content_categories" "test" {}

resource "bloxone_td_category_filter" "example" {
  name       = "example_category_filter"
  categories = [data.bloxone_td_content_categories.test.results.0.category_name]

  # Other optional fields
  description = "Example of an Category Filter"
  tags = {
    site = "Site A"
  }
}