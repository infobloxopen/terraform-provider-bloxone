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

type IpamsvcBulkCopyIPSpaceModel struct {
	CopyDhcpOptions types.Bool   `tfsdk:"copy_dhcp_options"`
	CopyObjects     types.List   `tfsdk:"copy_objects"`
	Recursive       types.Bool   `tfsdk:"recursive"`
	SkipOnError     types.Bool   `tfsdk:"skip_on_error"`
	Target          types.String `tfsdk:"target"`
}

var IpamsvcBulkCopyIPSpaceAttrTypes = map[string]attr.Type{
	"copy_dhcp_options": types.BoolType,
	"copy_objects":      types.ListType{ElemType: types.StringType},
	"recursive":         types.BoolType,
	"skip_on_error":     types.BoolType,
	"target":            types.StringType,
}

var IpamsvcBulkCopyIPSpaceResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcBulkCopyIPSpaceResourceSchemaAttributes,
}

var IpamsvcBulkCopyIPSpaceResourceSchemaAttributes = map[string]schema.Attribute{
	"copy_dhcp_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether dhcp options for IPv4 should be copied or not when objects (_ipam/address_block_ and _ipam/subnet_ only) are copied.  Defaults to _false_.`,
	},
	"copy_objects": schema.ListAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"recursive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether child objects should be copied or not.  Defaults to _false_.`,
	},
	"skip_on_error": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.`,
	},
	"target": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcBulkCopyIPSpace(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcBulkCopyIPSpace {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcBulkCopyIPSpaceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcBulkCopyIPSpaceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcBulkCopyIPSpace {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcBulkCopyIPSpace{
		CopyDhcpOptions: m.CopyDhcpOptions.ValueBoolPointer(),
		CopyObjects:     ExpandFrameworkListString(ctx, m.CopyObjects, diags),
		Recursive:       m.Recursive.ValueBoolPointer(),
		SkipOnError:     m.SkipOnError.ValueBoolPointer(),
		Target:          m.Target.ValueString(),
	}
	return to
}

func flattenIpamsvcBulkCopyIPSpace(ctx context.Context, from *ipam.IpamsvcBulkCopyIPSpace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcBulkCopyIPSpaceAttrTypes)
	}
	m := IpamsvcBulkCopyIPSpaceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcBulkCopyIPSpaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcBulkCopyIPSpaceModel) flatten(ctx context.Context, from *ipam.IpamsvcBulkCopyIPSpace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcBulkCopyIPSpaceModel{}
	}

	m.CopyDhcpOptions = types.BoolPointerValue(from.CopyDhcpOptions)
	m.CopyObjects = FlattenFrameworkListString(ctx, from.CopyObjects, diags)
	m.Recursive = types.BoolPointerValue(from.Recursive)
	m.SkipOnError = types.BoolPointerValue(from.SkipOnError)
	m.Target = types.StringValue(from.Target)

}
