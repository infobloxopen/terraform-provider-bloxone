variable "hosts" {
  description = "Map of hostnames with their roles, routing protocols, BGP, and OSPF configurations."
  type = map(object({
    role               = string
    routing_protocols  = list(string)
    bgp_config = optional(object({
      asn            = string
      holddown_secs  = number
      neighbors      = list(object({
        asn        = string
        ip_address = string
      }))
    }))
    ospf_config = optional(object({
      area                = string
      area_type           = string
      authentication_type = string
      interface           = string
      authentication_key  = string
      hello_interval      = number
      dead_interval       = number
      retransmit_interval = number
      transmit_delay      = number
    }))
  }))
}

variable "ha_name" {
  description = "Name of the HA group."
  type        = string
  default = null
}

variable "service" {
  description = "The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`)."
  type    = string
  default = "dhcp"
}

variable "anycast_ip_address" {
  description = "Anycast IP address."
  type    = string
}

variable "anycast_config_name" {
  description = "Name of the Anycast configuration."
  type    = string
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