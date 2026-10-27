package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ApplicationDataContentModel struct {
	Scripts types.List `tfsdk:"scripts"`
}

var ApplicationDataContentAttrTypes = map[string]attr.Type{
	"scripts": types.ListType{ElemType: types.ObjectType{AttrTypes: ApplicationScriptAttrTypes}},
}

var ApplicationDataContentResourceSchemaAttributes = map[string]schema.Attribute{
	"scripts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ApplicationScriptResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The scripts of the application.",
	},
}

func ExpandApplicationDataContent(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.ApplicationDataContent {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ApplicationDataContentModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ApplicationDataContentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.ApplicationDataContent {
	if m == nil {
		return nil
	}
	to := &cdc.ApplicationDataContent{
		Scripts: flex.ExpandFrameworkListNestedBlock(ctx, m.Scripts, diags, ExpandApplicationScript),
	}
	return to
}

func FlattenApplicationDataContent(ctx context.Context, from *cdc.ApplicationDataContent, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApplicationDataContentAttrTypes)
	}
	m := ApplicationDataContentModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ApplicationDataContentAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ApplicationDataContentModel) Flatten(ctx context.Context, from *cdc.ApplicationDataContent, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ApplicationDataContentModel{}
	}
	m.Scripts = flex.FlattenFrameworkListNestedBlock(ctx, from.Scripts, ApplicationScriptAttrTypes, diags, FlattenApplicationScript)
}
