# Data Connector (CDC) service running on an on-prem host (Required as parent)
data "bloxone_infra_services" "cdc" {
  filters = {
    service_type = "cdc"
  }
}

locals {
  # cdc_services takes the bare service ID, without the "infra/service/" prefix
  cdc_service_id = trimprefix(data.bloxone_infra_services.cdc.results.0.id, "infra/service/")
}

# Built-in sources, present in every account (Required as parent)
data "bloxone_cdc_sources" "cloud" {
  filters = {
    type = "SOURCE_BLOXONE"
  }
}

data "bloxone_cdc_sources" "schedule" {
  filters = {
    type = "SOURCE_SCHEDULE"
  }
}

# NIOS source (Required as parent)
resource "bloxone_cdc_source_nios" "example" {
  name    = "example_cdc_flow_source_nios"
  address = "flow-nios.example.com"
}

# Destinations (Required as parent)
resource "bloxone_cdc_destination_syslog" "example" {
  name    = "example_cdc_flow_destination_syslog"
  address = "flow-syslog.example.com"
}

resource "bloxone_cdc_destination_splunk" "example" {
  name         = "example_cdc_flow_destination_splunk"
  address      = "flow-splunk.example.com:9997"
  indexer_name = "main"
}

resource "bloxone_cdc_application" "example" {
  name = "example_cdc_flow_application"

  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
}

resource "bloxone_cdc_destination_application" "example" {
  name        = "example_cdc_flow_destination_application"
  application = bloxone_cdc_application.example.id
}

# ETL filter (Required as parent)
resource "bloxone_cdc_etl_filter_threat_class_property" "example" {
  name = "example_cdc_flow_etl_filter_threat_class_property"
  data = ["Malicious"]
}

# NIOS source sending DNS query and response logs to a syslog server
resource "bloxone_cdc_flow" "example_nios" {
  name              = "example_cdc_flow_nios"
  cdc_services      = [local.cdc_service_id]
  source            = bloxone_cdc_source_nios.example.id
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.example.id]
}

# Infoblox Cloud source sending threat feed hits to Splunk, filtered by threat class
resource "bloxone_cdc_flow" "example_cloud" {
  name              = "example_cdc_flow_cloud"
  cdc_services      = [local.cdc_service_id]
  source            = data.bloxone_cdc_sources.cloud.results.0.id
  source_data_types = ["TD_THREAT_FEEDS_HITS_LOG"]
  destinations      = [bloxone_cdc_destination_splunk.example.id]
  description       = "Example CDC flow with an ETL filter"
  enabled           = true

  # filter_expression refers to the ETL filters by name
  etl_filters       = [bloxone_cdc_etl_filter_threat_class_property.example.id]
  filter_expression = bloxone_cdc_etl_filter_threat_class_property.example.name

  tags = {
    site = "Site A"
  }
}

# Schedule source running the automation script every 2 days
resource "bloxone_cdc_flow" "example_schedule" {
  name              = "example_cdc_flow_schedule"
  cdc_services      = [local.cdc_service_id]
  source            = data.bloxone_cdc_sources.schedule.results.0.id
  source_data_types = []
  destinations      = [bloxone_cdc_destination_application.example.id]
  schedule_option   = "custom"
  script_schedule   = "0 0 0 */2 * * *"
}
