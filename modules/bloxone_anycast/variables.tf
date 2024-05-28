variable "hosts" {
  description = "Map of hostnames or IP addresses for the Anycast configuration"
  type        = map(string)
  default     = {
    host1 = "ujjwal"
    host2 = "anycast_real"
  }
}

variable "dhcp_hosts" {
  description = "Map of hostnames or IP addresses for DHCP hosts"
  type        = map(string)
  default     = {
    host1 = "ujjwal"
    host2 = "anycast_real"
  }
}

variable "name" {
  description = "Name of the Anycast service"
  type        = string
  default = "anycast-service"
}

variable "service" {
  type    = string
  default = "DHCP"
}

variable "anycast_ip_address" {
  type    = string
  default = "10.10.10.5"
}

variable "anycast_config_name" {
  type    = string
  default = "anycast-config-1"
}

variable "routing_protocols" {
  description = "List of routing protocols to be configured (e.g., BGP, OSPF)"
  type        = list(string)
  default = ["BGP", "OSPF"]
}

variable "asn" {
  type    = string
  default = "6500"
}

variable "holddown_secs" {
  type    = number
  default = 180
}

variable "bgp_neighbors" {
  type = list(object({
    asn        = string
    ip_address = string
  }))
  default = [
    {
      asn        = "6501"
      ip_address = "172.28.4.198"
    }
  ]
}

variable "ospf_config" {
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