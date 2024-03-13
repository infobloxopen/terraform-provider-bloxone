package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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

var IpamsvcOptionCodeResourceSchemaAttributes = map[string]schema.Attribute{
	"array": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether the option value is an array of the type or not.",
	},
	"code": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The option code.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"option_space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"source": schema.StringAttribute{
		Computed: true,
		MarkdownDescription: "The source for the option code. Valid values are:\n" +
			"  * _dhcp_server_\n" +
			"  * _reserved_\n" +
			"  * _blox_one_ddi_\n" +
			"  * _customer_\n\n" +
			"  Defaults to _customer_.",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("address4", "address6", "boolean", "empty", "fqdn", "int8", "int16", "int32", "text", "uint8", "uint16", "uint32"),
		},
		MarkdownDescription: "The option type for the option code. Valid values are:\n" +
			"  * _address4_\n" +
			"  * _address6_\n" +
			"  * _boolean_\n" +
			"  * _empty_\n" +
			"  * _fqdn_\n" +
			"  * _int8_\n" +
			"  * _int16_\n" +
			"  * _int32_\n" +
			"  * _text_\n" +
			"  * _uint8_\n" +
			"  * _uint16_\n" +
			"  * _uint32_",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandIpamsvcOptionCode(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionCode {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcOptionCodeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcOptionCodeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionCode {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcOptionCode{
		Array:       flex.ExpandBoolPointer(m.Array),
		Code:        flex.ExpandInt64(m.Code),
		Comment:     flex.ExpandStringPointer(m.Comment),
		Name:        flex.ExpandString(m.Name),
		OptionSpace: flex.ExpandString(m.OptionSpace),
		Type:        flex.ExpandString(m.Type),
	}
	return to
}

func FlattenIpamsvcOptionCode(ctx context.Context, from *ipam.IpamsvcOptionCode, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionCodeAttrTypes)
	}
	m := IpamsvcOptionCodeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionCodeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionCodeModel) Flatten(ctx context.Context, from *ipam.IpamsvcOptionCode, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionCodeModel{}
	}
	m.Array = types.BoolPointerValue(from.Array)
	m.Code = flex.FlattenInt64(from.Code)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.OptionSpace = flex.FlattenString(from.OptionSpace)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Type = flex.FlattenString(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
