provider "google" {
  project = var.project_id
  default_labels = {
    goog-partner-solution = "isol_plb32_0014m00001h31caqay_vcrx7g5d6abwbyhqip2tvj7lspfpallp"
  }
}

locals {
  network_interfaces = [ for i, n in var.networks : {
    network     = n,
    subnetwork  = length(var.sub_networks) > i ? element(var.sub_networks, i) : null
    external_ip = length(var.external_ips) > i ? element(var.external_ips, i) : "NONE"
    }
  ]
  
  # TODO check for google-logging-enable & google-monitoring-enable
  # m = { 
  #  google-logging-enable = var.google-logging-enable == true ? 0 : 1
  #  google-monitoring-enable = var.google-monitoring-enable == true ? 0 : 1
  # }

  labels = jsondecode(var.labels)
  nios_x_tags = jsondecode(var.nios_x_tags)
  
  metadata = merge(
    jsondecode(var.metadata),
    {
      user-data = templatefile("${path.module}/userdata.tftpl", {
        join_token = var.join_token
        http_proxy = var.http_proxy
        tags       = local.nios_x_tags
      })
    }
  )
}

resource "google_compute_instance" "instance" {
  name                = var.goog_cm_deployment_name
  machine_type        = var.machine_type
  labels              = local.labels
  #deletion_protection = var.deletion_protection
  zone                = var.zone

  boot_disk {
    initialize_params {
      image  = var.source_image
      type   = var.boot_disk_type
      size   = var.boot_disk_size
    }
  }
  
  dynamic "network_interface" {
    for_each = local.network_interfaces
    content {
      network = network_interface.value.network
      subnetwork = network_interface.value.subnetwork

      dynamic "access_config" {
        for_each = network_interface.value.external_ip == "NONE" ? [] : [1]
        content {
          nat_ip = network_interface.value.external_ip == "EPHEMERAL" ? null : network_interface.value.external_ip
        }
      }
    }
  }

  service_account {
    email = "default"
    scopes = compact([
      "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring.write"
      ,var.enable_cloud_api == true ? "https://www.googleapis.com/auth/cloud-platform" : null
    ])
  }

  metadata = local.metadata
}
