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

type IpamsvcLeasesCommandModel struct {
	Address types.List   `tfsdk:"address"`
	Command types.String `tfsdk:"command"`
	Range   types.List   `tfsdk:"range"`
	Subnet  types.List   `tfsdk:"subnet"`
}

var IpamsvcLeasesCommandAttrTypes = map[string]attr.Type{
	"address": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcLeaseAddressAttrTypes}},
	"command": types.StringType,
	"range":   types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcLeaseRangeAttrTypes}},
	"subnet":  types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcLeaseSubnetAttrTypes}},
}

var IpamsvcLeasesCommandResourceSchema = schema.Schema{
	MarkdownDescription: `The __LeasesCommand__ (_dhcp/leases_command_) is used to perform an action on a lease or a set of leases defined by the list of IP addresses or Subnet or Range. Host(s) owning the lease(s) must be available for this action to succeed. The command is executed asynchronously.`,
	Attributes:          IpamsvcLeasesCommandResourceSchemaAttributes,
}

var IpamsvcLeasesCommandResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcLeaseAddressResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of IP addresses to execute the \&quot;command\&quot; on. It can be 1 or more IP addresses.`,
	},
	"command": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The command to perform on the lease(s).  Valid values are:  | command       | description | | :------       | ----------- | | _clear_       | Removes selected lease(s) from the DHCP server(s). This will NOT affect the client that issued the lease. | | _resend-ddns_ | Resends a request to kea-dhcp-ddns to update DNS for selected lease(s). |`,
	},
	"range": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcLeaseRangeResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of ranges to execute the \&quot;command\&quot; on. For now it is limited to 1 range.`,
	},
	"subnet": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcLeaseSubnetResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of subnets to execute the \&quot;command\&quot; on. For now it is limited to 1 subnet.`,
	},
}

func expandIpamsvcLeasesCommand(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcLeasesCommand {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcLeasesCommandModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcLeasesCommandModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcLeasesCommand {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcLeasesCommand{
		Address: ExpandFrameworkListNestedBlock(ctx, m.Address, diags, expandIpamsvcLeaseAddress),
		Command: m.Command.ValueString(),
		Range:   ExpandFrameworkListNestedBlock(ctx, m.Range, diags, expandIpamsvcLeaseRange),
		Subnet:  ExpandFrameworkListNestedBlock(ctx, m.Subnet, diags, expandIpamsvcLeaseSubnet),
	}
	return to
}

func flattenIpamsvcLeasesCommand(ctx context.Context, from *ipam.IpamsvcLeasesCommand, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcLeasesCommandAttrTypes)
	}
	m := IpamsvcLeasesCommandModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcLeasesCommandAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcLeasesCommandModel) flatten(ctx context.Context, from *ipam.IpamsvcLeasesCommand, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcLeasesCommandModel{}
	}

	m.Address = FlattenFrameworkListNestedBlock(ctx, from.Address, IpamsvcLeaseAddressAttrTypes, diags, flattenIpamsvcLeaseAddress)
	m.Command = types.StringValue(from.Command)
	m.Range = FlattenFrameworkListNestedBlock(ctx, from.Range, IpamsvcLeaseRangeAttrTypes, diags, flattenIpamsvcLeaseRange)
	m.Subnet = FlattenFrameworkListNestedBlock(ctx, from.Subnet, IpamsvcLeaseSubnetAttrTypes, diags, flattenIpamsvcLeaseSubnet)

}
