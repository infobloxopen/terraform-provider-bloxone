package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedECSBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedECSBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedECSBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.ConfigECSBlock{
					EcsEnabled:    true,
					EcsForwarding: true,
					EcsPrefixV4:   20,
					EcsPrefixV6:   20,
					EcsZones: []*models.ConfigECSZone{
						{
							Access:       swag.String("unit-test-access"),
							Fqdn:         swag.String("unit-test-fqdn"),
							ProtocolFqdn: "unit-test-fqdn",
						},
					},
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"ecs_enabled":    true,
							"ecs_forwarding": true,
							"ecs_prefix_v4":  int64(20),
							"ecs_prefix_v6":  int64(20),
							"ecs_zones": []interface{}{
								map[string]interface{}{
									"access":        swag.String("unit-test-access"),
									"fqdn":          swag.String("unit-test-fqdn"),
									"protocol_fqdn": "unit-test-fqdn",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedECSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigInheritedECSBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.ConfigInheritedECSBlock
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
							"ecs_enabled":    true,
							"ecs_forwarding": true,
							"ecs_prefix_v4":  int64(20),
							"ecs_prefix_v6":  int64(20),
							"ecs_zones": []interface{}{
								map[string]interface{}{
									"access":        swag.String("unit-test-access"),
									"fqdn":          swag.String("unit-test-fqdn"),
									"protocol_fqdn": "unit-test-fqdn",
								},
							},
						},
					},
				},
			},
			expect: &models.ConfigInheritedECSBlock{
				Action: "inherit",
				Source: "unit-test-source",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedECSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
