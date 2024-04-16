resource "bloxone_td_access_code" "example" {
  name       = "example_access_code"
  activation = timestamp()
  expiration = timeadd(timestamp(), "24h")
  rules = [
    {
      action        = "",
      data          = "terraform_test",
      description   = "",
      redirect_name = "",
      type          = "custom_list"
    }
  ]
  # Other optional fields
  description = "Example Access Code"

}