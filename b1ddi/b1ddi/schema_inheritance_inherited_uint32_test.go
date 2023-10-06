package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritanceInheritedUInt32(t *testing.T) {
	cases := map[string]struct {
		input    interface{}
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput_InheritanceInheritedUInt32": {
			input: &models.InheritanceInheritedUInt32{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       int64(20),
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        int64(20),
				},
			},
		},
		"FullInput_Inheritance2InheritedUInt32": {
			input: &models.Inheritance2InheritedUInt32{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       int64(20),
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        int64(20),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritanceInheritedUInt32(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritanceInheritedUInt32(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.InheritanceInheritedUInt32
	}{
		"NilInput": {
			input:    []interface{}{},
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        20,
				},
			},
			expected: &models.InheritanceInheritedUInt32{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       int64(20),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritanceInheritedUInt32(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritance2InheritedUInt32(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.Inheritance2InheritedUInt32
	}{
		"NilInput": {
			input:    []interface{}{},
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value":        20,
				},
			},
			expected: &models.Inheritance2InheritedUInt32{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value:       int64(20),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritance2InheritedUInt32(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
