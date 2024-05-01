data "bloxone_infra_services" "example" {
  name = "example_dfp_service"
}

resource "bloxone_td_internal_domain_list" "example_list" {
  name             = "example_internal_domain_list"
  internal_domains = ["example.domain.com"]
}

resource "bloxone_dfp_service" "example" {
  service_id = data.bloxone_infra_services.example.results.0.id

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