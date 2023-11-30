resource "bloxone_keys_tsig" "example" {
  name = "example_tsig_key"

  # Other optional fields
  comment   = "key created through terraform"
  algorithm = "hmac_sha256"
  secret    = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
  tags = {
    location = "site1"
  }
}
