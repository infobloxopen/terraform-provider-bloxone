terraform {
  required_version = ">= 1.5"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
      version = ">= 1.1.0"
    }
    google = {
      source = "hashicorp/google"
      version = ">= 3.5.0"
    }
  }
}
