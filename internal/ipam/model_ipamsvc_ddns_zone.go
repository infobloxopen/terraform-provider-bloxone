package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
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

var IpamsvcDDNSZoneResourceSchema = schema.Schema{
	MarkdownDescription: `A __DDNSZone__ object (_dhcp/ddns_zone_) represents a DNS zone that can accept dynamic DNS updates from DHCP.`,
	Attributes:          IpamsvcDDNSZoneResourceSchemaAttributes,
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
		Attributes:          IpamsvcTSIGKeyResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
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

func expandIpamsvcDDNSZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDDNSZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDDNSZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDDNSZoneModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDDNSZone {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDDNSZone{
		Fqdn:           m.Fqdn.ValueStringPointer(),
		GssTsigEnabled: m.GssTsigEnabled.ValueBoolPointer(),
		Nameservers:    ExpandFrameworkListNestedBlock(ctx, m.Nameservers, diags, expandIpamsvcNameserver),
		TsigEnabled:    m.TsigEnabled.ValueBoolPointer(),
		TsigKey:        expandIpamsvcTSIGKey(ctx, m.TsigKey, diags),
		View:           m.View.ValueStringPointer(),
		Zone:           m.Zone.ValueString(),
	}
	return to
}

func flattenIpamsvcDDNSZone(ctx context.Context, from *ipam.IpamsvcDDNSZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSZoneAttrTypes)
	}
	m := IpamsvcDDNSZoneModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSZoneModel) flatten(ctx context.Context, from *ipam.IpamsvcDDNSZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSZoneModel{}
	}

	m.Fqdn = types.StringPointerValue(from.Fqdn)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Nameservers = FlattenFrameworkListNestedBlock(ctx, from.Nameservers, IpamsvcNameserverAttrTypes, diags, flattenIpamsvcNameserver)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = flattenIpamsvcTSIGKey(ctx, from.TsigKey, diags)
	m.View = types.StringPointerValue(from.View)
	m.ViewName = types.StringPointerValue(from.ViewName)
	m.Zone = types.StringValue(from.Zone)

}
