# Terraform Provider for Infoblox BloxOne

The Terraform Provider for Infoblox BloxOne allows you to manage your Infoblox BloxOne resources such as DNS records, DHCP/IPAM configurations, Threat Defense policies, and more using Terraform. This provider uses the [universal-ddi-go-client](https://github.com/infobloxopen/universal-ddi-go-client) for all API calls to interact with the Infoblox BloxOne Cloud Services Portal (CSP).


## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
    - [Authentication](#authentication)
    - [Installation](#installation)
- [Usage Examples](#usage-examples)
- [Available Resources and DataSources](#available-resources-and-datasources)
- [Importing Existing Resources](#importing-existing-resources)
- [Unified Nameservers](#unified-nameservers)
- [Documentation](#documentation)
- [Debugging](#logging-and-debugging)
- [Support](#support)

## Requirements

- [Go](https://golang.org/doc/install) >= 1.25.1
- [Terraform](https://www.terraform.io/downloads.html) >= 1.12.1
- An active [Infoblox BloxOne](https://www.infoblox.com/products/bloxone-ddi/) subscription

## Version Compatibility Matrix

| Provider Version | Go Version | Terraform Version |
|-----------------|------------|-------------------|
| 1.6.0           | >= 1.25.1  | >= 1.12.1         |
| 1.5.x and earlier | >= 1.21.0 | >= 1.5.0         |

> **Note:** Starting from v1.6.0, the provider was migrated from the `bloxone-go-client` SDK to the `universal-ddi-go-client` SDK. Environment variable names have changed — see the [migration notes](#migration-notes) below.

**Important Notes:**
- **v1.6.0+** uses `INFOBLOX_PORTAL_URL` and `INFOBLOX_PORTAL_KEY` environment variables.
- **v1.5.x and earlier** used `BLOXONE_CSP_URL` and `BLOXONE_API_KEY`. These are still supported as deprecated fallbacks but will be removed in a future release.
- For migration from the B1DDI provider, refer to the [Migration Guide](docs/guides/migration.md).

## Getting Started

### Authentication

To use the BloxOne Terraform Provider, you need a BloxOne API key. You can configure it in two ways:

1. **Environment Variables** (recommended):
   ```shell
   export INFOBLOX_PORTAL_URL="https://csp.infoblox.com"
   export INFOBLOX_PORTAL_KEY="<your-api-key>"
   ```

2. **Provider Configuration**:
   ```terraform
   provider "bloxone" {
     csp_url = "https://csp.infoblox.com"
     api_key = "<your-api-key>"

     # Optional: default tags applied to all resources
     default_tags = {
       managed_by = "terraform"
       site       = "Site A"
     }
   }
   ```

For instructions on generating an API key, see [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

### Installation

To use the provider in your Terraform configuration, add the following `required_providers` block:

```terraform
terraform {
  required_providers {
    bloxone = {
      source  = "infobloxopen/bloxone"
      version = ">= 1.6.0"
    }
  }
}
```

Then run:

```shell
terraform init
```

For step-by-step guides, see:
- [Quickstart: Managing DNS](docs/guides/quickstart-dns.md)
- [Quickstart: Managing DHCP and IPAM](docs/guides/quickstart-dhcp.md)

## Usage Examples

Detailed examples for each resource and data source are available in the `examples` directory of the repository. Each resource and data source has its own directory with sample configurations.

For example:
- Resources examples: [`examples/resources/bloxone_*`](examples/resources/)
- Data sources examples: [`examples/data-sources/bloxone_*`](examples/data-sources/)
- Provider configuration: [`examples/provider/`](examples/provider/)

Please refer to these examples for detailed usage patterns and configurations.

## Available Resources and DataSources

The object groups available in this provider are categorized as follows:

- **DHCP** — DHCP servers, hosts, HA groups, option codes/groups/spaces, fixed addresses
- **DNS** — Auth zones, forward zones, views, ACLs, NSGs, delegations, and all record types (A, AAAA, CAA, CNAME, DNAME, HTTPS, MX, NAPTR, NS, PTR, SRV, SVCB, TXT)
- **IPAM** — Address blocks, subnets, IP addresses, ranges
- **Anycast** — Anycast configurations and hosts
- **Cloud Discovery** — Cloud discovery providers
- **DFP** — DNS Forwarding Proxy services
- **Federation** — Federated blocks and realms
- **Infrastructure** — Infrastructure hosts and join tokens
- **Keys** — TSIG and KSK/ZSK key management
- **Threat Defense** — Access policies, application filters, category filters, custom lists, internal domains, named lists, network lists, security policies

For a full list of available resources and data sources, refer to the [Terraform Registry documentation](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs).

## Importing Existing Resources

Existing BloxOne resources can be imported into Terraform state using their resource ID. For example:

```shell
terraform import bloxone_dns_auth_zone.example <resource-id>
```

Each resource's documentation page lists the import syntax specific to that resource.

## Migration Notes

If you are upgrading from an earlier version or migrating from the B1DDI provider:

- Refer to the [Migration Guide](docs/guides/migration.md) for step-by-step upgrade instructions.
- After upgrading to v1.6.0+, update environment variable names from `BLOXONE_CSP_URL`/`BLOXONE_API_KEY` to `INFOBLOX_PORTAL_URL`/`INFOBLOX_PORTAL_KEY`.

## Unified Nameservers

### Deprecation Notice: DNS Server Group (NSG) Assignment

As part of the Unified Nameservers initiative, the following DNS Server Group configurations are deprecated and will no longer be allowed. Requests that use them will be rejected:

- **Nested DNS Server Groups** — A DNS Server Group cannot be nested inside an authoritative DNS Server Group (AuthNSG).
- **Dual-role DNS Server Groups** — The same DNS Server Group cannot be assigned to both a primary zone and a secondary zone.
- **Multiple DNS Server Groups on an authoritative zone** — Only one DNS Server Group may be assigned to an authoritative zone (AuthZone).
- **DNS Server Groups combined with internal secondaries** — A DNS Server Group and internal secondaries cannot coexist on an authoritative zone.

Review your Terraform configurations for any of the above patterns and update them before they are enforced.

## Documentation

Full provider documentation, including schema references for all resources and data sources, is available on the [Terraform Registry](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs).

Local documentation is available in the [`docs/`](docs/) directory.

## Logging and Debugging

To enable detailed logging for troubleshooting, set the `TF_LOG` environment variable:

```shell
export TF_LOG=DEBUG
terraform apply
```

For provider-level logging, set:

```shell
export TF_LOG_PROVIDER=DEBUG
```

Logs are written to stderr by default. To redirect to a file:

```shell
export TF_LOG_PATH=./terraform.log
```

## Support

If you have any questions or issues, you can reach out using the following channels:

- **GitHub Issues**: Submit bugs or enhancement requests on the [GitHub Issues Page](https://github.com/infobloxopen/terraform-provider-bloxone/issues)
- **Infoblox Support**: Contact [Infoblox Support](https://info.infoblox.com/contact-form/) for product-related questions
