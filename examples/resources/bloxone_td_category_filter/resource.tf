resource "bloxone_td_category_filter" "example" {
  name       = "example_category_filter"
  categories = ["Tutoring", "College"]

  # Other optional fields
  description = "Example of a Category Filter"
  tags = {
    site = "Site A"
  }
}
