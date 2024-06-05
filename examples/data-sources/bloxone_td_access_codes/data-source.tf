# Get Access Codes filtered by an attribute
data "bloxone_td_access_codes" "example_by_attribute" {
  filters = {
    "name" = "example_access_code"
  }
}

# Get all Access Codes
data "bloxone_td_access_codes" "example_all" {}
