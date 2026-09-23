package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-bloxone/internal/types"
)

type ZoneFilterModel struct {
	Action    types.String                     `tfsdk:"action"`
	Wildcards internaltypes.UnorderedListValue `tfsdk:"wildcards"`
}

var ZoneFilterAttrTypes = map[string]attr.Type{
	"action":    types.StringType,
	"wildcards": internaltypes.UnorderedListOfStringType,
}

var ZoneFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("include", "exclude"),
		},
		MarkdownDescription: "Action to take on matching zones. Allowed values: \"include\", \"exclude\".",
	},
	"wildcards": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "List of zone wildcard patterns to include or exclude.",
	},
}

func ExpandZoneFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.ZoneFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.ZoneFilter {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.ZoneFilter{
		Action:    flex.ExpandStringPointer(m.Action),
		Wildcards: flex.ExpandFrameworkListString(ctx, m.Wildcards, diags),
	}
	return to
}

func FlattenZoneFilter(ctx context.Context, from *clouddiscovery.ZoneFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneFilterAttrTypes)
	}
	m := ZoneFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneFilterModel) Flatten(ctx context.Context, from *clouddiscovery.ZoneFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneFilterModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Wildcards = flex.FlattenFrameworkUnorderedListNotNull(ctx, types.StringType, from.Wildcards, diags)
}
