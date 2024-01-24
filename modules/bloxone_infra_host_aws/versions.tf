terraform {
  required_version = ">= 1.0"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
    }
    aws = {
      source  = "hashicorp/aws"
    }
  }
}
