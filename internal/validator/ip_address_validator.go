package validator

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ipAddressValidator{}

// ipAddressValidator validates that the provided value is a plain IPv4 or IPv6 address without a prefix length.
type ipAddressValidator struct{}

func IsValidIPAddress() validator.String {
	return ipAddressValidator{}
}

func (v ipAddressValidator) Description(_ context.Context) string {
	return "value must be a valid IPv4 or IPv6 address without a prefix length"
}

func (v ipAddressValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipAddressValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if _, err := netip.ParseAddr(value); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Address Value",
			fmt.Sprintf("Value %q is not a valid IPv4 or IPv6 address. Specify the address without a prefix length (for example \"10.0.0.0\" or \"2001:db8::\"). Error: %s", value, err),
		)
	}
}
