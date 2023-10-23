package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcBulkCopyErrorModel struct {
	Description types.String `tfsdk:"description"`
	Id          types.String `tfsdk:"id"`
	Message     types.String `tfsdk:"message"`
}

var IpamsvcBulkCopyErrorAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"id":          types.StringType,
	"message":     types.StringType,
}

var IpamsvcBulkCopyErrorResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcBulkCopyErrorResourceSchemaAttributes,
}

var IpamsvcBulkCopyErrorResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description of the resource that was requested to be copied.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The reason why the copy failed.`,
	},
}

func expandIpamsvcBulkCopyError(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcBulkCopyError {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcBulkCopyErrorModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcBulkCopyErrorModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcBulkCopyError {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcBulkCopyError{
		Description: m.Description.ValueStringPointer(),
		Message:     m.Message.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcBulkCopyError(ctx context.Context, from *ipam.IpamsvcBulkCopyError, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcBulkCopyErrorAttrTypes)
	}
	m := IpamsvcBulkCopyErrorModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcBulkCopyErrorAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcBulkCopyErrorModel) flatten(ctx context.Context, from *ipam.IpamsvcBulkCopyError, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcBulkCopyErrorModel{}
	}

	m.Description = types.StringPointerValue(from.Description)
	m.Id = types.StringPointerValue(from.Id)
	m.Message = types.StringPointerValue(from.Message)

}
