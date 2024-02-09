resource "time_offset" "one_week" {
  offset_days = 7
}

resource "bloxone_infra_join_token" "example" {
  name = "example_join_token"

  # Other optional fields
  description = "Join token for Site A"
  expires_at  = time_offset.one_week.rfc3339
  tags = {
    site = "Site A"
  }
}
