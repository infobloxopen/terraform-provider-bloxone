package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritanceInheritedString(t *testing.T) {
	cases := map[string]struct {
		input    interface{}
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput_InheritanceInheritedString": {
			input: &models.InheritanceInheritedString{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-value",
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        "unit-test-value",
				},
			},
		},
		"FullInput_Inheritance2InheritedString": {
			input: &models.Inheritance2InheritedString{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-value",
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        "unit-test-value",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritanceInheritedString(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritanceInheritedString(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.InheritanceInheritedString
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        "unit-test-value",
				},
			},
			expected: &models.InheritanceInheritedString{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritanceInheritedString(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritance2InheritedString(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.Inheritance2InheritedString
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        "unit-test-value",
				},
			},
			expected: &models.Inheritance2InheritedString{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       "unit-test-value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritance2InheritedString(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
