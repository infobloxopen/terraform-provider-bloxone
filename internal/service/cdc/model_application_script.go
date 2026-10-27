package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ApplicationScriptModel struct {
	Content  types.String `tfsdk:"content"`
	Filename types.String `tfsdk:"filename"`
}

var ApplicationScriptAttrTypes = map[string]attr.Type{
	"content":  types.StringType,
	"filename": types.StringType,
}

var ApplicationScriptResourceSchemaAttributes = map[string]schema.Attribute{
	"content": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The content of the script, e.g. `file(\"script_1.py\")`.",
	},
	"filename": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The file name of the script.",
	},
}

func ExpandApplicationScript(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.ApplicationScript {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ApplicationScriptModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ApplicationScriptModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.ApplicationScript {
	if m == nil {
		return nil
	}
	to := &cdc.ApplicationScript{
		Content:  flex.ExpandStringPointer(m.Content),
		Filename: flex.ExpandStringPointer(m.Filename),
	}
	return to
}

func FlattenApplicationScript(ctx context.Context, from *cdc.ApplicationScript, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApplicationScriptAttrTypes)
	}
	m := ApplicationScriptModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ApplicationScriptAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ApplicationScriptModel) Flatten(ctx context.Context, from *cdc.ApplicationScript, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ApplicationScriptModel{}
	}
	m.Content = flex.FlattenStringPointerWithNilAsEmpty(from.Content)
	m.Filename = flex.FlattenStringPointer(from.Filename)
}
