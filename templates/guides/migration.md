---
page_title: "Migrating from the B1DDI provider"
subcategory: "Guides"
description: |-
  This guide outlines the migration process from the B1DDI provider to the BloxOne provider.
---

# Migrating from the B1DDI provider

This guide covers the changes introduced in the BloxOne provider and outlines the steps you may need to take to upgrade your configuration.

The BloxOne provider replaces the [B1DDI provider](https://registry.terraform.io/providers/infobloxopen/b1ddi/latest) and is not backwards compatible. This means you will need to update your configuration to use the new provider.

## Backup

Before making changes to your state, it's a good idea to back up your state file. Any state modification commands made using the CLI will automatically create a backup. 
If you prefer to manually back up your state file, you can copy your `terraform.tfstate` file to a backup location.

Having a backup ensures that you have a snapshot of your infrastructure's state at a specific moment, allowing you to revert or refer to it if necessary.

## Add new provider

You will need the new provider to be added to your configuration. For example:

```terraform 
terraform {
    required_providers {
        bloxone = {
          source = "infobloxopen/bloxone"
          version = "0.1.0"
        }
    }
}
```

You will have to run `terraform init` to download the new provider.

## Replace resource types in configuration

The resource types have changed in the new provider. The following table shows the old and new resource types.

| B1DDI Provider        | BloxOne Provider           |
|-----------------------|----------------------------|
| b1ddi_ip_space        | bloxone_ipam_ip_space      |
| b1ddi_address_block   | bloxone_ipam_address_block |
| b1ddi_subnet          | bloxone_ipam_subnet        |
| b1ddi_range           | bloxone_ipam_range         |
| b1ddi_fixed_address   | bloxone_dhcp_fixed_address |
| b1ddi_address         | bloxone_ipam_address       |
| b1ddi_dns_view        | bloxone_dns_view           |
| b1ddi_dns_forward_nsg | bloxone_dns_forward_nsg    |
| b1ddi_dns_auth_zone   | bloxone_dns_auth_zone      |
| b1ddi_dns_auth_nsg    | bloxone_dns_auth_nsg       |
| b1ddi_dns_record      | bloxone_dns_record*        |

> NOTE: The _b1ddi_dns_record_ in the B1DDI provider used to be a single resource type that could be used to create any type of DNS record. 
> In the BloxOne provider, this has been split into multiple resource types, one for each record type. 
> For example, _b1ddi_dns_record_ is now _bloxone_dns_a_record_, _bloxone_dns_aaaa_record_, _bloxone_dns_caa_record_, etc. 
> 
> The _bloxone_dns_record_ is for resource records that are not supported explicitly by the provider. For example, if you want to create a DNS record of type _URI_, you can use _bloxone_dns_record_.

To migrate your configuration, you will need to replace the old resource types with the new resource types. For example:

```terraform
resource "b1ddi_ip_space" "example" {
    name = "example"
}
```
will become
```terraform
resource "bloxone_ipam_ip_space" example {
    name = "example"
}
```

## Replace resources in state

If your configuration was already in use, you will need to replace the provider in your state file. 
It is recommended that you start fresh and re apply your configuration to a new state file.

If you want to preserve your state, you will have to move the resources to the new provider with the new name. 
To do this you will have to remove the old resource from state, and import the new resource into state. 

-> Do not run `terraform plan` until you have moved all resources to the new provider in the state file.

To import the new resource, you will need the ID of all existing resource. You can use the `terraform show` command to get the IDs of all resources in your state file. For example:

#### Get Resource IDs
```shell
terraform show -json | jq -c '.values.root_module.resources[] | {"resource":.address, "id":.values.id}'
```

#### Remove old resource from state
To remove the old resource from state, you can use the `terraform state rm` command. For example:

```shell
terraform state rm b1ddi_ip_space.example
```

#### Import new resource into state
To import the new resource into state, you can use the `terraform import` command. For example:

```shell
terraform import bloxone_ipam_ip_space.example ipam/ip_space/5f26be86-abef-11ee-babd-a2f371a672c6
```
If you are using Terraform v1.5.0 or later, you can also use the [import block](https://developer.hashicorp.com/terraform/language/import) in your configuration to import the resource into state. For example:

```terraform 
import {
    id = "ipam/ip_space/5f26be86-abef-11ee-babd-a2f371a672c6"
    to = bloxone_ipam_ip_space.example
}
```

## Plan and Apply

Once you have replaced all the resources in your configuration and state, you can run `terraform plan` to see what changes will be made to your infrastructure.
There should be no changes to your infrastructure if you have replaced all the resources correctly.

In case there are changes, you will need to make the necessary changes to your configuration to match the new provider.
Some of the changes you may need to make are listed below.
 - **Unsupported block type**: Configuration written as blocks will have to be rewritten as values. For example, if you have a block like this:
    ```terraform
    internal_secondaries {
        host = "dns/host/989a0d20-c030-11ee-a93d-0b6e6ea305e3"
    }
    ```
    you will have to rewrite it with an equal sign :
    ```terraform
    internal_secondaries = [ 
      { 
        host = "dns/host/989a0d20-c030-11ee-a93d-0b6e6ea305e3"
      }
    ]
    ```
 - **Read Only attributes**: Some attributes are read only in the BloxOne provider. For example, the _type_ attribute in _bloxone_dns_a_record_ is read only. If you have a configuration like this:
    ```terraform
    resource "bloxone_dns_a_record" "example" {
        name = "domain.com"
        type = "A"
    }
    ```
    you will have to remove _type_ from the config as it is read only.
 - **Default values**: Some default values may have changed in the new provider. For example, the _hostname_rewrite_regex_ attribute in _bloxone_dns_forward_nsg_ has been changed from `[^a-zA-Z0-9.-]` to `[^a-zA-Z0-9_.]`. 
    If your configuration, like the one below, doesn't specify a value for _hostname_rewrite_regex_:
    ```terraform
    resource "bloxone_dns_forward_nsg" "example" {
        name = "example"
    }
    ```
    you'll need to add a _hostname_rewrite_regex_ value to your configuration. This value should match the existing or old default value, otherwise, the plan will indicate a change.
    ```terraform
    resource "bloxone_dns_forward_nsg" "example" {
        name = "example"
        hostname_rewrite_regex = "[^a-zA-Z0-9.-]"
    }
    ```
