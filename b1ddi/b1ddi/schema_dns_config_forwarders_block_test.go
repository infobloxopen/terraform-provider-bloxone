package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigForwardersBlock(t *testing.T) {
	cases := map[string]struct {
		input    *models.ConfigForwardersBlock
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigForwardersBlock{
				Forwarders: []*models.ConfigForwarder{
					{
						Address:      swag.String("unit-test-addr"),
						Fqdn:         swag.String("unit.test.fqdn"),
						ProtocolFqdn: "unit-test-protocol",
					},
				},
				ForwardersOnly: true,
			},
			expected: []interface{}{
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
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigForwardersBlock(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandConfigForwardersBlock(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.ConfigForwardersBlock
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullConfig": {
			input: []interface{}{
				map[string]interface{}{
					"forwarders": []interface{}{
						map[string]interface{}{
							"address":       "unit-test-addr",
							"fqdn":          "unit.test.fqdn",
							"protocol_fqdn": "unit-test-protocol",
						},
					},
					"forwarders_only": true,
				},
			},
			expected: &models.ConfigForwardersBlock{
				Forwarders: []*models.ConfigForwarder{
					{
						Address: swag.String("unit-test-addr"),
						Fqdn:    swag.String("unit.test.fqdn"),
					},
				},
				ForwardersOnly: true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigForwardersBlock(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
