package validator

import (
	"context"
	"fmt"
	"net/netip"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = fqdnOrIPAddressValidator{}

// fqdnLabel matches one DNS label: 1-63 letters, digits or hyphens, not
// starting or ending with a hyphen.
var fqdnLabel = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?$`)

// fqdnOrIPAddressValidator validates that the provided value is either a plain
// IPv4/IPv6 address or a fully qualified domain name without a trailing dot.
type fqdnOrIPAddressValidator struct{}

func IsValidFQDNOrIPAddress() validator.String {
	return fqdnOrIPAddressValidator{}
}

func (v fqdnOrIPAddressValidator) Description(_ context.Context) string {
	return "value must be a valid FQDN or IPv4/IPv6 address"
}

func (v fqdnOrIPAddressValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v fqdnOrIPAddressValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if _, err := netip.ParseAddr(value); err == nil {
		return
	}
	if isValidFQDN(value) {
		return
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid FQDN or IP Address Value",
		fmt.Sprintf("Value %q is neither a valid FQDN nor a valid IPv4 or IPv6 address. Specify a domain name without a trailing dot (for example \"nios.example.com\") or an address without a prefix length (for example \"10.0.0.1\").", value),
	)
}

// isValidFQDN reports whether value is a domain name of at most 253
// characters made of valid labels. A trailing dot is rejected, matching the
// Infoblox Portal validation.
func isValidFQDN(value string) bool {
	if value == "" || len(value) > 253 {
		return false
	}
	start := 0
	for i := 0; i <= len(value); i++ {
		if i == len(value) || value[i] == '.' {
			if !fqdnLabel.MatchString(value[start:i]) {
				return false
			}
			start = i + 1
		}
	}
	return true
}
