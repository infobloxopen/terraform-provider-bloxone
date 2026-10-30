package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SourceBloxoneModel struct {
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Tags        types.Map         `tfsdk:"tags"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
}

var SourceBloxoneAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"enabled":     types.BoolType,
	"id":          types.Int64Type,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"updated_at":  timetypes.RFC3339Type{},
}

var SourceBloxoneResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional: true,
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
	},
	"id": schema.Int64Attribute{
		Computed: true,
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
}

func ExpandSourceBloxone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.SourceBloxone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceBloxoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SourceBloxoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.SourceBloxone {
	if m == nil {
		return nil
	}
	to := &cdc.SourceBloxone{
		Description: flex.ExpandStringPointer(m.Description),
		Enabled:     flex.ExpandBoolPointer(m.Enabled),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenSourceBloxone(ctx context.Context, from *cdc.SourceBloxone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceBloxoneAttrTypes)
	}
	m := SourceBloxoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceBloxoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceBloxoneModel) Flatten(ctx context.Context, from *cdc.SourceBloxone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SourceBloxoneModel{}
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
