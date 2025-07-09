<!-- BEGIN_TF_DOCS -->
# Terraform Module to create BloxOne Host in GCP

This module will provision a GCP virtual machine that uses a BloxOne image.
The virtual machine will be configured to join a BloxOne Cloud Services Platform (CSP) with the provided join token.
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
      source = "infobloxopen/bloxone"
    # Other parameters
    }
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  project     = "<gcp-project-id>"
  credentials = file("<path-to-service-account-key>.json")  
  region      = "selected-region"      
  zone        = "selected-zone" 
}

provider "bloxone" {
  csp_url = "<csp-url>" 
  api_key = "<api-key>"           
}

module "bloxone_infra_host_gcp" {
  source = "github.com/infobloxopen/terraform-provider-bloxone//modules/bloxone_infra_host_gcp"

  name         = "bloxone-vm"
  source_image = "bloxone-v381"

  machine_type = "machine-type"

  network_interfaces = [
    {
      network          = "gcp-external-network"
      subnetwork       = "gcp-external-subnet"
      assign_public_ip = true
    },
    {
      network          = "gcp-internal-network"
      subnetwork       = "gcp-internal-subnet"
    }
  ]

  gcp_instance_labels = {
    environment = "dev"
  }

  gcp_disk_labels = {
    environment = "dev"
    module      = "bloxone"
  }

  tags = {
    location = "office1"
  }

  metadata = {
    purpose = "IaaC-testing"
  }

  services = {
    dhcp = "start"
    dns = "start"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5 |
| <a name="requirement_bloxone"></a> [bloxone](#requirement\_bloxone) | >= 1.1.0 |
| <a name="requirement_google"></a> [google](#requirement\_google) | >= 3.5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_bloxone"></a> [bloxone](#provider\_bloxone) | >= 1.1.0 |
| <a name="provider_google"></a> [google](#provider\_google) | >= 3.5.0 |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Resources

| Name | Type |
|------|------|
| [bloxone_infra_join_token.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_join_token) | resource |
| [bloxone_infra_service.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/resources/infra_service) | resource |
| [google_compute_instance.this](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance) | resource |
| [random_uuid.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/uuid) | resource |
| [bloxone_infra_hosts.this](https://registry.terraform.io/providers/infobloxopen/bloxone/latest/docs/data-sources/infra_hosts) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_deletion_protection"></a> [deletion\_protection](#input\_deletion\_protection) | Whether the BloxOne Host should have deletion protection enabled. | `bool` | `false` | no |
| <a name="input_disk_size"></a> [disk\_size](#input\_disk\_size) | The size of the data disk in GB. | `number` | `59` | no |
| <a name="input_disk_type"></a> [disk\_type](#input\_disk\_type) | The type of the data disk. | `string` | `"pd-standard"` | no |
| <a name="input_gcp_instance_labels"></a> [gcp\_instance\_labels](#input\_gcp\_instance\_labels) | The labels to associate with the virtual machine. For `tags` to be used for the BloxOne Host, use the `tags` variable. | `map(string)` | `{}` | no |
| <a name="input_join_token"></a> [join\_token](#input\_join\_token) | The join token to use for the BloxOne Host. If not provided, a join token will be created. | `string` | `null` | no |
| <a name="input_machine_type"></a> [machine\_type](#input\_machine\_type) | The machine type to use for the virtual machine | `string` | `"e2-standard-4"` | no |
| <a name="input_name"></a> [name](#input\_name) | The name of the virtual machine | `string` | n/a | yes |
| <a name="input_network_interfaces"></a> [network\_interfaces](#input\_network\_interfaces) | List of network interfaces to be attached to the virtual machine. | <pre>list(object({<br/>    network          = string<br/>    subnetwork       = string<br/>    assign_public_ip = optional(bool)<br/>  }))</pre> | n/a | yes |
| <a name="input_service_account"></a> [service\_account](#input\_service\_account) | The service account to use for the BloxOne Host. | <pre>object({<br/>    email  = string<br/>    scopes = list(string)<br/>  })</pre> | `null` | no |
| <a name="input_services"></a> [services](#input\_services) | The services to provision on the BloxOne Host. The services must be a map of valid service type with values of "start" or "stop". Valid service types are "dhcp", "dns", "anycast", "dfp". | `map(string)` | n/a | yes |
| <a name="input_source_image"></a> [source\_image](#input\_source\_image) | The source image to use for the virtual machine. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | The tags to use for the BloxOne Host. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). If not provided, the default timeouts will be used. | <pre>object({<br/>    create = string<br/>    update = string<br/>    read   = string<br/>  })</pre> | `null` | no |
| <a name="input_wait_for_state"></a> [wait\_for\_state](#input\_wait\_for\_state) | If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API. | `bool` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_google_compute_instance"></a> [google\_compute\_instance](#output\_google\_compute\_instance) | The GCP instance object for the instance |
| <a name="output_host"></a> [host](#output\_host) | The `bloxone_infra_host` object for the instance |
| <a name="output_services"></a> [services](#output\_services) | The `bloxone_infra_service` objects for the instance. May be empty if no services were specified. |
<!-- END_TF_DOCS -->
