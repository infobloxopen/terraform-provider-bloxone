resource "bloxone_td_named_list" "example" {
  name = "example_named_list"
  items_described = [
    {
      item        = "tf-domain.com"
      description = "Example Domain"
    }
  ]
  type = "custom_list"
}

resource "bloxone_td_access_code" "example" {
  name       = "example_access_code"
  activation = timestamp()
  expiration = timeadd(timestamp(), "24h")
  rules = [
    {
      data = bloxone_td_named_list.example.name,
      type = bloxone_td_named_list.example.type
    }
  ]
  # Other optional fields
  description = "Example Access Code"

}