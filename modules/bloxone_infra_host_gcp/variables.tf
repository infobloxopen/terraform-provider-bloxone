variable "join_token" {
  description = "The join token to use for the BloxOne Host. If not provided, a join token will be created."
  type        = string
  default     = null
}

variable "name" {
  description = "The name of the virtual machine"
  type        = string
}

variable "machine_type" {
  description = "The machine type to use for the virtual machine"
  type        = string
  default     = "e2-standard-4"
}

variable "network_interfaces" {
  description = "List of network interfaces to be attached to the virtual machine."
  type        = list(object({
    network          = string
    subnetwork       = string
    assign_public_ip = optional(bool)
  }))
}

variable "disk_type" {
  description = "The type of the data disk."
  type        = string
  default     = "pd-standard"
}

variable "disk_size" {
  description = "The size of the data disk in GB."
  type        = number
  default     = 59
}

variable "source_image" {
  description = "The source image to use for the virtual machine."
  type        = string
}

variable "gcp_instance_labels" {
  description = "The labels to associate with the virtual machine. For `tags` to be used for the BloxOne Host, use the `tags` variable."
  type        = map(string)
  default     = {}
}

variable "tags" {
  description = "The tags to use for the BloxOne Host."
  type        = map(string)
  default     = {}
}

variable "service_account" {
  description = "The service account to use for the BloxOne Host."
  type        = object({
    email  = string
    scopes = list(string)
  })
  default = null
}

variable "deletion_protection" {
  description = "Whether the BloxOne Host should have deletion protection enabled."
  type        = bool
  default     = false
}

variable "services" {
  description = "The services to provision on the BloxOne Host. The services must be a map of valid service type with values of \"start\" or \"stop\". Valid service types are \"dhcp\", \"dns\", \"anycast\", \"dfp\"."
  type        = map(string)
  validation {
    condition     = length(keys(var.services)) == length([for k in keys(var.services) : k if contains(["dhcp", "dns", "anycast", "dfp"], k)]) && alltrue([for v in values(var.services) : contains(["start", "stop"], v)])
    error_message = "services must be a map of valid service type with values of start or stop"
  }
}

variable "timeouts" {
  description = "The timeouts to use for the BloxOne Host. The timeout value is a string that can be parsed as a duration consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). If not provided, the default timeouts will be used."
  type        = object({
    create = string
    update = string
    read   = string
  })
  default = null
}
