package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwPoPRegionModel struct {
	Addresses types.List   `tfsdk:"addresses"`
	Id        types.Int32  `tfsdk:"id"`
	Location  types.String `tfsdk:"location"`
	Region    types.String `tfsdk:"region"`
}

var AtcfwPoPRegionAttrTypes = map[string]attr.Type{
	"addresses": types.ListType{ElemType: types.StringType},
	"id":        types.Int32Type,
	"location":  types.StringType,
	"region":    types.StringType,
}

var AtcfwPoPRegionResourceSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "PoP Region's IP addresses",
	},
	"id": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The PoP Region's serial, unique ID",
	},
	"location": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "PoP Region's location",
	},
	"region": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "PoP Region's name",
	},
}

func FlattenAtcfwPoPRegion(ctx context.Context, from *fw.PoPRegion, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwPoPRegionAttrTypes)
	}
	m := AtcfwPoPRegionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwPoPRegionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwPoPRegionModel) Flatten(ctx context.Context, from *fw.PoPRegion, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwPoPRegionModel{}
	}
	m.Addresses = flex.FlattenFrameworkListString(ctx, from.Addresses, diags)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Location = flex.FlattenStringPointer(from.Location)
	m.Region = flex.FlattenStringPointer(from.Region)
}
