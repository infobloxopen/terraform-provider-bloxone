resource "bloxone_td_named_list" "example" {
  name = "example_named_list"
  items_described = [
  {
    item = "tf-domain.com"
    description = "Exaample Domain"
  }
]
  # Other optional fields
  description = "Example Named List"
  tags = {
    location = "site1"
  }
  threat_level = "MEDIUM"
  confidence_level = "HIGH"

}
