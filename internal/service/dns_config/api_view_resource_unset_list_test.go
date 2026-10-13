package dns_config_test

// Regression tests for https://github.com/infobloxopen/terraform-provider-bloxone/issues/138
//
// They run the real provider against an in-process mock of /api/ddi/v1/dns/view that
// mimics the BloxOne PATCH semantics: fields omitted from the request body are left
// untouched on the server. No CSP credentials are required.

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

type mockViewAPI struct {
	mu      sync.Mutex
	views   map[string]map[string]any
	nextID  int
	patches []string
}

func (m *mockViewAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()

	const prefix = "/api/ddi/v1/dns/view"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.NotFound(w, r)
		return
	}
	id := strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, prefix), "/")
	fullID := "dns/view/" + id

	var body map[string]any
	if b, _ := io.ReadAll(r.Body); len(b) > 0 {
		_ = json.Unmarshal(b, &body)
		if r.Method == http.MethodPatch {
			m.patches = append(m.patches, string(b))
		}
	}

	switch {
	case r.Method == http.MethodPost && id == "":
		m.nextID++
		fullID = fmt.Sprintf("dns/view/%d", m.nextID)
		body["id"] = fullID
		fillServerFields(body)
		m.views[fullID] = body
	case r.Method == http.MethodGet:
	case r.Method == http.MethodPatch:
		// PATCH merge: only keys present in the body are changed.
		for k, v := range body {
			m.views[fullID][k] = v
		}
		fillServerFields(m.views[fullID])
	case r.Method == http.MethodDelete:
		delete(m.views, fullID)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	v, ok := m.views[fullID]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"result": v})
}

// fillServerFields sets the fields the real API computes on every write.
func fillServerFields(body map[string]any) {
	body["compartment_id"] = ""
	body["created_at"] = "2024-03-12T19:18:12Z"
	body["updated_at"] = "2024-03-12T19:18:12Z"
	if _, ok := body["tags"]; !ok {
		body["tags"] = map[string]any{}
	}
	body["dnssec_root_keys"] = []any{map[string]any{
		"algorithm": 8, "protocol_zone": ".", "public_key": "AwEAAaz/...74bU=", "sep": true, "zone": ".",
	}}
	za, _ := body["zone_authority"].(map[string]any)
	if za == nil {
		za = map[string]any{}
	}
	za["mname"], za["protocol_mname"] = "ns.b1ddi", "ns.b1ddi"
	za["rname"], za["protocol_rname"] = "hostmaster", "hostmaster"
	body["zone_authority"] = za
}

// simulateUIEdit mimics saving the view in the portal UI, which writes the
// allow/any default into filter_aaaa_acl.
func (m *mockViewAPI) simulateUIEdit() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, v := range m.views {
		v["filter_aaaa_acl"] = []any{map[string]any{"access": "allow", "element": "any", "address": ""}}
	}
}

func reproViewConfig(url, extra string) string {
	return fmt.Sprintf(`
provider "bloxone" {
  csp_url = %q
  api_key = "dummy"
}

resource "bloxone_dns_view" "test" {
  name = "ari-test-iac"
  match_clients_acl = [{ access = "allow", element = "any" }]
  match_destinations_acl = [{ access = "allow", element = "ip", address = "10.192.48.0/28" }]
  %s
}
`, url, extra)
}

// A list that exists on the server but is absent from the config must be cleared
// on apply, not reported as an inconsistent result.
func TestViewResource_Issue138_UnsetFilterAaaaAcl(t *testing.T) {
	mock := &mockViewAPI{views: map[string]map[string]any{}}
	srv := httptest.NewServer(mock)
	defer srv.Close()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: reproViewConfig(srv.URL, ""),
				Check:  resource.TestCheckNoResourceAttr("bloxone_dns_view.test", "filter_aaaa_acl.#"),
			},
			// View saved in the UI -> API now holds filter_aaaa_acl = allow/any.
			{
				PreConfig: mock.simulateUIEdit,
				Config:    reproViewConfig(srv.URL, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("bloxone_dns_view.test", "filter_aaaa_acl.#"),
					func(*terraform.State) error {
						last := mock.patches[len(mock.patches)-1]
						if !strings.Contains(last, `"filter_aaaa_acl":[]`) {
							return fmt.Errorf("expected PATCH to clear filter_aaaa_acl, got: %s", last)
						}
						return nil
					},
				),
			},
		},
	})
}

// An explicit empty list in config must not produce an inconsistent result.
func TestViewResource_Issue138_ExplicitEmptyFilterAaaaAcl(t *testing.T) {
	mock := &mockViewAPI{views: map[string]map[string]any{}}
	srv := httptest.NewServer(mock)
	defer srv.Close()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: reproViewConfig(srv.URL, "")},
			{
				PreConfig: mock.simulateUIEdit,
				Config:    reproViewConfig(srv.URL, "filter_aaaa_acl = []"),
				Check:     resource.TestCheckResourceAttr("bloxone_dns_view.test", "filter_aaaa_acl.#", "0"),
			},
			{
				Config: reproViewConfig(srv.URL, ""),
				Check:  resource.TestCheckNoResourceAttr("bloxone_dns_view.test", "filter_aaaa_acl.#"),
			},
		},
	})
}

// Removing a previously configured ACL from the config must clear it on the server.
func TestViewResource_RemoveQueryAcl(t *testing.T) {
	mock := &mockViewAPI{views: map[string]map[string]any{}}
	srv := httptest.NewServer(mock)
	defer srv.Close()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: reproViewConfig(srv.URL, `query_acl = [{ access = "allow", element = "any" }]`),
				Check:  resource.TestCheckResourceAttr("bloxone_dns_view.test", "query_acl.#", "1"),
			},
			{
				Config: reproViewConfig(srv.URL, ""),
				Check:  resource.TestCheckNoResourceAttr("bloxone_dns_view.test", "query_acl.#"),
			},
		},
	})
}
