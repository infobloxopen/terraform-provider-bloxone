package ipamfederation

import "testing"

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
	if !flattenPoolName(nil).IsNull() {
		t.Error("nil name should flatten to null")
	}
	n := "p(ipam-pool-1)"
	if got := flattenPoolName(&n); got.ValueString() != "p" {
		t.Errorf("got %q", got.ValueString())
	}
}
