package types_test

import (
	"context"
	"testing"

	internaltypes "github.com/infobloxopen/terraform-provider-bloxone/internal/types"
)

func TestIPAddressValue_StringSemanticEquals(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want bool
	}{
		{"ipv4 same", "10.0.0.0", "10.0.0.0", true},
		{"ipv4 different", "10.0.0.0", "10.0.1.0", false},
		{"ipv6 case", "2001:db8::", "2001:DB8::", true},
		{"ipv6 compression", "2001:db8:0:0::", "2001:db8::", true},
		{"ipv6 different", "2001:db8::", "2001:db9::", false},
		{"ipv4 vs ipv6", "10.0.0.0", "2001:db8::", false},
		{"invalid", "not-an-ip", "10.0.0.0", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := internaltypes.NewIPAddressValue(tc.a).StringSemanticEquals(context.Background(), internaltypes.NewIPAddressValue(tc.b))
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if got != tc.want {
				t.Errorf("StringSemanticEquals(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
