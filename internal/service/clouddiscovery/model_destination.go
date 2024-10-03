package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

type DestinationModel struct {
	Config          types.Object `tfsdk:"config"`
	DestinationType types.String `tfsdk:"destination_type"`
}

var DestinationAttrTypes = map[string]attr.Type{
	"config":           types.ObjectType{AttrTypes: DestinationConfigAttrTypes},
	"destination_type": types.StringType,
}

var DestinationResourceSchemaAttributes = map[string]schema.Attribute{
	"config": schema.SingleNestedAttribute{
		Attributes:          DestinationConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Destination configuration.",
	},
	"destination_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			internalvalidator.StringNotNull(),
			stringvalidator.OneOf("DNS", "IPAM/DHCP", "ACCOUNTS"),
		},
		MarkdownDescription: "Destination type: DNS or IPAM/DHCP or ACCOUNTS.",
	},
}

func ExpandDestination(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.Destination {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.Destination {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.Destination{
		Config:          ExpandDestinationConfig(ctx, m.Config, diags),
		DestinationType: flex.ExpandString(m.DestinationType),
	}
	return to
}

func FlattenDestination(ctx context.Context, from *clouddiscovery.Destination, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationAttrTypes)
	}
	m := DestinationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationModel) Flatten(ctx context.Context, from *clouddiscovery.Destination, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationModel{}
	}
	m.Config = FlattenDestinationConfig(ctx, from.Config, diags)
	m.DestinationType = flex.FlattenString(from.DestinationType)
}
