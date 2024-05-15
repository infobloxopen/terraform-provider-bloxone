resource "bloxone_td_application_filter" "example" {
  name             = "example_application_filter"
  criteria  = [
    {
      name = "Microsoft 365"
    }
]
  # Other optional fields
  description = "Example of an Application Filter"
  tags = {
    site = "Site A"
  }
}