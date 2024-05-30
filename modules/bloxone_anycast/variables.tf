variable "hosts" {
  description = "Map of hostnames or IP addresses for the Anycast configuration."
  type        = map(string)
}

variable "name" {
  description = "Name of the Anycast service."
  type        = string
  default = "anycast-service"
}

variable "service" {
  description = "The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, `dfp`)."
  type    = string
  default = "DHCP"
}

variable "anycast_ip_address" {
  description = "Anycast IP address."
  type    = string
  default = "10.10.10.5"
}

variable "anycast_config_name" {
  description = "Name of the Anycast configuration."
  type    = string
  default = "anycast-config-1"
}

variable "routing_protocols" {
  description = "List of routing protocols to be configured (e.g., BGP, OSPF)."
  type        = list(string)
  default = ["BGP", "OSPF"]
}

variable "bgp_config" {
  description = "BGP configuration"
  type = object({
    asn            = string
    holddown_secs  = number
    neighbors      = list(object({
      asn        = string
      ip_address = string
    }))
  })
  default = {
    asn           = "6500"
    holddown_secs = 180
    neighbors     = [
      {
        asn        = "6501"
        ip_address = "172.28.4.198"
      }
    ]
  }
}

variable "ospf_config" {
  description = "OSPF configuration."
  type = object({
    area                = string
    area_type           = string
    authentication_type = string
    interface           = string
    authentication_key  = string
    hello_interval      = number
    dead_interval       = number
    retransmit_interval = number
    transmit_delay      = number
  })
  default = {
    area                = "10.10.0.1"
    area_type           = "STANDARD"
    authentication_type = "Clear"
    interface           = "eth0"
    authentication_key  = "YXV0aGV"
    hello_interval      = 10
    dead_interval       = 40
    retransmit_interval = 5
    transmit_delay      = 1
  }
}