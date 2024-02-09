package ipam

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
    "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
    "github.com/hashicorp/terraform-plugin-framework/attr"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

    "github.com/infobloxopen/bloxone-go-client/ipam"

    "github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcHAGroupModel struct {
    AnycastConfigId types.String      `tfsdk:"anycast_config_id"`
    Comment         types.String      `tfsdk:"comment"`
    CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
    Hosts           types.List        `tfsdk:"hosts"`
    Id              types.String      `tfsdk:"id"`
    IpSpace         types.String      `tfsdk:"ip_space"`
    Mode            types.String      `tfsdk:"mode"`
    Name            types.String      `tfsdk:"name"`
    Status          types.String      `tfsdk:"status"`
    Tags            types.Map         `tfsdk:"tags"`
    TagsAll         types.Map         `tfsdk:"tags_all"`
    UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
    CollectStats    types.Bool        `tfsdk:"collect_stats"`
}

var IpamsvcHAGroupAttrTypes = map[string]attr.Type{
    "anycast_config_id": types.StringType,
    "comment":           types.StringType,
    "created_at":        timetypes.RFC3339Type{},
    "hosts":             types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHAGroupHostAttrTypes}},
    "id":                types.StringType,
    "ip_space":          types.StringType,
    "mode":              types.StringType,
    "name":              types.StringType,
    "status":            types.StringType,
    "tags":              types.MapType{ElemType: types.StringType},
    "tags_all":          types.MapType{ElemType: types.StringType},
    "updated_at":        timetypes.RFC3339Type{},
    "collect_stats":     types.BoolType,
}

var IpamsvcHAGroupResourceSchemaAttributes = map[string]schema.Attribute{
    "anycast_config_id": schema.StringAttribute{
        Optional:            true,
        MarkdownDescription: "The resource identifier.",
    },
    "comment": schema.StringAttribute{
        Optional: true,
        Validators: []validator.String{
            stringvalidator.LengthBetween(0, 1024),
        },
        MarkdownDescription: "The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.",
    },
    "created_at": schema.StringAttribute{
        CustomType:          timetypes.RFC3339Type{},
        Computed:            true,
        MarkdownDescription: "Time when the object has been created.",
    },
    "hosts": schema.ListNestedAttribute{
        NestedObject: schema.NestedAttributeObject{
            Attributes: IpamsvcHAGroupHostResourceSchemaAttributes,
        },
        Required:            true,
        MarkdownDescription: "The list of hosts.",
    },
    "id": schema.StringAttribute{
        Computed: true,
        PlanModifiers: []planmodifier.String{
            stringplanmodifier.UseStateForUnknown(),
        },
        MarkdownDescription: "The resource identifier.",
    },
    "ip_space": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The resource identifier.",
    },
    "mode": schema.StringAttribute{
        Required: true,
        Validators: []validator.String{
            stringvalidator.OneOf("active-active", "active-passive", "advanced-active-passive", "anycast"),
        },
        MarkdownDescription: "The mode of the HA group. Valid values are:\n" +
                "  * _active-active_: Both on-prem hosts remain active.\n" +
                "  * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over.\n" +
                "  * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.",
    },
    "name": schema.StringAttribute{
        Required: true,
        Validators: []validator.String{
            stringvalidator.LengthBetween(1, 256),
        },
        MarkdownDescription: "The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.",
    },
    "status": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.",
    },
    "tags": schema.MapAttribute{
        ElementType:         types.StringType,
        Optional:            true,
        MarkdownDescription: "The tags for the HA group.",
    },
    "tags_all": schema.MapAttribute{
        ElementType:         types.StringType,
        Computed:            true,
        MarkdownDescription: "The tags for the HA group including default tags.",
    },
    "updated_at": schema.StringAttribute{
        CustomType:          timetypes.RFC3339Type{},
        Computed:            true,
        MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
    },
    "collect_stats": schema.BoolAttribute{
        Computed:    true,
        Optional:    true,
        Description: "collect_stats gets the HA group stats(state, status, heartbeat) if set to true. Defaults to false",
        Default:     booldefault.StaticBool(false),
    },
}

func ExpandIpamsvcHAGroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroup {
    if o.IsNull() || o.IsUnknown() {
        return nil
    }
    var m IpamsvcHAGroupModel
    diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
    if diags.HasError() {
        return nil
    }
    return m.Expand(ctx, diags)
}

func (m *IpamsvcHAGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroup {
    if m == nil {
        return nil
    }
    to := &ipam.IpamsvcHAGroup{
        AnycastConfigId: flex.ExpandStringPointer(m.AnycastConfigId),
        Comment:         flex.ExpandStringPointer(m.Comment),
        Hosts:           flex.ExpandFrameworkListNestedBlock(ctx, m.Hosts, diags, ExpandIpamsvcHAGroupHost),
        IpSpace:         flex.ExpandStringPointer(m.IpSpace),
        Mode:            flex.ExpandStringPointer(m.Mode),
        Name:            flex.ExpandString(m.Name),
        Status:          flex.ExpandStringPointer(m.Status),
        Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
    }
    return to
}

func FlattenIpamsvcHAGroup(ctx context.Context, from *ipam.IpamsvcHAGroup, diags *diag.Diagnostics) types.Object {
    if from == nil {
        return types.ObjectNull(IpamsvcHAGroupAttrTypes)
    }
    m := IpamsvcHAGroupModel{}
    m.Flatten(ctx, from, diags)
    t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupAttrTypes, m)
    diags.Append(d...)
    return t
}

func (m *IpamsvcHAGroupModel) Flatten(ctx context.Context, from *ipam.IpamsvcHAGroup, diags *diag.Diagnostics) {
    if from == nil {
        return
    }
    if m == nil {
        *m = IpamsvcHAGroupModel{}
    }
    m.AnycastConfigId = flex.FlattenStringPointer(from.AnycastConfigId)
    m.Comment = flex.FlattenStringPointer(from.Comment)
    m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
    m.Hosts = flex.FlattenFrameworkListNestedBlock(ctx, from.Hosts, IpamsvcHAGroupHostAttrTypes, diags, FlattenIpamsvcHAGroupHost)
    m.Id = flex.FlattenStringPointer(from.Id)
    m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
    m.Mode = flex.FlattenStringPointer(from.Mode)
    m.Name = flex.FlattenString(from.Name)
    m.Status = flex.FlattenStringPointer(from.Status)
    m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
    m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
