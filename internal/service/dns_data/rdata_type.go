package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type rdata struct {
}

func (r rdata) TerraformType(ctx context.Context) tftypes.Type {
	//TODO implement me
	panic("implement me")
}

func (r rdata) ValueFromTerraform(ctx context.Context, value tftypes.Value) (attr.Value, error) {
	//TODO implement me
	panic("implement me")
}

func (r rdata) ValueType(ctx context.Context) attr.Value {
	//TODO implement me
	panic("implement me")
}

func (r rdata) Equal(t attr.Type) bool {
	//TODO implement me
	panic("implement me")
}

func (r rdata) String() string {
	//TODO implement me
	panic("implement me")
}

func (r rdata) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (r rdata) ValueFromObject(ctx context.Context, value basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	//TODO implement me
	panic("implement me")
}
