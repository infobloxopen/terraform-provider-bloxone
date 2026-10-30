package validator

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = hostPortListValidator{}

// hostPortListValidator validates that the provided value is a comma-separated
// list of "host:port" entries, where host is an FQDN or an IPv4/IPv6 address
// (IPv6 in brackets, e.g. "[2001:db8::1]:9997") and port is between 1 and 65535.
type hostPortListValidator struct{}

func IsValidHostPortList() validator.String {
	return hostPortListValidator{}
}

func (v hostPortListValidator) Description(_ context.Context) string {
	return "value must be a comma-separated list of host:port entries, where host is an FQDN or IP address and port is between 1 and 65535"
}

func (v hostPortListValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v hostPortListValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	for _, entry := range strings.Split(value, ",") {
		if err := validateHostPort(entry); err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Host and Port Value",
				fmt.Sprintf("Value %q is not a valid comma-separated list of host:port entries (for example \"10.10.10.1:9997,splunk.example.com:9997\"): entry %q %s.", value, entry, err),
			)
			return
		}
	}
}

func validateHostPort(entry string) error {
	host, port, err := net.SplitHostPort(entry)
	if err != nil {
		return fmt.Errorf("must be in host:port form")
	}
	if _, err := netip.ParseAddr(host); err != nil && !isValidFQDN(host) {
		return fmt.Errorf("has a host that is neither a valid FQDN (without a trailing dot) nor an IP address")
	}
	if p, err := strconv.Atoi(port); err != nil || p < 1 || p > 65535 {
		return fmt.Errorf("has a port that is not between 1 and 65535")
	}
	return nil
}
