resource "bloxone_td_network_list" "example" {
  name = "example_network_list"
  items = ["156.2.3.0/24","10.24.32.0/24"]

  # Other optional fields
  description = "Example Network List"
}
