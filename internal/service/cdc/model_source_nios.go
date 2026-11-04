package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SourceNiosModel struct {
	Address            types.String      `tfsdk:"address"`
	CaCertificate      types.String      `tfsdk:"ca_certificate"`
	Certificate        types.String      `tfsdk:"certificate"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at"`
	Description        types.String      `tfsdk:"description"`
	Enabled            types.Bool        `tfsdk:"enabled"`
	Id                 types.Int64       `tfsdk:"id"`
	InsecureMode       types.Bool        `tfsdk:"insecure_mode"`
	Name               types.String      `tfsdk:"name"`
	NiosPassword       types.String      `tfsdk:"nios_password"`
	NiosUsername       types.String      `tfsdk:"nios_username"`
	RpzLogInsecureMode types.Bool        `tfsdk:"rpz_log_insecure_mode"`
	ScpPassword        types.String      `tfsdk:"scp_password"`
	ScpUsername        types.String      `tfsdk:"scp_username"`
	Tags               types.Map         `tfsdk:"tags"`
	TagsAll            types.Map         `tfsdk:"tags_all"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at"`
}

var SourceNiosAttrTypes = map[string]attr.Type{
	"address":               types.StringType,
	"ca_certificate":        types.StringType,
	"certificate":           types.StringType,
	"created_at":            timetypes.RFC3339Type{},
	"description":           types.StringType,
	"enabled":               types.BoolType,
	"id":                    types.Int64Type,
	"insecure_mode":         types.BoolType,
	"name":                  types.StringType,
	"nios_password":         types.StringType,
	"nios_username":         types.StringType,
	"rpz_log_insecure_mode": types.BoolType,
	"scp_password":          types.StringType,
	"scp_username":          types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"tags_all":              types.MapType{ElemType: types.StringType},
	"updated_at":            timetypes.RFC3339Type{},
}

// ca_certificate, certificate, nios_password and scp_password are write-only on
// the CDC backend: it returns "********" when set and "" when unset, so they are
// never read back (see Flatten) and cannot be recovered on import.
var SourceNiosResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IP address or FQDN of the NIOS Grid Master. Must be unique across NIOS sources.",
	},
	"ca_certificate": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The PEM-encoded CA certificate used to verify the NIOS Grid Master (\"CA Certificate\" in the UI). This value is not returned by the API.",
	},
	"certificate": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The PEM-encoded certificate for RPZ logs (\"Certificate for RPZ Logs\" in the UI). This value is not returned by the API.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description of the NIOS source.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether the NIOS source is enabled. Defaults to `false`.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"insecure_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates whether to skip TLS certificate verification when connecting to the NIOS Grid Master.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the NIOS source. Must be unique across NIOS sources.",
	},
	"nios_password": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The password of the NIOS user. This value is not returned by the API.",
	},
	"nios_username": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The username used to connect to the NIOS Grid Master.",
	},
	"rpz_log_insecure_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether to skip TLS certificate verification when collecting RPZ logs. Defaults to `false`.",
	},
	"scp_password": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The password of the SCP user. This value is not returned by the API.",
	},
	"scp_username": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The username used for SCP transfer of DNS query logs from NIOS.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the NIOS source.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the NIOS source, including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandSourceNios(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.SourceNios {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceNiosModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SourceNiosModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.SourceNios {
	if m == nil {
		return nil
	}
	to := &cdc.SourceNios{
		Address:            flex.ExpandStringPointer(m.Address),
		CaCertificate:      flex.ExpandStringPointer(m.CaCertificate),
		Certificate:        flex.ExpandStringPointer(m.Certificate),
		Description:        flex.ExpandStringPointer(m.Description),
		Enabled:            flex.ExpandBoolPointer(m.Enabled),
		InsecureMode:       flex.ExpandBoolPointer(m.InsecureMode),
		Name:               flex.ExpandStringPointer(m.Name),
		NiosPassword:       flex.ExpandStringPointer(m.NiosPassword),
		NiosUsername:       flex.ExpandStringPointer(m.NiosUsername),
		RpzLogInsecureMode: flex.ExpandBoolPointer(m.RpzLogInsecureMode),
		ScpPassword:        flex.ExpandStringPointer(m.ScpPassword),
		ScpUsername:        flex.ExpandStringPointer(m.ScpUsername),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenSourceNios(ctx context.Context, from *cdc.SourceNios, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceNiosAttrTypes)
	}
	m := SourceNiosModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, SourceNiosAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceNiosModel) Flatten(ctx context.Context, from *cdc.SourceNios, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SourceNiosModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.InsecureMode = types.BoolPointerValue(from.InsecureMode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NiosUsername = flex.FlattenStringPointer(from.NiosUsername)
	m.RpzLogInsecureMode = types.BoolPointerValue(from.RpzLogInsecureMode)
	m.ScpUsername = flex.FlattenStringPointer(from.ScpUsername)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
