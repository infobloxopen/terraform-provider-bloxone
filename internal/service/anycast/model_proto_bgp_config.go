package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoBgpConfigModel struct {
	Asn           types.Int64  `tfsdk:"asn"`
	AsnText       types.String `tfsdk:"asn_text"`
	Fields        types.Object `tfsdk:"fields"`
	HolddownSecs  types.Int64  `tfsdk:"holddown_secs"`
	KeepAliveSecs types.Int64  `tfsdk:"keep_alive_secs"`
	LinkDetect    types.Bool   `tfsdk:"link_detect"`
	Neighbors     types.List   `tfsdk:"neighbors"`
	Preamble      types.String `tfsdk:"preamble"`
}

var ProtoBgpConfigAttrTypes = map[string]attr.Type{
	"asn":             types.Int64Type,
	"asn_text":        types.StringType,
	"fields":          types.ObjectType{AttrTypes: ProtobufFieldMaskAttrTypes},
	"holddown_secs":   types.Int64Type,
	"keep_alive_secs": types.Int64Type,
	"link_detect":     types.BoolType,
	"neighbors":       types.ListType{ElemType: types.ObjectType{AttrTypes: ProtoBgpNeighborAttrTypes}},
	"preamble":        types.StringType,
}

var ProtoBgpConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"asn": schema.Int64Attribute{
		Required: true,
	},
	"asn_text": schema.StringAttribute{
		// TODO: keeping computed for now, but should verify if this can be provided as an alternative to asn
		Computed: true,

		MarkdownDescription: "Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid",
	},
	"fields": schema.SingleNestedAttribute{
		Attributes: ProtobufFieldMaskResourceSchemaAttributes,
		Optional:   true,
	},
	"holddown_secs": schema.Int64Attribute{
		Required: true,
	},
	"keep_alive_secs": schema.Int64Attribute{
		Optional: true,
		Computed: true,
	},
	"link_detect": schema.BoolAttribute{
		Optional: true,
	},
	"neighbors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ProtoBgpNeighborResourceSchemaAttributes,
		},
		Optional: true,
	},
	"preamble": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Any predefined BGP configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.",
	},
}

func ExpandProtoBgpConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtoBgpConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoBgpConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoBgpConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtoBgpConfig {
	if m == nil {
		return nil
	}
	to := &anycast.ProtoBgpConfig{
		Asn:           flex.ExpandInt64Pointer(m.Asn),
		AsnText:       flex.ExpandStringPointer(m.AsnText),
		Fields:        ExpandProtobufFieldMask(ctx, m.Fields, diags),
		HolddownSecs:  flex.ExpandInt64Pointer(m.HolddownSecs),
		KeepAliveSecs: flex.ExpandInt64Pointer(m.KeepAliveSecs),
		LinkDetect:    flex.ExpandBoolPointer(m.LinkDetect),
		Neighbors:     flex.ExpandFrameworkListNestedBlock(ctx, m.Neighbors, diags, ExpandProtoBgpNeighbor),
		Preamble:      flex.ExpandStringPointer(m.Preamble),
	}
	return to
}

func FlattenProtoBgpConfig(ctx context.Context, from *anycast.ProtoBgpConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoBgpConfigAttrTypes)
	}
	m := ProtoBgpConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoBgpConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoBgpConfigModel) Flatten(ctx context.Context, from *anycast.ProtoBgpConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoBgpConfigModel{}
	}
	m.Asn = flex.FlattenInt64Pointer(from.Asn)
	m.AsnText = flex.FlattenStringPointer(from.AsnText)
	m.Fields = FlattenProtobufFieldMask(ctx, from.Fields, diags)
	m.HolddownSecs = flex.FlattenInt64Pointer(from.HolddownSecs)
	m.KeepAliveSecs = flex.FlattenInt64Pointer(from.KeepAliveSecs)
	m.LinkDetect = types.BoolPointerValue(from.LinkDetect)
	m.Neighbors = flex.FlattenFrameworkListNestedBlock(ctx, from.Neighbors, ProtoBgpNeighborAttrTypes, diags, FlattenProtoBgpNeighbor)
	m.Preamble = flex.FlattenStringPointer(from.Preamble)
}
