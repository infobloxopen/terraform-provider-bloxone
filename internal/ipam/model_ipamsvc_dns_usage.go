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

type IpamsvcDNSUsageModel struct {
	AbsoluteName types.String `tfsdk:"absolute_name"`
	Address      types.String `tfsdk:"address"`
	DnsRdata     types.String `tfsdk:"dns_rdata"`
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Record       types.String `tfsdk:"record"`
	Space        types.String `tfsdk:"space"`
	Type         types.String `tfsdk:"type"`
	View         types.String `tfsdk:"view"`
	Zone         types.String `tfsdk:"zone"`
}

var IpamsvcDNSUsageAttrTypes = map[string]attr.Type{
	"absolute_name": types.StringType,
	"address":       types.StringType,
	"dns_rdata":     types.StringType,
	"id":            types.StringType,
	"name":          types.StringType,
	"record":        types.StringType,
	"space":         types.StringType,
	"type":          types.StringType,
	"view":          types.StringType,
	"zone":          types.StringType,
}

var IpamsvcDNSUsageResourceSchema = schema.Schema{
	MarkdownDescription: `The __DNSUsage__ object tracks DNS usage of a resource record on an address.`,
	Attributes:          IpamsvcDNSUsageResourceSchemaAttributes,
}

var IpamsvcDNSUsageResourceSchemaAttributes = map[string]schema.Attribute{
	"absolute_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The absolute name of the resource record in associated zone.`,
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The address of the referenced record.`,
	},
	"dns_rdata": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The DNS rdata of the referenced record.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The name in zone of the referenced record.`,
	},
	"record": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of the referenced record.`,
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcDNSUsage(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDNSUsage {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDNSUsageModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDNSUsageModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDNSUsage {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDNSUsage{
		Record: m.Record.ValueStringPointer(),
		Space:  m.Space.ValueStringPointer(),
		View:   m.View.ValueStringPointer(),
		Zone:   m.Zone.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcDNSUsage(ctx context.Context, from *ipam.IpamsvcDNSUsage, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDNSUsageAttrTypes)
	}
	m := IpamsvcDNSUsageModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDNSUsageAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDNSUsageModel) flatten(ctx context.Context, from *ipam.IpamsvcDNSUsage, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDNSUsageModel{}
	}

	m.AbsoluteName = types.StringPointerValue(from.AbsoluteName)
	m.Address = types.StringPointerValue(from.Address)
	m.DnsRdata = types.StringPointerValue(from.DnsRdata)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringPointerValue(from.Name)
	m.Record = types.StringPointerValue(from.Record)
	m.Space = types.StringPointerValue(from.Space)
	m.Type = types.StringPointerValue(from.Type)
	m.View = types.StringPointerValue(from.View)
	m.Zone = types.StringPointerValue(from.Zone)

}
