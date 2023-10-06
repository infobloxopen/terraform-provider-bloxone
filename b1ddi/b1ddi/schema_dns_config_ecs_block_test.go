package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigECSBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigECSBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigECSBlock{
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
			expect: []interface{}{
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
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigECSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
