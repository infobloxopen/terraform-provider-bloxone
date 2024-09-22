package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ResourceModel struct {
	Excluded types.Bool   `tfsdk:"excluded"`
	Id       types.String `tfsdk:"id"`
}

var ResourceAttrTypes = map[string]attr.Type{
	"excluded": types.BoolType,
	"id":       types.StringType,
}

var ResourceResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded": schema.BoolAttribute{
		Optional: true,
		Computed: true,
	},
	"id": schema.StringAttribute{
		Optional: true,
		Computed: true,
	},
}

func ExpandResource(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.Resource {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ResourceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ResourceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.Resource {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.Resource{
		Excluded: flex.ExpandBoolPointer(m.Excluded),
	}
	return to
}

func FlattenResource(ctx context.Context, from *clouddiscovery.Resource, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ResourceAttrTypes)
	}
	m := ResourceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ResourceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ResourceModel) Flatten(ctx context.Context, from *clouddiscovery.Resource, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ResourceModel{}
	}
	m.Excluded = types.BoolPointerValue(from.Excluded)
	m.Id = flex.FlattenStringPointer(from.Id)
}
