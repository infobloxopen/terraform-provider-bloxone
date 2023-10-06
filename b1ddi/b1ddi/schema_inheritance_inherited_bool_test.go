package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritanceInheritedBool(t *testing.T) {
	cases := map[string]struct {
		input    interface{}
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput_InheritanceInheritedBool": {
			input: &models.InheritanceInheritedBool{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       true,
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        true,
				},
			},
		},
		"FullInput_Inheritance2InheritedBool": {
			input: &models.Inheritance2InheritedBool{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       true,
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritanceInheritedBool(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritanceInheritedBool(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.InheritanceInheritedBool
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullConfig": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        true,
				},
			},
			expected: &models.InheritanceInheritedBool{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritanceInheritedBool(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritance2InheritedBool(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.Inheritance2InheritedBool
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullConfig": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        true,
				},
			},
			expected: &models.Inheritance2InheritedBool{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritance2InheritedBool(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
