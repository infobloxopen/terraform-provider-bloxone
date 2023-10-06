package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigRootNS(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigRootNS
		expect map[string]interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: map[string]interface{}(nil),
		},
		"FullInput": {
			input: &models.ConfigRootNS{
				Address:      swag.String("unit-test-address"),
				Fqdn:         swag.String("unit-test-fqdn"),
				ProtocolFqdn: "unit-test-fqdn",
			},
			expect: map[string]interface{}{
				"address":       swag.String("unit-test-address"),
				"fqdn":          swag.String("unit-test-fqdn"),
				"protocol_fqdn": "unit-test-fqdn",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigRootNS(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigRootNS(t *testing.T) {
	cases := map[string]struct {
		input  map[string]interface{}
		expect *models.ConfigRootNS
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: map[string]interface{}{
				"address":       "unit-test-address",
				"fqdn":          "unit-test-fqdn",
				"protocol_fqdn": "unit-test-fqdn",
			},
			expect: &models.ConfigRootNS{
				Address: swag.String("unit-test-address"),
				Fqdn:    swag.String("unit-test-fqdn"),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigRootNS(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
