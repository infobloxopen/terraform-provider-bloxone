package validator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

func TestIsValidHostPortList(t *testing.T) {
	cases := []struct {
		name    string
		value   types.String
		wantErr bool
	}{
		{"fqdn with port", types.StringValue("splunk.example.com:9997"), false},
		{"ipv4 with port", types.StringValue("10.0.0.1:9997"), false},
		{"ipv6 with port", types.StringValue("[2001:db8::1]:9997"), false},
		{"list", types.StringValue("10.10.10.1:9997,10.10.10.2:9997"), false},
		{"null", types.StringNull(), false},
		{"unknown", types.StringUnknown(), false},
		{"empty", types.StringValue(""), true},
		{"missing port", types.StringValue("splunk.example.com"), true},
		{"trailing dot", types.StringValue("splunk.example.com.:9997"), true},
		{"port out of range", types.StringValue("splunk.example.com:70000"), true},
		{"port zero", types.StringValue("splunk.example.com:0"), true},
		{"non-numeric port", types.StringValue("splunk.example.com:abc"), true},
		{"space after comma", types.StringValue("10.10.10.1:9997, 10.10.10.2:9997"), true},
		{"empty entry", types.StringValue("10.10.10.1:9997,"), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			internalvalidator.IsValidHostPortList().ValidateString(context.Background(),
				validator.StringRequest{Path: path.Root("address"), ConfigValue: tc.value}, resp)
			if resp.Diagnostics.HasError() != tc.wantErr {
				t.Errorf("HasError = %v, want %v (%v)", resp.Diagnostics.HasError(), tc.wantErr, resp.Diagnostics)
			}
		})
	}
}
