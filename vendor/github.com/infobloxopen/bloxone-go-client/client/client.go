// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/bloxone-go-client/infra_provision"
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
func NewAPIClient(conf Configuration) (*APIClient, error) {
	c, err := conf.internal()
	if err != nil {
		return nil, err
	}
	return &APIClient{
		IPAddressManagementAPI: ipam.NewAPIClient(c),
		DNSConfigurationAPI:    dns_config.NewAPIClient(c),
		DNSDataAPI:             dns_data.NewAPIClient(c),
		HostActivationAPI:      infra_provision.NewAPIClient(c),
		InfraManagementAPI:     infra_mgmt.NewAPIClient(c),
	}, nil
}
