package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcOptionGroupModel struct {
	Comment     types.String      `tfsdk:"comment"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions types.List        `tfsdk:"dhcp_options"`
	Id          types.String      `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Protocol    types.String      `tfsdk:"protocol"`
	Tags        types.Map         `tfsdk:"tags"`
	TagsAll     types.Map         `tfsdk:"tags_all"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcOptionGroupAttrTypes = map[string]attr.Type{
	"comment":      types.StringType,
	"created_at":   timetypes.RFC3339Type{},
	"dhcp_options": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"id":           types.StringType,
	"name":         types.StringType,
	"protocol":     types.StringType,
	"tags":         types.MapType{ElemType: types.StringType},
	"tags_all":     types.MapType{ElemType: types.StringType},
	"updated_at":   timetypes.RFC3339Type{},
}

var IpamsvcOptionGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP options for the option group. May be either a specific option or a group of options.",
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
		MarkdownDescription: "The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"protocol": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ip4", "ip6"),
		},
		MarkdownDescription: "The type of protocol (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the option group in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the option group in JSON format including default tag.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func (m *IpamsvcOptionGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.OptionGroup {
	if m == nil {
		return nil
	}
	to := &ipam.OptionGroup{
		Comment:     flex.ExpandStringPointer(m.Comment),
		DhcpOptions: flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		Name:        flex.ExpandString(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.Protocol = flex.ExpandStringPointer(m.Protocol)
	}
	return to
}

func FlattenIpamsvcOptionGroupDataSource(ctx context.Context, from *ipam.OptionGroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionGroupAttrTypes)
	}
	m := IpamsvcOptionGroupModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionGroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionGroupModel) Flatten(ctx context.Context, from *ipam.OptionGroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionGroupModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
