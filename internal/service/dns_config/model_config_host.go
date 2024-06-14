package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
)

func ExpandConfigHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Host {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}
