---
page_title: "Managing DHCP service with the BloxOne Terraform Provider"
subcategory: "Guides"
description: |-
  This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage IPAM and DHCP resources.
---

# Managing DHCP service with the BloxOne Terraform Provider

This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage IPAM and DHCP resources.

## Configuring the Provider

The provider needs to be configured with an API key and the URL of the Infoblox Cloud Services Portal (CSP). You can get the API Key from the Infoblox Cloud Services Portal (CSP) by following the steps outlined in this guide - [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

Create a directory for the Terraform configuration and create a file named `main.tf` with the following content:

````terraform
terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = ">= 0.1.0"
    }
  }
  
  required_version = ">= 1.0.0"
}

provider "bloxone" {
  csp_url = "https://csp.infoblox.com"
  api_key = "<BloxOne API Key>"
  default_tags = {
	managed_by = "terraform"
  }
}
````

!> Warning: Hard-coded credentials are not recommended in any configuration file. It is recommended to use environment variables.

You can also use the following environment variables to configure the provider:
`BLOXONE_CSP_URL` and `BLOXONE_API_KEY`.

Initialize the provider by running the following command. This will download the provider and initialize the working directory.

```shell
terraform init
```

## Configuring Resources

### IPAM and DHCP Resources
In this example, you will use the following resources to create an IP space, address block, and subnet.

- [bloxone_ipam_ip_space](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/ipam_ip_space)
- [bloxone_ipam_address_block](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/ipam_address_block)
- [bloxone_ipam_subnet](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/ipam_subnet)

Add the following to the `main.tf` file:

````terraform
// Create an IP space
resource "bloxone_ipam_ip_space" "this" {
  name = "example"
}

// Create an address block within the IP space
resource "bloxone_ipam_address_block" "this" {
  space   = bloxone_ipam_ip_space.this.id
  address = "10.0.0.0"
  cidr    = "16"
}

// Create a subnet within the address block
resource "bloxone_ipam_subnet" "this" {
  space   = bloxone_ipam_ip_space.this.id
  address = "10.0.0.0"
  cidr    = "24"
}

````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

Further, you will create a range and a fixed address reservation within the subnet.

You will use the following resources to create these
- [bloxone_ipam_range](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/ipam_range)
- [bloxone_dhcp_fixed_address](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dhcp_fixed_address)

You will use the following data sources to get the option codes in the default DHCPv4 option space
- [bloxone_dhcp_option_spaces](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dhcp_option_spaces)
- [bloxone_dhcp_option_codes](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dhcp_option_codes)

Add the following code to your main.tf:

````terraform
// Get the default option space for DHCPv4
data "bloxone_dhcp_option_spaces" "dhcp4" {
  filters = {
    name = "dhcp4"
  }
}

// Get the option codes for the default option space
data "bloxone_dhcp_option_codes" "dhcp4" {
  filters = {
    option_space = data.bloxone_dhcp_option_spaces.dhcp4.results.0.id
  }
}

locals {
  // Create a map of option code names to option code IDs
  // This will be used to look up the option code ID by name
  dhcp4_option_code_lookup = zipmap(data.bloxone_dhcp_option_codes.dhcp4.results[*].name, data.bloxone_dhcp_option_codes.dhcp4.results[*].id)
}

// Create a range within the subnet
resource "bloxone_ipam_range" "this" {
  space = bloxone_ipam_ip_space.this.id
  start = "10.0.0.5"
  end   = "10.0.0.100"
  dhcp_options = [
    {
      option_code  = local.dhcp4_option_code_lookup["domain-name-servers"]
      option_value = "10.0.0.1"
      type         = "option"
    },
    {
      option_code  = local.dhcp4_option_code_lookup["domain-name"]
      option_value = "domain.com"
      type         = "option"
    }
  ]

  depends_on = [bloxone_ipam_subnet.this]
}

// Create a Fixed Address within the subnet
resource "bloxone_dhcp_fixed_address" "this" {
  ip_space      = bloxone_ipam_ip_space.this.id
  address       = "10.0.0.6"
  match_type    = "mac"
  match_value   = "00:11:22:33:44:55"
  dhcp_options  = [
    {
      option_code  = local.dhcp4_option_code_lookup["time-offset"]
      option_value = "-5"
      type         = "option"
    }
  ]
  
  depends_on = [bloxone_ipam_subnet.this]
}
````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

### BloxOne Host on AWS with DHCP service

As a final step, you will also configure a BloxOne Host on AWS with DHCP service. 
You will use the following module to create these
- [bloxone_infra_host_aws](https://github.com/infobloxopen/terraform-provider-bloxone/tree/master/modules/bloxone_infra_host_aws)

The module requires the [AWS terraform provider](https://registry.terraform.io/providers/hashicorp/aws/latest) to be configured.
To configure the AWS provider, add the following code to your main.tf:

````terraform
provider "aws" {
  region     = "us-west-2"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}
```` 

!> Warning: Hard-coded credentials are not recommended in any configuration file. It is recommended to use the AWS credentials file or environment variables. 

You can also use the following environment variables to configure the provider:
`AWS_REGION`, `AWS_ACCESS_KEY_ID`, and `AWS_SECRET_ACCESS_KEY`.

To create an EC2 instance with DHCP service, you will need to have the following information:
- key_name: The name of the key pair to use for the instance
- subnet_id: The ID of the subnet to launch the instance into
- vpc_security_group_ids: A list of security group IDs to associate with the instance

Add the following code to your main.tf to create an EC2 instance with DHCP service:

````terraform

// Create a BloxOne Host on AWS with DHCP service
module "bloxone_infra_host_aws" {
  source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_aws"
  
  key_name               = "my-key"
  subnet_id              = "subnet-id"
  vpc_security_group_ids = ["vpc-security-group-id"]

  services = {
    dhcp = "start"
  }
}
````

You will need the ID for the DHCP host to assign the subnet to the BloxOne Host. 
To get the ID, you can use the [bloxone_dhcp_hosts](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dhcp_hosts) data source. 
Add the following code to your main.tf:

````terraform
data "bloxone_dhcp_hosts" "this" {
  filters = {
    name = module.bloxone_infra_host_aws.host.display_name
  }

  retry_if_not_found  = true
  depends_on          = [module.bloxone_infra_host_aws]
}
````
The `retry_if_not_found` attribute is set to true to allow the data source to retry if the host is not found immediately. The `depends_on` attribute is used to ensure that the data source is read after the BloxOne Host is created.


You will also have to modify the `bloxone_ipam_subnet` resource to assign the BloxOne Host to serve the subnet. To do this, replace the `bloxone_ipam_subnet` resource with the following code:

````terraform
resource "bloxone_ipam_subnet" "this" {
  space     = bloxone_ipam_ip_space.this.id
  address   = "10.0.0.0"
  cidr      = "24"
  dhcp_host = one(data.bloxone_dhcp_hosts.this.results).id
}
````

Here, the `dhcp_host` attribute has been added to the subnet resource and set to the ID of the DHCP host.

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

## Applying the Configuration

To create the resources, run the following command:

```shell
terraform apply
```

## Next steps

You can also use the BloxOne Terraform Provider to manage other resources such as DNS zones, DNS records. For more information, see the [BloxOne Terraform Provider documentation](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs).
