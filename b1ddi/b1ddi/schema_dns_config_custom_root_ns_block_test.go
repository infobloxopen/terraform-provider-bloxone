package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigCustomRootNSBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigCustomRootNSBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigCustomRootNSBlock{
				CustomRootNs: []*models.ConfigRootNS{
					{
						Address:      swag.String("unit-test-address"),
						Fqdn:         swag.String("unit-test-fqdn"),
						ProtocolFqdn: "unit-test-fqdn",
					},
				},
				CustomRootNsEnabled: true,
			},
			expect: []interface{}{
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
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigCustomRootNSBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
