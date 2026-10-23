package validator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

func TestIsValidIPAddress(t *testing.T) {
	cases := []struct {
		name    string
		value   types.String
		wantErr bool
	}{
		{"ipv4", types.StringValue("10.40.0.0"), false},
		{"ipv6", types.StringValue("2001:db8::"), false},
		{"null", types.StringNull(), false},
		{"unknown", types.StringUnknown(), false},
		{"ipv4 with prefix", types.StringValue("10.40.0.0/18"), true},
		{"ipv6 with prefix", types.StringValue("2001:db8::/32"), true},
		{"garbage", types.StringValue("abc"), true},
		{"empty", types.StringValue(""), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			internalvalidator.IsValidIPAddress().ValidateString(context.Background(),
				validator.StringRequest{Path: path.Root("address"), ConfigValue: tc.value}, resp)
			if resp.Diagnostics.HasError() != tc.wantErr {
				t.Errorf("HasError = %v, want %v (%v)", resp.Diagnostics.HasError(), tc.wantErr, resp.Diagnostics)
			}
		})
	}
}
