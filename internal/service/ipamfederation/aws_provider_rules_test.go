package ipamfederation

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/option"
)

// fakeFederationAPI answers the realm and pool reads with the given provider type.
func fakeFederationAPI(t *testing.T, providerType string) *universalddiclient.APIClient {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"result":{"id":"x","name":"x","federated_realm":"federation/federated_realm/x","protocol":"ip4","region":"","provider":%q}}`, providerType)
	}))
	t.Cleanup(srv.Close)
	return universalddiclient.NewAPIClient(option.WithCSPUrl(srv.URL), option.WithAPIKey("test"))
}

var networkComplianceValue = tftypes.NewValue(
	tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"default_netmask_length": tftypes.Number, "maximum_netmask_length": tftypes.Number, "minimum_netmask_length": tftypes.Number,
	}},
	map[string]tftypes.Value{
		"default_netmask_length": tftypes.NewValue(tftypes.Number, 26),
		"maximum_netmask_length": tftypes.NewValue(tftypes.Number, 28),
		"minimum_netmask_length": tftypes.NewValue(tftypes.Number, 25),
	},
)

// objectWith returns a value of the schema's type with every attribute null except the given ones.
func objectWith(ctx context.Context, t *testing.T, s interface {
	Type() attr.Type
}, set map[string]tftypes.Value) tftypes.Value {
	t.Helper()
	objType := s.Type().TerraformType(ctx).(tftypes.Object)
	vals := map[string]tftypes.Value{}
	for name, typ := range objType.AttributeTypes {
		vals[name] = tftypes.NewValue(typ, nil)
	}
	for name, v := range set {
		vals[name] = v
	}
	return tftypes.NewValue(objType, vals)
}

func TestFederatedPoolModifyPlan(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		providerType string
		nc           bool
		realmUnknown bool
		wantErr      bool
	}{
		{"AWS realm with network_compliance", "AWS", true, false, true},
		{"AZURE realm is not supported", "AZURE", false, false, true},
		{"NIOS realm is not supported", "NIOS", false, false, true},
		{"NIOS_X realm without network_compliance", "NIOS_X", false, false, false},
		{"NIOS_X realm with network_compliance", "NIOS_X", true, false, false},
		{"AWS realm without network_compliance", "AWS", false, false, false},
		{"realm not known yet", "AWS", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FederatedPoolResource{client: fakeFederationAPI(t, tt.providerType)}
			var sresp resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &sresp)

			set := map[string]tftypes.Value{
				"name":            tftypes.NewValue(tftypes.String, "p"),
				"protocol":        tftypes.NewValue(tftypes.String, "ip4"),
				"federated_realm": tftypes.NewValue(tftypes.String, "federation/federated_realm/x"),
			}
			if tt.realmUnknown {
				set["federated_realm"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
			}
			if tt.nc {
				set["network_compliance"] = networkComplianceValue
			}
			raw := objectWith(ctx, t, sresp.Schema, set)
			req := resource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: sresp.Schema, Raw: raw},
				Plan:   tfsdk.Plan{Schema: sresp.Schema, Raw: raw},
				State:  tfsdk.State{Schema: sresp.Schema, Raw: tftypes.NewValue(raw.Type(), nil)},
			}
			resp := &resource.ModifyPlanResponse{}
			r.ModifyPlan(ctx, req, resp)
			if resp.Diagnostics.HasError() != tt.wantErr {
				t.Fatalf("error = %v, want %v (%v)", resp.Diagnostics.HasError(), tt.wantErr, resp.Diagnostics)
			}
			if tt.wantErr && !strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tt.providerType) {
				t.Errorf("error should name the provider type: %v", resp.Diagnostics)
			}
		})
	}
}

func TestFederatedPoolModifyPlan_Destroy(t *testing.T) {
	ctx := context.Background()
	r := &FederatedPoolResource{client: fakeFederationAPI(t, "AWS")}
	var sresp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &sresp)
	null := tftypes.NewValue(sresp.Schema.Type().TerraformType(ctx), nil)
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(ctx, resource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: sresp.Schema, Raw: null},
		Plan:   tfsdk.Plan{Schema: sresp.Schema, Raw: null},
		State:  tfsdk.State{Schema: sresp.Schema, Raw: null},
	}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
}

func TestFederatedBlockModifyPlan(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		providerType string
		nc           bool
		poolUnknown  bool
		wantErr      bool
	}{
		{"AWS pool with network_compliance", "AWS", true, false, true},
		{"NIOS_X pool with network_compliance", "NIOS_X", true, false, false},
		{"AWS pool without network_compliance", "AWS", false, false, false},
		{"pool not known yet", "AWS", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FederatedBlockResource{client: fakeFederationAPI(t, tt.providerType)}
			var sresp resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &sresp)

			set := map[string]tftypes.Value{
				"cidr":              tftypes.NewValue(tftypes.Number, 24),
				"federated_realm":   tftypes.NewValue(tftypes.String, "federation/federated_realm/x"),
				"federated_pool_id": tftypes.NewValue(tftypes.String, "federation/federated_pool/x"),
			}
			if tt.poolUnknown {
				set["federated_pool_id"] = tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
			}
			if tt.nc {
				set["network_compliance"] = networkComplianceValue
			}
			raw := objectWith(ctx, t, sresp.Schema, set)
			resp := &resource.ModifyPlanResponse{}
			r.ModifyPlan(ctx, resource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: sresp.Schema, Raw: raw},
				Plan:   tfsdk.Plan{Schema: sresp.Schema, Raw: raw},
				State:  tfsdk.State{Schema: sresp.Schema, Raw: tftypes.NewValue(raw.Type(), nil)},
			}, resp)
			if resp.Diagnostics.HasError() != tt.wantErr {
				t.Fatalf("error = %v, want %v (%v)", resp.Diagnostics.HasError(), tt.wantErr, resp.Diagnostics)
			}
		})
	}
}

func TestFederatedModifyPlan_ExistingResourcesAreNotChecked(t *testing.T) {
	ctx := context.Background()
	nc := networkComplianceValue

	// An existing AWS pool and block whose configuration differs from the state: a destroy must not be blocked.
	pool := &FederatedPoolResource{client: fakeFederationAPI(t, "AWS")}
	var ps resource.SchemaResponse
	pool.Schema(ctx, resource.SchemaRequest{}, &ps)
	poolRaw := objectWith(ctx, t, ps.Schema, map[string]tftypes.Value{
		"name": tftypes.NewValue(tftypes.String, "p"), "protocol": tftypes.NewValue(tftypes.String, "ip4"),
		"federated_realm": tftypes.NewValue(tftypes.String, "federation/federated_realm/x"), "network_compliance": nc,
		"provider_type": tftypes.NewValue(tftypes.String, "AWS"),
	})
	presp := &resource.ModifyPlanResponse{}
	pool.ModifyPlan(ctx, resource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: ps.Schema, Raw: poolRaw},
		Plan:   tfsdk.Plan{Schema: ps.Schema, Raw: poolRaw},
		State:  tfsdk.State{Schema: ps.Schema, Raw: objectWith(ctx, t, ps.Schema, map[string]tftypes.Value{"provider_type": tftypes.NewValue(tftypes.String, "AWS")})},
	}, presp)
	if presp.Diagnostics.HasError() {
		t.Errorf("existing pool must not be checked: %v", presp.Diagnostics)
	}

	block := &FederatedBlockResource{client: fakeFederationAPI(t, "AWS")}
	var bs resource.SchemaResponse
	block.Schema(ctx, resource.SchemaRequest{}, &bs)
	blockRaw := objectWith(ctx, t, bs.Schema, map[string]tftypes.Value{
		"cidr": tftypes.NewValue(tftypes.Number, 24), "federated_realm": tftypes.NewValue(tftypes.String, "federation/federated_realm/x"),
		"federated_pool_id": tftypes.NewValue(tftypes.String, "federation/federated_pool/x"), "network_compliance": nc,
	})
	bresp := &resource.ModifyPlanResponse{}
	block.ModifyPlan(ctx, resource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: bs.Schema, Raw: blockRaw},
		Plan:   tfsdk.Plan{Schema: bs.Schema, Raw: blockRaw},
		State:  tfsdk.State{Schema: bs.Schema, Raw: objectWith(ctx, t, bs.Schema, map[string]tftypes.Value{"cidr": tftypes.NewValue(tftypes.Number, 24)})},
	}, bresp)
	if bresp.Diagnostics.HasError() {
		t.Errorf("existing block must not be checked: %v", bresp.Diagnostics)
	}
}

func TestFederatedPoolUpdate_AWSAllowsNameAndDescriptionButNotTags(t *testing.T) {
	ctx := context.Background()
	tagsType := tftypes.Map{ElementType: tftypes.String}
	tags := func(site string) tftypes.Value {
		return tftypes.NewValue(tagsType, map[string]tftypes.Value{"site": tftypes.NewValue(tftypes.String, site)})
	}
	attrs := func(name, description, site string) map[string]tftypes.Value {
		return map[string]tftypes.Value{
			"id": tftypes.NewValue(tftypes.String, "federation/federated_pool/x"), "name": tftypes.NewValue(tftypes.String, name),
			"description": tftypes.NewValue(tftypes.String, description), "protocol": tftypes.NewValue(tftypes.String, "ip4"),
			"provider_type": tftypes.NewValue(tftypes.String, "AWS"), "tags": tags(site),
		}
	}
	tests := []struct {
		name                         string
		planName, planDesc, planSite string
		wantErr                      bool
	}{
		{"name changed", "renamed", "", "A", false},
		{"description changed", "p", "new", "A", false},
		{"tags changed", "p", "", "B", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FederatedPoolResource{client: fakeFederationAPI(t, "AWS")}
			var sresp resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &sresp)
			state := objectWith(ctx, t, sresp.Schema, attrs("p", "", "A"))
			resp := &resource.UpdateResponse{State: tfsdk.State{Schema: sresp.Schema, Raw: state}}
			r.Update(ctx, resource.UpdateRequest{
				Plan:  tfsdk.Plan{Schema: sresp.Schema, Raw: objectWith(ctx, t, sresp.Schema, attrs(tt.planName, tt.planDesc, tt.planSite))},
				State: tfsdk.State{Schema: sresp.Schema, Raw: state},
			}, resp)
			if resp.Diagnostics.HasError() != tt.wantErr {
				t.Fatalf("error = %v, want %v (%v)", resp.Diagnostics.HasError(), tt.wantErr, resp.Diagnostics)
			}
		})
	}
}
