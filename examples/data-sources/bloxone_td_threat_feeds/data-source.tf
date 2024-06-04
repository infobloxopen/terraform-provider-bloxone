# Get Threat Feeds filtered by an attribute
data "bloxone_td_threat_feeds" "example_by_attribute" {
  filters = {
    "name" = "Cryptocurrency"
  }
}

# Get all Threat Feeds
data "bloxone_td_threat_feeds" "example_all" {}
