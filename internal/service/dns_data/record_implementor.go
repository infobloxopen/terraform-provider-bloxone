package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type recordModelCommon interface {
	flattenRData(ctx context.Context, m map[string]interface{}, diags *diag.Diagnostics) types.Object
	expandRData(ctx context.Context, m types.Object, diags *diag.Diagnostics) map[string]interface{}
	schemaAttributes() map[string]schema.Attribute
	attributeTypes() map[string]attr.Type
	recordType() string
	description() string
}

type recordResourceImplementor interface {
	recordModelCommon
	resourceName() string
}

type recordDataSourceImplementor interface {
	recordModelCommon
	dataSourceName() string
}
