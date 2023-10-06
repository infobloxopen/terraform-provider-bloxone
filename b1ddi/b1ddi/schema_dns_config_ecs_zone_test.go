package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigECSZone(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigECSZone
		expect map[string]interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: &models.ConfigECSZone{
				Access:       swag.String("unit-test-access"),
				Fqdn:         swag.String("unit-test-fqdn"),
				ProtocolFqdn: "unit-test-fqdn",
			},
			expect: map[string]interface{}{
				"access":        swag.String("unit-test-access"),
				"fqdn":          swag.String("unit-test-fqdn"),
				"protocol_fqdn": "unit-test-fqdn",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigECSZone(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigECSZone(t *testing.T) {
	cases := map[string]struct {
		input  map[string]interface{}
		expect *models.ConfigECSZone
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: map[string]interface{}{
				"access":        "unit-test-access",
				"fqdn":          "unit-test-fqdn",
				"protocol_fqdn": "unit-test-fqdn",
			},
			expect: &models.ConfigECSZone{
				Access: swag.String("unit-test-access"),
				Fqdn:   swag.String("unit-test-fqdn"),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigECSZone(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
