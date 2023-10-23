package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcOptionCodeModel struct {
	Array       types.Bool        `tfsdk:"array"`
	Code        types.Int64       `tfsdk:"code"`
	Comment     types.String      `tfsdk:"comment"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	Id          types.String      `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	OptionSpace types.String      `tfsdk:"option_space"`
	Source      types.String      `tfsdk:"source"`
	Type        types.String      `tfsdk:"type"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcOptionCodeAttrTypes = map[string]attr.Type{
	"array":        types.BoolType,
	"code":         types.Int64Type,
	"comment":      types.StringType,
	"created_at":   timetypes.RFC3339Type{},
	"id":           types.StringType,
	"name":         types.StringType,
	"option_space": types.StringType,
	"source":       types.StringType,
	"type":         types.StringType,
	"updated_at":   timetypes.RFC3339Type{},
}

var IpamsvcOptionCodeResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionCode__ (_dhcp/option_code_) defines a DHCP option code.`,
	Attributes:          IpamsvcOptionCodeResourceSchemaAttributes,
}

var IpamsvcOptionCodeResourceSchemaAttributes = map[string]schema.Attribute{
	"array": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether the option value is an array of the type or not.`,
	},
	"code": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: `The option code.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"option_space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The source for the option code.  Valid values are:  * _dhcp_server_  * _reserved_  * _blox_one_ddi_  * _customer_  Defaults to _customer_.`,
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The option type for the option code.  Valid values are: * _address4_ * _address6_ * _boolean_ * _empty_ * _fqdn_ * _int8_ * _int16_ * _int32_ * _text_ * _uint8_ * _uint16_ * _uint32_`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func expandIpamsvcOptionCode(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionCode {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionCodeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionCodeModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionCode {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionCode{
		Array:       m.Array.ValueBoolPointer(),
		Code:        int64(m.Code.ValueInt64()),
		Comment:     m.Comment.ValueStringPointer(),
		Name:        m.Name.ValueString(),
		OptionSpace: m.OptionSpace.ValueString(),
		Type:        m.Type.ValueString(),
	}
	return to
}

func flattenIpamsvcOptionCode(ctx context.Context, from *ipam.IpamsvcOptionCode, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionCodeAttrTypes)
	}
	m := IpamsvcOptionCodeModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionCodeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionCodeModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionCode, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionCodeModel{}
	}

	m.Array = types.BoolPointerValue(from.Array)
	m.Code = types.Int64Value(int64(from.Code))
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringValue(from.Name)
	m.OptionSpace = types.StringValue(from.OptionSpace)
	m.Source = types.StringPointerValue(from.Source)
	m.Type = types.StringValue(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
