# Get DFP Services filtered by an attribute
data "bloxone_dfp_services" "test" {
  filters = {
    service_name = "example_dfp_service"
  }
}

# Get all DFP Services
data "bloxone_dfp_services" "example_all" {}
