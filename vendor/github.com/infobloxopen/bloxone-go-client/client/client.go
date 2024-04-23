// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/bloxone-go-client/infra_provision"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/bloxone-go-client/keys"
)

// APIClient is an aggregation of different BloxOne API clients.
type APIClient struct {
	IPAddressManagementAPI *ipam.APIClient
	DNSConfigurationAPI    *dns_config.APIClient
	DNSDataAPI             *dns_data.APIClient
	HostActivationAPI      *infra_provision.APIClient
	InfraManagementAPI     *infra_mgmt.APIClient
	KeysAPI                *keys.APIClient
	DNSForwardingProxyAPI  *dfp.APIClient
	FWAPI                  *fw.APIClient
	AnycastAPI             *anycast.APIClient
}

// NewAPIClient creates a new BloxOne API Client.
func NewAPIClient(conf Configuration) (*APIClient, error) {
	ipamConf, err := conf.internal(ipam.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	dnsConfigConf, err := conf.internal(dns_config.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	dnsDataConf, err := conf.internal(dns_data.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	infraProvisionConf, err := conf.internal(infra_provision.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	infraMgmtConf, err := conf.internal(infra_mgmt.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	keysConf, err := conf.internal(keys.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	dfpConf, err := conf.internal(dfp.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	fwConf, err := conf.internal(fw.ServiceBasePath)
	if err != nil {
		return nil, err
	}
	anycastConf, err := conf.internal(anycast.ServiceBasePath)
	if err != nil {
		return nil, err
	}

	return &APIClient{
		IPAddressManagementAPI: ipam.NewAPIClient(ipamConf),
		DNSConfigurationAPI:    dns_config.NewAPIClient(dnsConfigConf),
		DNSDataAPI:             dns_data.NewAPIClient(dnsDataConf),
		HostActivationAPI:      infra_provision.NewAPIClient(infraProvisionConf),
		InfraManagementAPI:     infra_mgmt.NewAPIClient(infraMgmtConf),
		KeysAPI:                keys.NewAPIClient(keysConf),
		DNSForwardingProxyAPI:  dfp.NewAPIClient(dfpConf),
		FWAPI:                  fw.NewAPIClient(fwConf),
		AnycastAPI:             anycast.NewAPIClient(anycastConf),
	}, nil
}
