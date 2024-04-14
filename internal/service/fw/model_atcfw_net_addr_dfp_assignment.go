package fw

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwNetAddrDfpAssignmentModel struct {
	AddrNet         types.String `tfsdk:"addr_net"`
	DfpIds          types.List   `tfsdk:"dfp_ids"`
	DfpServiceIds   types.List   `tfsdk:"dfp_service_ids"`
	End             types.String `tfsdk:"end"`
	ExternalScopeId types.String `tfsdk:"external_scope_id"`
	HostId          types.String `tfsdk:"host_id"`
	IpSpaceId       types.String `tfsdk:"ip_space_id"`
	ScopeType       types.String `tfsdk:"scope_type"`
	Start           types.String `tfsdk:"start"`
}

var AtcfwNetAddrDfpAssignmentAttrTypes = map[string]attr.Type{
	"addr_net":          types.StringType,
	"dfp_ids":           types.ListType{ElemType: types.Int64Type},
	"dfp_service_ids":   types.ListType{ElemType: types.StringType},
	"end":               types.StringType,
	"external_scope_id": types.StringType,
	"host_id":           types.StringType,
	"ip_space_id":       types.StringType,
	"scope_type":        types.StringType,
	"start":             types.StringType,
}

var AtcfwNetAddrDfpAssignmentResourceSchemaAttributes = map[string]schema.Attribute{
	"addr_net": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "network address in IPv4 CIDR (address/bitmask length) string format",
	},
	"dfp_ids": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Computed:            true,
		MarkdownDescription: "The list of identifiers of DFPs that have association with this scope.",
	},
	"dfp_service_ids": schema.ListAttribute{
		ElementType: types.StringType,
		Computed:    true,
	},
	"end": schema.StringAttribute{
		Optional: true,
	},
	"external_scope_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "external scope ID, UUID",
	},
	"host_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Host reference, UUID",
	},
	"ip_space_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPSpace reference, UUID",
	},
	"scope_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("UNKNOWN"),
	},
	"start": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Start and end pair of addresses used for range scope type",
	},
}

func ExpandAtcfwNetAddrDfpAssignment(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwNetAddrDfpAssignment {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwNetAddrDfpAssignmentModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwNetAddrDfpAssignmentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwNetAddrDfpAssignment {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwNetAddrDfpAssignment{
		AddrNet:         flex.ExpandStringPointer(m.AddrNet),
		End:             flex.ExpandStringPointer(m.End),
		ExternalScopeId: flex.ExpandStringPointer(m.ExternalScopeId),
		HostId:          flex.ExpandStringPointer(m.HostId),
		IpSpaceId:       flex.ExpandStringPointer(m.IpSpaceId),
		Start:           flex.ExpandStringPointer(m.Start),
		ScopeType:       (*fw.NetAddrDfpAssignmentScopeType)(flex.ExpandStringPointer(m.ScopeType)),
	}
	return to
}

func FlattenAtcfwNetAddrDfpAssignment(ctx context.Context, from *fw.AtcfwNetAddrDfpAssignment, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwNetAddrDfpAssignmentAttrTypes)
	}
	m := AtcfwNetAddrDfpAssignmentModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwNetAddrDfpAssignmentAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwNetAddrDfpAssignmentModel) Flatten(ctx context.Context, from *fw.AtcfwNetAddrDfpAssignment, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwNetAddrDfpAssignmentModel{}
	}
	m.AddrNet = flex.FlattenStringPointer(from.AddrNet)
	m.DfpServiceIds = flex.FlattenFrameworkListString(ctx, from.DfpServiceIds, diags)
	m.End = flex.FlattenStringPointer(from.End)
	m.ExternalScopeId = flex.FlattenStringPointer(from.ExternalScopeId)
	m.HostId = flex.FlattenStringPointer(from.HostId)
	m.IpSpaceId = flex.FlattenStringPointer(from.IpSpaceId)
	m.Start = flex.FlattenStringPointer(from.Start)
	m.DfpIds = flex.FlattenFrameworkListInt32(ctx, from.DfpIds, diags)
	m.ScopeType = flex.FlattenStringPointer((*string)(from.ScopeType))
}
