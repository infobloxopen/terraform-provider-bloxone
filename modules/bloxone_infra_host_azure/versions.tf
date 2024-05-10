terraform {
  required_version = ">= 1.5"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
      version = ">= 1.1.0"
    }
    azurerm = {
      source = "hashicorp/azurerm"
      version = ">= 3.0.0"
    }
  }
}
