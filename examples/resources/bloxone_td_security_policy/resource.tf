resource "bloxone_td_security_policy" "example" {
  name = "example_security_policy"

  # Other optional fields
  rules = [
    {
      action = "action_allow",
      data   = "terraform_example",
      type   = "custom_list"
    }
  ]

  description = "Example Security Policy"

}
