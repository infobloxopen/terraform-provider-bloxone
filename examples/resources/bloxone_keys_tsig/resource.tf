resource "bloxone_keys_tsig" "example" {
  name = "example_tsig_key"

  # Other optional fields
  comment   = "key created through terraform"
  algorithm = "hmac_sha256"
  secret    = "your-base64-encoded-secret-here"
  tags = {
    location = "site1"
  }
}
