# Get all DNS Hosts
data "bloxone_dns_hosts" "all_hosts" {}

# Get DNS Host by name
data "bloxone_dns_hosts" "dns_host_by_name" {
  filters = {
    name = "dns_host_by_name"
  }
}

# Get DNS Hosts with the specific tags
data "bloxone_dns_hosts" "all_dns_hosts_by_tag" {
  tag_filters = {
    location = "site1"
  }
}
