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

type DestinationHttpModel struct {
	Address           types.String      `tfsdk:"address"`
	Authentication    types.String      `tfsdk:"authentication"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	CredentialId      types.String      `tfsdk:"credential_id"`
	Description       types.String      `tfsdk:"description"`
	Enabled           types.Bool        `tfsdk:"enabled"`
	Id                types.Int64       `tfsdk:"id"`
	InsecureMode      types.Bool        `tfsdk:"insecure_mode"`
	Name              types.String      `tfsdk:"name"`
	OutputDataFormat  types.String      `tfsdk:"output_data_format"`
	Tags              types.Map         `tfsdk:"tags"`
	TransportProtocol types.String      `tfsdk:"transport_protocol"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationHttpAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"authentication":     types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"credential_id":      types.StringType,
	"description":        types.StringType,
	"enabled":            types.BoolType,
	"id":                 types.Int64Type,
	"insecure_mode":      types.BoolType,
	"name":               types.StringType,
	"output_data_format": types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"transport_protocol": types.StringType,
	"updated_at":         timetypes.RFC3339Type{},
}

var DestinationHttpResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
	},
	"authentication": schema.StringAttribute{
		Optional: true,
	},
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"credential_id": schema.StringAttribute{
		Optional: true,
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
	"insecure_mode": schema.BoolAttribute{
		Optional: true,
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"transport_protocol": schema.StringAttribute{
		Optional: true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
}

func ExpandDestinationHttp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationHttp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationHttpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationHttpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationHttp {
	if m == nil {
		return nil
	}
	to := &cdc.DestinationHttp{
		Address:           flex.ExpandStringPointer(m.Address),
		Authentication:    flex.ExpandStringPointer(m.Authentication),
		CredentialId:      flex.ExpandStringPointer(m.CredentialId),
		Description:       flex.ExpandStringPointer(m.Description),
		Enabled:           flex.ExpandBoolPointer(m.Enabled),
		InsecureMode:      flex.ExpandBoolPointer(m.InsecureMode),
		Name:              flex.ExpandStringPointer(m.Name),
		OutputDataFormat:  flex.ExpandStringPointer(m.OutputDataFormat),
		Tags:              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransportProtocol: flex.ExpandStringPointer(m.TransportProtocol),
	}
	return to
}

func FlattenDestinationHttp(ctx context.Context, from *cdc.DestinationHttp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationHttpAttrTypes)
	}
	m := DestinationHttpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationHttpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationHttpModel) Flatten(ctx context.Context, from *cdc.DestinationHttp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationHttpModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Authentication = flex.FlattenStringPointer(from.Authentication)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CredentialId = flex.FlattenStringPointer(from.CredentialId)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.InsecureMode = types.BoolPointerValue(from.InsecureMode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutputDataFormat = flex.FlattenStringPointer(from.OutputDataFormat)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransportProtocol = flex.FlattenStringPointer(from.TransportProtocol)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
