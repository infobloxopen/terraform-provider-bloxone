package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritanceInheritedIdentifier(t *testing.T) {
	cases := map[string]struct {
		input  *models.InheritanceInheritedIdentifier
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.InheritanceInheritedIdentifier{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-identifier",
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        "unit-test-identifier",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritanceInheritedIdentifier(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandInheritanceInheritedIdentifier(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.InheritanceInheritedIdentifier
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
					"value":        "unit-test-identifier",
				},
			},
			expect: &models.InheritanceInheritedIdentifier{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-identifier",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritanceInheritedIdentifier(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
