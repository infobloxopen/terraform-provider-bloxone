package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

type DestinationReportingModel struct {
	Address                   types.String      `tfsdk:"address"`
	CreatedAt                 timetypes.RFC3339 `tfsdk:"created_at"`
	Description               types.String      `tfsdk:"description"`
	Enabled                   types.Bool        `tfsdk:"enabled"`
	Id                        types.Int64       `tfsdk:"id"`
	Name                      types.String      `tfsdk:"name"`
	Password                  types.String      `tfsdk:"password"`
	ReportingApplianceAddress types.String      `tfsdk:"reporting_appliance_address"`
	Tags                      types.Map         `tfsdk:"tags"`
	UpdatedAt                 timetypes.RFC3339 `tfsdk:"updated_at"`
	Username                  types.String      `tfsdk:"username"`
}

var DestinationReportingAttrTypes = map[string]attr.Type{
	"address":                     types.StringType,
	"created_at":                  timetypes.RFC3339Type{},
	"description":                 types.StringType,
	"enabled":                     types.BoolType,
	"id":                          types.Int64Type,
	"name":                        types.StringType,
	"password":                    types.StringType,
	"reporting_appliance_address": types.StringType,
	"tags":                        types.MapType{ElemType: types.StringType},
	"updated_at":                  timetypes.RFC3339Type{},
	"username":                    types.StringType,
}

var DestinationReportingResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			internalvalidator.IsValidFQDNOrIPAddress(),
		},
		MarkdownDescription: "The FQDN or IP address of the NIOS Grid Master.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthAtMost(256),
		},
		MarkdownDescription: "The description of the NIOS reporting destination. May contain 0 to 256 characters.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the NIOS reporting destination is enabled. Defaults to _true_.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the NIOS reporting destination.",
	},
	"password": schema.StringAttribute{
		Required:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The password used to connect to the NIOS Grid Master.",
	},
	"reporting_appliance_address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			internalvalidator.IsValidFQDNOrIPAddress(),
		},
		MarkdownDescription: "The FQDN or IP address of the NIOS reporting appliance.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the NIOS reporting destination.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"username": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The username used to connect to the NIOS Grid Master.",
	},
}

func ExpandDestinationReporting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationReporting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationReportingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationReportingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationReporting {
	if m == nil {
		return nil
	}
	to := &cdc.DestinationReporting{
		Address:                   flex.ExpandStringPointer(m.Address),
		Description:               flex.ExpandStringPointer(m.Description),
		Enabled:                   flex.ExpandBoolPointer(m.Enabled),
		Name:                      flex.ExpandStringPointer(m.Name),
		Password:                  flex.ExpandStringPointer(m.Password),
		ReportingApplianceAddress: flex.ExpandStringPointer(m.ReportingApplianceAddress),
		Tags:                      flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Username:                  flex.ExpandStringPointer(m.Username),
	}
	return to
}

func FlattenDestinationReporting(ctx context.Context, from *cdc.DestinationReporting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationReportingAttrTypes)
	}
	m := DestinationReportingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationReportingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationReportingModel) Flatten(ctx context.Context, from *cdc.DestinationReporting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationReportingModel{}
	}
	// password is never read back from the API: the backend masks secrets the
	// same way it does for ca_certificate on the syslog destination.
	m.Address = flex.FlattenStringPointer(from.Address)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ReportingApplianceAddress = flex.FlattenStringPointer(from.ReportingApplianceAddress)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Username = flex.FlattenStringPointer(from.Username)
}
