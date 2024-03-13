variable "join_token" {
  description = "The join token to use for the BloxOne Host. If not provided, a join token will be created."
  type = string
  default = null
}

variable "ami" {
  description = "The AMI to use for the BloxOne Host. If not provided, the latest AMI will be used."
  type = string
  default = null
}

variable "instance_type" {
  description = "The instance type to use for the BloxOne Host. Infoblox recommends you choose an instance type that has minimum resources of 8 CPU and 16 GB of RAM."
  default     = "c5a.2xlarge"
  type        = string
}

variable "subnet_id" {
  description = "The subnet to use for the EC2 instance. The subnet must be in the same VPC as the security group."
  type        = string
}

variable "vpc_security_group_ids" {
  description = "The security group to use for EC2 instance. The security group must be in the same VPC as the subnet."
  type        = list(string)
}

variable "key_name" {
  description = "The key name to use for EC2 instance. The key must be in the same region as the subnet."
  type        = string
}

variable "tags" {
  description = "The tags to use for the BloxOne Host. For tags to use in AWS EC2, use `aws_tags`."
  type        = map(string)
  default     = {}
}

variable "aws_instance_tags" {
  description = "The tags to use for the AWS EC2 instance. For tags to use in BloxOne resources, use `tags`."
  type        = map(string)
  default     = {}
}

variable "services" {
  description = "The services to provision on the BloxOne Host. The services must be a map of valid service type with values of \"start\" or \"stop\". Valid service types are \"dhcp\" and \"dns\"."
  type        = map(string)
  validation {
    condition     = length(keys(var.services)) == length([for k in keys(var.services) : k if contains(["dhcp", "dns"], k)]) && alltrue([for v in values(var.services) : contains(["start", "stop"], v)])
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
