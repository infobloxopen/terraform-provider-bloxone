output "services" {
  value       = bloxone_infra_service.this
  description = "The `bloxone_infra_service` objects for the instance. May be empty if no services were specified."
}

output "host" {
  value       = data.bloxone_infra_hosts.this.results.0
  description = "The `bloxone_infra_host` object for the instance"
}

output "gcp_instance" {
  value       = google_compute_instance.this
  description = "The GCP instance object for the instance"
}
