terraform {
  required_version = ">= 1.5.0"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
      version = ">= 1.0.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0.0"
    }
  }
}
