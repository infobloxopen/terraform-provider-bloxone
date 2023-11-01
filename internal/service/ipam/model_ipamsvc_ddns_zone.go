package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcDDNSZoneModel struct {
	Fqdn           types.String `tfsdk:"fqdn"`
	GssTsigEnabled types.Bool   `tfsdk:"gss_tsig_enabled"`
	Nameservers    types.List   `tfsdk:"nameservers"`
	TsigEnabled    types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey        types.Object `tfsdk:"tsig_key"`
	View           types.String `tfsdk:"view"`
	ViewName       types.String `tfsdk:"view_name"`
	Zone           types.String `tfsdk:"zone"`
}

var IpamsvcDDNSZoneAttrTypes = map[string]attr.Type{
	"fqdn":             types.StringType,
	"gss_tsig_enabled": types.BoolType,
	"nameservers":      types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcNameserverAttrTypes}},
	"tsig_enabled":     types.BoolType,
	"tsig_key":         types.ObjectType{AttrTypes: IpamsvcTSIGKeyAttrTypes},
	"view":             types.StringType,
	"view_name":        types.StringType,
	"zone":             types.StringType,
}

var IpamsvcDDNSZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Zone FQDN.  If _zone_ is defined, the _fqdn_ field must be empty.`,
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.`,
	},
	"nameservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcNameserverResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The Nameservers in the zone.  Each nameserver IP should be unique across the list of nameservers.`,
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if TSIG key should be used for the update.  Defaults to _false_.`,
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes: IpamsvcTSIGKeyResourceSchemaAttributes,
		Optional:   true,
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"view_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The name of the view.`,
	},
	"zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandIpamsvcDDNSZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDDNSZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcDDNSZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcDDNSZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDDNSZone {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcDDNSZone{
		Fqdn:           m.Fqdn.ValueStringPointer(),
		GssTsigEnabled: m.GssTsigEnabled.ValueBoolPointer(),
		Nameservers:    flex.ExpandFrameworkListNestedBlock(ctx, m.Nameservers, diags, ExpandIpamsvcNameserver),
		TsigEnabled:    m.TsigEnabled.ValueBoolPointer(),
		TsigKey:        ExpandIpamsvcTSIGKey(ctx, m.TsigKey, diags),
		View:           m.View.ValueStringPointer(),
		Zone:           m.Zone.ValueString(),
	}
	return to
}

func FlattenIpamsvcDDNSZone(ctx context.Context, from *ipam.IpamsvcDDNSZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSZoneAttrTypes)
	}
	m := IpamsvcDDNSZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSZoneModel) Flatten(ctx context.Context, from *ipam.IpamsvcDDNSZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSZoneModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Nameservers = flex.FlattenFrameworkListNestedBlock(ctx, from.Nameservers, IpamsvcNameserverAttrTypes, diags, FlattenIpamsvcNameserver)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = FlattenIpamsvcTSIGKey(ctx, from.TsigKey, diags)
	m.View = flex.FlattenStringPointer(from.View)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.Zone = flex.FlattenString(from.Zone)

}
