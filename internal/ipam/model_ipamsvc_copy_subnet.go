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

type IpamsvcCopySubnetModel struct {
	Comment         types.String `tfsdk:"comment"`
	CopyDhcpOptions types.Bool   `tfsdk:"copy_dhcp_options"`
	Id              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Recursive       types.Bool   `tfsdk:"recursive"`
	SkipOnError     types.Bool   `tfsdk:"skip_on_error"`
	Space           types.String `tfsdk:"space"`
}

var IpamsvcCopySubnetAttrTypes = map[string]attr.Type{
	"comment":           types.StringType,
	"copy_dhcp_options": types.BoolType,
	"id":                types.StringType,
	"name":              types.StringType,
	"recursive":         types.BoolType,
	"skip_on_error":     types.BoolType,
	"space":             types.StringType,
}

var IpamsvcCopySubnetResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcCopySubnetResourceSchemaAttributes,
}

var IpamsvcCopySubnetResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the copied subnet. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"copy_dhcp_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether dhcp options should be copied or not when _ipam/subnet_ object is copied.  Defaults to _false_.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name for the copied subnet. May contain 1 to 256 characters. Can include UTF-8.`,
	},
	"recursive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether child objects should be copied or not.  Defaults to _false_.`,
	},
	"skip_on_error": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.`,
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcCopySubnet(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcCopySubnet {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcCopySubnetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcCopySubnetModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcCopySubnet {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcCopySubnet{
		Comment:         m.Comment.ValueStringPointer(),
		CopyDhcpOptions: m.CopyDhcpOptions.ValueBoolPointer(),
		Name:            m.Name.ValueStringPointer(),
		Recursive:       m.Recursive.ValueBoolPointer(),
		SkipOnError:     m.SkipOnError.ValueBoolPointer(),
		Space:           m.Space.ValueString(),
	}
	return to
}

func flattenIpamsvcCopySubnet(ctx context.Context, from *ipam.IpamsvcCopySubnet, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcCopySubnetAttrTypes)
	}
	m := IpamsvcCopySubnetModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcCopySubnetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcCopySubnetModel) flatten(ctx context.Context, from *ipam.IpamsvcCopySubnet, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcCopySubnetModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CopyDhcpOptions = types.BoolPointerValue(from.CopyDhcpOptions)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringPointerValue(from.Name)
	m.Recursive = types.BoolPointerValue(from.Recursive)
	m.SkipOnError = types.BoolPointerValue(from.SkipOnError)
	m.Space = types.StringValue(from.Space)

}
