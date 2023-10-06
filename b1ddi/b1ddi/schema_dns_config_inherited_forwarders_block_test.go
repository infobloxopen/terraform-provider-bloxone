package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedForwardersBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedForwardersBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedForwardersBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.ConfigForwardersBlock{
					Forwarders: []*models.ConfigForwarder{
						{
							Address:      swag.String("unit-test-addr"),
							Fqdn:         swag.String("unit.test.fqdn"),
							ProtocolFqdn: "unit-test-protocol",
						},
					},
					ForwardersOnly: true,
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"forwarders": []interface{}{
								map[string]interface{}{
									"address":       swag.String("unit-test-addr"),
									"fqdn":          swag.String("unit.test.fqdn"),
									"protocol_fqdn": "unit-test-protocol",
								},
							},
							"forwarders_only": true,
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedForwardersBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigInheritedForwardersBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.ConfigInheritedForwardersBlock
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
							"forwarders": []interface{}{
								map[string]interface{}{
									"address":       swag.String("unit-test-addr"),
									"fqdn":          swag.String("unit.test.fqdn"),
									"protocol_fqdn": "unit-test-protocol",
								},
							},
							"forwarders_only": true,
						},
					},
				},
			},
			expect: &models.ConfigInheritedForwardersBlock{
				Action: "inherit",
				Source: "unit-test-source",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedForwardersBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
