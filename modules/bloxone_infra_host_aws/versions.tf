terraform {
  required_version = ">= 1.5"
  required_providers {
    bloxone = {
      source = "infobloxopen/bloxone"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.9"
    }
  }
}
