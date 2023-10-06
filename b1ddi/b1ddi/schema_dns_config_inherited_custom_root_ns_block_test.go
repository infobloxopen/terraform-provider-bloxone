package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedCustomRootNSBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedCustomRootNSBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedCustomRootNSBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.ConfigCustomRootNSBlock{
					CustomRootNs: []*models.ConfigRootNS{
						{
							Address:      swag.String("unit-test-address"),
							Fqdn:         swag.String("unit-test-fqdn"),
							ProtocolFqdn: "unit-test-fqdn",
						},
					},
					CustomRootNsEnabled: true,
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"custom_root_ns": []map[string]interface{}{
								{
									"address":       swag.String("unit-test-address"),
									"fqdn":          swag.String("unit-test-fqdn"),
									"protocol_fqdn": "unit-test-fqdn",
								},
							},
							"custom_root_ns_enabled": true,
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedCustomRootNSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigInheritedCustomRootNSBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.ConfigInheritedCustomRootNSBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"custom_root_ns": []map[string]interface{}{
								{
									"address":       swag.String("unit-test-address"),
									"fqdn":          swag.String("unit-test-fqdn"),
									"protocol_fqdn": "unit-test-fqdn",
								},
							},
							"custom_root_ns_enabled": true,
						},
					},
				},
			},
			expect: &models.ConfigInheritedCustomRootNSBlock{
				Action: "inherit",
				Source: "unit-test-source",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedCustomRootNSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
