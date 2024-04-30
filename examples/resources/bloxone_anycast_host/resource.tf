data "bloxone_infra_hosts" "anycast_hosts" {
  filters = {
    #pool_id = data.bloxone_infra_services.anycast_services.results.0.pool_id
    display_name = "anycast_real"
  }
}

# Create an anycast config profile with onprem hosts
resource "bloxone_anycast_config" "test_onprem_hosts" {
  anycast_ip_address = "10.10.10.1"
  name               = "Anycast_config_example"
  service            = "DNS" # service set to DNS

}

# Adding an anycast host with BGP routing protocol
resource "bloxone_anycast_host" "test_anycast_host" {
  id = one(data.bloxone_infra_hosts.anycast_hosts.results).legacy_id

  # Adding the anycast config profile and enabling BGP routing protocol
  anycast_config_refs = [
    {
      anycast_config_name = bloxone_anycast_config.test_onprem_hosts.name
      routing_protocols   = ["BGP", "OSPF"]
    }
  ]

  # Adding the BGP configuration
  config_bgp = {
    asn           = "6500"
    asn_text      = "6500"
    holddown_secs = 180
    neighbors = [
      {
        asn        = "6501"
        ip_address = "10.20.0.3"
      }
    ]
  }

  # Adding the BGP configuration
  config_ospf = {
    area_type           = "STANDARD"
    area                = "10.0.0.1"
    authentication_type = "Clear"
    interface           = "eth0"
    authentication_key  = "YXV0aGV"
    hello_interval      = 10
    dead_interval       = 40
    retransmit_interval = 5
    transmit_delay      = 1
  }
}
