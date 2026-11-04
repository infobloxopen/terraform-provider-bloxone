package cdc

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	TagsAll           types.Map         `tfsdk:"tags_all"`
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
	"tags_all":           types.MapType{ElemType: types.StringType},
	"transport_protocol": types.StringType,
	"updated_at":         timetypes.RFC3339Type{},
}

var DestinationHttpResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(regexp.MustCompile(`^https?://[^\s/]+`), "must be an http:// or https:// URL"),
		},
		MarkdownDescription: "The URL of the HTTP endpoint, e.g. `https://http.example.com/ingest`.",
	},
	"authentication": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The authentication method used to connect to the HTTP endpoint.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"credential_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The identifier of an existing credential used to authenticate to the HTTP endpoint.",
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthAtMost(256),
		},
		MarkdownDescription: "The description of the HTTP destination. May contain 0 to 256 characters.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the HTTP destination is enabled. Defaults to _true_.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"insecure_mode": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the HTTP endpoint certificate is not verified. Always _false_: insecure mode is not supported for HTTP destinations.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the HTTP destination.",
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ASIM"),
		Validators: []validator.String{
			stringvalidator.OneOf("ASIM", "CIM"),
		},
		MarkdownDescription: "The format of the data sent to the HTTP endpoint. Allowed values:\n\n  * _ASIM_: Microsoft Sentinel\n  * _CIM_: Splunk CIM\n\nDefaults to _ASIM_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the HTTP destination.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the HTTP destination, including default tags.",
	},
	"transport_protocol": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("HTTP"),
		Validators: []validator.String{
			stringvalidator.OneOf("HTTP", "HTTPS"),
		},
		MarkdownDescription: "The transport protocol used to connect to the HTTP endpoint. Allowed values:\n\n  * _HTTP_\n  * _HTTPS_\n\nDefaults to _HTTP_.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
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

func DataSourceFlattenDestinationHttp(ctx context.Context, from *cdc.DestinationHttp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationHttpAttrTypes)
	}
	m := DestinationHttpModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
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
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransportProtocol = flex.FlattenStringPointer(from.TransportProtocol)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
