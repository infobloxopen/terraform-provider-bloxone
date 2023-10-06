// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/dns_config"
	"github.com/infobloxopen/b1ddi-go-client/dns_data"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
)

// Client is an aggregation of different BloxOne DDI API clients.
type Client struct {
	IPAddressManagementAPI *ipamsvc.IPAddressManagementAPI
	DNSConfigurationAPI    *dns_config.DNSConfigurationAPI
	DNSDataAPI             *dns_data.DNSDataAPI
}

// NewClient creates a new BloxOne DDI API Client.
func NewClient(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{
		IPAddressManagementAPI: ipamsvc.New(transport, formats),
		DNSConfigurationAPI:    dns_config.New(transport, formats),
		DNSDataAPI:             dns_data.New(transport, formats),
	}
}

// B1DDIAPIKey provides a header for the BloxOne DDI API authentication.
//
// See https://docs.infoblox.com/display/BloxOneDDI/BloxOne+DDI+API+Guide learn how to get the API key.
func B1DDIAPIKey(apiKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", "Token "+apiKey)
	})
}
