// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/bloxone-go-client/infra_provision"
	"github.com/infobloxopen/bloxone-go-client/internal"
	"github.com/infobloxopen/bloxone-go-client/ipam"
)

// APIClient is an aggregation of different BloxOne API clients.
type APIClient struct {
	IPAddressManagementAPI *ipam.APIClient
	DNSConfigurationAPI    *dns_config.APIClient
	DNSDataAPI             *dns_data.APIClient
	HostActivationAPI      *infra_provision.APIClient
	InfraManagementAPI     *infra_mgmt.APIClient
}

// NewAPIClient creates a new BloxOne API Client.
func NewAPIClient(host string, apiKey string) *APIClient {
	conf := internal.NewConfiguration()
	conf.Host = host
	conf.AddDefaultHeader("Authorization", "Token "+apiKey)

	return &APIClient{
		IPAddressManagementAPI: ipam.NewAPIClient(conf),
		DNSConfigurationAPI:    dns_config.NewAPIClient(conf),
		DNSDataAPI:             dns_data.NewAPIClient(conf),
		HostActivationAPI:      infra_provision.NewAPIClient(conf),
		InfraManagementAPI:     infra_mgmt.NewAPIClient(conf),
	}
}
