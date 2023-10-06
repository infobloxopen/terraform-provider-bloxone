package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigDNSSECValidationBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigDNSSECValidationBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigDNSSECValidationBlock{
				DnssecEnableValidation: true,
				DnssecEnabled:          true,
				DnssecTrustAnchors: []*models.ConfigTrustAnchor{
					{
						Algorithm:    swag.Int64(8),
						ProtocolZone: "unit-test-protocol-zone",
						PublicKey:    swag.String("unit-test-public-key"),
						Sep:          true,
						Zone:         swag.String("unit-test-zone"),
					},
				},
				DnssecValidateExpiry: true,
			},
			expect: []interface{}{
				map[string]interface{}{
					"dnssec_enable_validation": true,
					"dnssec_enabled":           true,
					"dnssec_trust_anchors": []interface{}{
						map[string]interface{}{
							"algorithm":     swag.Int64(8),
							"protocol_zone": "unit-test-protocol-zone",
							"public_key":    swag.String("unit-test-public-key"),
							"sep":           true,
							"zone":          swag.String("unit-test-zone"),
						},
					},
					"dnssec_validate_expiry": true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigDNSSECValidationBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
