package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritedDHCPConfigFilterList(t *testing.T) {
	cases := map[string]struct {
		input  *models.InheritedDHCPConfigFilterList
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.InheritedDHCPConfigFilterList{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []string{
					"unit-test-value",
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []string{
						"unit-test-value",
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritedDHCPConfigFilterList(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandInheritedDHCPConfigFilterList(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.InheritedDHCPConfigFilterList
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullConfig": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []string{
						"unit-test-value",
					},
				},
			},
			expect: &models.InheritedDHCPConfigFilterList{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []string{
					"unit-test-value",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritedDHCPConfigFilterList(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
