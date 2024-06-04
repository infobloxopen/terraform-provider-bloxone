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

resource "bloxone_td_security_policy" "example" {
  name = "example_security_policy"

  # Other optional fields
  rules = [
    {
      action = "action_allow",
      data   = bloxone_td_named_list.example.name,
      type   = bloxone_td_named_list.example.type
    }
  ]
  description    = "Example Security Policy"
  ecs            = true
  onprem_resolve = true
  safe_search    = false
  tags = {
    site = "Site A"
  }
}
