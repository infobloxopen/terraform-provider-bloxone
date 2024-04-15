#Get anycast configuration by tag filter
data "bloxone_anycast_ac_configs" "example_filter" {
  tag_filters = {
    tag1 = "value1"
  }
}

#Get anycast configuration by service
data "bloxone_anycast_ac_configs" "example_service" {
  service = "DNS"
}

#Get all the anycast configuration
data "bloxone_anycast_ac_configs" "example_all" {
}
