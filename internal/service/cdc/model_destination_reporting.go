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
		Optional: true,
	},
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
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
	"password": schema.StringAttribute{
		Optional: true,
	},
	"reporting_appliance_address": schema.StringAttribute{
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
	"username": schema.StringAttribute{
		Optional: true,
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
	m.Address = flex.FlattenStringPointer(from.Address)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.ReportingApplianceAddress = flex.FlattenStringPointer(from.ReportingApplianceAddress)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Username = flex.FlattenStringPointer(from.Username)
}
