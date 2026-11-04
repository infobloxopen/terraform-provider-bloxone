package validator_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

func TestIsValidFQDNOrIPAddress(t *testing.T) {
	cases := []struct {
		name    string
		value   types.String
		wantErr bool
	}{
		{"fqdn", types.StringValue("nios.example.com"), false},
		{"single label", types.StringValue("nios"), false},
		{"hyphen inside label", types.StringValue("nios-gm.example.com"), false},
		{"ipv4", types.StringValue("10.0.0.1"), false},
		{"ipv6", types.StringValue("2001:db8::1"), false},
		{"null", types.StringNull(), false},
		{"unknown", types.StringUnknown(), false},
		{"empty", types.StringValue(""), true},
		{"trailing dot", types.StringValue("nios.example.com."), true},
		{"empty label", types.StringValue("nios..example.com"), true},
		{"leading hyphen", types.StringValue("-nios.example.com"), true},
		{"invalid character", types.StringValue("nios_gm.example.com"), true},
		{"ipv4 with prefix", types.StringValue("10.0.0.0/24"), true},
		{"label too long", types.StringValue(strings.Repeat("a", 64) + ".example.com"), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			internalvalidator.IsValidFQDNOrIPAddress().ValidateString(context.Background(),
				validator.StringRequest{Path: path.Root("address"), ConfigValue: tc.value}, resp)
			if resp.Diagnostics.HasError() != tc.wantErr {
				t.Errorf("HasError = %v, want %v (%v)", resp.Diagnostics.HasError(), tc.wantErr, resp.Diagnostics)
			}
		})
	}
}
