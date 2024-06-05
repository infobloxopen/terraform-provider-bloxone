/**
 * # Terraform Module to create BloxOne Host in GCP
 *
 * This module will provision a GCP virtual machine that uses a BloxOne image.
 * The virtual machine will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
 * If a join token is not provided, a new one will be created.
 *
 * The BloxOne Host created in the CSP is created automatically, and cannot be managed through terraform.
 * A `bloxone_infra_hosts` data source is provided to retrieve the host information from the CSP.
 * The data source will use the `tags` variable to filter the hosts.
 * A `tf_module_host_id` tag will be added to the tags variable so that the data source can uniquely find the host.
 *
 * This module will also create a BloxOne Infra Service for each service type provided in the `services` variable.
 * The service will be named `<service_type>_<host_display_name>`.
 *
/**
 * ## Example Usage
 *
 * ```hcl
 * module "bloxone_infra_host_gcp" {
 *   source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_gcp"
 *
 *   name         = "bloxone-vm"
 *   source_image = "bloxone-v381"
 *
 *   network_interfaces = [
 *     {
 *       network          = "gcp-external-network"
 *       subnetwork       = "gcp-external-subnet"
 *       assign_public_ip = true
 *     },
 *     {
 *       network          = "gcp-internal-network"
 *       subnetwork       = "gcp-internal-subnet"
 *     }
 *   ]
 *
 *   gcp_instance_labels = {
 *     environment = "dev"
 *   }
 *
 *   tags = {
 *     location = "office1"
 *   }
 *
 *   services = {
 *     dhcp = "start"
 *     dns = "start"
 *   }
 * }
 * ```
 * 
 */

resource "random_uuid" "this" {}

locals {
  join_token = var.join_token == null ? bloxone_infra_join_token.this[0].join_token : var.join_token
  tags = merge(
    var.tags,
    {
      "tf_module_host_id" = "bloxone-host-${random_uuid.this.result}"
    }
  )
}

resource "bloxone_infra_join_token" "this" {
  count = var.join_token == null ? 1 : 0
  name  = "jt-${random_uuid.this.result}"
}

resource "google_compute_instance" "this" {
  name                = var.name
  machine_type        = var.machine_type
  labels              = var.gcp_instance_labels
  deletion_protection = var.deletion_protection

  boot_disk {
    initialize_params {
      image  = var.source_image
      type   = var.disk_type
      size   = var.disk_size
      labels = var.gcp_instance_labels
    }
  }

  dynamic "network_interface" {
    for_each = var.network_interfaces
    content {
      network    = network_interface.value.network
      subnetwork = network_interface.value.subnetwork

      dynamic "access_config" {
        for_each = network_interface.value.assign_public_ip == true ? [1] : []
        content {
          nat_ip = network_interface.value.assign_public_ip == true ? null : ""
        }
      }
    }
  }

  dynamic "service_account" {
    for_each = var.service_account == null ? [] : [1]
    content {
      email  = var.service_account.email
      scopes = var.service_account.scopes
    }
  }

  metadata = {
    user-data = templatefile("${path.module}/userdata.tftpl", {
      join_token = local.join_token
      tags       = local.tags
    })
  }
}

data "bloxone_infra_hosts" "this" {
  filters = {
    host_type = "3" // For BloxOne VM
  }
  tag_filters        = local.tags
  retry_if_not_found = true
  timeouts           = var.timeouts

  lifecycle {
    postcondition {
      condition     = self.results == null ? false : length(self.results) == 1
      error_message = "BloxOne Host not found in CSP."
    }
  }

  depends_on = [google_compute_instance.this]
}

resource "bloxone_infra_service" "this" {
  for_each       = var.services
  name           = format("%s_%s", each.key, data.bloxone_infra_hosts.this.results[0].display_name)
  pool_id        = data.bloxone_infra_hosts.this.results[0].pool_id
  service_type   = each.key
  desired_state  = each.value
  tags           = local.tags
  timeouts       = var.timeouts
  wait_for_state = var.wait_for_state
}
