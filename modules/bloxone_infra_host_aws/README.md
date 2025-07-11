<!-- BEGIN_TF_DOCS -->
# Terraform Module to create BloxOne Host in AWS

This module will provision an AWS EC2 instance that uses a BloxOne AMI.
The instance will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
If a join token is not provided, a new one will be created.

The BloxOne Host created in the CSP is created automatically, and cannot be managed through terraform.
A `bloxone_infra_hosts` data source is provided to retrieve the host information from the CSP.
The data source will use the `tags` variable to filter the hosts.
A `tf_module_host_id` tag will be added to the tags variable so that the data source can uniquely find the host.

This module will also create a BloxOne Infra Service for each service type provided in the `services` variable.
The service will be named `<service_type>_<host_display_name>`.

## Example Usage

```hcl
terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      # Other parameters...
    }
    aws = {
      source  = "hashicorp/aws"
    }
  }
}

provider "aws" {
  access_key = "<access-key>"
  secret_key = "<secret-key>"
  token      = "<session-token>"
  region     = "<selected-region>"   
}

provider "bloxone" {
  csp_url = "<csp-url>"
  api_key = "<api-key>"
}

module "bloxone_infra_host_aws" {
  source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_aws"

  key_name = "my-key"
  subnet_id = "subnet-id"
  vpc_security_group_ids = ["vpc-security-group-id"]

  services = {
    dhcp = "start"
    dns = "start"
  }

  aws_instance_tags= {
    Name = "bloxone-instance"
    Environment = "dev"
  }
  
  tags = {
    location = "office1"
    purpose  = "terraform-deployment"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0.0 |
| <a name="requirement_bloxone"></a> [bloxone](#requirement\_bloxone) | >= 1.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 5.0.0 |
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | >= 1.1.0 |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_instance.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [bloxone_infra_join_token.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_join_token) | resource |
| [bloxone_infra_service.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service) | resource |
| [random_uuid.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/uuid) | resource |
| [aws_ami.bloxone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami) | data source |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ami"></a> [ami](#input\_ami) | The AMI to use for the BloxOne Host. If not provided, the latest AMI will be used. | `string` | `null` | no |
| <a name="input_aws_instance_tags"></a> [aws\_instance\_tags](#input\_aws\_instance\_tags) | The tags to use for the AWS EC2 instance. For tags to use in BloxOne resources, use `tags`. | `map(string)` | `{}` | no |
| <a name="input_instance_type"></a> [instance\_type](#input\_instance\_type) | The instance type to use for the BloxOne Host. Infoblox recommends you choose an instance type that has minimum resources of 8 CPU and 16 GB of RAM. | `string` | `"c5a.2xlarge"` | no |
| <a name="input_join_token"></a> [join\_token](#input\_join\_token) | The join token to use for the BloxOne Host. If not provided, a join token will be created. | `string` | `null` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | The key name to use for EC2 instance. The key must be in the same region as the subnet. | `string` | n/a | yes |
| <a name="input_services"></a> [services](#input\_services) | The services to provision on the BloxOne Host. The services must be a map of valid service type with values of "start" or "stop". Valid service types are "dhcp", "dns", "anycast", "dfp". | `map(string)` | n/a | yes |
| <a name="input_subnet_id"></a> [subnet\_id](#input\_subnet\_id) | The subnet to use for the EC2 instance. The subnet must be in the same VPC as the security group. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | The tags to use for the BloxOne Host. For tags to use in AWS EC2, use `aws_tags`. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br/>    create = string<br/>    update = string<br/>    read   = string<br/>  })</pre> | `null` | no |
| <a name="input_vpc_security_group_ids"></a> [vpc\_security\_group\_ids](#input\_vpc\_security\_group\_ids) | The security group to use for EC2 instance. The security group must be in the same VPC as the subnet. | `list(string)` | n/a | yes |
| <a name="input_wait_for_state"></a> [wait\_for\_state](#input\_wait\_for\_state) | If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API. | `bool` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_instance"></a> [aws\_instance](#output\_aws\_instance) | The AWS instance object for the instance |
| <a name="output_host"></a> [host](#output\_host) | The `bloxone_infra_host` object for the instance |
| <a name="output_services"></a> [services](#output\_services) | The `bloxone_infra_service` objects for the instance. May be empty if no services were specified. |
<!-- END_TF_DOCS -->
