---
page_title: "Managing DNS service with the BloxOne Terraform Provider"
subcategory: "Guides"
description: |-
  This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage DNS resources.
---

# Managing DNS service with the BloxOne Terraform Provider

This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage DNS resources.

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

### Authoritative Zone


In this example, you will use the following resources to create an authoritative zone.

- [bloxone_dns_auth_zone](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dns_auth_zone)

Add the following to the `main.tf` file:

````terraform
// Create a DNS zone for the domain
resource "bloxone_dns_auth_zone" "this" {
  fqdn         = "domain.com."
  primary_type = "cloud"
}

````

Here the `view` attribute has not been set, so the default view will be used.

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

Further, you will create an A record and a CNAME record within the subnet.

You will use the following resources to create these
- [bloxone_dns_a_record](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dns_a_record)
- [bloxone_dns_cname_record](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dns_cname_record)

Add the following code to your main.tf:

````terraform
// Create an A record
resource "bloxone_dns_a_record" "this" {
  zone = bloxone_dns_auth_zone.this.id
  name_in_zone = "host"
  rdata = {
    address = "10.0.0.10"
  }
}


// Create a CNAME record
resource "bloxone_dns_cname_record" "this" {
  zone = bloxone_dns_auth_zone.this.id
  name_in_zone = "alias"
  rdata = {
    cname = "host.domain.com."
  }
}
````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

### BloxOne Host on AWS with DNS service

As a final step, you will also configure a BloxOne Host on AWS with DNS service. 
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

To create an EC2 instance with DNS service, you will need to have the following information:
- key_name: The name of the key pair to use for the instance
- subnet_id: The ID of the subnet to launch the instance into
- vpc_security_group_ids: A list of security group IDs to associate with the instance

Add the following code to your main.tf to create an EC2 instance with DNS service:

````terraform

// Create a BloxOne Host on AWS with DNS service
module "bloxone_infra_host_aws" {
  source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_aws"
  
  key_name               = "my-key"
  subnet_id              = "subnet-id"
  vpc_security_group_ids = ["vpc-security-group-id"]

  services = {
    dns = "start"
  }
}
````

You will need the ID for the DNS host to assign the subnet to the BloxOne Host. 
To get the ID, you can use the [bloxone_dns_hosts](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/dns_hosts) data source. 
Add the following code to your main.tf:

````terraform
data "bloxone_dns_hosts" "this" {
  filters = {
    name = module.bloxone_infra_host_aws.host.display_name
  }

  retry_if_not_found  = true
  depends_on          = [module.bloxone_infra_host_aws]
}
````
The `retry_if_not_found` attribute is set to true to allow the data source to retry if the host is not found immediately. The `depends_on` attribute is used to ensure that the data source is read after the BloxOne Host is created.


You will also have to modify the `bloxone_dns_auth_zone` resource to assign the BloxOne Host to serve the zone. To do this, replace the `bloxone_dns_auth_zone` resource with the following code:

````terraform
resource "bloxone_dns_auth_zone" "this" {
  fqdn         = "domain.com."
  primary_type = "cloud"

  internal_secondaries = [
    {
      host = one(data.bloxone_dns_hosts.this.results).id
    }
  ]
}
````

Here, the `internal_secondaries.host` attribute has been added to the zone resource and set to the ID of the DNS host.

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

You can also use the BloxOne Terraform Provider to manage other resources such as ACLs, NSGs, Subnets, Fixed Addresses. For more information, see the [BloxOne Terraform Provider documentation](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs).
