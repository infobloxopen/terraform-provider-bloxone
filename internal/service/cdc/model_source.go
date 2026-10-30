package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Tags        types.Map         `tfsdk:"tags"`
	Type        types.String      `tfsdk:"type"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
}

var SourceAttrTypes = map[string]attr.Type{
	"created_at":  timetypes.RFC3339Type{},
	"description": types.StringType,
	"enabled":     types.BoolType,
	"id":          types.Int64Type,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"type":        types.StringType,
	"updated_at":  timetypes.RFC3339Type{},
}

// SourceResourceSchemaAttributes is only used to build the read-only
// bloxone_cdc_sources data source schema; sources of all types are listed,
// including the built-in ones that cannot be managed by Terraform.
var SourceResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description of the source.",
	},
	"enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the source is enabled.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the source.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the source.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the source, e.g. _SOURCE_BLOXONE_ (Infoblox Cloud Source), _SOURCE_SCHEDULE_ (Schedule Source) or _SOURCE_NIOS_.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func FlattenSource(ctx context.Context, from *cdc.Source, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceAttrTypes)
	}
	m := SourceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceModel) Flatten(ctx context.Context, from *cdc.Source, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
