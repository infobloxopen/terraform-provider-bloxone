package dns_config

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/attr"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

    "github.com/infobloxopen/bloxone-go-client/dns_config"

    "github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigACLModel struct {
    Comment types.String `tfsdk:"comment"`
    Id      types.String `tfsdk:"id"`
    List    types.List   `tfsdk:"list"`
    Name    types.String `tfsdk:"name"`
    Tags    types.Map    `tfsdk:"tags"`
    TagsAll types.Map    `tfsdk:"tags_all"`
}

var ConfigACLAttrTypes = map[string]attr.Type{
    "comment":  types.StringType,
    "id":       types.StringType,
    "list":     types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
    "name":     types.StringType,
    "tags":     types.MapType{ElemType: types.StringType},
    "tags_all": types.MapType{ElemType: types.StringType},
}

var ConfigACLResourceSchemaAttributes = map[string]schema.Attribute{
    "comment": schema.StringAttribute{
        Optional:            true,
        MarkdownDescription: `Optional. Comment for ACL.`,
    },
    "id": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: `The resource identifier.`,
        PlanModifiers: []planmodifier.String{
            stringplanmodifier.UseStateForUnknown(),
        },
    },
    "list": schema.ListNestedAttribute{
        NestedObject: schema.NestedAttributeObject{
            Attributes: ConfigACLItemResourceSchemaAttributes,
        },
        Optional:            true,
        MarkdownDescription: `Optional. Ordered list of access control elements.  Elements are evaluated in order to determine access. If evaluation reaches the end of the list then access is denied.`,
    },
    "name": schema.StringAttribute{
        Required:            true,
        MarkdownDescription: `ACL object name.`,
    },
    "tags": schema.MapAttribute{
        ElementType:         types.StringType,
        Optional:            true,
        MarkdownDescription: `Tagging specifics.`,
    },
    "tags_all": schema.MapAttribute{
        ElementType:         types.StringType,
        Computed:            true,
        MarkdownDescription: `Tagging specifics includes the default tags.`,
    },
}

func ExpandConfigACL(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigACL {
    if o.IsNull() || o.IsUnknown() {
        return nil
    }
    var m ConfigACLModel
    diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        return nil
    }
    return m.Expand(ctx, diags)
}

func (m *ConfigACLModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigACL {
    if m == nil {
        return nil
    }
    to := &dns_config.ConfigACL{
        Comment: flex.ExpandStringPointer(m.Comment),
        List:    flex.ExpandFrameworkListNestedBlock(ctx, m.List, diags, ExpandConfigACLItem),
        Name:    flex.ExpandString(m.Name),
        Tags:    flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
    }
    return to
}

func DataSourceFlattenConfigACL(ctx context.Context, from *dns_config.ConfigACL, diags *diag.Diagnostics) types.Object {
    if from == nil {
        return types.ObjectNull(ConfigACLAttrTypes)
    }
    m := ConfigACLModel{}
    m.Flatten(ctx, from, diags)
    m.Tags = m.TagsAll
    t, d := types.ObjectValueFrom(ctx, ConfigACLAttrTypes, m)
    diags.Append(d...)
    return t
}

func (m *ConfigACLModel) Flatten(ctx context.Context, from *dns_config.ConfigACL, diags *diag.Diagnostics) {
    if from == nil {
        return
    }
    if m == nil {
        *m = ConfigACLModel{}
    }
    m.Comment = flex.FlattenStringPointer(from.Comment)
    m.Id = flex.FlattenStringPointer(from.Id)
    m.List = flex.FlattenFrameworkListNestedBlock(ctx, from.List, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
    m.Name = flex.FlattenString(from.Name)
    m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
}
