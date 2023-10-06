package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedZoneAuthorityMNameBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedZoneAuthorityMNameBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedZoneAuthorityMNameBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.ConfigZoneAuthorityMNameBlock{
					Mname:           "unit-test-mname",
					ProtocolMname:   "unit-test-mname",
					UseDefaultMname: true,
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"mname":             "unit-test-mname",
							"protocol_mname":    "unit-test-mname",
							"use_default_mname": true,
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedZoneAuthorityMNameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigInheritedZoneAuthorityMNameBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.ConfigInheritedZoneAuthorityMNameBlock
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
							"mname":             "unit-test-mname",
							"protocol_mname":    "unit-test-mname",
							"use_default_mname": true,
						},
					},
				},
			},
			expect: &models.ConfigInheritedZoneAuthorityMNameBlock{
				Action: "inherit",
				Source: "unit-test-source",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedZoneAuthorityMNameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
