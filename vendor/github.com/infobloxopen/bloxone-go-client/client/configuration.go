package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

const (
	ENVBloxOneCSPURL = "BLOXONE_CSP_URL"
	ENVBloxOneAPIKey = "BLOXONE_API_KEY"
)

const (
	HeaderClient        = "x-infoblox-client"
	HeaderSDK           = "x-infoblox-sdk"
	HeaderAuthorization = "Authorization"
)

const version = "0.1"
const sdkIdentifier = "golang-sdk"

// Configuration stores the configuration of the API client
type Configuration struct {
	// ClientName is the name of the client using the SDK.
	// Required.
	ClientName string

	// CSPURL is the URL for BloxOne Cloud Services Portal.
	// Can also be configured using the `BLOXONE_CSP_URL` environment variable.
	// Optional. Default is https://csp.infoblox.com
	CSPURL *url.URL

	// APIKey for accessing the BloxOne API.
	// Can also be configured by using the `BLOXONE_API_KEY` environment variable.
	// https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys
	// Required.
	APIKey string

	// HTTPClient to use for the SDK.
	// Optional. The default HTTPClient will be used if not provided.
	HTTPClient *http.Client
}

func (c Configuration) internal() (*internal.Configuration, error) {
	var err error
	cspURL := c.CSPURL
	if cspURL == nil {
		cspURL = &url.URL{Scheme: "https", Host: "csp.infoblox.com"}
		if v, ok := os.LookupEnv(ENVBloxOneCSPURL); ok {
			if cspURL, err = url.Parse(v); err != nil {
				return nil, err
			}
		}
	}
	if len(cspURL.Scheme) == 0 {
		cspURL.Scheme = "https"
	}

	apiKey := c.APIKey
	if len(apiKey) == 0 {
		var ok bool
		if apiKey, ok = os.LookupEnv(ENVBloxOneAPIKey); !ok {
			return nil, errors.New("APIKey is required")
		}
	}

	clientName := c.ClientName
	if len(clientName) == 0 {
		return nil, errors.New("ClientName is required")
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	defaultHeaders := map[string]string{
		HeaderAuthorization: "Token " + apiKey,
		HeaderClient:        clientName,
		HeaderSDK:           sdkIdentifier,
	}

	userAgent := fmt.Sprintf("bloxone-%s/%s", sdkIdentifier, version)

	return &internal.Configuration{
		Host:             cspURL.Host,
		Scheme:           cspURL.Host,
		DefaultHeader:    defaultHeaders,
		UserAgent:        userAgent,
		Debug:            false,
		OperationServers: nil,
		HTTPClient:       httpClient,
	}, nil
}
