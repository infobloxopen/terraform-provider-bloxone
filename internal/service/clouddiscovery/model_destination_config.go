package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
)

type DestinationConfigModel struct {
	Dns  types.Object `tfsdk:"dns"`
	Ipam types.Object `tfsdk:"ipam"`
}

var DestinationConfigAttrTypes = map[string]attr.Type{
	"dns":  types.ObjectType{AttrTypes: DNSConfigAttrTypes},
	"ipam": types.ObjectType{AttrTypes: IPAMConfigAttrTypes},
}

var DestinationConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"dns": schema.SingleNestedAttribute{
		Attributes: DNSConfigResourceSchemaAttributes,
		Optional:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.RequiresReplaceIf(
				func(ctx context.Context, request planmodifier.ObjectRequest, response *objectplanmodifier.RequiresReplaceIfFuncResponse) {
					if !request.ConfigValue.IsNull() && !request.StateValue.IsNull() {
						planVal := request.PlanValue.Attributes()
						stateVal := request.StateValue.Attributes()
						if !(planVal["view_id"].String() == stateVal["view_id"].String()) ||
							!(planVal["view_name"].String() == stateVal["view_name"].String()) ||
							!(planVal["consolidated_zone_data_enabled"].String() == stateVal["consolidated_zone_data_enabled"].String()) ||
							!(planVal["split_view_enabled"].String() == stateVal["split_view_enabled"].String()) {
							response.RequiresReplace = true
						}
					}
				},
				"If the value of the view_id , view_name , consolidated_zone_data_enabled or split_view_enabled is configured and changes, Terraform will destroy and recreate the resource.",
				"If the value of the view_id , view_name , consolidated_zone_data_enabled or split_view_enabled is configured and changes, Terraform will destroy and recreate the resource.",
			),
		},
		MarkdownDescription: "Destination Config for DNS",
	},
	"ipam": schema.SingleNestedAttribute{
		Attributes:          IPAMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Destination Config for IPAM/DHCP",
	},
}

func ExpandDestinationConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.DestinationConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.DestinationConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.DestinationConfig{
		Dns:  ExpandDNSConfig(ctx, m.Dns, diags),
		Ipam: ExpandIPAMConfig(ctx, m.Ipam, diags),
	}
	return to
}

func FlattenDestinationConfig(ctx context.Context, from *clouddiscovery.DestinationConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationConfigAttrTypes)
	}
	m := DestinationConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationConfigModel) Flatten(ctx context.Context, from *clouddiscovery.DestinationConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationConfigModel{}
	}
	m.Dns = FlattenDNSConfig(ctx, from.Dns, diags)
	m.Ipam = FlattenIPAMConfig(ctx, from.Ipam, diags)
}
