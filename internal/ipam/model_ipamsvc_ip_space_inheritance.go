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

type IpamsvcIPSpaceInheritanceModel struct {
	AsmConfig                       types.Object `tfsdk:"asm_config"`
	DdnsClientUpdate                types.Object `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.Object `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsEnabled                     types.Object `tfsdk:"ddns_enabled"`
	DdnsHostnameBlock               types.Object `tfsdk:"ddns_hostname_block"`
	DdnsTtlPercent                  types.Object `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateBlock                 types.Object `tfsdk:"ddns_update_block"`
	DdnsUpdateOnRenew               types.Object `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Object `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                      types.Object `tfsdk:"dhcp_config"`
	DhcpOptions                     types.Object `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.Object `tfsdk:"dhcp_options_v6"`
	HeaderOptionFilename            types.Object `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.Object `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.Object `tfsdk:"header_option_server_name"`
	HostnameRewriteBlock            types.Object `tfsdk:"hostname_rewrite_block"`
	VendorSpecificOptionOptionSpace types.Object `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcIPSpaceInheritanceAttrTypes = map[string]attr.Type{
	"asm_config":                          types.ObjectType{AttrTypes: IpamsvcInheritedASMConfigAttrTypes},
	"ddns_client_update":                  types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_conflict_resolution_mode":       types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_enabled":                        types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_hostname_block":                 types.ObjectType{AttrTypes: IpamsvcInheritedDDNSHostnameBlockAttrTypes},
	"ddns_ttl_percent":                    types.ObjectType{AttrTypes: InheritanceInheritedFloatAttrTypes},
	"ddns_update_block":                   types.ObjectType{AttrTypes: IpamsvcInheritedDDNSUpdateBlockAttrTypes},
	"ddns_update_on_renew":                types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_use_conflict_resolution":        types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"dhcp_config":                         types.ObjectType{AttrTypes: IpamsvcInheritedDHCPConfigAttrTypes},
	"dhcp_options":                        types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionListAttrTypes},
	"dhcp_options_v6":                     types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionListAttrTypes},
	"header_option_filename":              types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_address":        types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_name":           types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"hostname_rewrite_block":              types.ObjectType{AttrTypes: IpamsvcInheritedHostnameRewriteBlockAttrTypes},
	"vendor_specific_option_option_space": types.ObjectType{AttrTypes: InheritanceInheritedIdentifierAttrTypes},
}

var IpamsvcIPSpaceInheritanceResourceSchema = schema.Schema{
	MarkdownDescription: `The __IPSpaceInheritance__ object specifies how and which fields _IPSpace_ object inherits from the parent.`,
	Attributes:          IpamsvcIPSpaceInheritanceResourceSchemaAttributes,
}

var IpamsvcIPSpaceInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedASMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_client_update": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_conflict_resolution_mode": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_enabled": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_hostname_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDDNSHostnameBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_ttl_percent": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedFloatResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_update_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDDNSUpdateBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_update_on_renew": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ddns_use_conflict_resolution": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDHCPConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"dhcp_options_v6": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"header_option_filename": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"header_option_server_address": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"header_option_server_name": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"hostname_rewrite_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedHostnameRewriteBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"vendor_specific_option_option_space": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedIdentifierResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcIPSpaceInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcIPSpaceInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcIPSpaceInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcIPSpaceInheritanceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcIPSpaceInheritance {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcIPSpaceInheritance{
		AsmConfig:                       expandIpamsvcInheritedASMConfig(ctx, m.AsmConfig, diags),
		DdnsClientUpdate:                expandInheritanceInheritedString(ctx, m.DdnsClientUpdate, diags),
		DdnsConflictResolutionMode:      expandInheritanceInheritedString(ctx, m.DdnsConflictResolutionMode, diags),
		DdnsEnabled:                     expandInheritanceInheritedBool(ctx, m.DdnsEnabled, diags),
		DdnsHostnameBlock:               expandIpamsvcInheritedDDNSHostnameBlock(ctx, m.DdnsHostnameBlock, diags),
		DdnsTtlPercent:                  expandInheritanceInheritedFloat(ctx, m.DdnsTtlPercent, diags),
		DdnsUpdateBlock:                 expandIpamsvcInheritedDDNSUpdateBlock(ctx, m.DdnsUpdateBlock, diags),
		DdnsUpdateOnRenew:               expandInheritanceInheritedBool(ctx, m.DdnsUpdateOnRenew, diags),
		DdnsUseConflictResolution:       expandInheritanceInheritedBool(ctx, m.DdnsUseConflictResolution, diags),
		DhcpConfig:                      expandIpamsvcInheritedDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     expandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
		DhcpOptionsV6:                   expandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptionsV6, diags),
		HeaderOptionFilename:            expandInheritanceInheritedString(ctx, m.HeaderOptionFilename, diags),
		HeaderOptionServerAddress:       expandInheritanceInheritedString(ctx, m.HeaderOptionServerAddress, diags),
		HeaderOptionServerName:          expandInheritanceInheritedString(ctx, m.HeaderOptionServerName, diags),
		HostnameRewriteBlock:            expandIpamsvcInheritedHostnameRewriteBlock(ctx, m.HostnameRewriteBlock, diags),
		VendorSpecificOptionOptionSpace: expandInheritanceInheritedIdentifier(ctx, m.VendorSpecificOptionOptionSpace, diags),
	}
	return to
}

func flattenIpamsvcIPSpaceInheritance(ctx context.Context, from *ipam.IpamsvcIPSpaceInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIPSpaceInheritanceAttrTypes)
	}
	m := IpamsvcIPSpaceInheritanceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIPSpaceInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIPSpaceInheritanceModel) flatten(ctx context.Context, from *ipam.IpamsvcIPSpaceInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIPSpaceInheritanceModel{}
	}

	m.AsmConfig = flattenIpamsvcInheritedASMConfig(ctx, from.AsmConfig, diags)
	m.DdnsClientUpdate = flattenInheritanceInheritedString(ctx, from.DdnsClientUpdate, diags)
	m.DdnsConflictResolutionMode = flattenInheritanceInheritedString(ctx, from.DdnsConflictResolutionMode, diags)
	m.DdnsEnabled = flattenInheritanceInheritedBool(ctx, from.DdnsEnabled, diags)
	m.DdnsHostnameBlock = flattenIpamsvcInheritedDDNSHostnameBlock(ctx, from.DdnsHostnameBlock, diags)
	m.DdnsTtlPercent = flattenInheritanceInheritedFloat(ctx, from.DdnsTtlPercent, diags)
	m.DdnsUpdateBlock = flattenIpamsvcInheritedDDNSUpdateBlock(ctx, from.DdnsUpdateBlock, diags)
	m.DdnsUpdateOnRenew = flattenInheritanceInheritedBool(ctx, from.DdnsUpdateOnRenew, diags)
	m.DdnsUseConflictResolution = flattenInheritanceInheritedBool(ctx, from.DdnsUseConflictResolution, diags)
	m.DhcpConfig = flattenIpamsvcInheritedDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
	m.DhcpOptionsV6 = flattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptionsV6, diags)
	m.HeaderOptionFilename = flattenInheritanceInheritedString(ctx, from.HeaderOptionFilename, diags)
	m.HeaderOptionServerAddress = flattenInheritanceInheritedString(ctx, from.HeaderOptionServerAddress, diags)
	m.HeaderOptionServerName = flattenInheritanceInheritedString(ctx, from.HeaderOptionServerName, diags)
	m.HostnameRewriteBlock = flattenIpamsvcInheritedHostnameRewriteBlock(ctx, from.HostnameRewriteBlock, diags)
	m.VendorSpecificOptionOptionSpace = flattenInheritanceInheritedIdentifier(ctx, from.VendorSpecificOptionOptionSpace, diags)

}
