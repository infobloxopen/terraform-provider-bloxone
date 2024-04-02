resource "bloxone_td_access_code" "example" {
  name       = "example_access_code"
  activation = "2021-01-01T00:00:00Z"
  expiration = "2025-01-02T00:00:00Z"
  rules = [
    {
      action        = "",
      data          = "antimalware",
      description   = "",
      redirect_name = "",
      type          = "named_feed"
    }
  ]
  # Other optional fields
  description = "Example Access Code"

}
