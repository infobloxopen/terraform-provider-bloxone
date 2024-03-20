/**
 * # Terraform Module to create BloxOne Host in AWS
 *
 * This module will provision an AWS EC2 instance that uses a BloxOne AMI.
 * The instance will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
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
 * ## Example Usage
 *
 * ```hcl
 * module "bloxone_infra_host_aws" {
 *   source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_aws"
 *
 *   key_name = "my-key"
 *   subnet_id = "subnet-id"
 *   vpc_security_group_ids = ["vpc-security-group-id"]
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
  ami_id     = var.ami == null ? data.aws_ami.bloxone.id : var.ami
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

data "aws_ami" "bloxone" {
  most_recent = true
  filter {
    name   = "name"
    values = ["BloxOne381_MarketPlace-96cf85a8-a940-4dd0-80a5-80ab90fb1d1a"]
  }
}

resource "aws_instance" "this" {
  ami                    = local.ami_id
  instance_type          = var.instance_type
  key_name               = var.key_name
  vpc_security_group_ids = var.vpc_security_group_ids
  subnet_id              = var.subnet_id
  user_data = templatefile(
    "${path.module}/userdata.tftpl",
    {
      join_token : local.join_token
      tags : local.tags
    }
  )
  user_data_replace_on_change = true
  tags                        = var.aws_instance_tags
  metadata_options {
    instance_metadata_tags = "enabled"
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

  depends_on = [
    aws_instance.this
  ]
}

resource "bloxone_infra_service" "this" {
  for_each      = var.services
  name          = format("%s_%s", each.key, data.bloxone_infra_hosts.this.results[0].display_name)
  pool_id       = data.bloxone_infra_hosts.this.results[0].pool_id
  service_type  = each.key
  desired_state = each.value
  tags          = local.tags
  timeouts      = var.timeouts
}
