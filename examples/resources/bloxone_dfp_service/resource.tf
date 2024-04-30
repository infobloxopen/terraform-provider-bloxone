resource "bloxone_infra_host" "example_host" {
  display_name = "example_host"
}
resource "bloxone_infra_service" "example" {
  name           = "example_dfp_service"
  pool_id        = bloxone_infra_host.example_host.pool_id
  service_type   = "dfp"
  desired_state  = "start"
  wait_for_state = false

}

resource "bloxone_td_internal_domain_list" "example_list" {
  name             = "example_internal_domain_list"
  internal_domains = ["example.domain.com"]
}

resource "bloxone_dfp_service" "example" {
  service_id = bloxone_infra_service.example.id

  # Other optional fields
  internal_domain_lists = [bloxone_td_internal_domain_list.example_list.id]
  resolvers_all = [
    {
      address     = "1.1.1.1"
      is_fallback = true
      is_local    = false
      protocols   = ["DOT", "DO53"]
    }
  ]
}