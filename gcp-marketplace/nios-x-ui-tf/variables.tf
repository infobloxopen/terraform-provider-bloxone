variable "project_id" {
  description = "The ID of the project in which to provision resources."
  type        = string
}

// Marketplace requires this variable name to be declared
variable "goog_cm_deployment_name" {
  description = "The name of the deployment and VM instance."
  type        = string
}

variable "source_image" {
  description = "The image name for the disk for the VM instance."
  type        = string
  default     = "projects/infoblox-public-436917/global/images/nios-x-virtual-server-3-8-1"
}

variable "zone" {
  description = "(Optional) The zone that the machine should be created in. If it is not provided, the provider zone is used."
  type        = string
  default     = "us-central1-a"
}

variable "machine_type" {
  description = "The machine type to create, e.g. e2-small"
  type        = string
  default     = "n2-standard-8"
}

variable "boot_disk_type" {
  description = "The boot disk type for the VM instance."
  type        = string
  default     = "pd-ssd"
}

variable "boot_disk_size" {
  description = "The boot disk size for the VM instance in GBs"
  type        = number
  default     = 60
}

variable "networks" {
  description = "The network name to attach the VM instance."
  type        = list(string)
  default     = ["default"]
}

variable "sub_networks" {
  description = "The sub network name to attach the VM instance."
  type        = list(string)
  default     = []
}

variable "external_ips" {
  description = "The external IPs assigned to the VM for public access."
  type        = list(string)
  default     = ["EPHEMERAL"]
}

variable "join_token" {
  description = "(Required) The jointoken generated from Infoblox Portal NIOS-X Server authentication"
  type        = string
  default     = ""
}

variable "http_proxy" {
  description = "(Optional) HTTP Proxy of your environment"
  type        = string
  default     = ""
}

# Workaround as Deploy config UI doesn't support map(string) type
# "{\"key1\":\"value1\",\"key2\":\"value2\"}"
variable "nios_x_tags" {
  description = "The tags to use for the NIOS-X Server in Infoblox Portal."
  type        = string
  default     = "{\"DeployedFrom\":\"GCP-Marketplace\"}"
}

variable "enable_cloud_api" {
  description = "Allow full access to all of Google Cloud Platform APIs on the VM"
  type        = bool
  default     = true
}

# Workaround as Deploy config UI doesn't support map(string) type
# "{\"key1\":\"value1\",\"key2\":\"value2\"}"
variable "labels" {
  description = "(Optional) A map of key/value in string format to assign to the instance."
  type        = string
  default     = "{}"
}

variable "tags" {
  description = "(Optional) A list of network tags to attach to the instance."
  type        = list(string)
  default     = []
}

# Workaround as Deploy config UI doesn't support map(string) type
# "{\"key1\":\"value1\",\"key2\":\"value2\"}"
#  default     = {
#      google-logging-enable = "0"
#      google-monitoring-enable = "0"
#  }
variable "metadata" {
  description = "(Optional) Metadata key/value in string format to make available from within the VM instance."
  type        = string
  default     = "{\"google-logging-enable\":\"0\",\"google-monitoring-enable\":\"0\"}"
}
