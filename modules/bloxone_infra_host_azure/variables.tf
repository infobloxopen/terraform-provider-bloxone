variable "join_token" {
  description = "The join token to use for the BloxOne Host. If not provided, a join token will be created."
  type        = string
  default     = null
}

variable "vm_name" {
  description = "The name of the virtual machine"
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group in which the resources will be created"
  type        = string
}

variable "location" {
  description = "The location where the resources will be created"
  type        = string
}

variable "vm_size" {
  description = "Size of the Virtual Machine based on Azure sizing"
  type        = string
  default     = "Standard_F8s"
}

variable "vm_network_interface_ids" {
  description = "The network interface ids that will be associated to the BloxOne Host"
  type        = list(string)
}

variable "tags" {
  description = "The tags to use for the BloxOne Host."
  type        = map(string)
  default     = {}
}

variable "azure_instance_tags" {
  description = "The tags to use for the Azure virtual machine."
  type        = map(string)
  default     = {}
}

variable "os_disk_caching" {
  description = "The caching type to use for the OS disk."
  type        = string
  default     = "ReadWrite"
}

variable "os_disk_storage_account_type" {
  description = "The storage account type to use for the OS disk."
  type        = string
  default     = "Standard_LRS"
}

variable "source_image_reference_publisher" {
  description = "The publisher of the image that you want to deploy"
  type        = string
  default     = "infoblox"
}

variable "source_image_reference_offer" {
  description = "The offer of the image that you want to deploy"
  type        = string
  default     = "infoblox-bloxone-34"
}

variable "source_image_reference_sku" {
  description = "The sku of the image that you want to deploy"
  type        = string
  default     = "infoblox-bloxone"
}

variable "source_image_reference_version" {
  description = "The version of the image that you want to deploy."
  type        = string
  default     = "3.8.1"
}

variable "plan_name" {
  description = "The name of the plan to use for the BloxOne Host."
  type        = string
  default     = "infoblox-bloxone"
}

variable "plan_product" {
  description = "The product to use for the BloxOne Host."
  type        = string
  default     = "infoblox-bloxone-34"
}

variable "plan_publisher" {
  description = "The publisher to use for the BloxOne Host."
  type        = string
  default     = "infoblox"
}

variable "admin_username" {
  description = "The username to use for the BloxOne Host."
  type        = string
  default     = "infobloxadmin"
}

variable "ssh_public_key_path" {
  description = "The path to the SSH public key to use for the BloxOne Host."
  type        = string
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
  type = object({
    create = string
    update = string
    read   = string
  })
  default = null
}

variable "wait_for_state" {
  description = "If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API."
  type        = bool
  default     = null
}