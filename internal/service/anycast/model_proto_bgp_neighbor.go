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

type ProtoBgpNeighborModel struct {
	Asn         types.Int64  `tfsdk:"asn"`
	AsnText     types.String `tfsdk:"asn_text"`
	IpAddress   types.String `tfsdk:"ip_address"`
	MaxHopCount types.Int64  `tfsdk:"max_hop_count"`
	Multihop    types.Bool   `tfsdk:"multihop"`
	Password    types.String `tfsdk:"password"`
}

var ProtoBgpNeighborAttrTypes = map[string]attr.Type{
	"asn":           types.Int64Type,
	"asn_text":      types.StringType,
	"ip_address":    types.StringType,
	"max_hop_count": types.Int64Type,
	"multihop":      types.BoolType,
	"password":      types.StringType,
}

var ProtoBgpNeighborResourceSchemaAttributes = map[string]schema.Attribute{
	"asn": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: `Autonomous system number of this BGP/anycast enabled on-prem host.`,
	},
	"asn_text": schema.StringAttribute{
		Computed: true,
		Optional: true,
		MarkdownDescription: `Autonomous system as text (supported in ASDOT or ASPLAIN format) Optional, requires the asn field to be set to the equivalent integer value of the ASDOT/ASPLAIN string contained in this field or be unset/zero.
		Example:

		| ASDOT       | ASPLAIN     | INTEGER     | VALID/INVALID |
		|-------------|-------------|-------------|---------------|
		| 0.1         | 1           | 1           | Valid         |
		| 1           | 1           | 1           | Valid         |
		| 65535       | 65535       | 65535       | Valid         |
		| 0.65535     | 65535       | 65535       | Valid         |
		| 1.0         | 65536       | 65536       | Valid         |
		| 1.1         | 65537       | 65537       | Valid         |
		| 1.65535     | 131071      | 131071      | Valid         |
		| 65535.0     | 4294901760  | 4294901760  | Valid         |
		| 65535.1     | 4294901761  | 4294901761  | Valid         |
		| 65535.65535 | 4294967295  | 4294967295  | Valid         |
		| 0.65536     |             |             | Invalid       |
		| 65535.655536|             |             | Invalid       |
		| 65536.0     |             |             | Invalid       |
		| 65536.65535 |             |             | Invalid       |
		|             | 4294967296  |             | Invalid       | 
`,
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv4 address of the BGP neighbor",
	},
	"max_hop_count": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Max hop count, if BGP multihop is enabled.`,
	},
	"multihop": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `BGP multihop enabled or not.`,
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `BGP protocol access password for this BGP neighbor, max 25 characters long.`,
	},
}

func ExpandProtoBgpNeighbor(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.BgpNeighbor {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoBgpNeighborModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoBgpNeighborModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.BgpNeighbor {
	if m == nil {
		return nil
	}
	to := &anycast.BgpNeighbor{
		Asn:         flex.ExpandInt64Pointer(m.Asn),
		AsnText:     flex.ExpandStringPointer(m.AsnText),
		IpAddress:   flex.ExpandStringPointer(m.IpAddress),
		MaxHopCount: flex.ExpandInt64Pointer(m.MaxHopCount),
		Multihop:    flex.ExpandBoolPointer(m.Multihop),
		Password:    flex.ExpandStringPointer(m.Password),
	}
	return to
}

func FlattenProtoBgpNeighbor(ctx context.Context, from *anycast.BgpNeighbor, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoBgpNeighborAttrTypes)
	}
	m := ProtoBgpNeighborModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoBgpNeighborAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoBgpNeighborModel) Flatten(ctx context.Context, from *anycast.BgpNeighbor, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoBgpNeighborModel{}
	}
	m.Asn = flex.FlattenInt64Pointer(from.Asn)
	m.AsnText = flex.FlattenStringPointer(from.AsnText)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.MaxHopCount = flex.FlattenInt64Pointer(from.MaxHopCount)
	m.Multihop = types.BoolPointerValue(from.Multihop)
	m.Password = flex.FlattenStringPointer(from.Password)
}
