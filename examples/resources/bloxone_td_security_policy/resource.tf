resource "bloxone_td_security_policy" "example" {
  name = "example_security_policy"

  # Other optional fields
  rules = [
    {
      action = "action_allow",
      data   = "custom_list_example",
      type   = "custom_list"
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
