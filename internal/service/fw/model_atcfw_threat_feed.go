package fw

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwThreatFeedModel struct {
	ConfidenceLevel types.String `tfsdk:"confidence_level"`
	Description     types.String `tfsdk:"description"`
	Key             types.String `tfsdk:"key"`
	Name            types.String `tfsdk:"name"`
	Source          types.String `tfsdk:"source"`
	ThreatLevel     types.String `tfsdk:"threat_level"`
}

var AtcfwThreatFeedAttrTypes = map[string]attr.Type{
	"confidence_level": types.StringType,
	"description":      types.StringType,
	"key":              types.StringType,
	"name":             types.StringType,
	"source":           types.StringType,
	"threat_level":     types.StringType,
}

var AtcfwThreatFeedResourceSchemaAttributes = map[string]schema.Attribute{
	"confidence_level": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Confidence Level of the Feed.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The brief description for the thread feed.",
	},
	"key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key of the threat feed.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the thread feed.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("Infoblox"),
		MarkdownDescription: "The source of the threat feed.",
	},
	"threat_level": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Threat Level of the Feed.",
	},
}

func ExpandAtcfwThreatFeed(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.ThreatFeed {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwThreatFeedModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwThreatFeedModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.ThreatFeed {
	if m == nil {
		return nil
	}
	to := &fw.ThreatFeed{
		Source: ExpandThreatFeedSource(m.Source),
	}
	return to
}

func FlattenAtcfwThreatFeed(ctx context.Context, from *fw.ThreatFeed, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwThreatFeedAttrTypes)
	}
	m := AtcfwThreatFeedModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwThreatFeedAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwThreatFeedModel) Flatten(ctx context.Context, from *fw.ThreatFeed, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwThreatFeedModel{}
	}
	m.ConfidenceLevel = flex.FlattenStringPointer(from.ConfidenceLevel)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ThreatLevel = flex.FlattenStringPointer(from.ThreatLevel)
	m.Source = FlattenThreatFeedSource(from.Source)
}

func ExpandThreatFeedSource(s types.String) *fw.ThreatFeedSource {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	source := fw.ThreatFeedSource(s.ValueString())
	return &source
}

func FlattenThreatFeedSource(s *fw.ThreatFeedSource) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(string(*s))
}
