package ipamfederation

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestFederatedPoolExpand_NetworkCompliance(t *testing.T) {
	ctx := context.Background()
	configured := types.ObjectValueMust(NetworkComplianceAttrTypes, map[string]attr.Value{
		"default_netmask_length": types.Int64Value(26),
		"maximum_netmask_length": types.Int64Value(28),
		"minimum_netmask_length": types.Int64Value(25),
	})
	tests := []struct {
		name         string
		providerType types.String
		nc           types.Object
		isCreate     bool
		wantSent     bool
		wantErr      bool
	}{
		{"NIOS_X sends it", types.StringValue("NIOS_X"), configured, true, true, false},
		{"provider unset sends it", types.StringUnknown(), configured, true, true, false},
		{"AWS configured on create is rejected", types.StringValue("AWS"), configured, true, false, true},
		{"AWS not configured on create", types.StringValue("AWS"), types.ObjectUnknown(NetworkComplianceAttrTypes), true, false, false},
		{"AWS null on create", types.StringValue("AWS"), types.ObjectNull(NetworkComplianceAttrTypes), true, false, false},
		{"AWS update never sends it", types.StringValue("AWS"), configured, false, false, false},
		{"NIOS_X update sends it", types.StringValue("NIOS_X"), configured, false, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FederatedPoolModel{
				Name:              types.StringValue("p"),
				ProviderType:      tt.providerType,
				NetworkCompliance: tt.nc,
				Tags:              types.MapNull(types.StringType),
			}
			var diags diag.Diagnostics
			to := m.Expand(ctx, &diags, tt.isCreate)
			if diags.HasError() != tt.wantErr {
				t.Fatalf("error = %v, want %v (%v)", diags.HasError(), tt.wantErr, diags)
			}
			if tt.wantErr {
				return
			}
			if sent := to.NetworkCompliance != nil; sent != tt.wantSent {
				t.Errorf("network_compliance sent = %v, want %v", sent, tt.wantSent)
			}
		})
	}
}

func TestFederatedPoolExpand_UpdateOfAWSPoolSendsOnlyNameAndDescription(t *testing.T) {
	ctx := context.Background()
	tags := types.MapValueMust(types.StringType, map[string]attr.Value{"site": types.StringValue("A")})
	for _, tt := range []struct {
		providerType               string
		wantOnlyNameAndDescription bool
	}{
		{"AWS", true}, {"NIOS_X", false}, {"", false},
	} {
		t.Run(tt.providerType, func(t *testing.T) {
			m := &FederatedPoolModel{
				Name:              types.StringValue("p"),
				Description:       types.StringValue("d"),
				Tags:              tags,
				ProviderType:      types.StringValue(tt.providerType),
				NetworkCompliance: types.ObjectNull(NetworkComplianceAttrTypes),
			}
			var diags diag.Diagnostics
			to := m.Expand(ctx, &diags, false)
			if diags.HasError() {
				t.Fatal(diags)
			}
			onlyNameAndDescription := to.Name != nil && to.Description != nil && to.Tags == nil && to.NetworkCompliance == nil
			if onlyNameAndDescription != tt.wantOnlyNameAndDescription {
				t.Errorf("only name and description sent = %v, want %v (%+v)", onlyNameAndDescription, tt.wantOnlyNameAndDescription, to)
			}
		})
	}
}
