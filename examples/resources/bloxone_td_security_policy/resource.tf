resource "bloxone_td_security_policy" "example" {
  name       = "example_security_policy"

  # Other optional fields
  rules = [
    {
      action = "allow",
      data = "terraform_test",
      type = "custom_list"
    }
  ]

  description = "Example Security Policy"

}