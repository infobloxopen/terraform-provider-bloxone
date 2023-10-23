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

type IpamsvcCopyIPSpaceModel struct {
	Comment         types.String `tfsdk:"comment"`
	CopyDhcpOptions types.Bool   `tfsdk:"copy_dhcp_options"`
	Id              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	SkipOnError     types.Bool   `tfsdk:"skip_on_error"`
}

var IpamsvcCopyIPSpaceAttrTypes = map[string]attr.Type{
	"comment":           types.StringType,
	"copy_dhcp_options": types.BoolType,
	"id":                types.StringType,
	"name":              types.StringType,
	"skip_on_error":     types.BoolType,
}

var IpamsvcCopyIPSpaceResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcCopyIPSpaceResourceSchemaAttributes,
}

var IpamsvcCopyIPSpaceResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the copied IP space. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"copy_dhcp_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether dhcp options should be copied or not when _ipam/ip_space_ object is copied.  Defaults to _false_.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name for the copied IP space. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"skip_on_error": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether copying should skip an object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.`,
	},
}

func expandIpamsvcCopyIPSpace(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcCopyIPSpace {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcCopyIPSpaceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcCopyIPSpaceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcCopyIPSpace {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcCopyIPSpace{
		Comment:         m.Comment.ValueStringPointer(),
		CopyDhcpOptions: m.CopyDhcpOptions.ValueBoolPointer(),
		Name:            m.Name.ValueString(),
		SkipOnError:     m.SkipOnError.ValueBoolPointer(),
	}
	return to
}

func flattenIpamsvcCopyIPSpace(ctx context.Context, from *ipam.IpamsvcCopyIPSpace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcCopyIPSpaceAttrTypes)
	}
	m := IpamsvcCopyIPSpaceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcCopyIPSpaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcCopyIPSpaceModel) flatten(ctx context.Context, from *ipam.IpamsvcCopyIPSpace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcCopyIPSpaceModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CopyDhcpOptions = types.BoolPointerValue(from.CopyDhcpOptions)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringValue(from.Name)
	m.SkipOnError = types.BoolPointerValue(from.SkipOnError)

}
