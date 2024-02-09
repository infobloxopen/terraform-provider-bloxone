provider "bloxone" {
  csp_url = "https://csp.infoblox.com"
  api_key = "<BloxOne DDI API Key>"
  default_tags = {
    managed_by = "terraform"
  }
}
