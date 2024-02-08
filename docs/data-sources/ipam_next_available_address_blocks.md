---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_ipam_next_available_address_blocks Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves the next available address blocks in the specified address block.
---

# bloxone_ipam_next_available_address_blocks (Data Source)

Retrieves the next available address blocks in the specified address block.

## Example Usage

```terraform
data "bloxone_ipam_address_blocks" "example_by_attribute" {
  filters = {
    name = "example_address_block"
  }
}

// 'address_block_count' allows you to get the number of next available address blocks in the address block specified by 'id'
// If not defined, count would default to 1
data "bloxone_ipam_next_available_address_blocks" "example_next_available_ab" {
  id                  = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
  address_block_count = 5
}

data "bloxone_ipam_next_available_address_blocks" "example_next_available_ab_default_count" {
  id = data.bloxone_ipam_address_blocks.example_by_attribute.results.0.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cidr` (Number) The cidr value of address blocks to be created.
- `id` (String) An application specific resource identity of a resource.

### Optional

- `address_block_count` (Number) Number of address blocks to generate. Default 1 if not set.

### Read-Only

- `results` (List of String) List of next available address block's addresses in the specified resource.