# Get GSS-TSIG/Kerberos Keys filtered by an attribute
data "bloxone_keys_kerberoses" "example_by_attribute" {
  filters = {
    "principal" = "DNS/ns.b1ddi.example.com"
  }
}

# Get GSS-TSIG/Kerberos Keys filtered by tag
data "bloxone_keys_kerberoses" "example_by_tag" {
  tag_filters = {
    "tag1" = "value1"
  }
}

# Get all GSS-TSIG/Kerberos Keys
data "bloxone_keys_kerberoses" "example_all" {}
