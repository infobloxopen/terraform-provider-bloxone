---
page_title: "Managing DHCP service with the BloxOne Terraform Provider"
subcategory: "Guides"
description: |-
  This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage IPAM and DHCP resources.
---

# Managing Policy Based DFP service using the BloxOne Terraform Provider

This guide provides step-by-step instructions for using the BloxOne Terraform Provider to manage Security Policies and various Threat Defense objects associated with it.

## Configuring the Provider

The provider needs to be configured with an API key and the URL of the Infoblox Cloud Services Portal (CSP). You can get the API Key from the Infoblox Cloud Services Portal (CSP) by following the steps outlined in this guide - [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

Create a directory for the Terraform configuration and create a file named `main.tf` with the following content:

````terraform
terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = ">= 1.0.0"
    }
  }
  required_version = ">= 1.5.0"
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

### BloxOne Host on AWS with DFP service

As the first step, you will also configure a BloxOne Host on AWS with DFP service.
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
    dfp = "start"
  }
}
````

You will need the pool ID of the AWS host to create the Infra Service block for DFP. `explain how to get the pool ID`
To create the Infra service block , we use the following resource :
- [bloxone_infra_service](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service)

`Add the following code to your main.tf:

````terraform
resource "bloxone_infra_service" "example" {
  name         = "example_dfp_service"
  pool_id      = data.bloxone_infra_hosts.dfp_host.results.0.pool_id
  service_type = "dfp"
  desired_state = "start"
  wait_for_state = false
}
````

`explain all 2 blocks below`
Further , we define the following:
- [bloxone_td_internal_domain_list](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/td_internal_domain_list)
- [bloxone_dfp_service](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/dfp_service)


````terraform
resource "bloxone_td_internal_domain_list" "example_list" {
  name = "example_internal_domain_list"
  internal_domains = ["example.domain.com"]
}

# Create the DFP Service
resource "bloxone_dfp_service" "example" {
  service_id = bloxone_infra_service.example.id

  # Other optional fields
  internal_domain_lists = [bloxone_td_internal_domain_list.example_list.id]
  resolvers_all = [
    {
      address = "1.1.1.1"
      is_fallback = true
      is_local = false
      protocols = ["DO53"]
    }
  ]
}
````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

### IPAM and DHCP Resources
In this example, you will use the following resources to create a Named/Custom List, Access/Bypass Code, and a Network List/External Network.
`add alternate names for the resources below`
- [bloxone_td_named_list](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/td_named_list)
- [bloxone_td_access_code](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/td_access_code)
- [bloxone_td_network_list](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/td_network_list)

Add the following to the `main.tf` file:

````terraform
# Create the Named List
resource "bloxone_td_named_list" "example" {
  name = "example_named_list"
  items_described = [
    {
      item        = "tf-domain.com"
      description = "Example Domain"
    }
  ]
  type = "custom_list"
}

# Create the Access Code using the Named List
resource "bloxone_td_access_code" "example" {
  name       = "example_access_code"
  activation = timestamp()
  expiration = timeadd(timestamp(), "24h")
  rules = [
    {
      data = bloxone_td_named_list.example.name,
      type = bloxone_td_named_list.example.type
    }
  ]
  # Other optional fields
  description = "Example Access Code"
}

# Create the Network List
resource "bloxone_td_network_list" "example" {
  name  = "example_network_list"
  items = ["156.2.3.0/24"]

  # Other optional fields
  description = "Example Network List"
}


````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

Finally, you will create the Security Policy that uses Named List, Access Code, and Network List created earlier.

- [bloxone_td_security_policy](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/td_security_policy)

Add the following code to your main.tf:

````terraform
# Create the Security Policy using the Named List, Network List, and Access Code
resource "bloxone_td_security_policy" "example" {
  name = "example_security_policy"

  # Other optional fields
  rules = [
    {
      action = "action_allow",
      data   = bloxone_td_named_list.example.name,
      type   = bloxone_td_named_list.example.type
    }
  ]
  description    = "Example Security Policy"
  dfps = [bloxone_dfp_service.example.id]
  ecs            = true
  onprem_resolve = true
  safe_search    = false
  tags = {
    site = "Site A"
  }
  network_lists = [bloxone_td_network_list.example.id]
  access_codes  = [bloxone_td_access_code.example.id]
}
````

`explain everthing above`

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

You can also use the BloxOne Terraform Provider to manage other resources such as DNS and DHCP/IPAM resources. For more information, see the [BloxOne Terraform Provider documentation](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs).
