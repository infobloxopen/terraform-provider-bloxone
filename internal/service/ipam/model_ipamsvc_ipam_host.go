package ipam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcIpamHostModel struct {
	Addresses           types.List        `tfsdk:"addresses"`
	AutoGenerateRecords types.Bool        `tfsdk:"auto_generate_records"`
	Comment             types.String      `tfsdk:"comment"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	HostNames           types.List        `tfsdk:"host_names"`
	Id                  types.String      `tfsdk:"id"`
	Name                types.String      `tfsdk:"name"`
	Tags                types.Map         `tfsdk:"tags"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcIpamHostAttrTypes = map[string]attr.Type{
	"addresses":             types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHostAddressAttrTypes}},
	"auto_generate_records": types.BoolType,
	"comment":               types.StringType,
	"created_at":            timetypes.RFC3339Type{},
	"host_names":            types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHostNameAttrTypes}},
	"id":                    types.StringType,
	"name":                  types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"updated_at":            timetypes.RFC3339Type{},
}

var IpamsvcIpamHostResourceSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcHostAddressResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of all addresses associated with the IPAM host, which may be in different IP spaces.`,
	},
	"auto_generate_records": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `This flag specifies if resource records have to be auto generated for the host.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"host_names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcHostNameResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The name records to be generated for the host.  This field is required if _auto_generate_records_ is true.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the IPAM host in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func ExpandIpamsvcIpamHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcIpamHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcIpamHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcIpamHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcIpamHost {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcIpamHost{
		Addresses:           flex.ExpandFrameworkListNestedBlock(ctx, m.Addresses, diags, ExpandIpamsvcHostAddress),
		AutoGenerateRecords: flex.ExpandBoolPointer(m.AutoGenerateRecords),
		Comment:             flex.ExpandStringPointer(m.Comment),
		HostNames:           flex.ExpandFrameworkListNestedBlock(ctx, m.HostNames, diags, ExpandIpamsvcHostName),
		Name:                flex.ExpandString(m.Name),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenIpamsvcIpamHost(ctx context.Context, from *ipam.IpamsvcIpamHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIpamHostAttrTypes)
	}
	m := IpamsvcIpamHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIpamHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIpamHostModel) Flatten(ctx context.Context, from *ipam.IpamsvcIpamHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIpamHostModel{}
	}
	m.Addresses = flex.FlattenFrameworkListNestedBlock(ctx, from.Addresses, IpamsvcHostAddressAttrTypes, diags, FlattenIpamsvcHostAddress)
	m.AutoGenerateRecords = types.BoolPointerValue(from.AutoGenerateRecords)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.HostNames = flex.FlattenFrameworkListNestedBlock(ctx, from.HostNames, IpamsvcHostNameAttrTypes, diags, FlattenIpamsvcHostName)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
