package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DestinationModel struct {
	Config          types.Object      `tfsdk:"config"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	DeletedAt       timetypes.RFC3339 `tfsdk:"deleted_at"`
	DestinationType types.String      `tfsdk:"destination_type"`
	Id              types.String      `tfsdk:"id"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationAttrTypes = map[string]attr.Type{
	"config":           types.ObjectType{AttrTypes: DestinationConfigAttrTypes},
	"created_at":       timetypes.RFC3339Type{},
	"deleted_at":       timetypes.RFC3339Type{},
	"destination_type": types.StringType,
	"id":               types.StringType,
	"updated_at":       timetypes.RFC3339Type{},
}

var DestinationResourceSchemaAttributes = map[string]schema.Attribute{
	"config": schema.SingleNestedAttribute{
		Attributes:          DestinationConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Destination configuration.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been created.",
	},
	"deleted_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"destination_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			internalvalidator.StringNotNull(),
			stringvalidator.OneOf("DNS", "IPAM/DHCP", "ACCOUNTS"),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIf(
				func(ctx context.Context, request planmodifier.StringRequest, response *stringplanmodifier.RequiresReplaceIfFuncResponse) {
					if !request.ConfigValue.IsNull() && !request.StateValue.IsNull() {
						if request.ConfigValue.String() != request.StateValue.String() {
							response.RequiresReplace = true
						}
					}

				},
				"If the value of the destination_type is configured and changes, Terraform will destroy and recreate the resource.",
				"If the value of the destination_type is configured and changes, Terraform will destroy and recreate the resource.",
			),
		},
		MarkdownDescription: "Destination type: DNS or IPAM/DHCP.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Auto-generated unique destination ID. Format BloxID.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been updated.",
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
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DeletedAt = timetypes.NewRFC3339TimePointerValue(from.DeletedAt)
	m.DestinationType = flex.FlattenString(from.DestinationType)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
