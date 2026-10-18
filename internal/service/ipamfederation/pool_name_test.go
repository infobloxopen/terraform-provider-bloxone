package ipamfederation

import (
	"testing"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func TestStripPoolNameSuffix(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"suffix removed", "my-pool(ipam-pool-0073c8bb836f02aec)", "my-pool"},
		{"no suffix", "my-pool", "my-pool"},
		{"user parentheses kept", "pool(a)", "pool(a)"},
		{"user parentheses plus suffix", "pool(a)(ipam-pool-01e6852c8a2268ca7)", "pool(a)"},
		{"suffix only", "(ipam-pool-abc123)", ""},
		{"suffix not at end", "(ipam-pool-abc123)-pool", "(ipam-pool-abc123)-pool"},
		{"other provider text kept", "pool(ipam-scope-abc123)", "pool(ipam-scope-abc123)"},
		{"unclosed", "pool(ipam-pool-abc", "pool(ipam-pool-abc"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripPoolNameSuffix(tt.in); got != tt.want {
				t.Errorf("stripPoolNameSuffix(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFlattenPoolName(t *testing.T) {
	if !flattenPoolName(nil, ipamfederation.PROVIDERTYPE_AWS).IsNull() {
		t.Error("nil name should flatten to null")
	}
	tests := []struct {
		name     string
		in       string
		provider ipamfederation.ProviderType
		want     string
	}{
		{"AWS pool: suffix removed", "p(ipam-pool-1)", ipamfederation.PROVIDERTYPE_AWS, "p"},
		{"AWS pool: no suffix", "p", ipamfederation.PROVIDERTYPE_AWS, "p"},
		{"NIOS_X pool: name chosen by the user is never altered", "p(ipam-pool-1)", ipamfederation.PROVIDERTYPE_NIOS_X, "p(ipam-pool-1)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenPoolName(&tt.in, tt.provider); got.ValueString() != tt.want {
				t.Errorf("got %q, want %q", got.ValueString(), tt.want)
			}
		})
	}
}
