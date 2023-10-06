package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigTrustAnchor(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigTrustAnchor
		expect map[string]interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: map[string]interface{}(nil),
		},
		"FullInput": {
			input: &models.ConfigTrustAnchor{
				Algorithm:    swag.Int64(8),
				ProtocolZone: "unit-test-protocol-zone",
				PublicKey:    swag.String("unit-test-public-key"),
				Sep:          true,
				Zone:         swag.String("unit-test-zone"),
			},
			expect: map[string]interface{}{
				"algorithm":     swag.Int64(8),
				"protocol_zone": "unit-test-protocol-zone",
				"public_key":    swag.String("unit-test-public-key"),
				"sep":           true,
				"zone":          swag.String("unit-test-zone"),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigTrustAnchor(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigTrustAnchor(t *testing.T) {
	cases := map[string]struct {
		input  map[string]interface{}
		expect *models.ConfigTrustAnchor
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: map[string]interface{}{
				"algorithm":     8,
				"protocol_zone": "unit-test-protocol-zone",
				"public_key":    "unit-test-public-key",
				"sep":           true,
				"zone":          "unit-test-zone",
			},
			expect: &models.ConfigTrustAnchor{
				Algorithm: swag.Int64(8),
				PublicKey: swag.String("unit-test-public-key"),
				Sep:       true,
				Zone:      swag.String("unit-test-zone"),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigTrustAnchor(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
