terraform {
  required_version = ">= 1.0"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
    }
    azurerm = {
      source = "hashicorp/azurerm"
    }
  }
}

provider "azurerm" {
  features {}
}
