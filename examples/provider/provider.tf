provider "bloxone" {
  csp_url = "https://csp.infoblox.com"
  api_key = "<BloxOne DDI API Key>"

  # Other Optional Fields
  default_tags = {
    managed_by = "terraform"
    location   = "us-west-1"
  }
}
