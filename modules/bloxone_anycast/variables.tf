variable "hosts" {
  description = "Map of hostnames or IP addresses for the Anycast configuration."
  type        = map(object({
    role             = string,
    routing_protocols = list(string)  # This will now handle multiple protocols
  }))
  default = {}
}

variable "ha_name" {
  description = "Name of the Anycast service."
  type        = string
}

variable "view_name" {
  description = "Name of the Anycast service."
  type        = string
}

variable "service" {
  description = "The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`)."
  type    = string
  default = "DHCP"
}

variable "anycast_ip_address" {
  description = "Anycast IP address."
  type    = string
}

variable "anycast_config_name" {
  description = "Name of the Anycast configuration."
  type    = string
}

# variable "routing_protocols" {
#   description = "List of routing protocols to be configured (e.g., BGP, OSPF)."
#   type        = list(string)
#   default = ["BGP", "OSPF"]
# }

variable "bgp_configs" {
  description = "Map of BGP configurations per host."
  type = map(object({
    asn            = string
    holddown_secs  = number
    neighbors      = list(object({
      asn        = string
      ip_address = string
    }))
  }))
  default = {}
}

variable "ospf_configs" {
  description = "Map of OSPF configurations per host."
  type = map(object({
    area_type           = string
    area                = string
    authentication_type = string
    interface           = string
    authentication_key  = string
    hello_interval      = number
    dead_interval       = number
    retransmit_interval = number
    transmit_delay      = number
  }))
  default = {}
}

variable "fqdn" {
    description = "FQDN of the Anycast service."
    type        = string
}

variable "primary_type"{
    description = "Primary type of the Anycast service."
    type        = string
}